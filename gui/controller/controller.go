package controller

import (
	"flag"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/amlwwalker/pickleit/gui/model"
	"github.com/amlwwalker/pickleit/logic"
	"github.com/amlwwalker/pickleit/utilities"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/qml"
	"github.com/therecipe/qt/xml"
)

var once sync.Once
var instance *Controller

type Controller struct {
	core.QObject
	SystemSettings     *UXSettings
	VersionInformation *UXVersion
	information        chan logic.OperatingMessage
	Tray               *Tray
	manager            *logic.Manager
	PluginManager      *logic.PluginManager
	fileData           *xml.QDomDocument
	engine             *qml.QQmlApplicationEngine
	persistence        bool
	persistentLog      *os.File
	//qml bridging...

	_ func() `constructor:"init"`

	_ *core.QAbstractItemModel `property:"fileModel"`
	_ *core.QAbstractItemModel `property:"diffModel"`
	_ *core.QAbstractListModel `property:"eventModel"`
	_ *core.QAbstractListModel `property:"detailModel"`

	//<-view
	_ func() `signal:"pickleItVersion"`

	//<-file
	_ func(row int) `signal:"changeFile"`

	//<-diff
	_ func(index int)                            `signal:"deleteDiff"`
	_ func(index int)                            `signal:"patchFile"`
	_ func(index *core.QModelIndex)              `signal:"showDiffDetails"`
	_ func(column int, order core.Qt__SortOrder) `signal:"sortTableView"`

	//->detail
	// _ func()                                                           `signal:"showImageLabel"`
	_ func()                           `signal:"clearDetails"` //clear the details panel when a file change selection occurs
	_ func(profileLabelText string)    `signal:"showFileProfile"`
	_ func(events []*model.DiffDetail) `signal:"displayEventDetails"`
	_ func()                           `signal:"updateSelectedDate"`
	// _ func(id, name, description, imagePath string, selectedDate *core.QDateTime) `signal:"displayDiffDetails"`
	// _ func(*core.QDate)                                           `signal:"displayQDateDetails"`
	_ func(index int, description string) `signal:"requestSaveDescription"` //hacked while correct mapping not configured. TODO: configure data mapping in eventModel
	_ func(imagePath string)              `signal:"requestExpandImage"`
	_ func(imagePath string)              `signal:"expandImage"`
	_ func()                              `signal:"showExplanation"`

	//diff<->dialog
	// _ func()                        `signal:"deleteDiffCommand"`
	// _ func()                        `signal:"deleteDiffRequest"`
	_ func(index int) `signal:"patchFileShowRequest"` //these need model agnostic indexes
	_ func(index int) `signal:"deleteDiffShowRequest"`

	//diff<->plugin
	_ func(code string) `signal:"pushNewStackCode"`

	//calendar
	_ func(*core.QDate) []*model.DiffDetail `slot:"eventsForDate"`
	_ *core.QDate                           `property:"selectedDate"`
	_ func()                                `slot:"eventUpdate"`
	_ func()                                `slot:"updateCalendarEvents"`

	//detail
	_ func(diffs []*model.DiffDetail, reset bool) `signal:"detailUpdate"`
	//dialog<->view
	// _ func(index *core.QModelIndex) `signal:"patchFileShowRequest"`
	_ func(filepath string) `signal:"addFile"`
	_ func()                `signal:"diffAdded"`

	// _ func()                                                     `signal:"addDiffShowRequest"`
	_ func(file, subjectHash, object, diffPath string, year int) `signal:"addDiff"`

	_ func() `signal:"beginWatchingRequest"`
	_ func() `signal:"reloadWatchingRequest"`
	_ func() `signal:"stopWatchingRequest"`

	//<-plugin
	_ func(code string) `signal:"changeStackView"`

	//->notifications
	_ func(notification string)       `signal:"pushNotification"`
	_ func(value float64, reset bool) `signal:"setDeterminateProgress"`

	//->settings
	_ func(id, name, description string) `signal:"updateSettings"`
}

// Write - the interface so that controller implements an io.Writer interface for logging
func (c *Controller) Write(p []byte) (n int, err error) {
	fmt.Fprintf(os.Stdout, "[gui-controller] "+string(p))
	if c.persistence {
		c.persistentLog.WriteString(string(p) + "\r\n")
	}
	return len(p), nil
}

// Adhere to the WriteCloser interface
func (cGui *Controller) Close() error {
	return nil
}

