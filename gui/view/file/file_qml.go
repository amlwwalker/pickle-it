// +build qml

package file

import (
	"fmt"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/quick"

	"github.com/amlwwalker/pickleit/gui/controller"
)

func init() {
	fileController_QmlRegisterType2("File", 1, 0, "FileController")
}

type fileController struct {
	quick.QQuickItem
	_ func() `constructor:"init"`

	_ *core.QAbstractItemModel `property:"listModel"`

	//qml->controller
	_ func(row int) `signal:"changeFile"`
}

func (f *fileController) init() {
	f.SetListModel(controller.Instance().FileModel())

	//qml<-controller
	f.ListModel().ConnectRowsInserted(func(*core.QModelIndex, int, int) {
		fmt.Println("detected row inserted")
		model := f.ListModel()
		f.SetListModel(nil)
		f.SetListModel(model)
	})

	//qml->controller
	f.ConnectChangeFile(controller.Instance().ChangeFile)
}
