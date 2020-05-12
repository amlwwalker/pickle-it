import QtQuick 2.7 //Image
import QtQuick.Controls 2.1 //GroupBox
import QtQuick.Layouts 1.3 //GridLayout

import Theme 1.0
import Plugins 1.0
import "." as T               //needed for name clash with std Controls.Button
// the plugin controller is designed to load in dynamic plugins to the front end. This will be where options are displayed.
PluginsController {
  GroupBox {
    anchors.fill: parent
    anchors.topMargin: 5
    background: Rectangle {
      color: Theme.white
      border.color: Theme.gray
      anchors.fill: parent
    }
    GridLayout {
      Layout.fillHeight: true
      Component {
          id: pluginDelegate
          Item {
              height: 60
              Column {
                  T.Button {
                    id: index
                    Layout.fillWidth: true
                    text: name
                    code: modelCode
                    image: imageSource
                  }
              }
          }
      }
      ListModel {
        id: model
          ListElement {
              name: "Calendar"
              modelCode: "calendar"
              imageSource: "qrc:/qml/assets/events.png"
          }
          ListElement {
              name: "Listing"
              modelCode: "listing"
              imageSource: "qrc:/qml/assets/list.png"
          }
      }
      ListView {
          height: 200
          id: pluginListView
          model: model
          delegate: pluginDelegate
          focus: true
      }
    }
  }
}