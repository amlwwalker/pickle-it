package fsmanager

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/amlwwalker/pickleit/fsmanager/compressor"
	"github.com/amlwwalker/fdelta"
)

// var originalFile string
// var newFile string
// var patchFile string
// var appliedFile string

func openReader(file io.Reader) ([]byte, error) {
	buffer, err := ioutil.ReadAll(file)
	return buffer, err
}
func openFile(path string) ([]byte, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("File reading error", err)
	}
	return data, err
}

func writeFile(path string, data []byte) error {
	err := ioutil.WriteFile(path, data, 0644)
	return err
}

func getOriginalBytes(originalFile string) ([]byte, error) {
	originalBytes, err := openFile(originalFile)
	if err != nil {
		return []byte{}, err
	}
	return originalBytes, nil
}
func createDelta(newFile string, originalBytes []byte) ([]byte, error) {
	newBytes, err := openFile(newFile)
	if err != nil {
		return []byte{}, err
	}
	delta := fdelta.Create(originalBytes, newBytes)
	fmt.Println("size of delta ", len(delta))
	return delta, nil
}
func compressDelta(delta []byte) ([]byte, error) {
	if binaryToGobBuffer, err := compressor.BytesToGob(delta); err != nil {
		return []byte{}, err
	} else if compressedData, err := compressor.CompressBinary(&binaryToGobBuffer); err != nil {
		return []byte{}, err
	} else {
		return compressedData.Bytes(), nil
	}
}
func storeDelta(patchFile string, delta []byte) ([]byte, error) {
	if compressedData, err := compressDelta(delta); err != nil {
		return []byte{}, err
	} else if err := writeFile(patchFile, compressedData); err != nil {
		return []byte{}, err
	} else {
		return compressedData, nil
	}
}
func retrieveDelta(patchFile string) ([]byte, error) {
	compressedData, err := openFile(patchFile)
	if err != nil {
		return []byte{}, err
	}
	return compressedData, nil
}

func decompressDelta(compressedData []byte) ([]byte, error) {
	var compressedBuffer bytes.Buffer
	compressedBuffer.Write(compressedData)
	if decompressionReader, err := compressor.DecompressBinary(compressedBuffer); err != nil {
		return []byte{}, err
	} else if res, err := compressor.GobToBytes(decompressionReader); err != nil {
		return []byte{}, err
	} else {
		return res, nil
	}
}

func applyPatchToFile(originalbytes, delta []byte) ([]byte, error) {
	if patchedBytes, err := fdelta.Apply(originalbytes, delta); err != nil {
		return []byte{}, err
	} else {
		return patchedBytes, nil
	}
}
