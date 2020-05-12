
import QtQuick.Controls 2.4 //Menu
import QtQuick 2.2

import ImageScroller 1.0

ImageController {
  property var eventId
  ListView {
    id: view
      y: 5
      x: 30
      width: parent.width// - 10
      height: parent.height// - 10// - 10
    model: imageModel
    delegate: Image {
      id: img
      source: display.path
      width: view.width// - 10
      height: view.height// - 10
      fillMode: Image.PreserveAspectFit
      smooth: true
      MouseArea {
          id: ma
          anchors.fill: parent
          onClicked: requestExpandImage(display.path)
      }
      
    }

    orientation: ListView.Horizontal
    snapMode: ListView.SnapToItem
  }
}