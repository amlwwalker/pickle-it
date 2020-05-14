// Copyright 2015 The gocui Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/amlwwalker/pickleit/utilities"

	"github.com/amlwwalker/pickleit/logic"
)

var (
	buildTagging string
	buildVersion string
	dbName       string
	buildHash    string
	buildDate    string
	production   string
	version      utilities.Version
)

func init() {
	fmt.Println("build version = ", buildVersion)
	if dbName == "" {
		fmt.Println("need a database name")
		os.Exit(2)
	}
	version = utilities.Version{
		Tag:     buildTagging,
		Version: buildVersion,
		DBName:  dbName,
		Hash:    buildHash,
		Date:    buildDate,
	}
	if production == "FALSE" {
		version.Production = false
	} else {
		version.Production = true
	}
	fmt.Printf("version information %+v\r\n", version)
}
func main() {
	logLevelPtr := flag.Int("log", 6, "Set the log level (1 - 6)")
	fmt.Println("current version " + version.Version)
	flag.Parse()
	var cGui clientGui
	informer := make(chan logic.OperatingMessage)
	//got to be careful now as type asserting for an io.Writer to log to
	manager, err := logic.NewManager(version, *logLevelPtr, "[%{module}] [%{file} - %{line}] [%{level}] %{message}", informer, &cGui) //set the log level
	if err != nil {
		fmt.Println("Error starting new manager", err)
	}
	defer manager.TearDown()
	cGui.newUI(manager, *logLevelPtr)
}
