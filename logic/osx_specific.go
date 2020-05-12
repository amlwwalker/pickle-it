//!build +darwin
package logic

import (
	"fmt"

	"github.com/everdev/mack"
)

func checkOSSpecifics() error {
	fmt.Println("OSX attempting to call system event")
	_, err := mack.Tell("System Events", "display dialog \"Pickle It is testing if it has all it needs. You may need all PickleIt to be able to make changes.\"")
	if err != nil {
		fmt.Println("couldn't display popup")
		return err
	} else {
		return nil
	}
}
