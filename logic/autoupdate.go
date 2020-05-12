package logic

// import (
// 	"bufio"
// 	"errors"
// 	"fmt"
// 	"os"

// 	"github.com/blang/semver"
// 	"github.com/rhysd/go-github-selfupdate/selfupdate"
// )

// //the repository is the github location of releases
// //version is the semantic version of the latest tag. See https://github.com/markchalloner/git-semver
// func retrieveUpdate(repository, version string) (*selfupdate.Release, error) {
// 	selfupdate.EnableLog()
// 	latest, found, err := selfupdate.DetectLatest(repository)
// 	if err != nil {
// 		return &selfupdate.Release{}, err
// 	}
// 	previous := semver.MustParse(version)
// 	fmt.Println("parsed version ", version, previous)
// 	if !found || latest.Version.LTE(previous) {
// 		return &selfUpdate.Release{}, errors.New("Current version is the latest")
// 	}
// 	return latest, nil
// }

// func checkWithUserWhetherToUpdate(version semver.Version) (error, bool) {
// 	fmt.Print("Do you want to update to", version, "? (y/n): ")
// 	input, err := bufio.NewReader(os.Stdin).ReadString('\n')
// 	if err != nil || (input != "y\n" && input != "n\n") {
// 		return errors.New("Invalid input"), false
// 	}
// 	if input == "n\n" {
// 		return nil, false
// 	}
// 	if input == "y\n" {
// 		return nil, true
// 	}
// 	return errors.New("no input"), false
// }

// func commenceUpdate(assetUrl string) error {
// 	exe, err := os.Executable()
// 	if err != nil {
// 		return errors.New("Could not locate executable path")
// 	}
// 	if err := selfupdate.UpdateTo(assetUrl, exe); err != nil {
// 		return err
// 	}
// 	//if its successful it should send a message to close the program and inform the user of the situation.
// 	return nil
// }