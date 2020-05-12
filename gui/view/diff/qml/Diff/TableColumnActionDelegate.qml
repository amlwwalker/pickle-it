import QtQuick 2.7          //Component
import QtQuick.Controls 1.4 //TableViewColumn
import QtQuick.Controls 2.1 //Label
import QtQuick.Layouts 1.3

import Theme 1.0            //Theme
// import FilesTemplate 1.0    //ActionButtonTemplate

import "."                  //ProgressBar

TableViewColumn {

  property QtObject tableView
  // property ActionButtonTemplate template: ActionButtonTemplate{}
  property bool currentPatch
  title: role
  resizable: false
  movable: false
  // width: parent.width * 0.25
  

  delegate: Component { GridLayout {
      rows: 2
      anchors.fill: parent
      TableColumnDelegateButton {
        id: holder
        Layout.fillWidth: true
      }

      TableColumnDelegateButton {
        id: deleteButton
        // anchors {
        //   top: parent.top
        //   // left: patchButton.right//downloadButton.visible ? downloadButton.right : progressBar.right
        //   // right: parent.right
        //   bottom: parent.bottom
        // }
        // anchors.centerIn: parent
        // width: parent.width * 0.2
        Layout.fillWidth: true
        source: "qrc:/qml/assets/ic_delete_forever_black_24px.svg"
        tooltipText: "delete patch"
        accent: Theme.gray
        onClicked: {
          console.log("delete: ", tableView.model.data(tableView.model.index(styleData.row, 0), Qt.UserRole + 1), tableView.model.index(styleData.row, 0), Qt.UserRole + 1)
          // template.deleteRequest(tableView.model.data(tableView.model.index(styleData.row, 0), Qt.UserRole + 1))
          // deleteDiff(tableView.model.index(styleData.row, 0))
          deleteDiffShowRequest(tableView.model.index(styleData.row, 0))
        }
      }
        TableColumnDelegateButton {
        id: patchButton
        // highlight: currentPatch
        // anchors {
        //   top: parent.top
        //   // left: parent.left
        //   // right: deleteButton.left
        //   bottom: parent.bottom
        //   }
          // anchors.centerIn: parent
        // width: parent.width * 0.2
        Layout.fillWidth: true
        source: "qrc:/qml/assets/patch.png"
        tooltipText: "apply patch"
        accent: Theme.dark
        onClicked: {
          console.log("patch: ", tableView.model.data(tableView.model.index(styleData.row, 0), Qt.UserRole + 1), tableView.model.index(styleData.row, 0), Qt.UserRole + 1)
          // template.deleteRequest(tableView.model.data(tableView.model.index(styleData.row, 0), Qt.UserRole + 1))
          patchFileShowRequest(tableView.model.index(styleData.row, 0))
        }
        // onClicked: template.showDownload(tableView.model.data(tableView.model.index(styleData.row, 0), Qt.UserRole + 1))

        // visible: styleData.value.available != null ? styleData.value.available : false
      }

    // ProgressBar {
    //   id: progressBar

    //   anchors {
    //     verticalCenter: parent.verticalCenter
    //     left: parent.left
    //   }
    //   width: parent.width * 0.75
    //   height: parent.height * 0.6

    //   // progressBarText: styleData.value.text != null ? styleData.value.text : ""
    //   value: styleData.value.value != null ? styleData.value.value : 0
    //   // error: styleData.value.error != null ? styleData.value.error : false

    //   visible: !downloadButton.visible
    // }
    // TableColumnDelegateButton {
    //   id: patchButton
    //   anchors {
    //     top: parent.top
    //     left: downloadButton.right//downloadButton.visible ? downloadButton.right : progressBar.right
    //     right: deleteButton.left
    //     bottom: parent.bottom
    //   }

    //   source: "qrc:/qml/assets/ic_insert_drive_file_black_24px.svg"
    //   accent: Theme.gray
    //   onClicked: template.deleteRequest(tableView.model.data(tableView.model.index(styleData.row, 0), Qt.UserRole + 1))
    // }

      //   TableColumnDelegateButton {
      //   id: uploadButton

      //   anchors {
      //     top: parent.top
      //     // left: patchButton.right
      //     // right: parent.right
      //     bottom: parent.bottom
      //     }
      //   width: parent.width * 0.2

      //   source: "qrc:/qml/assets/ic_cloud_upload_black_24px.svg"
      //   tooltipText: "Upload Version"
      //   accent: Theme.dark
      //   // onClicked: template.showDownload(tableView.model.data(tableView.model.index(styleData.row, 0), Qt.UserRole + 1))

      //   // visible: styleData.value.available != null ? styleData.value.available : false
      // }
      //   TableColumnDelegateButton {
      //   id: currentIndicator

      //   anchors {
      //     top: parent.top
      //     // left: patchButton.right
      //     // right: parent.right
      //     bottom: parent.bottom
      //     }
      //   width: parent.width * 0.2

      //   source: {
      //     console.log(styleData.row, tableView.model.index(styleData.row, 0))
      //     return "qrc:/qml/assets/electric-current-symbol.png" //return styleData.row == currentFlag ? "qrc:/qml/assets/eventindicator.png" : ""
      //   }
      //   tooltipText: "Currently checked out"
      //   accent: styleData.row == currentFlag ? Theme.lightGray : Theme.darkGray
      //   // onClicked: template.showDownload(tableView.model.data(tableView.model.index(styleData.row, 0), Qt.UserRole + 1))

      //   // visible: styleData.value.available != null ? styleData.value.available : false
      // }
    }
  }
}
