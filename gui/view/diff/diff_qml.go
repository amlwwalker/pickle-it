// +build qml

package diff

import (
	"fmt"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/quick"

	"github.com/amlwwalker/pickleit/gui/controller"
)

func init() {
	diffController_QmlRegisterType2("Diff", 1, 0, "DiffController")
}

type diffController struct {
	quick.QQuickItem

	_ func() `constructor:"init"`

	_ *core.QAbstractItemModel `property:"viewModel"`

	//qml<-controller
	_ func() `signal:"diffAdded"`
	_ func() `signal:"deleteDiffRequest"`
	_ func() `signal:"deleteDiffCommand"`

	//qml->controller
	// _ func()                                     `signal:"showImageLabel"`
	// _ func(index *core.QModelIndex)              `signal:"deleteDiff"`
	// _ func(index *core.QModelIndex)              `signal:"patchFile"`
	_ func(column int, order core.Qt__SortOrder) `signal:"sortTableView"`
	_ func(index *core.QModelIndex)              `signal:"showDiffDetails"`
	_ func(index *core.QModelIndex)              `signal:"deleteDiffShowRequest"`
	_ func(index *core.QModelIndex)              `signal:"patchFileShowRequest"`
}

func (d *diffController) init() {

	d.SetViewModel(controller.Instance().DiffModel())

	//<-controller
	controller.Instance().ConnectAddDiff(func(string, string, string, string, int) {
		fmt.Println("detected row inserted")
		d.DiffAdded()
	})
	// controller.Instance().ConnectDeleteDiffRequest(d.DeleteDiffRequest)
	// controller.Instance().ConnectDeleteDiffCommand(d.DeleteDiffCommand)

	//->controller
	// d.ConnectShowImageLabel(controller.Instance().ShowImageLabel)
	// d.ConnectDeleteDiff(func(index *core.QModelIndex) { fmt.Println("index ", index); controller.Instance().DeleteDiff(index) })
	// d.ConnectPatchFile(func(index *core.QModelIndex) { fmt.Println("index ", index); controller.Instance().PatchFile(index) })
	d.ConnectSortTableView(controller.Instance().SortTableView)
	d.ConnectShowDiffDetails(func(index *core.QModelIndex) { controller.Instance().ShowDiffDetails(index) })

	d.ConnectDeleteDiffShowRequest(func(index *core.QModelIndex) {
		fmt.Printf("del index %+v\r\n", index)
		var ok bool
		if parsedIndex := index.Data(int(core.Qt__UserRole) + 1).ToInt(&ok); !ok {
			fmt.Println("cant delete diff, couldn't convert to integer")
			return
		} else {
			controller.Instance().DeleteDiffShowRequest(parsedIndex)
		}
	})
	d.ConnectPatchFileShowRequest(func(index *core.QModelIndex) {
		fmt.Printf("patch index %+v\r\n", index)
		var ok bool
		if parsedIndex := index.Data(int(core.Qt__UserRole) + 1).ToInt(&ok); !ok {
			fmt.Println("cant delete diff, couldn't convert to integer")
			return
		} else {
			controller.Instance().PatchFileShowRequest(parsedIndex)
		}
	})
	// controller.Instance().ConnectDeleteDiffShowRequest(func(index *core.QModelIndex) {
	// 	fmt.Println("index ", index)
	// 	d.DeleteDiffShowRequest(index)
	// })
	controller.Instance().ConnectDiffAdded(d.DiffAdded)
}
