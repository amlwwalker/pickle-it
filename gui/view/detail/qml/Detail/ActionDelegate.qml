import QtQuick 2.7          //Image
import QtQuick.Controls 2.1 //GroupBox
import QtQuick.Layouts 1.3  //GridLayout

import Theme 1.0
Rectangle {
  property var currentPatch
  Row {
      width: 200
      x: 70
      y: -10
      DelegateButton {
        id: deleteButton
        source: "qrc:/qml/assets/ic_delete_forever_black_24px.svg"
        tooltipText: "delete patch"
        accent: Theme.gray
        onClicked: {
          console.log(tableView.model.data(tableView.model.index(styleData.row, 0), Qt.UserRole + 1), tableView.model.index(styleData.row, 0), Qt.UserRole + 1)
          // template.deleteRequest(tableView.model.data(tableView.model.index(styleData.row, 0), Qt.UserRole + 1))
        }
        width: parent.width * 0.15
      }
        DelegateButton {
        id: patchButton
        highlight: currentPatch
        // anchors {
        //   top: parent.top
        //   // left: parent.left
        //   // right: deleteButton.left
        //   bottom: parent.bottom
        //   }
        width: parent.width * 0.15

        source: "qrc:/qml/assets/patch.png"
        tooltipText: "apply patch"
        accent: Theme.dark
        // onClicked: template.showDownload(tableView.model.data(tableView.model.index(styleData.row, 0), Qt.UserRole + 1))

        // visible: styleData.value.available != null ? styleData.value.available : false
      }
        DelegateButton {
        id: uploadButton

        // anchors {
          // top: parent.top
          // // left: patchButton.right
          // // right: parent.right
          // bottom: parent.bottom
          // }
        width: parent.width * 0.15

        source: "qrc:/qml/assets/ic_cloud_upload_black_24px.svg"
        tooltipText: "Upload Version"
        accent: Theme.dark
        // onClicked: template.showDownload(tableView.model.data(tableView.model.index(styleData.row, 0), Qt.UserRole + 1))

        // visible: styleData.value.available != null ? styleData.value.available : false
      }

    }
}