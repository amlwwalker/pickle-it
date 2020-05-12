// +build qml

package view

import (
	"github.com/therecipe/qt/core"

	_ "github.com/amlwwalker/pickleit/gui/assets"
	_ "github.com/amlwwalker/pickleit/gui/calendarView"
	"github.com/amlwwalker/pickleit/gui/controller"
	_ "github.com/amlwwalker/pickleit/gui/model"
	_ "github.com/amlwwalker/pickleit/gui/view/detail"
	_ "github.com/amlwwalker/pickleit/gui/view/dialog"
	_ "github.com/amlwwalker/pickleit/gui/view/diff"
	_ "github.com/amlwwalker/pickleit/gui/view/file"
	_ "github.com/amlwwalker/pickleit/gui/view/imageScroller"
	_ "github.com/amlwwalker/pickleit/gui/view/notifications"
	_ "github.com/amlwwalker/pickleit/gui/view/plugins"
	_ "github.com/amlwwalker/pickleit/gui/view/settings"
	_ "github.com/amlwwalker/pickleit/gui/view/stack"
	_ "github.com/amlwwalker/pickleit/gui/view/status"
	_ "github.com/amlwwalker/pickleit/gui/view/theme"
)

func init() {
	viewController_QmlRegisterType2("View", 1, 0, "ViewController")
}

type viewController struct {
	core.QObject

	_ func() `constructor:"init"`

	//qml->controller
	_ func() `signal:"pickleItVersion"`
	// _ func(index *core.QModelIndex) `signal:"patchFileShowRequest"`
	_ func() `signal:"beginWatchingRequest"`
	_ func() `signal:"reloadWatchingRequest"`
	_ func() `signal:"stopWatchingRequest"`
	// _ func()                        `signal:"deleteDiffRequest"`
	_ func(imagePath string) `signal:"expandImage"`
	_ func()                 `signal:"explanationPopup"`
}

func (v *viewController) init() {

	//qml->controller

	v.ConnectPickleItVersion(controller.Instance().PickleItVersion)
	//qml->controller
	// v.ConnectPatchFileShowRequest(func(index *core.QModelIndex) {
	// 	controller.Instance().PatchFileShowRequest(index)
	// })

	v.ConnectReloadWatchingRequest(controller.Instance().ReloadWatchingRequest)
	v.ConnectBeginWatchingRequest(controller.Instance().BeginWatchingRequest)
	v.ConnectStopWatchingRequest(controller.Instance().StopWatchingRequest)

	// v.ConnectDeleteDiffRequest(controller.Instance().DeleteDiffRequest)
	controller.Instance().ConnectExpandImage(v.ExpandImage)
	controller.Instance().ConnectShowExplanation(v.ExplanationPopup)
}
