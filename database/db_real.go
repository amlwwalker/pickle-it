// +build !mock

package database

import (
	"fmt"

	"github.com/amlwwalker/pickleit/utilities"
	"github.com/asdine/storm"
)

// ConfigureDB sets up bolt and Storm according to the path of the database
// this is done here so that different databases can be configured in different scenarios
func (db *DB) ConfigureDB(dbPath string) error {
	var err error
	if db.DB, err = storm.Open(dbPath); err != nil {
		return err
	}
	var file utilities.File
	if err := db.One("Name", "-- All Files --", &file); err != nil {
		if err.Error() != "not found" {
			db.ErrorF("Error finding file by path %s", err)
			return err
		}
		db.WarningF("No file found. initialising the database")
		file.Name = "-- All Files --"
		file.Ignore = true //this is currently not used however could result in this file being ignored when file watching (etc) starts
		if err := db.Save(&file); err != nil {
			db.ErrorF("Error storing the diff %s", err)
			return err
		}
	}
	return nil
}

// CheckIfFileCurrentlyMonitored checks if the file is already being monitored. This is a read-only check
// to see whether the file was correctly initialised
// (BUG) The hash causes the same file to be in database multiple times!
func (db *DB) CheckIfFileCurrentlyMonitored(src string, hash [16]byte) (utilities.File, error) {
	var file utilities.File

	//TODO: check this actually works still (don't need hash passed to this anymore)
	if err := db.One("Path", src, &file); err != nil {
		if err.Error() != "not found" {
			db.ErrorF("Error finding file by path %s", err)
			return utilities.File{}, err
		}
		db.WarningF("no file found, %s", err)
		return utilities.File{}, err
	} else {
		return file, nil
	}
}

// RetrieveWatchedFiles all files that are in the database as "watched files"
// This can be used to trigger the same files to be watched again
func (db *DB) RetrieveWatchedFiles() ([]utilities.File, error) {
	var files []utilities.File
	if err := db.All(&files); err != nil {
		db.ErrorF("Error retrieving all watched files %s", err)
		return []utilities.File{}, err
	} else {
		return files, nil
	}
}

// RetrieveWatchedFiles all files that are in the database as "watched files"
// This can be used to trigger the same files to be watched again
func (db *DB) RetrieveAllDiffs() ([]utilities.DiffObject, error) {
	var diffs []utilities.DiffObject
	if err := db.All(&diffs); err != nil {
		db.ErrorF("Error retrieving all diffs %s", err)
		return []utilities.DiffObject{}, err
	} else {
		return diffs, nil
	}
}

// InitialiseFileInDatabase should be called before any file is copied/renamed/diff'd/patched,
// and this should be checked before any operation occurs on a file. Any loss of data is completely as a result
// of losing references
func (db *DB) InitialiseFileInDatabase(file utilities.File) (int, error) {
	if err := db.Save(&file); err != nil {
		db.ErrorF("Error initialising file in database %s", err)
		return file.ID, err
	}
	return file.ID, nil
}

// FindFileByPath is a search function looking for file details
func (db *DB) FindFileByPath(filePath string) (utilities.File, error) {
	var file utilities.File
	if err := db.One("Path", filePath, &file); err != nil {
		db.ErrorF("Error finding file by path %s", err)
		return utilities.File{}, err
	}
	return file, nil
}

// FindFileByID is a search function looking for file details
func (db *DB) FindFileByID(ID int) (utilities.File, error) {
	var file utilities.File
	if err := db.One("ID", ID, &file); err != nil {
		db.ErrorF("Error finding file by path %s", err)
		return utilities.File{}, err
	}
	return file, nil
}

// UpdateBase updates the current base file that diffs will compare to
func (db *DB) UpdateFileData(filePath, basePath string, hash [16]byte) error {
	if file, err := db.FindFileByPath(filePath); err != nil {
		db.ErrorF("Error updating the file base %s", err)
		return err
	} else {
		err := db.Update(&utilities.File{ID: file.ID, CurrentBase: basePath, CurrentHash: hash})
		return err
	}
}

