package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/amlwwalker/pickleit/logic"
	"github.com/amlwwalker/pickleit/utilities"
)

var clientVersion = utilities.Version{
	Tag: "initial",
	//Flavour:     string
	Version: "0.1",
	//DBName:      string
	//Hash        string
	Date:        "May 13, 2020",
	PersistLogs: true,
	Production:  false,
	//Virtual     bool
}

type tmpStruct struct{}

func (b *tmpStruct) Write(p []byte) (n int, err error) {
	fmt.Fprintf(os.Stdout, "%s", string(p))
	return len(p), nil
}

func (b *tmpStruct) Close() error {
	fmt.Fprintf(os.Stdout, "\r\n")
	return nil
}

// Terminal client testing a filewatcher for the application
// Flat should specify the log level
func main() {
	informer := make(chan logic.OperatingMessage) //TODO figure out wtf this does
	logLevelPtr := flag.Int("log", 6, "Set the log level (1 - 6)")
	diffPtr := flag.Bool("diff", true, "set to true to diff and false to patch")
	pathPtr := flag.String("file", "", "Set the file (path) to watch for changes")
	hardCopyPtr := flag.Bool("fsCopy", true, "Make a hard copy of files that are being watched")

	directionPtr := flag.String("direction", "forward", "Set the direction to patch (default forward)")
	patchPtr := flag.String("patch", "", "Choose the patch file, and source (must agree with patch direction) and the path file")
	flag.Parse()
	//manager := logic.NewManager(*logLevelPtr, "[%{module}] [%{line}] [%{level}] %{message}", os.Stdout) //set the log level

	manager, err := logic.NewManager(clientVersion, *logLevelPtr, "[%{module}] [%{file} - %{line}] [%{level}] %{message}", informer, os.Stdout)
	if err != nil {
		fmt.Println("Error starting manager", err)
	}
	manager.DemoLogging()
	defer manager.TearDown()

	if *pathPtr == "" {
		manager.Error("You must set a path to a file to watch/patch")
		os.Exit(1)
	}
	buf := &tmpStruct{}
	progressBar, _ := utilities.NewProgressbarTime(100, buf)
	progressBar.DemoBarTime(100)

	if *diffPtr { //Is going to watch and diff a file
		filePath, err := manager.AddFileToMonitor(*pathPtr, *hardCopyPtr)
		manager.InfoF("watching file %s ", filePath)
		if err != nil {
			manager.ErrorF("%s\r\n", err)
		}
		manager.BeginWatching() //the watcher needs the waitgroup
	} else { //Is going to patch a file. If this is the case, a patch needs to have been specified
		/*
			for patching, lets use the same file from the pathPtr to know what we are patching.
			* We then need a flag for the direction we are patching
			* a location to save the patched file
			* a name for the patched file
		*/
		manager.Warning("In future, this must be changed so that a patch is chosen based on other properties and its selected automatically based on direction.\r\nPerhaps, for instance, thumbnails and time are the selectors, or even a 'note to self', and the last known 'good' location, but definately not the patch directly.")

		if *patchPtr == "" {
			manager.Error("You must provide a patch file")
			os.Exit(1)
		}

		if *directionPtr == "forward" {
			//once we know the file, we can ask the database for the backup (if its forward patch)
			//if its backward then we just need the patch file (TODO: look this up aswell for a certain patch)
			manager.NoticeF("Patching forward for %s", *pathPtr)
			//check that the forward patch is where we expect it to be
			if err := manager.BeginForwardPatch(*pathPtr, *patchPtr, ""); err != nil {
				manager.ErrorF("There was an error forward patching file %s: %s", *pathPtr, err)
			}
		}
		//else {
		//	manager.NoticeF("Patching backward for", *pathPtr)
		//	if err := manager.BeginBackwardPatch(*pathPtr, *patchPtr); err != nil {
		//		manager.ErrorF("There was an error backward patching file %s: %s", *pathPtr, err)
		//	}
		//}

	}

	// if the WaitGroup runs out of things to wait for, the application will quit
	manager.Wait()
}
