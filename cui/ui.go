// emoji search https://www.emojihexa.com/search?
package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/amlwwalker/pickleit/utilities"

	"github.com/amlwwalker/pickleit/logic"
	logger "github.com/apsdehal/go-logger"
	"github.com/jroimartin/gocui"
	"github.com/ttacon/emoji"
)

var mu sync.Mutex // protects ctr

type clientGui struct {
	Manager *logic.Manager
	*logger.Logger
	logLevel                     int
	gui                          *gocui.Gui
	views                        map[string]*gocui.View
	menuItemControls             map[string]func(g *gocui.Gui) error
	watchedFilePaths             map[string]string
	patchFileIds                 map[string]int
	currentPatch                 *utilities.DiffObject
	currentlySelectedWatchedFile string
}

var (
	logCounter = 0
)

const (
	DELIMITER = " -> "
)

// Write - the interface so that clientGui implements an io.Writer interface for global logging
func (cGui *clientGui) Write(p []byte) (n int, err error) {
	logCounter = logCounter + 1
	// cGui.gui.Update(func(g *gocui.Gui) error {
	fmt.Fprintf(cGui.views["logging"], strconv.Itoa(logCounter)+" - "+string(p))
	// return nil
	// })
	return len(p), nil
}

// Adhere to the WriteCloser interface
func (cGui *clientGui) Close() error {
	fmt.Fprintf(cGui.views["logging"], "\r\n")
	return nil
}

//extractTextFromInput extracts a users input from a text box
// (clears the input box when compeleted)
func (c *clientGui) extractTextFromInput(g *gocui.Gui, v *gocui.View) error {
	//this extracts text from an input and closes the input (deletes it)
	vb := strings.TrimSuffix(v.ViewBuffer(), "\n")
	if filePath, err := c.Manager.AddFileToMonitor(vb, true); err != nil {
		c.ErrorF("%s", err)
	} else {
		c.InfoF("watching file %s", filePath)
	}
	delView(g, v)
	return nil
}

// menuItemController returns the controller from the map, depending on what
// option has been selected from the list
func (c *clientGui) menuItemController(g *gocui.Gui, v *gocui.View) error {
	var l string
	var err error

	_, cy := v.Cursor()
	if l, err = v.Line(cy); err != nil {
		return err
	}
	// Can come from the *clientGui if we want, however the interface requires that this is passed
	// g *gocui.Gui, v *gocui.View so why not use it :) ( - its the same object anyway)
	c.menuItemControls[l](g)

	return nil
}

// getPatches returns all the patches for the selected file
func (c *clientGui) getPatches(g *gocui.Gui, v *gocui.View) error {
	var l string
	var err error
	_, cy := v.Cursor()
	if l, err = v.Line(cy); err != nil {
		return err
	}
	c.InfoF("line %s", l)
	filePath := c.watchedFilePaths[strings.Split(l, DELIMITER)[1]]
	clearView(g, "patches")

	// Can come from the *clientGui if we want, however the interface requires that this is passed
	// g *gocui.Gui, v *gocui.View so why not use it :) ( - its the same object anyway)
	if patches, err := c.Manager.RetrievePatchesForFile(filePath, true); err != nil {
		c.WarningF("error retrieving patches for %s, %s", filePath, err)
	} else {
		var patchStrings []string
		for i := range patches {
			tmp := patches[i]
			diffSizeMB := float64(tmp.DiffSize) / 1024.0 / 1024.0
			formattedTime := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
				tmp.StartTime.Year(), tmp.StartTime.Month(), tmp.StartTime.Day(),
				tmp.StartTime.Hour(), tmp.StartTime.Minute(), tmp.StartTime.Second())
			direction := tmp.Direction
			var hash [16]byte
			var arrow string
			//the direction tells us whether the file is the subject or the object
			if direction { //then the file is the object (i.e forward diff), and we can get the hash of the object
				hash = tmp.ObjectHash
				arrow = emoji.Emojitize(":fast_forward:")
			} else {
				hash = tmp.SubjectHash
				arrow = emoji.Emojitize(":fast_backward:")
			}
			var patchDescription string
			if patches[i].Label != "" { //use label or hash
				patchDescription = arrow + DELIMITER + formattedTime + " - " + fmt.Sprintf("%.2f", diffSizeMB) + " - " + patches[i].Label
			} else {
				patchDescription = arrow + DELIMITER + formattedTime + " - " + fmt.Sprintf("%.2f", diffSizeMB) + " - " + fmt.Sprint(hash[:16])
			}
			c.patchFileIds[patchDescription] = patches[i].ID
			patchStrings = append(patchStrings, patchDescription)
		}
		updateViewElements(g, "patches", patchStrings)
	}
	return nil
}

