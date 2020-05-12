import QtQuick 2.7            //Image
import QtQuick.Layouts 1.3    //RowLayout
import QtQuick.Controls 2.1   //Button
import QtGraphicalEffects 1.0 //ColorOverlay


Button {
  id: button

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
        color: "green"
      }
    }

    Label {
      text: button.text
      color: "green"
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
