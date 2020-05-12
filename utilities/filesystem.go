package utilities

import (
	"bytes"
	"compress/gzip"
	"encoding/binary"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	logger "github.com/apsdehal/go-logger"
)

var log *logger.Logger

func init() {
	var err error
	log, err = logger.New("utilities logger", 1, os.Stdout)
	if err != nil {
		panic(err)
	}
	log.SetFormat("[%{module}] [%{level}] %{message}")
	log.Info("Utilities logger Created")
}

// CompressIntArray compresses an array of integers into a buffer
func CompressIntArray(arry []int64, compressionBuffer *bytes.Buffer) (float64, error) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, arry)
	if err != nil {
		return 0, err
	}
	//now compress it
	compressor := gzip.NewWriter(compressionBuffer)
	// if err != nil {
	// 	fmt.Println("writer level failed to set compression level")
	// }
	if _, err := compressor.Write(buf.Bytes()); err != nil {
		return 0, err
	}
	if err := compressor.Close(); err != nil {
		return 0, err
	}
	ratio := float64(len(compressionBuffer.Bytes())) / float64(len(buf.Bytes()))
	return ratio, nil
}

// ExpandToIntArray firstly unzips the byte array, then it
// converts the byte array back into an int array for use
func ExpandToIntArray(length int64, arry []byte, intArray *[]int64) error {
	buf := bytes.NewBuffer(arry)
	if reader, err := gzip.NewReader(buf); err != nil {
		fmt.Println("gzip failed ", err)
		return err
	} else {
		*intArray = make([]int64, length) //you must know the length of the original data if you are to do it this way.
		err := binary.Read(reader, binary.LittleEndian, intArray)
		if err != nil {
			fmt.Println("read failed ", err)
		}
		return nil
	}
}

// VerifySrcFile checks to see that the file is a regular file
// that the OS has meta information about and that can be read by
// the os.
func VerifySrcFile(src string) (string, error) {
	_, fileName := filepath.Split(src) //dirPath
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return fileName, errors.New("error on os.Stat " + err.Error())
	}

	if !sourceFileStat.Mode().IsRegular() {
		return fileName, errors.New("%s is not a regular file" + src)
	}
	return fileName, nil
}

func InitiateDirectory(directory string) {
	// For the keys-folder we need to check if the folder exists...
	checkDir, err := IsDirectory(directory)
	if err != nil {
		log.ErrorF("Error checking for "+directory+" directory: %s\r\n", err)
		panic(err)
	}

	if checkDir == true {
		log.Warning(directory + " already exists")
	} else {
		// Create the directory.
		log.Info("Creating " + directory)
		err = CreateDirectory(directory)
		if err != nil {
			log.ErrorF("Error creating the folder %s\r\n", err)
		}
	}
}

func IsDirectory(path string) (bool, error) {

	s, err := os.Stat(path) // returns an error if the path does not exist.
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		} else {
			return false, err // Different error...?
		}
	}

	if s.IsDir() {
		return true, nil
	}

	return false, nil // Redundancy

}

func CreateDirectory(path string) error {

	// Assumes checks have been done on if the directory exists...
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}

	return nil // Redundancy

}

func DeleteDirectory(path string) error {

	err := os.RemoveAll(path)
	return err

}

func StripFilePathBase(pathToFile, base string) string {
	return strings.Replace(pathToFile, base, "", -1)
}
