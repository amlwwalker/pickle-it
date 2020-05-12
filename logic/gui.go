// +build !mock

package logic

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/amlwwalker/pickleit/utilities"
	"github.com/boltdb/bolt"
)

func (m *Manager) GetFileCount() int {
	var files []utilities.File
	m.dB.All(&files)
	return len(files)
}

func (m *Manager) GetFileForID(ID int) *utilities.File {
	var file utilities.File
	m.dB.One("ID", ID, &file)
	return &file
}

func (m *Manager) GetFileForName(name string) *utilities.File {
	var file utilities.File
	m.dB.One("Path", name, &file)
	return &file
}

func (m *Manager) GetDiffCount() int {
	var diffs []utilities.DiffObject
	if err := m.dB.All(&diffs); err != nil {
		// db.ErrorF("Error retrieving all diffs %s", err)
		return 0
	} else {
		return len(diffs)
	}
}

// RetrieveDiffsForFileByPath returns the diffs for a file base on the file path. Diffs are very specific to a file,
// so this may not be as robust as searching by diff, however we may not have the diff available
func (m *Manager) GetDiffCountForFile(name string) int {
	var objDiffs []utilities.DiffObject
	fmt.Println("looking for name ", name)
	if err := m.dB.Find("Object", name, &objDiffs); err != nil && err.Error() != "not found" {
		fmt.Println("error is ", err)
		return 0
	}
	return len(objDiffs)
}

func (m *Manager) DeleteFile(ID int) {
	fmt.Println("DeleteFile has not been implemented yet, id ", ID)
}
func (m *Manager) DeleteDiff(ID int) {
	fmt.Println("DeleteDiff: diff with id ", ID)

	var diff utilities.DiffObject
	if err := m.dB.One("ID", ID, &diff); err != nil {
		fmt.Println("there was an error retrieving the diff with ID ", ID, " err ", err)
	} else if err := m.dB.DeleteStruct(&diff); err != nil {
		fmt.Println("there was an error deleting the diff with ID ", ID, " err ", err)
	}

}
func (m *Manager) UpdateDiffDescription(ID int, description string) error {
	return m.dB.UpdateDescription(ID, description)
}

func (m *Manager) CreateNewFile(fileId int, name string) {
	file := &utilities.File{fileId, "no-path", name, "no-bkp-location", "current-base", [16]byte{}, time.Now(), "12345", 0.12, false}

	if err := m.dB.Update(func(tx *bolt.Tx) error {

		fileBucket, err := tx.CreateBucketIfNotExists([]byte(strconv.Itoa(file.ID)))
		if err != nil {
			println("failed to create file bucket")
			return err
		}

		jsn, err := json.Marshal(file)
		if err != nil {
			return err
		}
		fileBucket.Put([]byte("json"), jsn)

		return nil
	}); err != nil {
		println("failed to CreateNewFile:", err.Error())
	}

}

func (m *Manager) CreateNewDiff(fileId, diffId int, subjectHash, object, diffPath string, year int) {
	// holder := []byte{}
	var diff utilities.DiffObject
	diff.Description = "description not set...."
	diff.Direction = true
	diff.Subject = "object"
	diff.Object = "subject"
	diff.StartTime = time.Now()
	diff.SubjectHash = [16]byte{}
	diff.ObjectHash = [16]byte{}
	if err := m.dB.Save(&diff); err != nil {
		fmt.Println("error saving diff ", err)
	}
}

func (m *Manager) GetNextFileID() int {
	count := m.GetFileCount()
	return count + 1
}

func (m *Manager) GetNextDiffID() int {
	count := m.GetDiffCount()
	return count + 1
}

// RetrieveWatchedFiles all files that are in the database as "watched files"
// This can be used to trigger the same files to be watched again
func (m *Manager) RetrieveAllDiffs() ([]utilities.DiffObject, error) {
	var diffs []utilities.DiffObject
	if err := m.dB.All(&diffs); err != nil {
		// db.ErrorF("Error retrieving all diffs %s", err)
		return []utilities.DiffObject{}, err
	} else {
		return diffs, nil
	}
}

// FindFileByPath is a search function looking for file details
func (m *Manager) FindFileByPath(filePath string) (utilities.File, error) {
	var file utilities.File
	if err := m.dB.One("Path", filePath, &file); err != nil {
		// db.ErrorF("Error finding file by path %s", err)
		return utilities.File{}, err
	}
	return file, nil
}
