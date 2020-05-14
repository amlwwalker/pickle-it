package logic

import (
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"os/user"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"

	db "github.com/amlwwalker/pickleit/database"
	fsmanager "github.com/amlwwalker/pickleit/fsmanager"
	utilities "github.com/amlwwalker/pickleit/utilities"
	logger "github.com/apsdehal/go-logger"
	"github.com/denisbrodbeck/machineid"
	radovskyb "github.com/radovskyb/watcher"
)

// Manager keeps track of all elements that are passed around the application.
// It is responsible for the interactions that other parts of the application
// may need with utilities.
//
// Logging from inside here can be dangerous as the io.Writer may not be truly configured yet
// so be careful of this as you can get null pointer exceptions
func NewManager(version utilities.Version, logLevel int, format string, informer chan OperatingMessage, logWriter io.WriteCloser) (*Manager, error) {
	//note re colors: https://stackoverflow.com/questions/1961209/making-some-text-in-printf-appear-in-green-and-red#1961222
	log, err := logger.New("client logger", 1, logWriter)
	if err != nil {
		panic(err) // Check for error, no easy way to recover without a logger.
	}
	log.SetFormat(format)
	log.SetLogLevel(logger.LogLevel(logLevel))
	var wg sync.WaitGroup

	watcher, err := fsmanager.NewWatcher(log, KEYFOLDER, DOWNLOADFOLDER, SYNCFOLDER, THUMBFOLDER, DIFFFOLDER)
	if err != nil {
		log.CriticalF("Error creating a watcher %s", err)
		return &Manager{}, err
	}
	patcher, err := fsmanager.NewPatcher(log, KEYFOLDER, DOWNLOADFOLDER, SYNCFOLDER, THUMBFOLDER, DIFFFOLDER)
	if err != nil {
		log.CriticalF("Error creating a patcher %s", err)
		return &Manager{}, err
	}
	fmt.Printf("manager version %+v\r\n", version)
	usr, err := user.Current()
	if err != nil {
		log.CriticalF("Error retrieving the current user %s", err)
		return &Manager{}, err
	}
	database, err := db.NewDB(filepath.Join(PATH, version.DBName), format, logLevel, logWriter)
	if err != nil {
		log.CriticalF("There was an error initialising the data base [%s]\r\n", err)
		return &Manager{}, err
	}
	versionFormat := "%{bigVersion}.%{littleVersion}.%{microVersion}_%{timeUnix}_%{client}_%{job}_%{creator}_%{owner}_%{hash}_%{message}"
	machineID, err := machineid.ID()
	if err != nil {
		log.WarningF("there was an error identifying the machine %s\r\n", err)
	} else {
		log.Infof("Machine identified as %s\r\n", machineID)
	}
	settings := &UserSettings{Usr: *usr, versionFormat: versionFormat, darkMode: false, machineID: machineID}

	m := Manager{
		version,
		settings,
		log,
		&wg,
		watcher,
		patcher,
		database,
		informer,
		logWriter,
	}
	return &m, nil
}

// TearDown should be called by the main loop to close everything that has been left open
// TODO: check if the watcher/patcher/waitgroup or logger need tearing down.
// Defer its running in main so that it happens last
func (m *Manager) TearDown() {
	defer m.watcher.Close()
	defer m.dB.Close()
	m.Informer <- Op_WatchStopped.Retrieve()
}

// StopWatching closes the current channels on the watcher, kills the file watcher and refreshes the instance
// with a new watcher before responding that the watcher has been closed.
func (m *Manager) StopWatching() {
	m.watcher.Close()
	m.watcher.Watcher = nil             //confirm the old watcher is not referenced anymore
	m.watcher.Watcher = radovskyb.New() //due to closing the channels we need to create a new instance of the file wather
	m.Informer <- Op_WatchStopped.Retrieve()
}

