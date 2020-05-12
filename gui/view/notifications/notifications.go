// +build qml

package notifications

import (
	"fmt"

	"github.com/amlwwalker/pickleit/gui/controller"
	"github.com/therecipe/qt/quick"
)

func init() {
	notificationsController_QmlRegisterType2("Notifications", 1, 0, "NotificationsController")
	fmt.Println("initialised")
}

type notificationsController struct {
	quick.QQuickItem
	_ func()                    `constructor:"init"`
	_ func(notification string) `signal:"pushNotification"`
}

func (s *notificationsController) init() {
	fmt.Println("notificationsController initialised")

	controller.Instance().ConnectPushNotification(func(notification string) {
		fmt.Println("receiving notification ", notification)
		s.PushNotification(notification)
	})
	//qml->controller
}
