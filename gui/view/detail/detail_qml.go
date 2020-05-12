// +build qml

package detail

import (
	"fmt"

	"github.com/amlwwalker/pickleit/gui/controller"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/quick"
)

func init() {
	detailController_QmlRegisterType2("Detail", 1, 0, "DetailController")
}

type detailController struct {
	quick.QQuickItem

	_ func()                   `constructor:"init"`
	_ *core.QAbstractItemModel `property:"detailModel"`
	_ *core.QDate              `property:"selectedDate"`
	//should this have a model (see the calendar example - it has a model.)
	//<-controller
	// _ func()                                                       `signal:"showImageLabel"`
	_ func() `signal:"clearDetails"`
	// _ func(profileLabelText string)                                    `signal:"showFileProfile"`
	_ func(id, name, description, imagePath string, selectedDate *core.QDateTime) `signal:"displayDiffDetails"`
	_ func(imagePath string)                                                      `signal:"requestExpandImage"`
	_ func(index int, description string)                                         `signal:"requestSaveDescription"` //hacked while eventModel not mapping data correctly
	_ func(index int)                                                             `signal:"deleteDiffShowRequest"`
	_ func(index int)                                                             `signal:"patchFileShowRequest"`
	_ func()                                                                      `signal:"updateSelectedDate"`
}

func (d *detailController) init() {
	d.SetDetailModel(controller.Instance().DetailModel())
	//<-controller
	d.ConnectRequestExpandImage(controller.Instance().RequestExpandImage)

	d.ConnectRequestSaveDescription(func(index int, description string) { //hack... see above *core.QModelIndex
		fmt.Println("index ", index)
		controller.Instance().RequestSaveDescription(index, description)
	})

	d.ConnectDeleteDiffShowRequest(func(index int) {
		fmt.Printf("del index %+v\r\n", index)
		// so at this point we have an index which will map to the index within the detailModel, however
		// does not map to the index of other models. For this reason then we need to get the actual index from the object
		// before we request something else to process it
		// entry := d.DetailModel().DetailFromIndex(index) //TODO: this shouldn't make requests directly to the model
		// var ok bool
		// if parsedIndex := index.Data(int(core.Qt__DisplayRole)).ToInt(&ok); !ok {
		// 	fmt.Println("cant delete diff, couldn't convert to integer")
		// 	return
		// } else {
		controller.Instance().DeleteDiffShowRequest(index)
		// }
	})
	d.ConnectPatchFileShowRequest(func(index int) {
		fmt.Printf("patch index %+v\r\n", index)
		// var ok bool
		// if parsedIndex := index.Data(int(core.Qt__DisplayRole)).ToInt(&ok); !ok {
		// 	fmt.Println("cant delete diff, couldn't convert to integer")
		// 	return
		// } else {
		controller.Instance().PatchFileShowRequest(index)
		// }
	})

	controller.Instance().ConnectUpdateSelectedDate(func() {
		d.SetSelectedDate(controller.Instance().SelectedDate())
	})

}
