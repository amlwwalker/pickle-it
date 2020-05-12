import QtQuick 2.7            //Image
import QtQuick.Layouts 1.3    //RowLayout
import QtQuick.Controls 2.1   //Button
import QtGraphicalEffects 1.0 //ColorOverlay

import Theme 1.0

Button {
  id: button
  font: Theme.font
  property string code
  property alias image : logo.source
  background: null

  contentItem: RowLayout {
    spacing: 20

    Item {
      Layout.fillWidth: true

      Image {
        id: logo
        anchors.centerIn: parent
      }

      ColorOverlay {
        anchors.fill: logo
        source: logo
        color: Theme.mainColor
      }
    }

    Label {
      text: button.text
      font: Theme.font
      color: Theme.mainColor
    }

    Item {
      Layout.fillWidth: true
    }
  }

  onClicked: {
    console.log(code)
    changeStackView(code)
  }
}
