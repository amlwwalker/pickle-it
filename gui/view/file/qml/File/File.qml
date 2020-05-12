import QtQuick 2.6
import QtQuick.Controls 2.3
import QtQuick.Layouts 1.3  //GridLayout

import File 1.0
import Theme 1.0

FileController {
  GroupBox {
    anchors.fill: parent
    anchors.topMargin: 5

    ComboBox {
      id: control
      anchors.fill: parent
      anchors.margins: 5
      model: listModel
      textRole: "display"

      indicator: Canvas {
          id: canvas
          x: control.width - width - control.rightPadding
          y: control.topPadding + (control.availableHeight - height) / 2
          width: 12
          height: 8
          contextType: "2d"

          Connections {
              target: control
              onPressedChanged: canvas.requestPaint()
          }

          onPaint: {
              context.reset();
              context.moveTo(0, 0);
              context.lineTo(width, 0);
              context.lineTo(width / 2, height);
              context.closePath();
              context.fillStyle = control.pressed ? Theme.mainColor : Theme.mainColorDarker;
              context.fill();
          }
      }

      contentItem: Text {
          leftPadding: 15
          rightPadding: control.indicator.width + control.spacing

          text: control.displayText
          font: control.font
          color: control.pressed ? Theme.mainColor : Theme.mainColorDarker
          verticalAlignment: Text.AlignVCenter
          elide: Text.ElideRight
      }

      background: Rectangle {
          implicitWidth: 120
          implicitHeight: 40
          border.color: control.pressed ? Theme.mainColor : Theme.mainColorDarker
          border.width: control.visualFocus ? 2 : 1
          radius: 2
      }

      popup: Popup {
          y: control.height - 1
          width: control.width
          implicitHeight: contentItem.implicitHeight
          padding: 1

          contentItem: ListView {
              clip: true
              implicitHeight: contentHeight
              model: control.popup.visible ? control.delegateModel : null
              currentIndex: control.highlightedIndex
              // highlight: Rectangle { color: "#69697C" }
              // highlightFollowsCurrentItem: true
              delegate:
              ItemDelegate {
                  width: control.width
                  contentItem: RowLayout {
                    Text {
                        text: modelData
                        color: "#21be2b"
                        font: control.font
                        elide: Text.ElideRight
                        verticalAlignment: Text.AlignVCenter
                    }
                    
                  }
                  // highlighted: control.highlightedIndex == index
              }
              // ScrollIndicator.vertical: ScrollIndicator { }
          }

          background: Rectangle {
              border.color: Theme.mainColorDarker
              radius: 2
          }
      }

      onActivated: {
          console.log("dropDownList Activated");
          var currentItem = delegateModel.items.get(currentIndex)
          console.log("Read Model Value: " + currentItem.model.display);
          console.log("changing to index ", index)
        changeFile(index)
      }
    }
  }
}