// RetrieveDiffsForFileByHash returns the diffs for a file. Diffs can be applied to a specific file (by its hash),
// so when looking for the diffs, the hash is a good place to start in terms of finding the diffs
func (db *DB) RetrieveDiffsForFileByHash(fileHash [16]byte, direction bool) ([]utilities.DiffObject, error) {
	var diffs []utilities.DiffObject
	var field string
	if direction {
		field = "ObjectHash"
	} else {
		field = "SubjectHash"
	}
	if err := db.Find(field, fileHash, &diffs); err != nil {
		return []utilities.DiffObject{}, err
	}
	return diffs, nil
}

// RetrieveDiffsForFileByPath returns the diffs for a file base on the file path. Diffs are very specific to a file,
// so this may not be as robust as searching by diff, however we may not have the diff available
func (db *DB) RetrieveDiffsForFileByPath(filePath string) ([]utilities.DiffObject, error) {
	var objDiffs []utilities.DiffObject
	var subDiffs []utilities.DiffObject
	if err := db.Find("Object", filePath, &objDiffs); err != nil && err.Error() != "not found" {
		db.ErrorF("Error finding diff by object %s", err)
		return []utilities.DiffObject{}, err
	}
	if err := db.Find("Subject", filePath, &subDiffs); err != nil && err.Error() != "not found" {
		db.ErrorF("Error finding diff by subject %s", err)
		return []utilities.DiffObject{}, err
	}
	return append(objDiffs, subDiffs...), nil
}

// StoreDiff just places the information about a diff in the database. Currently there is no protection
// to stop the entire diff entering the database (if fs is false), which may be very slow/bulky...
// TODO: decide what to do with diffs in memory
func (db *DB) StoreDiff(diff utilities.DiffObject) error {
	if err := db.Save(&diff); err != nil {
		db.ErrorF("Error storing the diff %s", err)
		return err
	}
	return nil
}

// FindDiffByPath is a search function looking for a diff
func (db *DB) FindDiffByPath(patchPath string) (utilities.DiffObject, error) {
	var diff utilities.DiffObject
	if err := db.One("DiffPath", patchPath, &diff); err != nil {
		db.ErrorF("Error finding diff by path %s", err)
		return utilities.DiffObject{}, err
	}
	return diff, nil
}

//RetrieveDiffsByID returns a diff based on the id it has in the database
func (db *DB) RetrieveDiffsByID(ID int) (utilities.DiffObject, error) {
	var diff utilities.DiffObject
	if err := db.One("ID", ID, &diff); err != nil {
		db.ErrorF("Error finding diff by ID %s", err)
		return utilities.DiffObject{}, err
	}
	return diff, nil
}

// ChangeLabel is a simple function to set the label on a patch (potentially deprecated. What's it doing?)
func (db *DB) ChangeLabel(patchID int, lb string) error {
	if err := db.Update(&utilities.DiffObject{ID: patchID, Label: lb}); err != nil {
		db.ErrorF("Error changing diff label %s", err)
		return err
	}
	return nil
}

// UpdateDescription is a simple function to set the label on a patch
func (db *DB) UpdateDescription(patchID int, description string) error {
	fmt.Println("attempting to path with id ", patchID, " description ", description)
	if err := db.Update(&utilities.DiffObject{ID: patchID, Description: description}); err != nil {
		db.ErrorF("Error changing diff label %s", err)
		return err
	}
	return nil
}

func (db *DB) StoreUXSettings(settings utilities.UXSettings) error {
	settings.ID = 1
	if err := db.Save(&settings); err != nil {
		db.ErrorF("Error storing the settings %s", err)
		return err
	}
	return nil
}
func (db *DB) RetrieveUXSettings() (utilities.UXSettings, error) {
	var settings utilities.UXSettings
	if err := db.One("ID", 1, &settings); err != nil {
		db.ErrorF("Error retrieving the settings %s", err)
		return utilities.UXSettings{}, err
	} else {
		return settings, nil
	}
}

// if file, err := db.FindFileByPath(filePath); err != nil {
// 	db.ErrorF("Error updating the file base %s", err)
// 	return err
// } else {
// 	err := db.Update(&utilities.File{ID: file.ID, CurrentBase: basePath, CurrentHash: hash})
// 	return err
// }
