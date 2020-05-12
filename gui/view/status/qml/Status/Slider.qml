import QtQuick 2.8
import QtQuick.Controls 2.1 as C
import QtQuick.Templates 2.1 as T
import QtQuick.Layouts 1.3  //ColumnLayout
import Theme 1.0

C.Slider {
    id: control
    // width: 297
    // height: 50
    stepSize: 1
    to: 18
    from: 10
    value: 14

    background: ColumnLayout {

      Rectangle {
          // horizontalAlignment: Text.AlignRight
          // verticalAlignment: Text.AlignVCenter
          implicitWidth: parent.width
          implicitHeight: 4
          // width: parent.width
          // height: implicitHeight
          radius: 2
          color: "#bdbebf"

          Rectangle {
              width: control.visualPosition * parent.width
              height: parent.height
              color: "#21be2b"
              radius: 2
          }
      }
    }

    handle: Rectangle {
        id: sliderHandle
        x: control.leftPadding + control.visualPosition * (control.availableWidth - width)
        y: control.topPadding + control.availableHeight / 2 - height / 2
        implicitWidth: 20
        implicitHeight: 20
        radius: 10
        color: control.pressed ? Theme.mainColorDarker : Theme.mainColor
        border.color: Theme.lightGray
    }
}
