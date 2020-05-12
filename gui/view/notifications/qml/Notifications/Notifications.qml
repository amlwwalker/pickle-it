import QtQuick 2.2

import Theme 1.0
import Notifications 1.0
import "."
NotificationsController {
  Rectangle {
    color: "transparent"
    height: 50
    width: parent.width
    ToastManager {
        id: toast
    }
    // Timer {
    //     interval: 1000
    //     repeat: true
    //     running: true
    //     property int i: 0
    //     onTriggered: {
    //         toast.show("This timer has triggered " + (++i) + " times!");
    //     }
    // }

    // Timer {
    //     interval: 3000
    //     repeat: true
    //     running: true
    //     property int i: 0
    //     onTriggered: {
    //         toast.show("This important message has been shown " + (++i) + " times.", 5000);
    //     }
    // }
  }
    onPushNotification: {
      console.log("received ", notification)
      toast.show(notification);
    }
}