package fsmanager

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"time"

	"github.com/amlwwalker/pickleit/utilities"
)

// ManageFileDiffing handles creating the diffs on the background routines and creating the information
// about each diff that is made.
//
// TODO: fs works however it takes a while to write the diffs to disk. It maybe a better idea to keep the diffs
// in memory (although they could get huge??) and then write them to disk at a later point in time.
// In any event, this works now.
//
// TODO: Be able to cancel a diff creation (for instance if the user resaves). Does this work? Should we block
// creating diffs within 5 minutes of creating one? Cancelling is probably better at this point.
// it might be nice to inform the user when diffs build up
func manageFileDiffing(ctx context.Context, subject, object, diffStorageLocation string, fs bool, screenshotChannel chan utilities.ScreenshotWrapper, diffChannel chan utilities.DiffObject, wg *sync.WaitGroup) error {

	var subjectHash, objectHash [16]byte
	var err error
	if subjectHash, err = utilities.UniqueFileHash(subject); err != nil {
		return err
	}
	if objectHash, err = utilities.UniqueFileHash(object); err != nil {
		return err
	}

	diffTime := time.Now()
	wg.Add(1)
	go func(messages chan<- utilities.DiffObject, ssChannel chan utilities.ScreenshotWrapper) {
		defer wg.Done()

		var dO utilities.DiffObject
		//doing this on routine to not lose anytime... does it change anything?
		dO.Description = ""
		dO.Direction = true
		dO.Subject = object
		dO.Object = subject
		dO.StartTime = diffTime
		dO.SubjectHash = objectHash //TODO: these being the wrong way round is a legacy thing. Swapping them needs testing, but should be fine
		dO.ObjectHash = subjectHash
		fmt.Println("creating diff object now")
		if diff, err := binaryDiff(ctx, &dO, diffStorageLocation, fs); err != nil { //binaryDiff
			fmt.Println("error from binary diff ", err)
			dO.E = err
		} else {
			dO.Diff = &diff
		}
		ssStruct := <-ssChannel
		fmt.Printf("received over ssChannel %+v\r\n", ssStruct)
		if ssStruct.ScreenshotError != nil {
			fmt.Println("screenshot failed, ", ssStruct.ScreenshotError)
		} else {
			fmt.Println("diff reeived screenshot ", ssStruct.Screenshot)
			dO.Screenshot = ssStruct.Screenshot
		}
		elapsed := time.Since(diffTime)
		dO.Message = "elapsed time:" + elapsed.String()
		messages <- dO
	}(diffChannel, screenshotChannel)
	return nil
}

//run instead of binaryDiff to turn it off
func dryrun(ctx context.Context, dO *utilities.DiffObject, diffStorageLocation string, fs bool) ([]byte, error) {
	return []byte{}, nil
}

// Diff manages the creation of the diffs but doesn't actually create the diffs itself.
// Sources are file system sources in this case and an array of diffs (io.Writers) are returned
// 1. This handles whether to save the diffs directly to the drive, and if so, will save to the
// specified location. If so, it will return the diffs.
// 2. Whether to save diffs in both directions
// 3. Creates a diff object that contains any necessary metadata about the diff files
// subject is the file that changed, object is file on record
func binaryDiff(ctx context.Context, dO *utilities.DiffObject, diffStorageLocation string, fs bool) ([]byte, error) {

	var fileName string
	if !dO.Direction { //if true, its forward, so use the object, else use the subject (i.e its backwards now)
		_, fileName = filepath.Split(dO.Object) // dirPath
	} else {
		_, fileName = filepath.Split(dO.Subject) // dirPath
	}
	dO.Watching = fileName
	// var sub io.Reader
	// if sub, err = os.Open(dO.Subject); err != nil {
	// 	return []byte{}, err
	// }
	// var obj io.Reader
	// if obj, err = os.Open(dO.Object); err != nil {
	// 	return []byte{}, err
	// }
	startTime := strconv.FormatInt(dO.StartTime.Unix(), 10)
	if fs { //if the wish is to store to the filesystem
		dO.DiffPath = filepath.Join(diffStorageLocation, fileName+"_"+startTime+"_"+dO.Description) + "_diff.patch"
		if writeDiff, err := os.Create(dO.DiffPath); err != nil {
			return []byte{}, err
		} else if deltaBytes, err := fdeltaDiff(ctx, dO.Subject, dO.Object); err != nil {
			return []byte{}, err
		} else {
			if bytesWritten, err := writeDiff.Write(deltaBytes); err != nil {
				return []byte{}, err
			} else {
				dO.DiffSize = int64(bytesWritten)
				return []byte{}, nil
			}
		}
	} else { //if we actually want the bytes we have to set fs to false (can do this above.)
		if deltaBytes, err := fdeltaDiff(ctx, dO.Subject, dO.Object); err != nil {
			return []byte{}, err
		} else {
			dO.DiffSize = int64(len(deltaBytes))
			return deltaBytes, nil
		}
	}
}

