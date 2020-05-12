import QtQuick 2.7            //Image
import QtQuick.Controls 2.1   //Button
import QtGraphicalEffects 1.0 //ColorOverlay

import Theme 1.0              //Theme

Button {
  id: button

  property var tooltipText: ""
  property var highlight
  property alias source: image.source
  property string accent

  ToolTip {
    visible: hovered
    delay: 1000
    contentItem: Text {
        color: Theme.mainColorDarker
        text: tooltipText
    }
    background: Rectangle {
        border.color: Theme.mainColorDarker
    }
  }
  background: null

  contentItem: Item {
    Image {
      id: image
      anchors.centerIn: parent
    }

    ColorOverlay {
      id: overlay
      anchors.fill: image
      source: image
      color: highlight || button.hovered ? Theme.mainColorDarker : Theme.charcoal 
    }
  }
}