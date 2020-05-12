package fsmanager

import (
	"io"
	"os"

	"github.com/amlwwalker/pickleit/utilities"
)

const (
	//largest file that we handle in memory. Else create a tmp file
	MEM_FILE_SIZE_BUFFER = 512000000
)

/*
	this file will manage all copy routines of the original
	to a tmp location for diff creation etc. We do not perform these operations
	on the original file incase of any performance issues experienced by the user
	or incase we corrupt a file.
	This does result in us having to keep previous versions available at all times
	so that we can create all the diffs required (bi-directional)

	According to https://opensource.com/article/18/6/copying-files-go
	we can safely use io.Copy to copy large files, however we could in theory, have multiple
	copy methods available for smaller files, and keep them in memory for the time required
	to create the patches
*/

// Copy checks the size of the file and decides what copy method to use
// Later, small files can be copied into memory instead
func (w *Watcher) CleverCopy(src, dst string) (int64, error) {
	var err error
	// var fileName string
	if _, err = utilities.VerifySrcFile(src); err != nil {
		return 0, err
	}
	var nBytes int64
	// if sourceFileStat.Size() > MEM_FILE_SIZE_BUFFER { //bytes
	if nBytes, err = copyFS(src, dst); err != nil {
		w.ErrorF("Error copying file to dst %s\r\n", err)
	}
	// }
	w.Info("Finished copying file")
	return nBytes, err
}

//ForceFSCopy will always copy the src file to the dst provided.
// The difference is the Copy function will try to manage the file in memory
// if it can. This function allows memory to be ignored
func (w *Watcher) ForceFSCopy(src, dst string) (int64, error) {
	var err error
	if _, err = utilities.VerifySrcFile(src); err != nil {
		return 0, err
	}
	var nBytes int64
	if nBytes, err = copyFS(src, dst); err != nil {
		w.ErrorF("Error copying file to dst %s\r\n", err)
	}
	return nBytes, err
}

// copyFS copies a file from a source to a destination.
// The file is not stored in memory
func copyFS(src, dst string) (int64, error) {
	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}

// copyFileIntoMemory will copy the src into memory
func copyFileIntoMemory(src string) (int64, []byte, error) {
	file, err := os.Open(src)
	if err != nil {
		return 0, []byte{}, err
	}
	defer file.Close()

	fileinfo, err := file.Stat()
	if err != nil {
		return 0, []byte{}, err
	}

	filesize := fileinfo.Size()
	buffer := make([]byte, filesize)

	_, err = file.Read(buffer) //bytesRead
	if err != nil {
		return 0, []byte{}, err
	}

	return filesize, buffer, nil
}

// creates a stream of chunks of data to process a file
// on the fly
func readFileChunksIntoMemory(src string) error {
	const BufferSize = 100
	file, err := os.Open(src)
	if err != nil {
		return err
	}
	defer file.Close()

	buffer := make([]byte, BufferSize)

	for {
		_, err := file.Read(buffer) //bytesread

		if err != nil {
			if err != io.EOF {
				return err
			}

			break
		}
	}
	return nil
}