//sub is the original
func fdeltaDiff(ctx context.Context, sub, obj string) ([]byte, error) {
	//now follow what is found in fdelta to retrieve the bytes and get back a delta
	//you can use the gob/compression code to save the files according to where in pickle it they are saved
	//TODO: currently the code is used to compress the bsdiff index, but we dont need that, just need to store the
	// delta on disk. This is currently already done somewhere, so can possibly add/swap out the delta and compressor code
	// so that it uses the new code.
	if originalBytes, err := getOriginalBytes(sub); err != nil {
		return []byte{}, err
	} else if deltaBytes, err := createDelta(obj, originalBytes); err != nil {
		return []byte{}, err
	} else if compressedDelta, err := compressDelta(deltaBytes); err != nil {
		return []byte{}, err
	} else {
		return compressedDelta, nil
	}
}

//// uniDirectionalDiff takes two file names, however returns the diffs as bytes
//// rather than writing the diffs out to a file
//// Technically binary diffs are optimized for one direction, so to create bi-directional
//// diffs, the trick is to call the differ in both directions.
//// This function juse manages creating the diffs and has no interest in the direction that it is occuring in.
//// An advanced version of this will be take io.Readers in so that it doesn't care what the data source is that it is diffing
//func uniDirectionalDiff(ctx context.Context, sub, obj io.Reader) ([]byte, error) {
//	var b bytes.Buffer
//
//	subObjDiff := bufio.NewWriter(&b)
//	if _, err := binarydist.Diff(ctx, sub, obj, subObjDiff); err != nil {
//		return nil, err
//	}
//	return b.Bytes(), nil
//}
//
//// uniDirectionalFSDiff creates a diff based on the input files
//// however saves it out to the filename provided
//// Technically binary diffs are optimized for one direction, so to create bi-directional
//// diffs, the trick is to call the differ in both directions.
//// This function juse manages creating the diffs and has no interest in the direction that it is occuring in.
//// An advanced version of this will be take io.Readers in so that it doesn't care what the data source is that it is diffing
//func uniDirectionalFSDiff(ctx context.Context, sub, obj io.Reader, filePath string) (int64, error) {
//	var subObjDiff io.Writer
//	var err error
//	if subObjDiff, err = os.Create(filePath); err != nil {
//		return 0, err
//	}
//	if bytesWritten, err := binarydist.Diff(ctx, sub, obj, subObjDiff); err != nil {
//		return 0, err
//	} else {
//		return bytesWritten, nil
//	}
//}
//
//// uniDirectionalFSDiffOptimized creates a diff based on the input files.
//// Major difference to the other Diff helpers is this one retrieves the index from a database.
//func uniDirectionalFSDiffOptimized(ctx context.Context, sub, obj io.Reader, filePath string, retrieveOptimalIndex func(index *[]int64) error) (int64, error) {
//	var subObjDiff io.Writer
//	var err error
//	if subObjDiff, err = os.Create(filePath); err != nil {
//		return 0, err
//	}
//	I := &[]int64{}
//	if err := retrieveOptimalIndex(I); err != nil {
//		return 0, err
//	}
//	if bytesWritten, err := binarydist.OptimizedDiff(ctx, I, sub, obj, subObjDiff); err != nil {
//		return 0, err
//	} else {
//		return bytesWritten, nil
//	}
//}
