package fsmanager

import (
	logger "github.com/apsdehal/go-logger"
	radovskyb "github.com/radovskyb/watcher"
)

type Watcher struct {
	*radovskyb.Watcher
	*logger.Logger
	Enabled        bool
	KEYFOLDER      string
	DOWNLOADFOLDER string
	SYNCFOLDER     string
	THUMBFOLDER    string
	DIFFFOLDER     string
}
type Patcher struct {
	*logger.Logger
	KEYFOLDER      string
	DOWNLOADFOLDER string
	SYNCFOLDER     string
	THUMBFOLDER    string
	DIFFFOLDER     string
}

// This is not currently in use
// however it may become what is stored in the database, with alot more information
// about the files and the diffs. Losing data is absolutely not OK, and may result in us keeping a db
// backup on the servers aswell to be assured that a user cannot lose any data. That could be what our backup offering
// is
type Diff struct {
	DiffHash       []byte
	DiffSize       int64
	FileHash       []byte
	FileSize       int64
	ParentFileHash []byte //hash of parent file
	ParentDiff     []byte //hash of parent diff
	ChildFileHash  []byte //hash of child file
	ChildDiffHash  []byte //hash of child diff
	// these could be version/stored name etc
	Id       []byte            //unique id of this hash
	ChildId  []byte            //unique id of the child
	ParentId []byte            //unique id of the parent
	MetaData map[string]string //any meta data about this diff or the file it creates
}
