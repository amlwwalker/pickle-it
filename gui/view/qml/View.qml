import QtQuick.Controls 2.4 //Menu
import QtQuick.Dialogs 1.2  //MessageDialog
import QtQuick.Layouts 1.3  //GridLayout
import QtGraphicalEffects 1.0 //ColorOverlay
import QtQuick.Controls.Styles 1.4
import QtQuick 2.7            //Image


import "." as Custom               //needed for name clash with std Controls

import View 1.0
import Theme 1.0
import Detail 1.0
import Dialog 1.0
import ImageScroller 1.0
import Plugins 1.0
import Status 1.0
import Stack 1.0
import File 1.0

ApplicationWindow {
  id: app
  visible: true
  title: "Pickle It"
  minimumWidth: 1200; minimumHeight: 768
  // visibility: "FullScreen"
  width: minimumWidth; height: minimumHeight
  property string imgPath
  property int explanationPage: 1
  property ViewController core: ViewController{
    onExpandImage: {
      explanation1.visible = false
      explanation2.visible = false

      // img.visible = true
      // imgPath = imagePath
      // console.log(imagePath)
      // popup.open()
      // swipeView.currentIndex = 1
      console.log("pop up popped")
      imgPath = imagePath
      img.height = popup.height - 30
      popup.open()
    }
    onExplanationPopup: {
      // swipeView.currentIndex = 0
      console.log("pop up popped")
      // imgPath = "qrc:/qml/assets/final-file-meme.png"
      explanation1.visible = true
      // img.height = popup.height / 2
      popup.open()
    }
  }

  PatchDialog {}
  DeleteDialog {}

  MessageDialog {
    id: aboutDialog
    title: "About Pickle It"
    text: "<p><b>Pickle It</b> is an application that allows any file to be monitored for changes and all previous versions to be available to you. It's the ultimate best friend while you are working away. You can literally never lose any work. Don't worry about maintaining versions, just allow your creative juices to flow and we'll Pickle the changes as they occur. Pickle It. Preserve your work, properly.</p>"
  }

  Popup {
    id: popup
    anchors.centerIn: parent
    width: parent.width - 50 * 2
    height: parent.height - 50 * 2
    leftMargin: 50
    topMargin: 50
    rightMargin: 50
    bottomMargin: 50
    modal: true
    focus: true
    closePolicy: Popup.CloseOnEscape | Popup.CloseOnPressOutsideParent | Popup.CloseOnPressOutside
    Button {
      id: closeButton
      anchors.right: parent.right
      anchors.top: parent.top
      background: Image {
        source: "qrc:/qml/assets/close_rounded.png"
        fillMode: Image.PreserveAspectFit
      }
      MouseArea {
        anchors.fill: parent
        cursorShape: Qt.PointingHandCursor
        onClicked: popup.close()
      }
    }
    Button {
      id: nextButton
      anchors.top: parent.top
      anchors.bottom: parent.bottom      
      anchors.right: parent.right
      background: Image {
        source: "qrc:/qml/assets/arrow-forward.png"
        fillMode: Image.PreserveAspectFit
      }
      MouseArea {
        anchors.fill: parent
        cursorShape: Qt.PointingHandCursor
        onClicked: {
          explanationPage++
          if (explanationPage >= 2) {
            nextButton.visible = false
          }
          lastButton.visible = true
          switch(explanationPage) {
            case 1:
              explanation1.visible = true
              explanation2.visible = false
              return
            case 2:
              explanation1.visible = false
              explanation2.visible = true
              return
            default: return
          }
        }
      }
    }
    Button {
      id: lastButton
      visible: false
      anchors.top: parent.top
      anchors.bottom: parent.bottom      
      anchors.left: parent.left
      background: Image {
        source: "qrc:/qml/assets/arrow-backward.png"
        fillMode: Image.PreserveAspectFit
      }
      MouseArea {
        anchors.fill: parent
        cursorShape: Qt.PointingHandCursor
        onClicked: {
          explanationPage--
          if (explanationPage <= 1) {
            lastButton.visible = false
          }
          nextButton.visible = true
          switch(explanationPage) {
            case 1:
              explanation1.visible = true
              explanation2.visible = false
              return
            case 2:
              explanation1.visible = false
              explanation2.visible = true
              return
            default: return
          }
        }
      }
    }
        Image {
          id: img
          source: imgPath
          fillMode: Image.PreserveAspectFit
          smooth: true
          anchors.fill: parent
        }
        GridLayout {
          id: explanation1
          anchors.fill: parent
          width: parent.width - 50
          rows: 3
          columns: 2
            Text {
              text: "Welcome, lets get you set up! PickleIt was built to try to stop us having loads of versions of a file, losing track of where you are, what you were doing and which version holds what."
              Layout.row: 0
              Layout.column: 0
              Layout.columnSpan: 1
              Layout.preferredHeight: 70
              Layout.fillWidth: true
              Layout.minimumWidth: 350
              Layout.maximumWidth: 500
              verticalAlignment: Text.AlignVCenter
              horizontalAlignment: Text.AlignHCenter
              wrapMode: Text.WordWrap; 
              font: Theme.font
            }
            Image {
              // id: img
              source: "qrc:/qml/assets/instruction/final-file-meme.png"
              fillMode: Image.PreserveAspectFit
              smooth: true
              Layout.row: 0
              Layout.column: 1
              Layout.columnSpan: 1
              Layout.preferredHeight: 200
              Layout.minimumWidth: 50
              Layout.rightMargin: 120
            }
            Text {
              text: "But first, if you see this popup, click Open System Preferences"
              Layout.row: 1
              Layout.column: 0
              Layout.columnSpan: 1
              Layout.preferredHeight: 70
              Layout.fillWidth: true
              Layout.minimumWidth: 350
              Layout.maximumWidth: parent.width - 60 
              verticalAlignment: Text.AlignVCenter
              horizontalAlignment: Text.AlignHCenter
              wrapMode: Text.WordWrap; 
              font: Theme.font
            }
          Image {
            source: "qrc:/qml/assets/instruction/acessibility-image.png"
            fillMode: Image.PreserveAspectFit
            smooth: true
            Layout.row: 1
            Layout.column: 1
            Layout.columnSpan: 1
            Layout.fillWidth: true
            Layout.preferredHeight: 100
            Layout.rightMargin: 120
          }
          Text {
            text: "and as shown in the animaton, click the padlock to allow changing settings, enter any password and tick the checkbox next to pickleIt. You may need to restart PickleIt to not see this message again, but from now everything is setup and you can enjoy your new, organised lifestyle."
            Layout.row: 2
            Layout.column: 0
            Layout.columnSpan: 1
            Layout.fillWidth: true
            Layout.preferredHeight: 100
            Layout.minimumWidth: 350
            Layout.maximumWidth: parent.width - 60 
            Layout.leftMargin: 12
            verticalAlignment: Text.AlignVCenter
            horizontalAlignment: Text.AlignHCenter
            wrapMode: Text.WordWrap; 
            font: Theme.font
          }

          AnimatedImage {
            source: "qrc:/qml/gifs/system-accessibility-osx.gif"
            fillMode: Image.PreserveAspectFit
            smooth: true
            Layout.row: 2
            Layout.column: 1
            Layout.columnSpan: 1
            Layout.fillWidth: true
            Layout.preferredHeight: 300
            Layout.rightMargin: 120
          }
        }
        GridLayout {
          id: explanation2
          visible: false
          anchors.fill: parent
          width: parent.width - 50
          rows: 3
          columns: 2
            Text {
              text: "Ok, so to use Pickle It, just drag a file onto the drop area like shown"
              Layout.row: 0
              Layout.column: 0
              Layout.columnSpan: 1
              Layout.preferredHeight: 70
              Layout.fillWidth: true
              Layout.minimumWidth: 350
              Layout.maximumWidth: 500
              verticalAlignment: Text.AlignVCenter
              horizontalAlignment: Text.AlignHCenter
              wrapMode: Text.WordWrap; 
              font: Theme.font
            }
            Image {
              // id: img
              source: "qrc:/qml/assets/instruction/dropzone.png"
              fillMode: Image.PreserveAspectFit
              smooth: true
              Layout.row: 0
              Layout.column: 1
              Layout.columnSpan: 1
              Layout.preferredHeight: 125
              Layout.minimumWidth: 50
              Layout.rightMargin: 120
            }
            Text {
              text: "Now all you need to do is get on with your work. Each time you save, a version will automatically be created for you"
              Layout.row: 1
              Layout.column: 0
              Layout.columnSpan: 1
              Layout.preferredHeight: 70
              Layout.fillWidth: true
              Layout.minimumWidth: 350
              Layout.maximumWidth: parent.width - 60 
              verticalAlignment: Text.AlignVCenter
              horizontalAlignment: Text.AlignHCenter
              wrapMode: Text.WordWrap; 
              font: Theme.font
            }
          Image {
            source: "qrc:/qml/assets/instruction/listview.png"
            fillMode: Image.PreserveAspectFit
            smooth: true
            Layout.row: 1
            Layout.column: 1
            Layout.columnSpan: 1
            Layout.fillWidth: true
            Layout.preferredHeight: 125
            Layout.rightMargin: 120
          }
          Text {
            text: "You can view the versions in either the list view or the calendar view. When you select one, you will see all the details about that version on the right hand side"
            Layout.row: 2
            Layout.column: 0
            Layout.columnSpan: 1
            Layout.fillWidth: true
            Layout.preferredHeight: 100
            Layout.minimumWidth: 350
            Layout.maximumWidth: parent.width - 60 
            Layout.leftMargin: 12
            verticalAlignment: Text.AlignVCenter
            horizontalAlignment: Text.AlignHCenter
            wrapMode: Text.WordWrap; 
            font: Theme.font
          }

          Image {
            source: "qrc:/qml/assets/instruction/calendarview.png"
            fillMode: Image.PreserveAspectFit
            smooth: true
            Layout.row: 2
            Layout.column: 1
            Layout.columnSpan: 1
            Layout.fillWidth: true
            Layout.preferredHeight: 175
            Layout.rightMargin: 120
          }
        }
}

  GridLayout {
    anchors.fill: parent
    rows: 2
    columns: 4
    File {
      
      Layout.row: 0
      Layout.column: 0
      Layout.columnSpan: 1
      Layout.fillWidth: true
      Layout.preferredHeight: 75
      Layout.minimumWidth: 650
      Layout.leftMargin: 12
    }
    Stack {
      Layout.row: 1
      Layout.column: 0
      Layout.rowSpan: 2
      Layout.columnSpan: 1
      Layout.minimumWidth: 650
      Layout.leftMargin: 12
      Layout.fillWidth: true
      Layout.fillHeight: true
    }

    ColumnLayout {
      Layout.row: 0
      Layout.column: 2
      Layout.rowSpan: 2
      Layout.columnSpan: 1
      Layout.maximumWidth: 500
      Layout.minimumWidth: 500
      Detail {
        Layout.fillWidth: true
        Layout.fillHeight: true
        Layout.topMargin: 5
        Layout.leftMargin: 12
        Layout.rightMargin: 6
        Layout.bottomMargin: 2
        Layout.minimumHeight: app.height / 4
      }
      Status {
          Layout.fillWidth: true
          Layout.fillHeight: true
          Layout.leftMargin: 12
          Layout.rightMargin: 6
      }
    }

  }
}
