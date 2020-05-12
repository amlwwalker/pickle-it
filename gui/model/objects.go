// +build !mock

package model

import (
	"github.com/amlwwalker/pickleit/logic"
	"github.com/amlwwalker/pickleit/utilities"
	"github.com/therecipe/qt/core"
)

func getModelArray(manager *logic.Manager) []*dbArrayStruct {
	o := make([]*dbArrayStruct, 0)
	//retrieve all the diffs
	diffs, _ := manager.RetrieveAllDiffs()
	//retrieve all the files for the diff (this is inefficient hopefully obviously)

	//this is a faker so that the drop down puts the first real entry in second place
	// holder := &dbArrayStruct{&utilities.File{
	// 	ID:   0,
	// 	Name: "<all>",
	// }, &utilities.DiffObject{}}
	// o = append(o, holder)
	for i := range diffs {
		fileForDiff, _ := manager.FindFileByPath(diffs[i].Object)
		tmp := &dbArrayStruct{&fileForDiff, &diffs[i]}
		o = append(o, tmp)
	}
	// fmt.Printf("o %+v\r\n", o)
	//for each diff, create an dbArrayStruct object
	//put the file and the diff in and add it to the array
	return o
}

type dbArrayStruct struct {
	File *utilities.File
	Diff *utilities.DiffObject
}

//DiffDetail is used by models that want to push a diff to display on the front end
type DiffDetail struct {
	core.QObject

	_ int             `property:"id"`
	_ string          `property:"name"`
	_ string          `property:"description"`
	_ string          `property:"startTime"`
	_ string          `property:"screenshot"`
	_ *core.QDateTime `property:"startDate"`
	_ *core.QDateTime `property:"endDate"`
}