func (cGui *clientGui) displayPatch(g *gocui.Gui, v *gocui.View) error {
	maxX, maxY := g.Size()

	if _, patch, err := cGui.retrievePatch(g, v); err != nil {
		cGui.WarningF("error retrieving patches for %s", err)
		return err
	} else {
		if v, err := createLayout(g, "currentPatch", "hello", 30, 5, maxX-30, maxY-5, true); err != nil {
			fmt.Println("error creating watching view ", err)
		} else {
			cGui.currentPatch = &patch
			cGui.views["currentPatch"] = v
			patchArray := []string{
				"Patch Information:",
				"",
				"Description: " + patch.Description,
				"DiffPath: " + patch.DiffPath,
				"DiffSize: " + fmt.Sprint(patch.DiffSize),
				"Direction: " + fmt.Sprint(patch.Direction),
				"E: " + fmt.Sprint(patch.E),
				"Fs: " + fmt.Sprint(patch.Fs),
				"ID: " + fmt.Sprint(patch.ID),
				"Label: " + patch.Label,
				"Message: " + patch.Message,
				"ObjectHash: " + fmt.Sprint(patch.ObjectHash[:16]),
				"SubjectHash: " + fmt.Sprint(patch.SubjectHash[:16]),
				"StartTime: " + fmt.Sprint(patch.StartTime),
				"Screenshot: " + patch.Screenshot,
				"Watching: " + patch.Watching,
				"",
				"",
				"",
				"HELP - Ctrl P to patch with this.",
				"HELP - Ctrl Q to cancel patching",
			}
			g.SetCurrentView("currentPatch")
			updateViewElements(g, "currentPatch", patchArray)
		}
	}
	return nil
}

// patchFile will recover the patch from the database by ID.
// once a patch has been selected, the patch will be applied to either the backup
// or the current file, depending on the direction of the patch, and the
// file will be restored to its correct location and named accordingly.
// Over writes of the current file will be optional however will cause branching in the backup
// system, and if allowed, these restored files will be of concern (as they won't be known to us)
func (cGui *clientGui) patchFile(g *gocui.Gui, v *gocui.View) error {
	patch := *cGui.currentPatch
	cGui.InfoF("Found patch with object hash %s", fmt.Sprint(patch.ObjectHash[:16]))

	//get rid of this view
	delView(g, v)
	g.SetCurrentView("menu")
	cGui.NoticeF("Patching forward for %s -> %s. Patch -> %s", patch.Object, patch.Subject, patch.DiffPath)
	//check that the forward patch is where we expect it to be

	if err := cGui.Manager.BeginForwardPatch(patch.Object, patch.DiffPath, ""); err != nil {
		cGui.ErrorF("There was an error forward patching file %s: %s", patch.DiffPath, err)
		return err
	}
	//TODO: We can't do backward patching, this will need to be removed I think
	// if patch.Direction { //forward patch
	// 	cGui.NoticeF("Patching forward for %s -> %s. Patch -> %s", patch.Object, patch.Subject, patch.DiffPath)
	// 	//check that the forward patch is where we expect it to be

	// 	if err := cGui.Manager.BeginForwardPatch(patch.Object, patch.DiffPath, ""); err != nil {
	// 		cGui.ErrorF("There was an error forward patching file %s: %s", patch.DiffPath, err)
	// 		return err
	// 	}
	// } else {
	// 	cGui.NoticeF("Patching backward for %s -> %s", patch.Subject, patch.Object)
	// 	if err := cGui.Manager.BeginBackwardPatch(patch.Subject, patch.DiffPath); err != nil {
	// 		cGui.ErrorF("There was an error backward patching file %s: %s", patch.DiffPath, err)
	// 		return err
	// 	}
	// }

	return nil
}

