// +build qml

package dialog

import (
	"github.com/amlwwalker/pickleit/gui/controller"
	"github.com/therecipe/qt/quick"
)

func init() {
	deleteDialogController_QmlRegisterType2("Dialog", 1, 0, "DeleteDialogController")
}

type deleteDialogController struct {
	quick.QQuickItem

	_ func() `constructor:"init"`

	//qml<-controller
	_ func(index int) `signal:"deleteDiffShowRequest"` //model agnostic index

	//qml->controller
	_ func(index int) `signal:"deleteDiff"`
}

func (d *deleteDialogController) init() {

	//qml<-controller
	// controller.Instance().ConnectDeleteDiffShowRequest(d.DeleteDiffShowRequest)

	//qml->controller
	d.ConnectDeleteDiff(controller.Instance().DeleteDiff)
	controller.Instance().ConnectDeleteDiffShowRequest(d.DeleteDiffShowRequest)
	// controller.Instance().DeleteDiffShowRequest(func(index *core.QModelIndex) {
	// 	a.ConnectDeleteDiffShowRequest(index)
	// })
	//qml->controller
	// a.ConnectPatchDiff(controller.Instance().PatchDiff)
}
