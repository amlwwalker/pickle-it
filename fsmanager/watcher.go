package fsmanager

import (
	"context"
	"path/filepath"
	"sync"
	"time"

	"github.com/amlwwalker/pickleit/utilities"

	logger "github.com/apsdehal/go-logger"
	radovskyb "github.com/radovskyb/watcher"
)

type key string
type Event struct {
	Name     string
	Progress int
	Total    int
}

// The watcher is responsible for not only seeing when a file changes,
// but also keeping track of
// * the file hash so that if it changes again any modifications can be handled
// * copying any versions and keeping them safe (even if temporary)
// * creating the diff of the file, in both directions if necessary
// * storing the details in the database
func NewWatcher(logger *logger.Logger, KEYFOLDER, DOWNLOADFOLDER, SYNCFOLDER, THUMBFOLDER, DIFFFOLDER string) (Watcher, error) {
	w := Watcher{
		radovskyb.New(),
		logger,
		true, //used to temporarily ignore events if necessary
		KEYFOLDER, DOWNLOADFOLDER, SYNCFOLDER, THUMBFOLDER, DIFFFOLDER,
	}
	return w, nil
}

func (w *Watcher) Ignore() bool {
	w.Enabled = false
	return w.Enabled
}
func (w *Watcher) Enable() bool {
	w.Enabled = true
	return w.Enabled
}
func (w *Watcher) IsEnabled() bool {
	return w.Enabled
}

// BeginWatcherRoutine kicks off the watcher. When the watcher noticies a file change,
// certain actions will be taken in case of event and error
// the routine will handle these whenever this occurs.
// If certain functions need to be called then this will
// need to be specified as part of the managers lambda functions
// TODO: Should return an error
func (w *Watcher) BeginWatcherRoutine(ctx context.Context, wg *sync.WaitGroup, diffChannel chan utilities.DiffObject, onFileChanged func(string) (utilities.File, error)) {

	//seems a bit barking, but we can now cancel any diff that is occuring on a file when it fires again
	cancelFunctions := make(map[string]func())
	for {
		select {
		// we have filtered already on the [Op]erations we want to listen for so no need to check here
		case event := <-w.Event:
			if !w.IsEnabled() {
				w.Infof("ignoring event and reenabling the watcher %s\r\n", event)
				w.Enable()
				continue
			}
			w.Infof("event fired ", event)
			//this is currently slow as it does a db lookup on the path.
			//TODO: On load (or whenever a file is added to the watcher, the db information for files being watched, could be cached in memory. This would be much faster)
			fileInfo, err := onFileChanged(event.Path) //could return the 'Event' object here
			syncFilePath := fileInfo.CurrentBase
			uniqueName := fileInfo.Unique
			// //begin taking screenshot if we are supposed to
			screenshotChannel := make(chan utilities.ScreenshotWrapper)
			go func(ssChannel chan utilities.ScreenshotWrapper) {
				w.Infof("beginning taking screenshot at ", time.Now())
				var ssStruct utilities.ScreenshotWrapper
				if screenshotFileName, err := takeScreenShot(w.THUMBFOLDER, uniqueName); err != nil {
					w.WarningF("could not take screenshot", err)
					ssStruct.ScreenshotError = err
				} else {
					ssStruct.Screenshot = filepath.Join(w.THUMBFOLDER, screenshotFileName)
					w.Infof("screenshot recorded ", ssStruct.Screenshot, " at ", time.Now())
				}
				ssChannel <- ssStruct
			}(screenshotChannel)

			// fileID := fileInfo.ID
			//we need the hash of the current base, not the hash of the original file
			// fileHash := fileInfo.CurrentHash //hash needs to come from
			if err != nil {
				w.ErrorF("path was not returned to sync path", err)
				continue
			}
			//cancel the event if it indeed is running...
			if cancelFunctions[event.Path] != nil {
				cancelFunctions[event.Path]()
				delete(cancelFunctions, event.Path)
			}

			//context for the current event. Calling cancel will cancel the routines
			//to kill a context you must have access to the cancel function.
			// i could add the cancel function to a map of them
			// if you want to kill it you call on the correct cancel function
			// that will kill the context. Fine.
			// If however you want to kill it from another place....
			// then you need a cancel channel, which inturn calls the equivelent cancel function.... ok
			// sounds about best i can do right now...
			// kind of bonkers right....
			cancelContext, cancel := context.WithCancel(ctx)
			cancelFunctions[event.Path] = cancel
			// good idea to not use strings as keys directly as can conflict across namespaces
			// this needs to be sorted out -- too many things called an event....
			// TODO: its totally bananas
			e := Event{
				Name:     event.Path,
				Progress: 0,
				Total:    100,
			}
			eventContext := context.WithValue(cancelContext, key(event.Path), e)
			if err := manageFileDiffing(eventContext, event.Path, syncFilePath, w.DIFFFOLDER, true, screenshotChannel, diffChannel, wg); err != nil {
				// I don't think this can be reached...
				w.WarningF("Error managing the diffing process %s", err)
			}
		case err := <-w.Watcher.Error:
			w.ErrorF("%s\r\n", err)
		case <-w.Closed:
			w.Notice("radovskyb closed")
			return
		}
	}
}
