package logic

//https://github.com/apsdehal/go-logger
import (
	"io"
	"os/user"
	"sync"
	"time"

	database "github.com/amlwwalker/pickleit/database"
	fsmanager "github.com/amlwwalker/pickleit/fsmanager"
	"github.com/amlwwalker/pickleit/utilities"
	logger "github.com/apsdehal/go-logger"
	"github.com/therecipe/qt/qml"
)

type Manager struct {
	Version  utilities.Version
	Settings *UserSettings
	*logger.Logger
	*sync.WaitGroup
	watcher              fsmanager.Watcher
	patcher              fsmanager.Patcher
	dB                   *database.DB
	Informer             chan OperatingMessage
	ProgressCommunicator io.WriteCloser
}

type CustomPlugin interface {
	Init()
	Name() string
	Description() string
}

type PluginManager struct {
	engine   *qml.QQmlApplicationEngine
	informer chan OperatingMessage
	path     string
	plugins  []string
}

type UserSettings struct {
	Usr            user.User
	versionFormat  string
	darkMode       bool
	licenseKey     string
	override       bool
	machineID      string
	systemSettings utilities.UXSettings
}
type VersioningFormat struct {
	bigVersion    int64
	littleVersion int64
	microVersion  int64
	currentTime   time.Time
	client        string
	job           string
	userId        string
	owner         string
	hash          string
	message       string
}
