import QtQuick 2.7          //Image
import QtQuick.Controls 2.1 //GroupBox
import QtQuick.Layouts 1.3  //GridLayout

import Theme 1.0
import "." as Custom              //needed for name clash with std Controls.Button
GroupBox {
    property var internalModel
    property ListView eView
    readonly property var detail: display
            width: parent.width - 10
            height: gLayout.height + 10
            Layout.fillHeight: true  
          GridLayout {
              id: gLayout
              Layout.fillWidth: true
              rows: 3
              columns: 4
              flow: GridLayout.TopToBottom

              Label { 
                text: "Name"
                font.family: "Helvetica"
                font.pixelSize: 12
                Layout.row: 0
                Layout.column: 0
                color: Theme.mainColor
                wrapMode: Text.WordWrap
              }
              Label { 
                text: "Description"
                font.family: "Helvetica"
                font.pixelSize: 12                
                Layout.row: 1
                Layout.column: 0
                color: Theme.mainColor
                wrapMode: Text.WordWrap
              }
            Label {
              text: "Created"
              font.family: "Helvetica"
              font.pixelSize: 12
              Layout.row: 2
              Layout.column: 0
              color: Theme.mainColor
              wrapMode: Text.WordWrap
            }

              Label { 
                text: {
                  var elipse = 20
                  var elipsed = ""
                  if (display.name.length > elipse) {
                    elipsed = display.name.substring(display.name.length - elipse,display.name.length)
                    elipsed = "..." + elipsed
                  }
                  return elipsed
                }
                Layout.row: 0
                Layout.column: 1
                MouseArea {
                    id: tooltipMa
                    anchors.fill: parent
                    hoverEnabled: true
                }
                ToolTip {
                  visible: display.name ? tooltipMa.containsMouse : false
                  delay: 500
                  contentItem: Text {
                      color: Theme.mainColorDarker
                      text: display.name
                  }
                  background: Rectangle {
                      border.color: Theme.mainColorDarker
                  }
                }
              }
              TextField {
                id: setDescription
                placeholderText: "no description yet...."
                text: display.description //"description" //diffDescription
                font.family: "Helvetica"
                font.pixelSize: 15
                Layout.row: 1
                Layout.column: 1
                Layout.columnSpan: 2
                background: Rectangle {
                    radius: 2
                    border.color: Theme.gray
                    border.width: 1
                    height: setDescription.height
                }
                color: Theme.darkGray
                wrapMode: Text.Wrap
                Layout.fillWidth: true
                onEditingFinished: {
                  //why doesn't this work?
                  // detailEventModel.index(detailEventModel.rowCount - 1, 0)
                  // i think its because the index the other end (in controller.go, doesn't have an understanding of the mapping of the Qt index to the field index, like the sortfilterview model does)
                  //so for now, hacking it by passing the index as an integer
                  console.log("[detail] patch ", display.id)
                  requestSaveDescription(display.id, setDescription.text)
                }
              }
            Text {
              id: titleLabel
              text: {
                console.log("display.startDate: ", display.startDate, " id ", display.id)
                return (display.startDate.getHours() < 10 ? '0' : '') + display.startDate.getHours() + ":" + (display.startDate.getMinutes() < 10 ? '0' : '') + display.startDate.getMinutes() + ":" + (display.startDate.getSeconds()  < 10 ? '0' : '') + display.startDate.getSeconds()
                
              }
              font.family: "Helvetica"
              font.pixelSize: 15
              color: Theme.gray
              Layout.row: 2
              Layout.column: 1
              wrapMode: Text.Wrap
            }
              Image {
                  Layout.row: 0
                  Layout.column: 3
                  Layout.rowSpan: 3
                  id: img
                  source: "file://"+display.screenshot
                  Layout.margins: 5
                  Layout.preferredWidth: 100
                  Layout.preferredHeight: 100
                  fillMode: Image.PreserveAspectFit
                  smooth: true
                  MouseArea {
                      id: imageMa
                      anchors.fill: parent
                      onClicked: {
                        console.log("image path ", "file://"+display.screenshot)
                        requestExpandImage("file://"+display.screenshot)
                      }
                      cursorShape: Qt.PointingHandCursor
                  }
              }
              RowLayout {
                Layout.row: 1
                Layout.column: 4
                Layout.columnSpan: 2
                Layout.fillWidth: true
                Custom.DelegateButton {
                  id: deleteButton
                  Layout.fillWidth: true
                  source: "qrc:/qml/assets/ic_delete_forever_black_24px.svg"
                  tooltipText: "delete patch"
                  accent: Theme.gray
                onClicked: {
                  console.log("delete: ", detail.name)
                  // console.log("detail ", internalModel.currentRow, internalModel.index(internalModel.currentRow, 0))
                  // console.log("delete: ", detailModel.index(detailModel.currentRow, 0))
                  deleteDiffShowRequest(detail.id)
                }
                }
                Custom.DelegateButton {
                  id: holder
                  Layout.fillWidth: true
                }
                Custom.DelegateButton {
                  id: patchButton
                  Layout.fillWidth: true
                  source: "qrc:/qml/assets/patch.png"
                  tooltipText: "apply patch"
                  accent: Theme.dark
                  onClicked: {

                  console.log("patch: ", detail.name)
                  // template.deleteRequest(eventsListView.model.data(eventsListView.model.index(eventsListView.model.currentRow, 0), Qt.UserRole + 1))
                  patchFileShowRequest(detail.id)
                  }
                }
              }
            }   
        }