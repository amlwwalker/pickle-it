// +build qml

package plugins

import (
	"github.com/amlwwalker/pickleit/gui/controller"
	"github.com/therecipe/qt/quick"
)

func init() {
	pluginsController_QmlRegisterType2("Plugins", 1, 0, "PluginsController")
}

type pluginsController struct {
	quick.QQuickItem

	_ func() `constructor:"init"`

	//qml->controller
	_ func(code string) `signal:"changeStackView"`
}

func (p *pluginsController) init() {
	//connect up buttons etc to the controller of the frontend

	//qml->controller
	p.ConnectChangeStackView(controller.Instance().ChangeStackView)
}
