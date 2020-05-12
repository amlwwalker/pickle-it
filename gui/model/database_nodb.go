// +build mock

package model

import "github.com/amlwwalker/pickleit/utilities"

var DEMO = true

func getFileCount() int {
	return len(db)
}

func getFileForID(ID int) *utilities.File {
	return db[ID]
}

func GetFileForName(name string) *utilities.File {
	for _, file := range db {
		if file.Name == name {
			return file
		}
	}
	return nil
}

func getDiffCount() int {
	var diffCount int
	for _, files := range db {
		diffCount += len(files.Diffs)
	}
	return diffCount
}

func GetDiffCountForFile(name string) int {
	return len(GetFileForName(name).Diffs)
}

func DeleteDiff(ID int) {
	for _, file := range db {
		for _, diff := range file.Diffs {
			if diff.ID == ID {
				delete(file.Diffs, diff.ID)
			}
		}
	}
}

func CreateNewFile(fileId int, name string) {
	db[fileId] = &utilities.File{fileId, "no-path", name, "no-bkp-location", make(map[int]*utilities.DiffObject)}
}

func CreateNewDiff(fileId, diffId int, subjectHash, object, diffPath string, year int) {
	db[fileId].Diffs[diffId] = &utilities.DiffObject{diffId, subjectHash, object, diffPath, year}
}

func GetNextFileID() int {
	var highestID int
	for _, file := range db {
		if file.ID > highestID {
			highestID = file.ID
		}
	}
	return highestID + 1
}

func GetNextDiffID() int {
	var highestID int
	for _, file := range db {
		for _, diff := range file.Diffs {
			if diff.ID > highestID {
				highestID = diff.ID
			}
		}
	}
	return highestID + 1
}

func getModelArray() []*dbArrayStruct {

	m := make(map[int]*dbArrayStruct)

	for _, file := range db {
		for _, diff := range file.Diffs {
			m[diff.ID] = &dbArrayStruct{file, diff}
		}
	}

	o := make([]*dbArrayStruct, 0)

	for i := 0; i <= GetNextDiffID(); i++ {
		if s, ok := m[i]; ok {
			o = append(o, s)
		}
	}

	return o
}
