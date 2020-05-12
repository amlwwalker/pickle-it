package compressor

import (
	"bytes"
	"compress/gzip"
	"encoding/gob"
	"fmt"
	"io"
)

type SomeStruct struct {
	A string
	B int64
	C float64
}

//1.
func StructToBytes(obj SomeStruct) (bytes.Buffer, error) {
	//now gob this
	var indexBuffer bytes.Buffer
	encoder := gob.NewEncoder(&indexBuffer)
	if err := encoder.Encode(obj); err != nil {
		return indexBuffer, err
	}
	return indexBuffer, nil
}

//1.
func BytesToGob(obj []byte) (bytes.Buffer, error) {
	//now gob this
	var indexBuffer bytes.Buffer
	encoder := gob.NewEncoder(&indexBuffer)
	if err := encoder.Encode(obj); err != nil {
		return indexBuffer, err
	}
	return indexBuffer, nil
}

//2.
func CompressBinary(binaryBuffer *bytes.Buffer) (bytes.Buffer, error) {
	//now compress it
	var compressionBuffer bytes.Buffer
	compressor := gzip.NewWriter(&compressionBuffer)
	_, err := compressor.Write(binaryBuffer.Bytes())
	err = compressor.Close()
	return compressionBuffer, err
}

//3.
func DecompressBinary(compressionBuffer bytes.Buffer) (*gzip.Reader, error) {
	//now decompress it
	dataReader := bytes.NewReader(compressionBuffer.Bytes())
	if reader, err := gzip.NewReader(dataReader); err != nil {
		fmt.Println("gzip failed ", err)
		return &gzip.Reader{}, err
	} else {
		err := reader.Close()
		return reader, err
	}
}

//4.
func GobToBytes(binaryBytes io.Reader) ([]byte, error) {
	decoder := gob.NewDecoder(binaryBytes)
	var tmp []byte
	if err := decoder.Decode(&tmp); err != nil {
		return tmp, err
	}
	return tmp, nil
}