// DemoLogging demos the capabilities of the logger under the current settings
// Nothing special here :)
func (m *Manager) DemoLogging() {
	m.Critical("1. Critical level")
	m.Error("2. Error level")
	m.Warning("3. Warning level")
	m.Notice("4. Notice level")
	m.Info("5. Info level")
	m.Debug("6. debugging level")
	VersionHelp()
}

// This adds a file for the watcher to keep an eye on
// however the file will also need to be backedup
// and added to the database.
// This changes all paths to absolute paths rather than relative
// when adding a file to monitor, this should check if the database
// is already expecting to monitor this file. If it is this function should
// do checks to make sure that it is successfully monitoring it, and that there
// is a historical breadcrumb trail to recreate all the versions that the database
// claims to have a copy of
func (m *Manager) AddFileToMonitor(file string, hardCopy bool) (string, error) {
	var err error
	// the filepath should be absolute, but this should be done dynamically
	if file, err = filepath.Abs(file); err != nil {
		return "", err
	}
	//TODO: what needs to happen is a channel for errors/progress is created
	//then pass that channel to a routine, and put all of the following in it
	// whenever an error returns, fire the string to the channel,
	// or send progress on the progress channel
	//however need to work out best way of returning the final result to the caller
	//- the way to do that is send the result on a third channel, for which is just the result
	//see commsManagment.go
	// f := NewFileManager()
	//DELAYED: this feature affects only large files and user experience. It can wait.

	var tmpFile utilities.File
	var filename string //we might aswell only verify the files validity once
	var hash [16]byte
	//check that the file actually exists
	if filename, err = utilities.VerifySrcFile(file); err != nil {
		//there was no source file or it was not recognisable as a file
		return "", err
	}
	//generate a unique file name from the hash and the moment it was created
	//a sampled (and therefore) fast, hash of the file for 'uniqueness'
	if hash, err = utilities.UniqueFileHash(file); err != nil {
		return "", err
	}

	if tmpFile, err = m.dB.CheckIfFileCurrentlyMonitored(file, hash); err != nil {
		if strings.Index(err.Error(), "not found") != -1 {
			//the file wasn't found, this is an ok error
			m.InfoF("The file was [%s], so continuing to create it in the database", err)
		} else {
			m.ErrorF("Error checking if file [%s] is monitored so will init file. Error: %s", tmpFile.Path, err)
		}
		tmpFile.CurrentHash = hash
		tmpFile.Name = filename
		tmpFile.Path = file
		tmpFile.CreatedAt = time.Now()
		tmpFile.Unique = base64.URLEncoding.EncodeToString([]byte(filename)) + "_" + base64.URLEncoding.EncodeToString((tmpFile.CurrentHash[:])) + "_" + strconv.FormatInt(tmpFile.CreatedAt.Unix(), 10) + "_" + filename
		tmpFile.BkpLocation = filepath.Join(SYNCFOLDER, tmpFile.Unique)
		tmpFile.CurrentBase = tmpFile.BkpLocation
		tmpFile.Ignore = false //we can have files in the database that are ignored. TODO: This was initially added so that 'All Files' would show up as a file (its a hack as it adds a dummy to the database)
		//we should now have a unique name for this file
		//if needs be, we can find out the real file name from the string
		//the hash will give us a reasononable indication of the similarity of the files
		//define filename of backup(s)
		if _, err := m.prepareDatabaseForFile(tmpFile); err != nil {
			return "", err
		} else {
			if err := m.copyFileToNewLocation(tmpFile.Path, tmpFile.BkpLocation, hardCopy); err != nil {
				m.ErrorF("There was an error copying the file to the backup location %s", err)
				return "", err
			}
			m.Informer <- Op_NewFile.Retrieve()
		}
	} else {
		m.DebugF("file [%s] is already in the database. Assuming sync file in place", tmpFile.Path)
		// we should check if the backup file exists, otherwise there is an issue
		if _, err := utilities.VerifySrcFile(tmpFile.BkpLocation); err != nil {
			//if the backup doesn't exist, something has gone quite wrong....
			m.DebugF("The backup file doesn't seem to exist at the expected location, %s", err)
			return "", err
		}
	}

	return tmpFile.Path, m.watcher.Add(tmpFile.Path)
}

