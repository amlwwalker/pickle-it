import QtQuick 2.7          //Component
import QtQuick.Controls 2.1 //Label

import Theme 1.0            //Theme

Component { Item {

  height: Theme.minHeight * 0.05
  visible: styleData.value != ""

  Rectangle {
    id: leftRound

    anchors {
      top: parent.top
      left: parent.left
      bottom: parent.bottom
    }

    width: styleData.column == 0 ? 50 : 0
    // radius: styleData.column == 0 ? 10 : 0
    color: Theme.white

    visible: styleData.column == 0
  }

  Rectangle {
    id: mainRectangle

    anchors {
      top: parent.top
      left: leftRound.horizontalCenter
      right: rightRound.horizontalCenter
      bottom: parent.bottom
    }

    color: Theme.white
  }

  Rectangle {
    id: rightRound

    anchors {
      top: parent.top
      right: parent.right
      bottom: parent.bottom
    }
    width: styleData.column == 3 ? 20 : 0
    // radius: styleData.column == 3 ? 10 : 0
    color: Theme.white

    visible: styleData.column == 3
  }

  Label {
    anchors.centerIn: parent

    text: styleData.value
    color: Theme.mainColor
    font.pointSize: 14
    font.bold: true
  }
} }

/*
import QtQuick 2.4
import QtQuick.Controls 1.4 //Menu

import Theme 1.0
Rectangle {
    height: textItem.implicitHeight * 1.2
    width: textItem.implicitWidth
    color: Theme.white
    Text {
        id: textItem
        anchors.fill: parent
        verticalAlignment: Text.AlignVCenter
        horizontalAlignment: styleData.textAlignment
        anchors.leftMargin: 12
        text: styleData.value
        elide: Text.ElideRight
        color: Theme.mainColorDarker
        renderType: Text.NativeRendering
    }
    Rectangle {
        anchors.right: parent.right
        anchors.top: parent.top
        anchors.bottom: parent.bottom
        anchors.bottomMargin: 1
        anchors.topMargin: 1
        width: 1
        // color: Theme.mainColorwhiteer
    }
}

*/