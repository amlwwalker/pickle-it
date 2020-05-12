package utilities

import "time"

type Version struct {
	Tag         string
	Flavour     string
	Version     string
	DBName      string
	Hash        string
	Date        string
	PersistLogs bool
	Production  bool
	Virtual     bool
}

type File struct {
	ID          int    `storm:"id,increment"`
	Path        string `storm:"index"`
	Name        string
	BkpLocation string
	CurrentBase string
	CurrentHash [16]byte `storm:"index,unique"`
	CreatedAt   time.Time
	Unique      string
	Version     float64
	Ignore      bool
}

type FileIndex struct {
	ID       int      `storm:"id,increment"`
	FileID   int      `storm:"index"`
	FileHash [16]byte `storm:"index,unique"`
	Index    []byte
	Length   int64
}

//screenshots need to be available in different parts of the application as they need to happen fast
type ScreenshotWrapper struct {
	ScreenshotError error
	Screenshot      string
}

// store the information for each diff that is made
type DiffObject struct {
	ID          int       `storm:"id,increment"`
	Subject     string    `storm:"index"`
	Object      string    `storm:"index"`
	SubjectHash [16]byte  `storm:"index"`
	ObjectHash  [16]byte  `storm:"index"`
	Watching    string    //name of the file being watched
	DiffPath    string    //path of the diff/patch
	Label       string    //store a comment if the user wants to (user written)
	Screenshot  string    //path to the screen shot when the diff was made
	Fs          bool      //whether it was written to the directly
	Direction   bool      //direction true == forward
	Description string    //record of forward or backward (just a quick helper)
	E           error     //a record of the error when it was created. Maybe able to optimize out later
	Diff        *[]byte   //the diff itself (incase we want to store in memory) - unused as of now
	DiffSize    int64     //the size of the diff in bytes
	StartTime   time.Time //when was the diff created (can take a while to create)
	Message     string    //any message we want to store against the diff while its created
}
type UXSettings struct {
	ID                       int `storm:"id"`
	AutoWatch                bool
	Statistics               bool
	Screenshot               bool
	NewFileReadySystemNotify bool
	PatchSystemNotify        bool
	OverrideExisting         bool
	Welcome                  bool
}
