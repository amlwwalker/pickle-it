// +build qml

package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/qml"
	"github.com/therecipe/qt/quick"
	"github.com/therecipe/qt/widgets"

	"github.com/amlwwalker/pickleit/gui/controller"

	_ "github.com/amlwwalker/pickleit/gui/view"

	"github.com/amlwwalker/pickleit/utilities"
)

var (
	buildTagging string
	flavour      string
	buildVersion string
	dbName       string
	buildHash    string
	buildDate    string
	production   string
	persistLogs  string
	virtual      string
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
		Flavour: flavour,
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
	if persistLogs == "FALSE" {
		version.PersistLogs = false
	} else {
		version.PersistLogs = true
	}
	if virtual == "TRUE" {
		version.Virtual = true
	} else {
		version.Virtual = false
	}
	fmt.Printf("version information %+v\r\n", version)
	if !version.Production {
		os.Setenv("QML_DISABLE_DISK_CACHE", "true")
	}
}

func main() {

	// configure UI
	var path string
	if version.Production {
		path = "qrc:/qml/view.qml"
	} else {
		path = filepath.Join(os.Getenv("GOPATH"), "src", "github.com", "amlwwalker", "pickleIt", "gui", "view", "qml", "main.qml")
	}
	core.QCoreApplication_SetOrganizationName("Pickle It") //needed to fix an QML Settings issue on windows
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	fmt.Printf("%+v\r\n", version)
	if version.Virtual { //check config for virtualisation state (manually set)
		quick.QQuickWindow_SetSceneGraphBackend(quick.QSGRendererInterface__Software) //needed to get the application working on VMs when using the windows docker images
	}

	qApp := widgets.NewQApplication(len(os.Args), os.Args)
	qApp.SetQuitOnLastWindowClosed(false)
	engine := qml.NewQQmlApplicationEngine(nil)
	if err := controller.Instance().InitWith(version, engine); err != nil {
		fmt.Printf("Error occured during initialising the controller. Can't move forward. Bailing. %s\r\n", err)
		panic(err)
	}
	defer controller.Instance().TearDown()

	//load any global objects you want available to qml
	engine.RootContext().SetContextProperty("systemSettings", controller.Instance().SystemSettings)
	//make the Version information available to the frontend
	engine.RootContext().SetContextProperty("versioning", controller.Instance().VersionInformation)

	engine.AddImportPath("qrc:/qml/")
	if !version.Production {
		engine.AddImportPath("./view/theme/qml")
		engine.AddImportPath("./view/detail/qml")
		engine.AddImportPath("./view/diff/qml")
		engine.AddImportPath("./view/file/qml")
		engine.AddImportPath("./view/dialog/qml")
		engine.AddImportPath("./view/notifications/qml")
		engine.AddImportPath("./view/plugins/qml")
		engine.AddImportPath("./view/settings/qml")
		engine.AddImportPath("./view/status/qml")
		engine.AddImportPath("./view/stack/qml")
		engine.AddImportPath("./calendarView/qml")
		engine.AddImportPath("./view/imageScroller/qml")
		engine.Load(core.NewQUrl3(path, 0))
	} else {
		engine.Load(core.NewQUrl3("qrc:/qml/main.qml", 0))
	}

	//if there is a qml error then index [0] won't exist and it will crash, so should check length of RootObjects.
	//TODO: Could perhaps default to client mode if UI is broken
	view := gui.NewQWindowFromPointer(engine.RootObjects()[0].Pointer()) //get a handle on the view as a OS object

	qApp.ConnectEvent(func(e *core.QEvent) bool {
		if e.Type() == core.QEvent__ApplicationActivate {
			view.Show()
		}
		return qApp.EventDefault(e)
	})

	//configuring the buttons on the systray, requires the view, so passing it from here
	controller.Instance().Tray.Build(func(bool) {
		fmt.Println("show view from systray")
		//could do a check here, but at the moment just do both
		//bring to front
		view.Raise()
		//show if minimized
		view.Show()
	})

	controller.Instance().PluginManager.InitialisePlugins()

	widgets.QApplication_Exec()
}