// BeginWatching Ultimately this begins the routine keeping track of file changes
// however eventually this will be responsible for the continued
// running of the manager and will do more than just manage the watcher routine
// TODO: REFACTOR THIS WHOLE THING. Passing functions around is weird behaviour.
func (m *Manager) BeginWatching() error {
	m.Informer <- Op_WatchCommencing.Retrieve()
	//here we can create a channel that the diffs are passed on to
	//when a diff is pushed onto the channel, we can look for it and add it to the database
	diffChannel := make(chan utilities.DiffObject)
	//this needs a redesign
	//parent context to all patching routines
	ctx := context.Background()
	m.WaitGroup.Add(3) // add wait groups; 1 for this. one for the watcherRoutine

	//filter the watcher to Write and Create
	m.watcher.FilterOps(radovskyb.Write, radovskyb.Create)
	//stick this on a watcher otherwise the next routine will never be reached
	go m.watcher.BeginWatcherRoutine(ctx, m.WaitGroup, diffChannel, m.OnFileChange)

	// TODO: This is causing it to go CPU crazy
	// it is perhaps too many routines and continuous loops and no way out...
	// manages responses from the diffing routine
	// TODO: could this be done by just waiting on the channel in a for loop?
	// perhaps that would help with memory intensivity?
	// Besides, still better to refactor the switch/case into one place if we can.
	// Although if we can just do in for loops and block the channel until the right time...?
	// https://stackoverflow.com/questions/55367231/golang-for-select-loop-consumes-100-of-cpu
	go func(diffChannel chan utilities.DiffObject) {
		defer m.WaitGroup.Done()
		for {
			select {
			case diff := <-diffChannel:
				if diff.E != nil {
					m.WarningF("Creating the diff caused an error %s", diff.E)
				} else {
					// from here we know the errors etc, but mainly we just want to write the diff object to the database
					if err := m.dB.StoreDiff(diff); err != nil {
						m.ErrorF("Error storing the diff to the database %s", err)
					} else {
						m.NoticeF("Diff stored %s : %s", diff.StartTime, diff.Screenshot)
						m.Informer <- Op_NewDiff.Retrieve()
					}
					// at this point we can consider making a new commit (base file).
					// what this should do is check the size of the diff that just came back.
					// if its great than x% of the size of the original file, then lets create a new base
					// file.
					// 1. get size of diff
					diffSize := diff.DiffSize
					// 2. get size of base file
					fi, err := os.Stat(diff.Subject)
					if err != nil {
						m.ErrorF("error getting size of file %s", err)
						continue
					}
					m.NoticeF("time to create diff %s", diff.Message)
					// get the size
					size := fi.Size()
					deltaComparison := float64(diffSize) / float64(size)
					m.InfoF("size comparison size: %0.2f, diffSize: %0.2f, result: %0.2f", float64(size), float64(diffSize), deltaComparison)
					// 3. compare sizes
					if deltaComparison > float64(0.1) { //TODO: store the base file limit in an env var
						m.NoticeF("creating a new commit file as percent size is %.2f", deltaComparison)
						//the diff is greater than 60% of the base file, so lets create a new base file
						// 4. decide to make a new base file.
						// -------------
						// this involves patching the basefile with this patch, and creating a new base file
						// that other diffs will be based on from this point forward.

						_, fileName := filepath.Split(diff.Object)
						unique := base64.URLEncoding.EncodeToString([]byte(fileName)) + "_" + base64.URLEncoding.EncodeToString((diff.ObjectHash[:])) + "_" + strconv.FormatInt(time.Now().Unix(), 10) + "_" + fileName
						restoreFile := filepath.Join(SYNCFOLDER, unique)
						//generate a unique file name from the hash and the moment it was created
						//a sampled (and therefore) fast, hash of the file for 'uniqueness'
						if err := m.BeginForwardPatch(diff.Object, diff.DiffPath, restoreFile); err != nil {
							m.ErrorF("There was an error patching the new base file, error: %s", err)
						} else if hash, err := utilities.UniqueFileHash(restoreFile); err != nil {
							m.ErrorF("There was an error gettig the hash for the file, error: %s", err)
						} else {
							// 2. find out how the base file is selected - do you change the file in the database
							// so that it links to the new base file. This means that when a patch is applied it will also
							// need to update the location of the base file it refers to.
							if err := m.dB.UpdateFileData(diff.Object, restoreFile, hash); err != nil {
								m.ErrorF("There was an error resetting the file base %s", err)
							}
							m.Informer <- Op_NewBase.Retrieve()
						}
					}
				}
			}
		}
	}(diffChannel)

	// Start the watching process - it'll check for changes every 100ms.
	go func() {
		if err := m.watcher.Start(time.Millisecond * 100); err != nil {
			m.ErrorF("error starting watcher %s\r\n", err)
		}
	}()
	return nil
}

