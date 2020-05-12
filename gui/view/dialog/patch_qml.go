// +build qml

package dialog

import (
	"github.com/amlwwalker/pickleit/gui/controller"
	"github.com/therecipe/qt/quick"
)

func init() {
	patchDialogController_QmlRegisterType2("Dialog", 1, 0, "PatchDialogController")
}

type patchDialogController struct {
	quick.QQuickItem

	_ func() `constructor:"init"`

	//qml<-controller
	_ func(index int) `signal:"patchFileShowRequest"` //call this to show the pop up

	//qml->controller
	_ func(index int) `signal:"patchFile"` //call this when user is convinced. This should really be an index
}

func (a *patchDialogController) init() {

	//qml<-controller

	controller.Instance().ConnectPatchFileShowRequest(a.PatchFileShowRequest)
	//qml->controller
	a.ConnectPatchFile(controller.Instance().PatchFile)
}