// Instance will always return the same controller. The controller manages the interactions with the front end
func Instance() *Controller {
	once.Do(func() {
		instance = NewController(nil)
	})
	return instance
}

func (c *Controller) init() {

	c.fileData = xml.NewQDomDocument()

	c.ConnectChangeFile(c.changeFile)

	//<-plugin
	c.ConnectChangeStackView(c.changeStackView)

	c.ConnectRequestExpandImage(func(imagePath string) {
		fmt.Println("expanding image ", imagePath)
		c.ExpandImage(imagePath)
	})
	//<-detail
	c.ConnectRequestSaveDescription(func(id int, description string) { //*core.QModelIndex
		var ok bool
		// id := index.Data(int(core.Qt__UserRole) + 1).ToInt(&ok)
		if !ok {
			fmt.Println("failed to convert the index to an ID")
		}
		fmt.Println("saving description for patch with id ", id, " and description ", description)
		//so now we have the index of the element, we can use that to convert back to the db index and save against that index

		if err := c.manager.UpdateDiffDescription(id, description); err != nil {
			c.PushNotification("Warning: " + err.Error())
		} else {
			c.PushNotification("Description updated")
		}
	})
	//<-diff
	c.ConnectDeleteDiff(c.deleteDiff)
	c.ConnectPatchFile(c.patchFile)
	c.ConnectShowDiffDetails(c.showDiffDetails)

	c.ConnectAddFile(c.addNewFile)
	c.ConnectBeginWatchingRequest(c.beginWatchingRequest)
	c.ConnectReloadWatchingRequest(c.reloadWatchingRequest)
	c.ConnectStopWatchingRequest(c.stopWatchingRequest)
	c.Tray = NewTray()

}

func (c *Controller) InitWith(version utilities.Version, engine *qml.QQmlApplicationEngine) error {

	// configure
	//make version information available to the UI
	c.VersionInformation = NewUXVersion(nil) //we need to instantiate it as VersionInformation is a pointer so that the binding can manage the Qt connection
	c.VersionInformation.SetTag(version.Tag)
	c.VersionInformation.SetFlavour(version.Flavour)
	c.VersionInformation.SetVersion(version.Version)
	c.VersionInformation.SetHash(version.Hash)
	c.VersionInformation.SetDate(version.Date)
	c.persistence = version.PersistLogs

	if c.persistence { //later can have a flag for persistent loggin
		if f, err := os.Create("/tmp/pickleit.log"); err != nil {
			fmt.Println("Persistence cannot be maintained. Panicking", err)
			panic(err)
		} else {
			c.persistentLog = f
		}
	}
	fmt.Printf("version information %s - %s - %s - %s%s\r\n", c.VersionInformation.Tag(), c.VersionInformation.Version(), c.VersionInformation.Hash(), c.VersionInformation.Date())
	logLevelPtr := flag.Int("log", 6, "Set the log level (1 - 6)")
	fmt.Println("current version " + version.Version)
	flag.Parse()
	c.information = make(chan logic.OperatingMessage)
	//got to be careful now as type asserting for an io.Writer to log to
	{

		var err error
		c.manager, err = logic.NewManager(version, *logLevelPtr, "[%{module}] [%{file} - %{line}] [%{level}] %{message}", c.information, c) //set the log level
		if err != nil {
			c.manager.CriticalF("Error occured during setup %s\r\n", err)
			return err
		}
	}
	c.engine = engine
	if pluginManager, err := logic.NewPluginManager(c.information, c.engine); err != nil {
		return err
	} else {
		c.PluginManager = pluginManager
	}
	//UX property settings

	lModel := model.NewListModel(nil)
	lModel.InitWith(c.manager)
	c.SetFileModel(model.ListModel)

	sModel := model.NewSortFilterModel(nil)
	sModel.InitWith(c.manager)
	c.SetDiffModel(model.SortFilterModel)
	//this is not confirmed approach as of now
	eModel := model.NewEventModel(nil) //configure the event model
	eModel.InitWith(c.manager)
	c.SetEventModel(model.EventModel)

	//<-calendar
	c.ConnectEventsForDate(model.EventModel.EventsForDate)
	c.ConnectEventUpdate(model.EventModel.Update)
	//calendar
	c.ConnectSelectedDate(model.EventModel.SelectedDate)
	c.ConnectSelectedDateChanged(func(date *core.QDate) {
		// fmt.Println("controller date changed to ", date.Year(), date.Month(), date.Day())
		c.SetSelectedDate(date) //do we need to set both of these
		model.EventModel.SetSelectedDate(date)
		// fmt.Println("eController selected date set to", c.SelectedDate().Year(), c.SelectedDate().Month(), c.SelectedDate().Day())
		// fmt.Println("eModel selected date set to", model.EventModel.SelectedDate().Year(), model.EventModel.SelectedDate().Month(), model.EventModel.SelectedDate().Day())
		// model.EventModel.BeginResetModel()
		// model.EventModel.EndResetModel()
		c.EventUpdate()          //this updates the model
		c.UpdateCalendarEvents() //this and the above should happen in conjunction with eacy other
		events := c.EventsForDate(date)
		c.DisplayEventDetails(events)
		c.UpdateSelectedDate()
	})

	//detail view
	//this is not confirmed approach as of now
	dModel := model.NewDetailModel(nil) //configure the event model
	dModel.InitWith(c.manager)
	c.SetDetailModel(model.DetailModel)
	c.ConnectDetailUpdate(model.DetailModel.Update)
	c.ConnectClearDetails(model.DetailModel.Clear)
	c.ConnectDisplayEventDetails(func(events []*model.DiffDetail) {
		c.ClearDetails() //clear the details panel when a file change selection occurs
		fmt.Printf("[walker] updating with events %+v \r\n", events)
		c.DetailUpdate(events, true)
	})
	//need the model instantiated before can connect the Sort functionality
	c.ConnectSortTableView(model.SortFilterModel.Sort)

	//assuming everything is ready to go lets do a demo log out
	go c.processInformationChannel()
	c.manager.DemoLogging() //this should push to the informer aswell
	// initialise UX settings
	c.initUXSettings()
	c.onReady()
	return nil
}

