// +build !mock

package model

// import (
// 	"encoding/json"
// 	"fmt"
// 	"os"
// 	"path/filepath"
// 	"strconv"
// 	"time"

// 	"github.com/amlwwalker/pickleit/utilities"
// 	"github.com/asdine/storm"
// 	"github.com/boltdb/bolt"
// )

// var DEMO = false
// var boltdb, _ = storm.Open(filepath.Join(os.Getenv("HOME"), "workhorse", "workhorse.db"))

// func init() {
// 	if DEMO {
// 		fmt.Println("init with boltdb in demo mode")
// 		if err := boltdb.Update(func(tx *bolt.Tx) error {

// 			for _, file := range db {
// 				fileBucket, err := tx.CreateBucketIfNotExists([]byte(strconv.Itoa(file.ID)))
// 				if err != nil {
// 					println("failed to create file bucket")
// 					return err
// 				}

// 				jsn, err := json.Marshal(file)
// 				if err != nil {
// 					return err
// 				}
// 				fileBucket.Put([]byte("json"), jsn)
// 			}

// 			return nil
// 		}); err != nil {
// 			println("failed to create boltdb:", err.Error())
// 		}
// 	}
// }

// func getFileCount() int {
// 	var files []utilities.File
// 	boltdb.All(&files)
// 	fmt.Println("file count ", len(files))
// 	return len(files)
// }

// func getFileForID(ID int) *utilities.File {
// 	var file utilities.File
// 	boltdb.One("ID", ID, &file)
// 	return &file
// }

// func GetFileForName(name string) *utilities.File {
// 	var file utilities.File
// 	boltdb.One("Path", name, &file)
// 	return &file
// }

// func getDiffCount() int {
// 	var diffs []utilities.DiffObject
// 	if err := boltdb.All(&diffs); err != nil {
// 		// db.ErrorF("Error retrieving all diffs %s", err)
// 		return 0
// 	} else {
// 		return len(diffs)
// 	}
// }

// // RetrieveDiffsForFileByPath returns the diffs for a file base on the file path. Diffs are very specific to a file,
// // so this may not be as robust as searching by diff, however we may not have the diff available
// func GetDiffCountForFile(name string) int {
// 	var objDiffs []utilities.DiffObject
// 	if err := boltdb.Find("Object", name, &objDiffs); err != nil && err.Error() != "not found" {
// 		return 0
// 	}
// 	return len(objDiffs)
// }

// func DeleteFile(ID int) {
// 	fmt.Println("DeleteFile has not been implemented yet, id ", ID)
// }
// func DeleteDiff(ID int) {
// 	fmt.Println("DeleteDiff: diff with id ", ID)

// 	var diff utilities.DiffObject
// 	boltdb.One("ID", ID, &diff)
// 	boltdb.DeleteStruct(diff)

// }

// func CreateNewFile(fileId int, name string) {
// 	file := &utilities.File{fileId, "no-path", name, "no-bkp-location", "current-base", [16]byte{}, time.Now(), "12345", 0.12}

// 	if err := boltdb.Update(func(tx *bolt.Tx) error {

// 		fileBucket, err := tx.CreateBucketIfNotExists([]byte(strconv.Itoa(file.ID)))
// 		if err != nil {
// 			println("failed to create file bucket")
// 			return err
// 		}

// 		jsn, err := json.Marshal(file)
// 		if err != nil {
// 			return err
// 		}
// 		fileBucket.Put([]byte("json"), jsn)

// 		return nil
// 	}); err != nil {
// 		println("failed to CreateNewFile:", err.Error())
// 	}

// }

// func CreateNewDiff(fileId, diffId int, subjectHash, object, diffPath string, year int) {
// 	holder := []byte{}
// 	diff := &utilities.DiffObject{diffId, subjectHash, object, [16]byte{}, [16]byte{}, "watching", "diffpath", "label", "screenshot", true, true, "description", nil, &holder, 1, time.Now(), "message"}
// 	if err := boltdb.Save(&diff); err != nil {
// 		fmt.Println("error saving diff ", err)
// 	}
// }

// func GetNextFileID() int {
// 	count := getFileCount()
// 	return count + 1
// }

// func GetNextDiffID() int {
// 	count := getDiffCount()
// 	return count + 1
// }

// // RetrieveWatchedFiles all files that are in the database as "watched files"
// // This can be used to trigger the same files to be watched again
// func retrieveAllDiffs() ([]utilities.DiffObject, error) {
// 	var diffs []utilities.DiffObject
// 	if err := boltdb.All(&diffs); err != nil {
// 		// db.ErrorF("Error retrieving all diffs %s", err)
// 		return []utilities.DiffObject{}, err
// 	} else {
// 		return diffs, nil
// 	}
// }

// // FindFileByPath is a search function looking for file details
// func findFileByPath(filePath string) (utilities.File, error) {
// 	var file utilities.File
// 	if err := boltdb.One("Path", filePath, &file); err != nil {
// 		// db.ErrorF("Error finding file by path %s", err)
// 		return utilities.File{}, err
// 	}
// 	return file, nil
// }
// func getModelArray() []*dbArrayStruct {
// 	o := make([]*dbArrayStruct, 0)
// 	//retrieve all the diffs
// 	diffs, _ := retrieveAllDiffs()
// 	//retrieve all the files for the diff (this is inefficient hopefully obviously)
// 	for i := range diffs {
// 		fileForDiff, _ := findFileByPath(diffs[i].Object)
// 		tmp := &dbArrayStruct{&fileForDiff, &diffs[i]}
// 		o = append(o, tmp)
// 	}
// 	// fmt.Printf("o %+v\r\n", o)
// 	//for each diff, create an dbArrayStruct object
// 	//put the file and the diff in and add it to the array
// 	return o
// }

// // func getModelArray() []*dbArrayStruct {

// // 	m := make(map[int]*dbArrayStruct)

// // 	//for each file
// // 	//get all the diffs related to it
// // 	//create a map of the id of the file to the array struct (file and diff)
// // 	for i := 0; i < getFileCount(); i++ {
// // 		file := getFileForID(i)
// // 		for _, diff := range file.Diffs {
// // 			m[diff.ID] = &dbArrayStruct{file, diff}
// // 		}
// // 	}

// // 	//then make an array of file:diff combinations
// // 	o := make([]*dbArrayStruct, 0)

// // 	//this is an ordered array of file and diff based on the order of the diffs in the database
// // 	//so what you need to do is create an array of dbArrayStruct objects that contains the file and the diff
// // 	//and then stick those in an array based on the order of the diffs in the database
// // 	//for each diff
// // 	for i := 0; i <= GetNextDiffID(); i++ {
// // 		//get the diff based on the ids in the map
// // 		if s, ok := m[i]; ok {
// // 			//now append the value from the map, based on the diff id to the
// // 			//array.
// // 			o = append(o, s)
// // 		}
// // 	}

// // 	return o
// // }
