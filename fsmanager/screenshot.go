package fsmanager

import (
	"errors"
	"path/filepath"
	"strconv"
	"time"

	"fmt"
	"strings"

	"github.com/everdev/mack"
	"github.com/go-vgo/robotgo"
	// "github.com/therecipe/qt/internal/examples/3rdparty/uglobalhotkey/UGlobalHotkey"
)

type window struct {
	name string
	w    int
	h    int
	x    int
	y    int
}

// takeScreenShot will return a screenshot for every available display.
//
// TODO: Going forward, we want to store the pngs against the diff
func takeScreenShot(storageLocation, uniqueName string) (string, error) {

	t := time.Now()
	formattedTime := t.Format("20060102150405")

	activeMeta := robotgo.GetActive()
	fmt.Printf("%+v\r\n", activeMeta)
	pidInt := int(robotgo.GetPID())
	pidString := strconv.Itoa(pidInt)

	/* #applescript
	tell application "System Events"
		set _P to a reference to (processes whose unix id is 79061)
		set _W to a reference to the first window of _P

		[_P's name, _W's size, _W's position]
	end tell
	*/

	response, err := mack.Tell("System Events", "set _P to a reference to (processes whose unix id is "+pidString+")", "set _W to a reference to the first window of _P", "[_P's name, _W's size, _W's position]")
	if err != nil {
		fmt.Println("couldnt get dimensions and position")
		return "", err
	} else {
		fmt.Println(response)
		if window, err := parseFocusedWindow(response); err != nil {
			fmt.Println("failed to capture the focused window. No screenshot captured ", err)
			return "", err
		} else {
			fileName := uniqueName + "_" + formattedTime + ".png"
			robotgo.SaveCapture(filepath.Join(storageLocation, fileName), window.x, window.y, window.w, window.h)
			return fileName, nil //we have to assume the capture was a success
		}
	}

}

func parseFocusedWindow(response string) (window, error) {

	splt := strings.Split(response, ", ")
	var w window
	if len(splt) != 5 {
		return w, errors.New("incorrect number of fields available")
	}
	w.name = splt[0]

	var err error
	if w.w, err = strconv.Atoi(splt[1]); err != nil {
		return w, err
	}
	if w.h, err = strconv.Atoi(splt[2]); err != nil {
		return w, err
	}
	if w.x, err = strconv.Atoi(splt[3]); err != nil {
		return w, err
	}
	if w.y, err = strconv.Atoi(splt[4]); err != nil {
		return w, err
	}
	fmt.Printf("%+v\r\n", w)
	return w, nil
}
