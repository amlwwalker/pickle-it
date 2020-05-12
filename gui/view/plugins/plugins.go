// +build !qml

package plugins

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"

	"github.com/amlwwalker/pickleit/gui/controller"
)

type pluginsController struct {
	widgets.QGroupBox

	_ func() `constructor:"init"`

	_ *core.QAbstractItemModel `property:"listModel"`

	//->controller
	_ func(row int) `signal:"changeFile"`

	fileView *widgets.QComboBox
}

func (a *pluginsController) init() {

	a.fileView = widgets.NewQComboBox(nil)
	a.fileView.SetModel(controller.Instance().FileModel())

	layout := widgets.NewQGridLayout2()
	layout.AddWidget(a.fileView, 0, 0, 0)
	a.SetLayout(layout)

	//

	//->controller
	a.fileView.ConnectCurrentIndexChanged(controller.Instance().ChangeFile)
}
