package fsmanager 

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/amlwwalker/pickleit/fsmanager/compressor"
	"github.com/amlwwalker/fdelta"
)
func main() {
	DefineFiles()
	originalBytes := GetOriginalBytes()
	delta := CreateDelta(originalBytes)
	StoreDelta(delta)
	retrievedDelta := RetrieveDelta()
	// var deltaBytes []byte
	fmt.Printf("res  : `%s`\n", len(retrievedDelta))
	//test loading the delta from disk
	appliedBytes, err := fdelta.Apply(originalBytes, retrievedDelta)
	if err != nil {
		panic(err)
	}
	fmt.Println("exporting delta")
	err = writeFile(appliedFile, appliedBytes)
	if err != nil {
		fmt.Println("error reading bytes [3]", err)
		os.Exit(1)
	}
	fmt.Printf("Origin : `%s`\n", originalFile)
	fmt.Printf("Target : `%s`\n", len(appliedBytes))
	fmt.Printf("Delta  : `%s`\n", len(delta))
	fmt.Printf("Result: `%s`\n", appliedFile)
}
