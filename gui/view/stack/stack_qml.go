package stack

import (
	"github.com/amlwwalker/pickleit/gui/controller"
	"github.com/therecipe/qt/quick"
)

func init() { stackController_QmlRegisterType2("Stack", 1, 0, "StackController") }

type stackController struct {
	quick.QQuickItem
	_ func()            `constructor:"init"`
	_ func(code string) `signal:"pushNewStackCode"`
}

func (s *stackController) init() {
	controller.Instance().ConnectPushNewStackCode(s.PushNewStackCode)
}
