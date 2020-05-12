import QtQuick 2.8
import QtQuick.Controls 2.1 as C
import QtQuick.Templates 2.1 as T
import QtQuick.Layouts 1.3  //ColumnLayout
import Theme 1.0

// Item {
  // property alias txtLabel: label
  // ColumnLayout {
  //   Text {
  //     id: label
  //     text: "Slider"
  //   }
    C.ProgressBar {
        id: progressBar
        // anchors.topMargin: Theme.baseSize
        // anchors.bottomMargin: Theme.baseSize
        // Layout.fillWidth: true
        value: 1.0
        background: Rectangle {
          implicitWidth: 200
          implicitHeight: 10
          radius: 5
          color: Theme.lightGray
          border.color: Theme.gray
        }
        contentItem: Item {
          Rectangle {
          color: Theme.mainColor
          border.color: Theme.mainColor
          width: progressBar.visualPosition * parent.width
          height: parent.height
          radius: 5
          }
        }
    }
  // }
// }