import QtQuick.Layouts 1.3  //ColumnLayout
import QtQuick.Controls 2.1
import QtQuick 2.2

import Theme 1.0

Rectangle {
  id:dropzone
  color: Theme.white
  border.color: Theme.lightGray
  border.width: 5
  radius: 5
  property var progressBar
  Text {
    text: "DROP ZONE"
    anchors.horizontalCenter: dropzone.horizontalCenter
    anchors.verticalCenter: dropzone.verticalCenter
    font.bold: true
    font.pixelSize: 20
    font.family: "Helvetica"
    color: Theme.darkGray
  }
    Image {
      anchors.horizontalCenter: dropzone.horizontalCenter
      anchors.verticalCenter: dropzone.verticalCenter

      height: 120
      width: 120
      opacity: 0.2
      fillMode: Image.PreserveAspectFit
      source: "qrc:/qml/assets/add-file.png"
    }
  Layout.fillHeight: true
  Layout.fillWidth: true
  DropArea {
    id: dropArea
    anchors.fill: parent
    onEntered: {
      console.log("entered")
    }
    onDropped: {
      console.log("source ", drop.urls[0])
      addFile(drop.urls[0])
      progressBar.indeterminate = true
    }
    onExited: {
      console.log("exited")
    }
    states: [
        State {
            name: "normal"
            when: !dropArea.containsDrag
        },
        State {
            name: "down"
            when: dropArea.containsDrag
            PropertyChanges {
                target: dropzone
                color: Theme.mainColorDarker
            }
        }
    ]
  }
}