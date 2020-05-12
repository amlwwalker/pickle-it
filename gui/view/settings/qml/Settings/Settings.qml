import QtQuick 2.7          //Rectangle
import QtQuick.Controls 2.1 //StackView
import QtQuick.Layouts 1.3  //GridLayout

import Settings 1.0
import Theme 1.0
import "." as T
SettingsController {
      Layout.fillWidth: true
      ColumnLayout {
          id: mainLayout
          anchors.right: parent.right
          anchors.left: parent.left
          GroupBox {
              id: gridBox
              Layout.fillWidth: true
              GridLayout {
                  id: gridLayout
                  rows: 5
                  flow: GridLayout.TopToBottom
                  anchors.fill: parent

                  Label { text: "Enable file watching on start"; Layout.fillWidth: true  }
                //   Label { text: "Save a screen shot of the file when creating versions"; Layout.fillWidth: true  }
                  Label { text: "Show notification when a new file is ready to be watched"; Layout.fillWidth: true  }
                  Label { text: "Show notification when a new patch is created"; Layout.fillWidth: true  }
                  Label { text: "Override existing file when restoring"; Layout.fillWidth: true  }
                  // Label { text: "Happy to send back statistics about how the app is performing"; Layout.fillWidth: true  }
                  Label { text: "Show welcome screen on start"; Layout.fillWidth: true  }

                  CheckBox { id: autoWatchCheck; checked: systemSettings.autoWatchCheck; onClicked: { systemSettings.autoWatchCheck = autoWatchCheck.checked; systemSettings.someSettingChanged()}}
                //   CheckBox { id: screenshotCheck; checked: systemSettings.screenshotCheck; onClicked: { systemSettings.screenshotCheck = screenshotCheck.checked; systemSettings.someSettingChanged()}}
                  CheckBox { id: newFileReadySystemNotifyCheck; checked: systemSettings.newFileReadySystemNotifyCheck; onClicked: { systemSettings.newFileReadySystemNotifyCheck = newFileReadySystemNotifyCheck.checked; systemSettings.someSettingChanged()}}
                  CheckBox { id: patchSystemNotifyCheck; checked: systemSettings.patchSystemNotifyCheck; onClicked: { systemSettings.patchSystemNotifyCheck = patchSystemNotifyCheck.checked; systemSettings.someSettingChanged()}}
                  CheckBox { id: overrideExistingCheck; checked: systemSettings.overrideExistingCheck; onClicked: { systemSettings.overrideExistingCheck = overrideExistingCheck.checked; systemSettings.someSettingChanged()}}
                  // CheckBox { id: statisticsCheck; checked: systemSettings.statisticsCheck; onClicked: { systemSettings.statisticsCheck = statisticsCheck.checked; systemSettings.someSettingChanged()}}
                  CheckBox { id: welcomeCheck; checked: systemSettings.welcomeCheck; onClicked: { systemSettings.welcomeCheck = welcomeCheck.checked; systemSettings.someSettingChanged()}}
              }
          }
          Text {
            text: "Information"
            font.bold: true
            font.underline: true
            font.pixelSize: 20
            font.family: "Helvetica"
          }
          TextEdit {
            text: "Build Flavour " + versioning.flavour
            selectByMouse: true
          }
          TextEdit {
            text: "Version " + versioning.version
            selectByMouse: true
          }
          TextEdit {
            text: "Build Hash " + versioning.hash
            selectByMouse: true
          }
          TextEdit {
            text: "Build Date " + versioning.date
            selectByMouse: true
          }
      }
}