// labelPatch is how a user can add a label to a patch.
// TODO: complete...
func (c *clientGui) labelPatch(g *gocui.Gui, v *gocui.View) error {
	c.retrievePatch(g, v)
	var l string
	var err error

	_, cy := v.Cursor()
	if l, err = v.Line(cy); err != nil {
		return err
	}
	c.NoticeF("l %s", l)
	// c.InfoF("to label patch with object hash %s", fmt.Sprint(patch.ObjectHash[:16]))
	patchID := c.patchFileIds[l]
	c.NoticeF("patchID %d", patchID)
	//the new label = replacing the hash with the label
	label := "LABEL !!!"
	//save the patch back
	if err := c.Manager.ChangePatchLabel(patchID, label); err != nil {
		//labelling failed
		c.ErrorF("changing the patch label failed %s", err)
	}
	// TODO: Reloading the patches requires knowing the current file selected.
	// the following causes it to runtime error
	//  else {
	// 	if err := c.getPatches(g, v); err != nil {
	// 		c.ErrorF("reloading the patch label failed %s", err)
	// 	}
	// }
	return nil
}

// retrievePatch - helper function to retrieve a patch from database api
func (c *clientGui) retrievePatch(g *gocui.Gui, v *gocui.View) (string, utilities.DiffObject, error) {
	var l string
	var err error

	_, cy := v.Cursor()
	if l, err = v.Line(cy); err != nil {
		return l, utilities.DiffObject{}, err
	}
	patchID := c.patchFileIds[l]
	// we need to look the patch up in the map so that we can get the patch data
	// once we have the patch data we can start applying the patch to the orig file
	// and restore it to the restoration path (usually the path of the orig)
	patch, err := c.Manager.RetrievePatchByID(patchID)
	return l, patch, err
}

//configureMenuItemControls sets up the menus.
func (c *clientGui) configureMenuItemControls() error {
	//add a new file
	c.menuItemControls["Add New File"] = func(g *gocui.Gui) error {
		c.Info("initialising adding new file to watcher...")
		//this should show a pop up where you can put the path of a file
		//that should in turn return the input and that should call the manager to add that file for watching
		getInput("Enter the full path of a file to watch", g)
		return nil
	}
	c.menuItemControls["Start Watching"] = func(g *gocui.Gui) error {
		// TODO: find out if watcher is currently watching.
		// TODO: get watcher to list out currently watched files
		c.Info("starting watching files")
		// this is possibly not necessary to run on another routine...
		// this has got the ui locked up now that we have
		// one wg running the show.
		// you can't quit while the watcher is running...
		// issues.
		if err := c.Manager.BeginWatching(); err != nil {
			c.ErrorF("failed to start radovskyb watcher %s", err)
		}
		return nil
	}

	c.menuItemControls["Reload Watchers"] = func(g *gocui.Gui) error {
		c.Info("updating watcher elements...")

		if filePaths, err := c.Manager.RetrieveWatchedFilePaths(); err != nil {
			c.ErrorF("error retrieving the file paths from the database %s", err)
			return err
		} else {
			//have to tidy up the strings...
			re := regexp.MustCompile("^.{0,20}") //trim off first 10 characters
			//need to save a map of the truncated file paths against the actual file paths

			var tmpFilePaths []string
			//takes time to process, so don't do every time
			thumbsUp := emoji.Emojitize(":thumbsup:")
			for i := range filePaths {
				//first lets add the files to monitor
				if _, err := c.Manager.AddFileToMonitor(filePaths[i], true); err != nil {
					c.ErrorF("%s", err)
				} else {
					c.NoticeF("watching file %s", filePaths[i])
					//if all was good lets add it to the list of files being watched
					if len(filePaths[i]) >= 20 {
						tmpPath := "..." + re.ReplaceAllString(filePaths[i], "$2")
						tmpFilePaths = append(tmpFilePaths, thumbsUp+DELIMITER+tmpPath)
						c.watchedFilePaths[tmpPath] = filePaths[i]
					}
				}
			}
			//at this point we need the manager to go and get the watched files from the database
			updateViewElements(g, "watching", tmpFilePaths)
		}
		return nil
	}
	c.menuItemControls["Reload Patchers"] = func(g *gocui.Gui) error {
		c.Info("clearing patcher elements")
		updateViewElements(g, "patches", []string{})
		return nil
	}
	return nil
}

