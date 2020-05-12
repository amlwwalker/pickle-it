// +build qml

package settings

import (
	"github.com/amlwwalker/pickleit/gui/controller"
	"github.com/therecipe/qt/quick"
)

func init() {
	settingsController_QmlRegisterType2("Settings", 1, 0, "SettingsController")
}

type settingsController struct {
	quick.QQuickItem

	_ func() `constructor:"init"`

	//<-controller
	// _ func()                                                       `signal:"showImageLabel"`
	// _ func()                                                       `signal:"clearDetails"`
	// _ func(profileLabelText string)                                `signal:"showFileProfile"`
	_ func(id, name, description string) `signal:"updateSettings"`
	// _ func(date *core.QDate) `signal:"displayQDateDetails"`
}

func (s *settingsController) init() {

	//<-controller
	controller.Instance().ConnectUpdateSettings(s.UpdateSettings)
	// controller.Instance().ConnectRetrieveSettings(s.RetrieveSettings)
	// controller.Instance().ConnectShowFileProfile(d.ShowFileProfile)
	// controller.Instance().ConnectDisplayDiffDetails(d.DisplayDiffDetails)

	// controller.Instance().ConnectDisplayDiffDetails(func(id, name, description string, selectedDate *core.QDateTime) {
	// 	s.DisplayDiffDetails(id, name, description, selectedDate)
	// })
	// d.ConnectDisplayQDateDetails(func (date *core.QDate) { controller.Instance().ConnectDisplayQDateDetails(date) })

}
