import QtQuick 2.7          //Rectangle
import QtQuick.Controls 2.1 //StackView
import QtQuick.Layouts 1.3  //GridLayout

import File 1.0
import Diff 1.0
import Settings 1.0
import Stack 1.0
import CalendarView 1.0
import Theme 1.0

StackController {
  Rectangle {
    anchors.fill: parent
    TabBar {
      id: bar
      Layout.fillWidth: true
      Layout.fillHeight: true      
      width: parent.width
      background: Rectangle {
          color: "transparent"
      }
      TabButton {
          text: qsTr("Calendar")
      }
      TabButton {
          text: qsTr("Listing")
      }
      // TabButton {
      //     text: qsTr("Share/Sync")
      // }
      // TabButton {
      //     text: qsTr("Plugins")
      // }
      TabButton {
          text: qsTr("Settings")
      }
  }
  StackLayout {
    width: parent.width
    implicitHeight: parent.height - bar.height
    anchors.top: bar.bottom
    
    currentIndex: bar.currentIndex
    Item {
        id: calendarTab
        CalendarView { 
          id: calendar 
          width: parent.width
          height: parent.height
        }
    }
    Item {
      id: listingTab
        Diff {
          id: listing
          width: parent.width
          height: parent.height
        }
    }
    // Item {
    //     id: shareTab
    //     Column {
    //       anchors.margins: 20
    //       anchors.fill: parent
    //         spacing: 15
    //         Label {
    //           text: "Back up and share your pickles"
    //           font.pixelSize: 24
    //           color: Theme.mainColor
    //           font.family: "Helvetica"
    //         }
    //         Label {
    //           text: "Military grade security for all your work"
    //           font.pixelSize: 16
    //           color: Theme.mainColor
    //           font.family: "Helvetica"
    //         }
    //         Label {
    //           text: "Send pickles to colleagues and collaborators"
    //           font.pixelSize: 16
    //           color: Theme.mainColor
    //           font.family: "Helvetica"
    //         }
    //         Label {
    //           text: "Backup all your versions to the cloud. Never lose a thing"
    //           font.pixelSize: 16
    //           color: Theme.mainColor
    //           font.family: "Helvetica"
    //         }
    //         Label {
    //           text: "Just send a colleague a link and hey presto. They have your pickle."
    //           font.pixelSize: 16
    //           color: Theme.mainColor
    //           font.family: "Helvetica"
    //         }
    //         Label {
    //           text: "This feature is in the pipe line. Please watch this space"
    //           font.pixelSize: 14
    //           color: Theme.mainColor
    //           font.family: "Helvetica"
    //           wrapMode: Text.WordWrap
    //         }
    //     }
    // }
    // Item {
    //     id: pluginsTab
    //     Column {
    //       anchors.margins: 20
    //       anchors.fill: parent
    //         spacing: 15
    //         Label {
    //           text: "Anyone can build plugins for PickleIt"
    //           font.pixelSize: 24
    //           color: Theme.mainColor
    //           font.family: "Helvetica"
    //         }
    //         Label {
    //           text: "If there is a way you would like to visualise your work, you can!"
    //           font.pixelSize: 16
    //           color: Theme.mainColor
    //           font.family: "Helvetica"
    //         }
    //         Label {
    //           text: "Want to see and share how frequently you make changes?"
    //           font.pixelSize: 16
    //           color: Theme.mainColor
    //           font.family: "Helvetica"
    //         }
    //         Label {
    //           text: "Visualise your work on a heatmap perhaps?"
    //           font.pixelSize: 16
    //           color: Theme.mainColor
    //           font.family: "Helvetica"
    //         }
    //         Label {
    //           text: "See who is the main contributor to a piece of work?"
    //           font.pixelSize: 16
    //           color: Theme.mainColor
    //           font.family: "Helvetica"
    //         }
    //         Label {
    //           text: "This feature is in the pipe line. Please watch this space"
    //           font.pixelSize: 14
    //           color: Theme.mainColor
    //           font.family: "Helvetica"
    //           wrapMode: Text.WordWrap
    //         }
    //     }
    // }
    Item {
        id: settingsTab
        // Column {
          anchors.margins: 20
        //   anchors.fill: parent
        Settings { 
          id: settings 
          width: parent.width
          height: parent.height
        }
          // Row {
          //   spacing: 20
          //   Label {
          //     text: "Settings to come"
          //   }
          //   Label {
          //     text: "Everything set to default for time being"
          //   }
          // }
        // }
    }
  }
  }
}