//initLayout is required by the gocui api
func (cGui *clientGui) initLayout(g *gocui.Gui) error {
	//define some sizes
	loggingAreaHeight := 20
	titlebarWidth := 1
	columnWidth := 45
	maxX, maxY := g.Size()
	maxY = maxY - loggingAreaHeight
	maxX--

	//we don't add the title bar to the list of available views
	if _, err := createLayout(g, "titlebar", "Work Horse Client Gui", 0, 0, maxX, titlebarWidth, false); err != nil {
		fmt.Println("error creating watching view ", err)
	}

	//create the menu. This must be exist in all UIs
	if v, err := g.SetView("menu", 0, titlebarWidth, 30, maxY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Menu"
		v.Highlight = true
		v.SelBgColor = gocui.ColorGreen
		v.SelFgColor = gocui.ColorBlack
		fmt.Fprintln(v, "Add New File")
		fmt.Fprintln(v, "Start Watching")
		fmt.Fprintln(v, "Reload Watchers")
		fmt.Fprintln(v, "Reload Patchers")
		cGui.views["menu"] = v
		if _, err := g.SetCurrentView("menu"); err != nil {
			return err
		}
	}

	if v, err := createLayout(g, "watching", "Watching Files", 30, titlebarWidth, 25*columnWidth/10, maxY, true); err != nil {
		fmt.Println("error creating watching view ", err)
	} else {
		cGui.views["watching"] = v
	}
	if v, err := createLayout(g, "patches", "File Patches", 25*columnWidth/10, titlebarWidth, maxX, maxY, true); err != nil {
		fmt.Println("error creating watching view ", err)
	} else {
		cGui.views["patches"] = v
	}
	if v, err := createLayout(g, "logging", "Logging Area", 0, maxY, maxX, maxY+loggingAreaHeight-1, false); err != nil {
		fmt.Println("error creating logging view ", err)
	} else {
		cGui.views["logging"] = v
	}

	return nil
}

//this must be the last thing to run as it must own the 'main loop'
func (c *clientGui) newUI(manager *logic.Manager, logLevel int) {
	c.menuItemControls = make(map[string]func(*gocui.Gui) error)
	c.views = make(map[string]*gocui.View)
	c.watchedFilePaths = make(map[string]string)
	c.patchFileIds = make(map[string]int)
	// now configure the client logger...
	// the view for logging needs to be instantiated before logger can write to it...
	// therefore you cannot use the logger inside this function
	log, err := logger.New("cui logger", 1, c)
	if err != nil {
		panic(err) // Check for error
	}
	if err != nil {
		panic(err) // Check for error
	}
	log.SetFormat("[%{module}] [%{file} - %{line}] [%{level}] %{message}")
	log.SetLogLevel(logger.LogLevel(logLevel))
	//so that we can log from elsewhere.
	c.Logger = log      //logger immediately available to the cui
	c.Manager = manager //the manager, in charge of all file management is available to the gui

	//configure the user interface
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		panic(err) // Check for error
	}
	defer g.Close()

	g.Cursor = true

	g.SetManagerFunc(c.initLayout)

	c.gui = g
	//setup the key bindings to the views
	if err := c.setKeyBindings(); err != nil {
		panic(err) // Check for error
	}
	c.configureMenuItemControls()
	go func() {
		time.Sleep(2 * time.Second)
		c.Manager.InfoF("PickleIt is ready...")
		c.Manager.NoticeF("Version %s", c.Manager.Version.Version)
	}()
	//this must be the loop that runs the application
	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		panic(err) // Check for error
	}
}

