package fsmanager

import (
	logger "github.com/apsdehal/go-logger"
)

// The watcher is responsible for not only seeing when a file changes,
// but also keeping track of
// * the file hash so that if it changes again any modifications can be handled
// * copying any versions and keeping them safe (even if temporary)
// * creating the diff of the file, in both directions if necessary
// * storing the details in the database
func NewPatcher(logger *logger.Logger, KEYFOLDER, DOWNLOADFOLDER, SYNCFOLDER, THUMBFOLDER, DIFFFOLDER string) (Patcher, error) {
	p := Patcher{
		logger,
		KEYFOLDER, DOWNLOADFOLDER, SYNCFOLDER, THUMBFOLDER, DIFFFOLDER,
	}
	return p, nil
}

// PatchFromFile takes the version of the file that was backed up
// and applies the specified patch to it, to get the latest file. This is incase the
// last save is the file you want to get.
func (p *Patcher) PatchFromFile(filePath, patchPath, restorePath string) error {
	if subject, err := openFile(filePath); err != nil {
		p.ErrorF("error on subject file: ", err)
	} else if patch, err := openFile(patchPath); err != nil {
		p.ErrorF("error on patch file: ", err)
	} else {
		return p.applyPatch(subject, patch, restorePath)
	}
	return nil
}

//applyPatch actively applies the patch to the subject. This could eventually
// be upgraded for different patching algorithms
func (p *Patcher) applyPatch(subject, patch []byte, restorePath string) error {
	if delta, err := decompressDelta(patch); err != nil {
		p.ErrorF("error decompressing delta", err)
	} else {
		if appliedBytes, err := applyPatchToFile(subject, delta); err != nil {
			p.ErrorF("error applying delta to original file", err)
			return err
		} else if err := writeFile(restorePath, appliedBytes); err != nil {
			p.ErrorF("error writing patchedFile", err)
			return err
		} else {
			return nil
		}
	}
	return nil
}
