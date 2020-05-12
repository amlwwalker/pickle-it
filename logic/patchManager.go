package logic

import (
	"path"
	"path/filepath"
	"strings"
)

//just an extension of the manager, but these functions are specifically for patching

// BeginForwardPatch is a helper function while patching has not been completely worked out
// and supports creating a new unique name for the file once it has been patched
func (m *Manager) BeginForwardPatch(filePath, patchPath, restoredFile string) error {
	// So we know where the file is
	// which means we should be able to find the file in the sync directory
	// and also its patches, based on the direction.
	// Lets get those

	//need to split the file name so you can rebuild it with the backup naming convention
	// using the manager VERSION_COUNTER variable before the extension
	// TODO: When you use a database, this naming convention should change drastically
	// as it must be mitigated against the possibility of two files, with different paths (or from)
	// different sources having the same name, and therefore resulting in collisions with the patches

	var err error
	var bkdUpFileName string
	if filePath, err = filepath.Abs(filePath); err != nil {
		m.ErrorF("Error making path absolute %s", err)
		return err
	}
	//the backup we are patching is now the subject of the patch as it will change depending on the base at the time the patch was made so we can't do it based on the location, we it based on the patch
	if dbFile, err := m.dB.FindDiffByPath(patchPath); err != nil {
		m.ErrorF("error finding file in database %s\r\n", err)
		return err
	} else {
		bkdUpFileName = dbFile.Subject
	}

	//you can supply a restoreFile location or it can be generated
	if restoredFile == "" {
		restorePath, fileName := filepath.Split(filePath)
		fileExtension := path.Ext(filePath)                       //get the file extension
		filePrefix := strings.TrimSuffix(fileName, fileExtension) //get the name without the extension
		if m.Settings.systemSettings.OverrideExisting {
			restoredFile = filepath.Join(restorePath, filePrefix+fileExtension)
			m.WarningF("writing patch to original path %s\r\n", restoredFile)
			op := Op_Message.Retrieve()
			op.CustomField = "overwritten file: " + restoredFile
			m.Informer <- op
		} else {
			restoredFile = filepath.Join(restorePath, filePrefix+"_restored_"+fileExtension)
			m.InfoF("writing patch to new path %s\r\n", restoredFile)
			op := Op_Message.Retrieve()
			op.CustomField = "patched to new file: " + restoredFile
			m.Informer <- op
		}
	}

	m.InfoF("file path %s", filePath)
	m.InfoF("restore path %s", restoredFile)
	m.InfoF("subject path %s", bkdUpFileName)
	m.InfoF("patch path %s", patchPath)
	//now patch the file
	// RetrieveDelta()
	watcherState := m.watcher.Ignore()
	m.InfoF("Temporarily ignoring changes while patching %b\r\n", watcherState)
	m.watcher.Ignore()
	err = m.patcher.PatchFromFile(bkdUpFileName, patchPath, restoredFile)
	return err
}