func (c *Controller) onReady() {
	go func() {
		//let everything initialise
		time.Sleep(2 * time.Second)
		//check whether to show welcome screen
		c.welcomeCheck()
		c.autoWatchCheck()
	}()
}

// ShowMessage is useful if you want to send messages to the status pop up on the destop
func (c *Controller) ShowMessage(title, message string) {
	c.Tray.ShowMessage(title, message)
}
func (c *Controller) TearDown() {
	c.manager.TearDown()
}

//TODO: is this redundant now? we aren't using a stackview anymore I dont think.....
func (c *Controller) changeStackView(code string) {
	fmt.Println("controller received code ", code)
	c.PushNewStackCode(code)
}
func (c *Controller) changeFile(row int) {
	file := model.ListModel.Index(row, 0, core.NewQModelIndex()).Data(int(core.Qt__DisplayRole)).ToString()
	if row > 0 {
		model.SortFilterModel.SetFilterFixedString(file)
		model.SortFilterModel.SetFilterKeyColumn(2)
		model.EventModel.SetFilterExpression(file)
		// c.showFileProfile(file)
	} else if row == 0 {
		model.SortFilterModel.SetFilterFixedString("") //set the patch view to show all
		model.EventModel.SetFilterExpression("")       //set the calendar to respond with everything
	}
	// model.EventModel.Update() //the calendar needs to update with new diffs, (so does the detail view aswell)
	c.EventUpdate()          //this updates the model
	c.UpdateCalendarEvents() //TODO: this and the above should happen in conjunction with eacy other
	//we can get the selected date, and retrieve the events again with the filter on, then update the details, or we can just clear for now.
	c.ClearDetails() //clear the details panel when a file change selection occurs
}

// func (c *Controller) showFileProfile(file string) {
// 	c.ShowFileProfile(fmt.Sprintf("File : %v \nNumber of Diffs: %v", file, c.manager.GetDiffCountForFile(file)))
// }

//when clicking on a diff, this takes the index and finds the corresponding diff to show details of

//this needs to get the date of the event

