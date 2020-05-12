// +build mock

package model

import (
	"github.com/amlwwalker/pickleit/utilities"
)

type dbArrayStruct struct {
	File *utilities.File
	Diff *utilities.DiffObject
}

var db = map[int]*utilities.File{
	0: &utilities.File{
		ID:   0,
		Name: "<all>",
	},

	1: &utilities.File{
		ID:          1,
		Name:        "moon-high-definition.png",
		Path:        "/somewhere/on/computer",
		BkpLocation: "/bkp/location/",
		// Diffs: map[int]*utilities.DiffObject{
		// 	1: &utilities.DiffObject{ID: 1, SubjectHash: "3J8H9K5Fhj0J", Object: "moon-high-definition.png", DiffPath: "/diff/path/1", Year: 1},
		// 	2: &utilities.DiffObject{ID: 2, SubjectHash: "a98sdhf9ash98", Object: "moon-high-definition.png", DiffPath: "/diff/path/2", Year: 11},
		// 	3: &utilities.DiffObject{ID: 3, SubjectHash: "23k4j2k3j4k2", Object: "moon-high-definition.png", DiffPath: "/diff/path/3", Year: 111},
		// },
	},

	2: &utilities.File{
		ID:          2,
		Name:        "ice-and-fire.png",
		Path:        "/somewhere/on/computer",
		BkpLocation: "/bkp/location/",
		// Diffs: map[int]*utilities.DiffObject{
		// 	1: &utilities.DiffObject{ID: 4, SubjectHash: "3J8H9K5Fhj0J", Object: "ice-and-fire.png", DiffPath: "/diff/path/2", Year: 2},
		// },
	},

	3: &utilities.File{
		ID:          3,
		Name:        "Paddington's day out.png",
		Path:        "/somewhere/on/computer",
		BkpLocation: "/bkp/location/",
		// Diffs: map[int]*utilities.DiffObject{
		// 	1: &utilities.DiffObject{ID: 5, SubjectHash: "3J8H9K5Fhj0J", Object: "Paddington's day out.png", DiffPath: "/diff/path/3", Year: 3},
		// },
	},
}
