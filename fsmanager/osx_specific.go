//!build +darwin
package fsmanager

import (
	"fmt"

	"github.com/everdev/mack"
	"github.com/go-vgo/robotgo"
)

//OSX requires certain allowances to take screenshots
func checkOsSpecific() error {
	bitmap := robotgo.CaptureScreen(10, 20, 30, 40)
	// use `defer robotgo.FreeBitmap(bit)` to free the bitmap
	defer robotgo.FreeBitmap(bitmap)

	return nil
}
func checkOSMessaging() error {
	fmt.Println("OSX attempting to call system event")
	_, err := mack.Tell("System Events", "display dialog \"Pickle It is testing if it has all it needs. You may need all PickleIt to be able to make changes.\"")
	if err != nil {
		fmt.Println("couldn't display popup")
		return err
	} else {
		return nil
	}
}