//WALKER this function should create a diffDetail object to push into the detail model.
func (c *Controller) showDiffDetails(index *core.QModelIndex) {

	id := index.Data(int(core.Qt__UserRole) + 1).ToInt(nil)
	name := index.Data(int(core.Qt__UserRole) + 3).ToString()             //Object
	startDateTime := index.Data(int(core.Qt__UserRole) + 15).ToDateTime() //this is the startTime according to the database
	description := index.Data(int(core.Qt__UserRole) + 12).ToString()
	screenshot := index.Data(int(core.Qt__UserRole) + 9).ToString()
	fmt.Println("[controller.go/showDiffDetails] screenshot path is " + screenshot)
	c.ClearDetails() //clear the details panel when a file change selection occurs
	fmt.Println("[controller.go/showDiffDetails] calling [DisplayDiffDetails] - no longer active.")
	// c.DisplayDiffDetails(id, name, description, screenshot, startDateTime)

	//now create a QDateTime object to append
	ev := model.NewDiffDetail(nil)
	// ev.SetName(fmt.Sprintf("event (%v) on the %v/%v/%v", diff.Watching, startDate.Day(), startDate.Month(), startDate.Year()))
	ev.SetName(name)
	ev.SetId(id)
	ev.SetScreenshot(screenshot)
	ev.SetDescription(description)
	// ev.SetStartTime(diff.StartTime.String())
	st := core.NewQDateTime()
	st.SetDate(core.NewQDate3(startDateTime.Date().Year(), int(startDateTime.Date().Month()), startDateTime.Date().Day()))
	st.SetTime(core.NewQTime3(startDateTime.Time().Hour(), startDateTime.Time().Minute(), startDateTime.Time().Second(), 0))
	ev.SetStartDate(st)

	et := core.NewQDateTime()
	et.SetDate(core.NewQDate3(startDateTime.Date().Year(), int(startDateTime.Date().Month()), startDateTime.Date().Day()))
	et.SetTime(core.NewQTime3(startDateTime.Time().Hour(), startDateTime.Time().Minute(), startDateTime.Time().Second(), 0))
	ev.SetEndDate(et)
	c.DetailUpdate([]*model.DiffDetail{ev}, true)
}

func (c *Controller) deleteDiff(index int) {
	// var ok bool
	// if parsedIndex := index.Data(int(core.Qt__UserRole) + 1).ToInt(&ok); !ok {
	// 	fmt.Println("cant delete diff, couldn't convert to integer")
	// 	return
	// } else {
	model.SortFilterModel.BeginResetModel()
	c.manager.DeleteDiff(index)
	model.SortFilterModel.EndResetModel()
	// }

	c.PushNotification("Patch deleted")
}

// func (c *Controller) removeDiffFromDatabase(index *core.QModelIndex) {

// 	//TODO
// 	//inserting or removing from this model (SortFilterModel) will NOT affect the sourceModel
// 	//(because calls are not going through?) and it will therefore lead to glitches
// 	//resetModel however affectes both models and works ... but it can be slow

// }

func (c *Controller) patchFile(index int) {
	//first the manager needs to retrieve the patch from the database
	//once we've done that, we can use this to patch the base file and restore it.
	// before we know what the patch is, we need to convert the Qt index into the patch information.
	// fmt.Printf("[patchFile] index = %+v ; %+v\r\n", *index, index)
	// fmt.Printf("attempting to convert %+v\r\n", index.Data(int(core.Qt__UserRole)+1).ToString())
	//something like this...
	// var ok bool
	// if patchID := index.Data(int(core.Qt__UserRole) + 1).ToInt(&ok); !ok {
	// 	fmt.Println("could not convert to int ", index.Data(int(core.Qt__UserRole)+1).ToString())
	// } else {
	fmt.Println("patchID ", index)
	if patch, err := c.manager.RetrievePatchByID(index); err != nil {
		fmt.Printf("There was an error retrieving patch file %s:\r\n", err)
	} else {
		//specifying a restore path means that you can put the patched file where you like rather than generating it next to the original
		if err := c.manager.BeginForwardPatch(patch.Object, patch.DiffPath, ""); err != nil {
			fmt.Printf("There was an error forward patching file %s: %s\r\n", patch.DiffPath, err)
			// return err
		} else {
			c.PushNotification("Patched!")
		}
	}
	// }
	//TODO: be good to mark this patch as the currently patched version so its obvious this is the current one worked on
	// return nil
}

func (c *Controller) addDiff(fileName string, subjectHash, object, diffPath string, year int) {

	//checks whether the file for the diff already exists
	var fileNameId int
	if a := c.manager.GetFileForName(fileName); a != nil {
		fileNameId = a.ID
	} else {
		fileNameId = c.manager.GetNextFileID()
		c.addNewFile(fileName)
	}
	//not sure how this should be formed yet. will wait
	// diffId :=
	c.addNewDiff(fileNameId, subjectHash, object, diffPath, year)
	// c.addMetaData(diffId, diffPath)
}

