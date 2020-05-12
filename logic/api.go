package logic

import (
	"os"
	"sort"

	"github.com/amlwwalker/pickleit/utilities"
)

//contains functions that other packages may want to access but we don't want direct access to the underlying code

func (m *Manager) StoreUXSettings(settings utilities.UXSettings) error {
	m.Settings.systemSettings = settings
	err := m.dB.StoreUXSettings(settings)
	return err
}
func (m *Manager) RetrieveUXSettings() (utilities.UXSettings, error) {
	if settings, err := m.dB.RetrieveUXSettings(); err != nil {
		return settings, err
	} else {
		m.Settings.systemSettings = settings
		return settings, nil
	}
}

// RetrieveWatchedFilePaths returns a list of the paths of all files that the database deems to be "being watched"
func (m *Manager) RetrieveWatchedFilePaths() ([]string, error) {
	var filePaths []string
	if files, err := m.dB.RetrieveWatchedFiles(); err != nil {
		return []string{}, err
	} else {
		for i := range files {
			filePaths = append(filePaths, files[i].Path)
		}
	}
	return filePaths, nil
}

// RetrievePatchesForFile will return any patches in the database for a specific file that has or is being watched
func (m *Manager) RetrievePatchesForFile(filePath string, ascending bool) ([]utilities.DiffObject, error) {
	var diffs []utilities.DiffObject
	var err error
	m.NoticeF("file path looking for %s", filePath)
	if diffs, err = m.dB.RetrieveDiffsForFileByPath(filePath); err != nil {
		m.ErrorF("error retrieving diffs for file %s : %s", filePath, err)
		return []utilities.DiffObject{}, err
	}
	//lets sort the diffs into date order (by rights backward/forward patches should still be next to each other)
	sort.Slice(diffs, func(i, j int) bool {
		if ascending { //sort into ascending order
			return diffs[i].StartTime.Before(diffs[j].StartTime)
		}
		//or descending....
		return diffs[i].StartTime.After(diffs[j].StartTime)
	})
	return diffs, nil
}

// RetrievePatchByID recovers the diff object from the database.
// This is usually to be called when wanting a specific patch, the then apply
func (m *Manager) RetrievePatchByID(ID int) (utilities.DiffObject, error) {
	if diff, err := m.dB.RetrieveDiffsByID(ID); err != nil {
		return utilities.DiffObject{}, err
	} else {
		return diff, nil
	}
}

// RetrieveWatchedFiles is a helpher function to retrieve a list of all the files
// currently being watched
func (m *Manager) RetrieveWatchedFiles() map[string]os.FileInfo {
	return m.watcher.WatchedFiles()
}

func (m *Manager) ChangePatchLabel(ID int, label string) error {
	err := m.dB.ChangeLabel(ID, label)
	return err
}