// setKeyBindings initialises the bindings for each menu
func (c *clientGui) setKeyBindings() error {
	//menu controls (general app controls that don't fit else where)
	if err := c.gui.SetKeybinding("menu", gocui.KeyCtrlSpace, gocui.ModNone, nextView); err != nil {
		return err
	}
	if err := c.gui.SetKeybinding("menu", gocui.KeyArrowDown, gocui.ModNone, cursorDown); err != nil {
		return err
	}
	if err := c.gui.SetKeybinding("menu", gocui.KeyArrowUp, gocui.ModNone, cursorUp); err != nil {
		return err
	}
	if err := c.gui.SetKeybinding("menu", gocui.KeyEnter, gocui.ModNone, c.menuItemController); err != nil {
		return err
	}
	//watching panel (lists all files being watched)
	if err := c.gui.SetKeybinding("watching", gocui.KeyCtrlSpace, gocui.ModNone, nextView); err != nil {
		return err
	}
	if err := c.gui.SetKeybinding("watching", gocui.KeyArrowDown, gocui.ModNone, cursorDown); err != nil {
		return err
	}
	if err := c.gui.SetKeybinding("watching", gocui.KeyArrowUp, gocui.ModNone, cursorUp); err != nil {
		return err
	}
	if err := c.gui.SetKeybinding("watching", gocui.KeyEnter, gocui.ModNone, c.getPatches); err != nil {
		return err
	}
	//patches panel (lists all the patches for the currently selected file)
	if err := c.gui.SetKeybinding("patches", gocui.KeyCtrlSpace, gocui.ModNone, nextView); err != nil {
		return err
	}
	if err := c.gui.SetKeybinding("patches", gocui.KeyArrowDown, gocui.ModNone, cursorDown); err != nil {
		return err
	}
	if err := c.gui.SetKeybinding("patches", gocui.KeyArrowUp, gocui.ModNone, cursorUp); err != nil {
		return err
	}
	if err := c.gui.SetKeybinding("patches", gocui.KeyEnter, gocui.ModNone, c.displayPatch); err != nil {
		return err
	}
	if err := c.gui.SetKeybinding("patches", gocui.KeyCtrlL, gocui.ModNone, c.labelPatch); err != nil {
		return err
	}

	if err := c.gui.SetKeybinding("main", gocui.KeyCtrlSpace, gocui.ModNone, nextView); err != nil {
		return err
	}

	if err := c.gui.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		return err
	}
	if err := c.gui.SetKeybinding("currentPatch", gocui.KeyArrowDown, gocui.ModNone, cursorDown); err != nil {
		return err
	}
	if err := c.gui.SetKeybinding("currentPatch", gocui.KeyArrowUp, gocui.ModNone, cursorUp); err != nil {
		return err
	}
	if err := c.gui.SetKeybinding("currentPatch", gocui.KeyCtrlP, gocui.ModNone, c.patchFile); err != nil {
		return err
	}
	if err := c.gui.SetKeybinding("currentPatch", gocui.KeyCtrlQ, gocui.ModNone, delView); err != nil {
		return err
	}

	//this is TBC
	if err := c.gui.SetKeybinding("input", gocui.KeyEnter, gocui.ModNone, c.extractTextFromInput); err != nil {
		return err
	}
	if err := c.gui.SetKeybinding("input", gocui.KeySpace, gocui.ModNone, delView); err != nil {
		return err
	}

	return nil
}