func (c *Controller) addNewFile(filepath string) {
	// fileId := c.manager.GetNextFileID()
	fmt.Println("received file path ", filepath)
	filepath = utilities.StripFilePathBase(filepath, "file://")
	fmt.Println("new file path ", filepath)
	model.ListModel.BeginInsertRows(core.NewQModelIndex(), model.ListModel.RowCount(core.NewQModelIndex())+1, model.ListModel.RowCount(core.NewQModelIndex())+1)
	// c.manager.CreateNewFile(fileId, name)
	fmt.Println("adding " + filepath + " to monitor")
	if path, err := c.manager.AddFileToMonitor(filepath, true); err != nil {
		fmt.Println("adding file "+path+" to watch errored ", err)
	}
	model.ListModel.EndInsertRows()
	// return fileId
}

func (c *Controller) addNewDiff(fileId int, subjectHash, object, diffPath string, year int) int {

	diffId := c.manager.GetNextDiffID()

	//TODO
	//inserting or removing from this model (SortFilterModel) will NOT affect the sourceModel
	//(because calls are not going through?) and it will therefore lead to glitches
	//resetModel however affectes both models and works ... but it can be slow

	//model.SortFilterModel.BeginInsertRows(core.NewQModelIndex(), d.model.RowCount(core.NewQModelIndex())+1, d.model.RowCount(core.NewQModelIndex())+1)
	model.SortFilterModel.BeginResetModel()
	c.manager.CreateNewDiff(fileId, diffId, subjectHash, object, diffPath, year)
	model.SortFilterModel.EndResetModel()
	//model.SortFilterModel.EndInsertRows()

	return diffId
}

func (c *Controller) reloadWatchingRequest() {
	fmt.Println("reloading watched files")
	if filePaths, err := c.manager.RetrieveWatchedFilePaths(); err != nil {
		c.manager.ErrorF("error retrieving the file paths from the database %s", err)
		// return err
	} else {
		for i := range filePaths {
			//first lets add the files to monitor
			if _, err := c.manager.AddFileToMonitor(filePaths[i], true); err != nil {
				c.manager.ErrorF("error adding file %s, %s\r\n", filePaths[i], err)
			} else {
				c.manager.NoticeF("watching file %s", filePaths[i])
			}
		}
	}
}
func (c *Controller) beginWatchingRequest() {
	// just begin watching files.
	if err := c.manager.BeginWatching(); err != nil {
		fmt.Println("error beginning to watch for file changes.")
	} else {
		fmt.Println("begin watching has begun")
	}
}

func (c *Controller) stopWatchingRequest() {
	c.manager.StopWatching()
}

//use this to retrieve information from the manager child of the gui.
func (c *Controller) processInformationChannel() {
	/*
		op codes:
			Op_NewDiff Op = iota
			Op_NewFile
			Op_WatchCommencing
			Op_WatchStopped
			Op_Message
	*/
	for {
		select {
		case op := <-c.information:
			switch op.Code {
			case logic.Op_NewDiff:
				model.SortFilterModel.Update()
				model.EventModel.Update() //the calendar needs to update with new diffs, (so does the detail view aswell)
				c.ClearDetails()          //for now lets clear the details panel of any diffs. //TODO: update the panel if the diffs for today's date change
				// TODO: events := c.EventsForDate(date)
				// TODO: c.DisplayEventDetails(events)
				c.PushNotification("New patch created")
				c.ShowMessage("New!", "New patch created")
			case logic.Op_NewFile:
				model.ListModel.Update()
				c.PushNotification("New file added")
				c.ShowMessage("New!", "New file added")
			case logic.Op_NewBase:
				fmt.Println("New Base: ", op.Code.String())
				c.PushNotification("New restore point created")
			case logic.Op_WatchCommencing:
				fmt.Println("Commencing Watching: ", op.Code.String())
				c.PushNotification("Commencing watching")
			case logic.Op_WatchStopped:
				fmt.Println("Stopped Watching: ", op.Code.String())
				c.PushNotification("Stopped watching")
			case logic.Op_Message:
				fmt.Println("Custom Message: ", op.Custom())
				c.PushNotification(op.Custom())
			default:
				fmt.Println("Hmm not sure: ", op.Code.String())
				c.PushNotification("Hmmm... " + op.Code.String())
			}
		}
	}
}
