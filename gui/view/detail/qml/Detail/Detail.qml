import QtQuick 2.7          //Image
import QtQuick.Controls 2.1 //GroupBox
import QtQuick.Layouts 1.3  //GridLayout

import Theme 1.0
import Detail 1.0
import ImageScroller 1.0
import "."
import "./moment.js" as M

DetailController {
  id: detailView
  property var parsedDateTime
  GroupBox {
    anchors.fill: parent
    // anchors.topMargin: 5
    background: Rectangle {
      color: Theme.white
      border.color: Theme.gray
      anchors.fill: parent
    }
    Component {
        id: eventListHeader
        Rectangle {
            Layout.fillWidth: true
            Layout.fillHeight: true
            Layout.margins: 40
          height: eventDayLabel.height + 20
        Row {
            id: eventDateRow
            width: parent.width
            height: eventDayLabel.height
            spacing: 30
            Label {
                id: eventDayLabel
                text: selectedDate ? selectedDate.getDate() : ""
                font.pointSize: 30
                color: "#aaa"
            }

            Column {
                height: eventDayLabel.height

                Label {
                    readonly property var options: { weekday: "long" }
                    text: selectedDate ? Qt.locale().standaloneDayName(selectedDate.getDay(), Locale.LongFormat) : "no event selected yet..."
                    font.pointSize: 18
                    color: "#aaa"
                }
                Label {
                    text: selectedDate ? Qt.locale().standaloneMonthName(selectedDate.getMonth()) + selectedDate.toLocaleDateString(Qt.locale(), " yyyy") : "pick a date or event for more details"
                    font.pointSize: 12
                    color: "#aaa"
                }
            }
        }
        }
    }
    // ListModel {
    //   id: detailEventModel
    // }
    ListView {
        id: eventsListView
        spacing: 2
        clip: true
        header: eventListHeader
        anchors.fill: parent
        model: detailModel//detailEventModel
        flickableDirection: Flickable.VerticalFlick
        boundsBehavior: Flickable.StopAtBounds
        ScrollBar.vertical: ScrollBar { 
          policy: ScrollBar.AlwaysOn
        }
        delegate: 
        // Rectangle {
        //   Text {
        //     text: {
        //       console.log("index", index, "model ", eventsListView.model.data(eventsListView.model.index(index, 0), Qt.DisplayRole))
        //       // console.log("patch: ", detailView.model.data(detailView.model.index(detailView.currentRow, 0), Qt.UserRole + 1), detailView.model.index(detailView.currentRow, 0), Qt.UserRole + 1)
        //       return display.description
        //     }
        //   }
        //         // DelegateButton {
        //         //   id: patchButton
        //         //   Layout.fillWidth: true
        //         //   source: "qrc:/qml/assets/patch.png"
        //         //   tooltipText: "apply patch"
        //         //   accent: Theme.dark
        //         //   onClicked: {
        //         //   patchFileShowRequest(eventsListView.model.index(index, 0))
        //         //   }
        //         // }
        // }
        DetailDelegate {
          internalModel: detailModel
          eView: eventsListView
        }
    }
  }
    // onClearDetails: {
    //   console.log("clearing")
    //   detailEventModel.clear()
    // }
    //  onDisplayDiffDetails: {
    //   var elipse = 20
    //   var elipsed = ""
    //   if (name.length > elipse) {
    //     elipsed = name.substring(name.length - elipse,name.length)
    //     elipsed = "..." + elipsed
    //   }
    //   console.log("displaying diff details for ", id)
    //   detailEventModel.append({patchId: id, tooltipText: name, diffName: elipsed, diffDescription: description, path: imagePath, diffTime: (selectedDate.getHours() < 10 ? '0' : '') + selectedDate.getHours() + ":" + (selectedDate.getMinutes() < 10 ? '0' : '') + selectedDate.getMinutes() + ":" + (selectedDate.getSeconds()  < 10 ? '0' : '') + selectedDate.getSeconds()})
    //   parsedDateTime = selectedDate

    // }
}
