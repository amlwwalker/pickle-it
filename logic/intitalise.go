package logic

import (
	"path/filepath"

	utils "github.com/amlwwalker/pickleit/utilities"
	"github.com/atrox/homedir"
)

const (
	storageDirectory = "pickleit"
)

var (
	PATH, KEYFOLDER, DOWNLOADFOLDER, SYNCFOLDER, DIFFFOLDER, THUMBFOLDER, LOGFOLDER, PLUGINFOLDER string
)

func init() {
	homeDirectory, err := homedir.Dir()
	if err != nil {
		panic(err)
	}
	PATH = filepath.Join(homeDirectory, storageDirectory)

	//where private and public keys are kept
	KEYFOLDER = filepath.Join(PATH, "keys")
	//where downloaded files start
	DOWNLOADFOLDER = filepath.Join(PATH, "downloads")
	//where file originals live
	SYNCFOLDER = filepath.Join(PATH, "sync")
	//where patches and last versions live
	DIFFFOLDER = filepath.Join(PATH, "diff")
	//where the thumbnails are stored
	THUMBFOLDER = filepath.Join(PATH, "thumb")
	//where the logs are stored
	LOGFOLDER = filepath.Join(PATH, "logs")
	//where plugins are stored
	PLUGINFOLDER = filepath.Join(PATH, "plugins")

	utils.InitiateDirectory(KEYFOLDER)
	utils.InitiateDirectory(DOWNLOADFOLDER)
	utils.InitiateDirectory(SYNCFOLDER)
	utils.InitiateDirectory(DIFFFOLDER)
	utils.InitiateDirectory(THUMBFOLDER)
	utils.InitiateDirectory(LOGFOLDER)
	utils.InitiateDirectory(PLUGINFOLDER)
}
