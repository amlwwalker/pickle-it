// +build qml

package status

import (
	"fmt"

	"github.com/amlwwalker/pickleit/gui/controller"
	"github.com/therecipe/qt/quick"
)

func init() {
	statusController_QmlRegisterType2("Status", 1, 0, "StatusController")
	fmt.Println("initialised")
}

type statusController struct {
	quick.QQuickItem
	_ func()                          `constructor:"init"`
	_ func(code string)               `signal:"pushNewStatusCode"`
	_ func(value float64, reset bool) `signal:"setDeterminateProgress"`
	_ func(filepath string)           `signal:"addFile"`

	// _ bool `property:"progressValue"` //not sure we need this
}

func (s *statusController) init() {
	fmt.Println("statusController initialised")

	//qml->controller
	s.ConnectAddFile(controller.Instance().AddFile)
	//when the controllers progress bar is called, it will call the SetDeterminateProgress here
	//this can be either a function that qml listens to, or could be values stored here
	controller.Instance().ConnectSetDeterminateProgress(func(value float64, reset bool) {
		s.SetDeterminateProgress(value, reset)
	})

	// fmt.Println("WARNING - [STATUS_QML] IS TESTING THE PROGRESS BAR")
	// go func() {
	// 	for i := 0.0; i < 10.0; i++ {
	// 		time.Sleep(time.Second * 1)
	// 		s.SetDeterminateProgress(0.1*i, false)
	// 	}
	// }()
}
