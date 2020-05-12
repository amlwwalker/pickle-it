package compressor

import (
	"fmt"
	"io/ioutil"
	"testing"
)

var testStruct = SomeStruct{
	A: "alex walker",
	B: 24,
	C: 1.234,
}

func TestMain(t *testing.T) {
	//use assert library to check for similarity or require library to check something exists
	t.Run("check that we can 'gob' the data correctly", func(t *testing.T) {
		if structBytes, err := StructToBytes(testStruct); err != nil {
			t.Error("failed to create the bytes from the struct provided ", err)
			t.FailNow()
		} else if binaryToGobBuffer, err := BytesToGob(structBytes.Bytes()); err != nil {
			t.Error("failed to create the gob from the bytes provided ", err)
			t.FailNow()
			//issue here with gob reading from an interface
		} else if compressedData, err := CompressBinary(&binaryToGobBuffer); err != nil {
			t.Error("failed to create the gob from the struct provided ", err)
			t.FailNow()
		} else if decompressionReader, err := DecompressBinary(compressedData); err != nil {
			t.Error("failed to decompress the binary data ", err)
			t.FailNow()
		} else if res, err := GobToBytes(decompressionReader); err != nil {
			t.Error("failed to convert bytes to struct ", err)
			t.FailNow()
		} else {
			fmt.Printf("result %+v\r\n", res)
		}
	})
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
