import QtQuick.Layouts 1.3  //ColumnLayout
import QtQuick.Controls 2.1
import QtQuick.Controls.Styles 1.4
import QtQuick 2.2

import Theme 1.0
import Status 1.0
import Notifications 1.0
import "." as Custom               //needed for name clash with std Controls

StatusController {
  // anchors.fill: parent

  ColumnLayout {
    Layout.fillHeight: true
    width: parent.width
    Custom.Switch {
        id: watchSwitch
        Layout.fillWidth: true
        title: "Watch Files"
        checked: systemSettings.autoWatchCheck
        onClicked: {
          if (checked) {
            //this should happen backend
            core.reloadWatchingRequest()
            core.beginWatchingRequest()
            watchSwitch.title = "Stop Watching"
            // progressAnimation.duration = 200
            // progressAnimation.running = true
            // toast.show("Watching files has started");
          } else {
            console.log("WARNING NEED TO STOP WATCHER!")
            watchSwitch.title = "Watch Files"
            core.stopWatchingRequest()
            // toast.show("Stopped watching files");
          }
        }
    }
    Custom.DropZone {
      width: parent.width
      height: 175
      progressBar: determinateProgress
    }
    Custom.ProgressBar {
      id: determinateProgress
      value: 0.0
      indeterminate: false
      Layout.fillWidth: true
    }
    // Custom.ProgressBar {
    //   id: indeterminateProgress
    //   indeterminate: true
    //   Layout.fillWidth: true
    // }
    // PropertyAnimation {
    //     id: progressAnimation
    //     target: determinateProgress
    //     property: "value"
    //     from: 0.0
    //     to: 1.0
    //     duration: 5000
    //     running: false
    //     loops: 1
    // }
    Notifications {
        id:notifications
        Layout.fillHeight: true
        Layout.fillWidth: true
    }  
  }
  onSetDeterminateProgress: {
    if (reset) {
      determinateProgress.value = 0.0
    }
    determinateProgress.value = value
  }
}