//TODO: refactor into a channel/routine
// OnFileChange is a callback function that will be passed to the watcher so that the watcher knows file to act on. This digs out the details of the file and returns them to the watcher.
// It returns the currently set base file as the subject that the diff will base itself off of
func (m *Manager) OnFileChange(fileChanged string) (utilities.File, error) {
	m.InfoF("on file changed %s", fileChanged)
	file, err := m.dB.FindFileByPath(fileChanged)
	return file, err
}

// OnProgressChanged manages progress of the diff, TODO: Implement. Requires binarydist, so is this an outdated function?
// func (m *Manager) OnProgressChanged(increment func(int) error, event *binarydist.Event) {
// 	// fmt.Fprintln(w, "callback for ", event.Name, " progress ", event.Progress)
// 	// increment(event.Progress)
// 	m.Notice("on progress changed called with ...")
// }

// prepareDatabaseForFile is responsible for keeping all references to the version of the file,
// the diff and the metadata of the diffs. Before any file is copied and stored, it should be managed by the database
//
// TODO: This will need to initialise a diff object in the database, currently created by the diff package,
// however going forward a diff maybe defined by the manager.
func (m *Manager) prepareDatabaseForFile(tmpFile utilities.File) (int, error) {
	if fileID, err := m.dB.InitialiseFileInDatabase(tmpFile); err != nil {
		m.ErrorF("Error checking if file [%s] is monitored. Error %s", tmpFile.Path, err)
		return 0, err
	} else {
		return fileID, nil
	}
}

// copyFileToNewLocation will check the size of the file that needs to be copied
// and decide whether it can manage it in memory. However it can be overridden if needs be
// based on whether it is required to copy it to a new location. In such a case
// the destination must be defined.
//
// TODO: Going forward this may use a totally custom naming convention to stop files
// appearing in search (with the actual filename in the filename (if you get me), it will appear in search)
//
// TODO: Going forward the file name of the backup should be a reference to the hash'd data
// incase two files being monitored have the same name.
// This will only be implemented when we have a database managing these details
func (m *Manager) copyFileToNewLocation(file, newLocation string, fsCopy bool) error {
	if fsCopy {
		//keep an original copy of the file available at all times
		if bytes, err := m.watcher.ForceFSCopy(file, newLocation); err != nil {
			return err
		} else {
			m.NoticeF("bytes %d ", bytes)
		}
	} else {
		if bytes, err := m.watcher.CleverCopy(file, newLocation); err != nil {
			return err
		} else {
			m.NoticeF("bytes %d ", bytes)
		}
	}
	return nil
}
