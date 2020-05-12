import QtQuick 2.2
import QtQuick.Controls 1.4 as Old
import QtQuick.Controls.Styles 1.4
import QtQuick.Controls.Private 1.0
import QtQuick.Layouts 1.3
import QtCharts 2.0

import CalendarView 1.0 
import File 1.0
import Theme 1.0
import "."

CalendarController {
    id: calController
    // visible: false
    // property alias selectedDate: calendar.selectedDate
    selectedDate: calendar.selectedDate
    ColumnLayout {
        id: mainLayout
        anchors.fill: parent
        Old.Calendar {
            id: calendar
            Layout.fillWidth: true
            Layout.fillHeight: true
            frameVisible: true
            // selectedDate: new Date(2014, 0, 1)
            focus: true
            // onClicked: {
            //     // console.log("selectedDate -->: ", selectedDate)
            //     // console.log("styleData.date ", styleData.date)
            //     // console.log("styleData.selected ", styleData.selected)
            //     console.log("styleData.date.getDate() ", selectedDate)
            //     // var events = calController.eventsForDate(styleData.date)
            //     // for (let i = 0; i < events.length; i++) {
            //     //     console.log("date ", styleData.date)
            //     //     console.log("id ", events[i].id)
            //     //     console.log("name ", events[i].name)
            //     //     console.log("description ", events[i].description)
            //     //     console.log("startTime ", events[i].startDate)
            //     //     displayEventDetails(styleData.date)
            //     // }
            // }
            style: CalendarStyle {
                gridVisible: true
                dayOfWeekDelegate: Rectangle {
                    color: gridVisible ? "#fcfcfc" : "transparent"
                    implicitHeight: Math.round(TextSingleton.implicitHeight * 2.25)
                    Old.Label {
                        text: control.locale.dayName(styleData.dayOfWeek, control.dayOfWeekFormat)
                        anchors.centerIn: parent
                        color: Theme.mainColor
                    }
                }
                dayDelegate: Item {
                    Rectangle {
                        id: rect
                        anchors.fill: parent
                        anchors{
                            left: parent.left
                            right: parent.right
                            top: parent.top
                            bottom: parent.bottom
                        }
                        border.color: styleData.selected ? Theme.mainColor : "transparent"
                        // {
                        //     var events = calController.eventsForDate(styleData.date)
                        //     for (let i = 0; i < events.length; i++) {
                        //         console.log("[border.color] date ", styleData.date)
                        //         // console.log("id ", events[i].id)
                        //         console.log("[border.color] name ", events[i].name)
                        //         // console.log("description ", events[i].description)
                        //         // console.log("startTime ", events[i].startDate)
                        //         displayEventDetails(styleData.date)
                        //     }
                        //     return styleData.selected ? Theme.mainColor : "transparent"
                        // }
                        border.width: 2
                        // color: {
                        //     // console.log("selectedDate -->: ", selectedDate)
                        //     // console.log("styleData.date ", styleData.date)
                        //     // console.log("styleData.selected ", styleData.selected)
                        //     // return styleData.date == selectedDate ? "red" : "transparent"
                        // }
                        // color: (styleData.visibleMonth && styleData.valid && styleData.date !== undefined && styleData.selected) ? "white" : "transparent";
                        clip: true
                        Text {
                            text: calController.eventsForDate(styleData.date).length > 0 ? calController.eventsForDate(styleData.date).length: ""
                            font.pixelSize: 40
                            color: Theme.darkGray
                            // anchors.verticalCenter: parent.verticalCenter 
                            // anchors.horizontalCenter: parent.horizontalCenter 
                            verticalAlignment: Text.AlignVCenter
                            horizontalAlignment: Text.AlignHCenter
                            anchors.fill: parent
                            // anchors{
                            //     left: parent.left
                            //     right: parent.right
                            //     top: parent.top
                            //     bottom: parent.bottom
                            // }
                        }
                        // GraphView {
                        //     data: styleData
                        //     visible: calController.eventsForDate(styleData.date).length
                        // }
                        Image {
                            visible: calController.eventsForDate(styleData.date).length
                            anchors.top: parent.top
                            anchors.right: parent.right
                            // anchors.margins: -1
                            // width: 12
                            // height: width
                            source: "qrc:/qml/assets/eventindicator.png"
                        }
                        Old.Label {
                            id: dayDelegateText
                            text: styleData.date.getDate()
                            // anchors.centerIn: parent
                            x: 5
                            horizontalAlignment: Text.AlignHCenter
                            font.pixelSize: 20//Math.min(parent.height/3, parent.width/3)
                            color: styleData.visibleMonth ? (styleData.selected ? "hotpink" : "black") : Theme.gray
                            // font.bold: styleData.selected
                        }

                        // MouseArea {
                        //     anchors.horizontalCenter: parent.horizontalCenter
                        //     anchors.verticalCenter: parent.verticalCenter
                        //     width: parent.width//styleData.selected ? parent.width / 2 : 0
                        //     height: parent.height//styleData.selected ? parent.height / 2 : 0
                        //     Rectangle {
                        //         anchors.fill: parent
                        //         color: "transparent"
                        //     }
                        //     onClicked: {
                        //         console.log(styleData.date.getDate())
                        //         var events = calController.eventsForDate(styleData.date)
                        //         for (let i = 0; i < events.length; i++) {
                        //             console.log("date ", styleData.date)
                        //             console.log("id ", events[i].id)
                        //             console.log("name ", events[i].name)
                        //             console.log("description ", events[i].description)
                        //             console.log("startTime ", events[i].startDate)
                        //             displayEventDetails(styleData.date)
                        //         }
                        //         // color: {
                        //         // styleData.selected = true
                        //         console.log("selectedDate -->: ", selectedDate)
                        //         console.log("styleData.date ", styleData.date)
                        //         console.log("styleData.selected ", styleData.selected)
                        //     // return styleData.date == selectedDate ? "red" : "transparent"
                        //     // }
                        //     }
                        // }
                    }
                }
                navigationBar: Rectangle {
                height: Math.round(TextSingleton.implicitHeight * 2.73)
                color: "#f9f9f9"

                Rectangle {
                    color: Qt.rgba(1,1,1,0.6)
                    height: 1
                    width: parent.width
                }

                Rectangle {
                    anchors.bottom: parent.bottom
                    height: 1
                    width: parent.width
                    color: "#ddd"
                }

                HoverButton {
                    id: previousMonth
                    width: parent.height
                    height: width
                    anchors.verticalCenter: parent.verticalCenter
                    anchors.left: parent.left
                    source: "./assets/leftanglearrow.png"
                    onClicked: control.showPreviousMonth()
                }
                Old.Label {
                    id: dateText
                    text: styleData.title
                    elide: Text.ElideRight
                    horizontalAlignment: Text.AlignHCenter
                    font.pixelSize: TextSingleton.implicitHeight * 1.25
                    color: Theme.mainColor
                    anchors.verticalCenter: parent.verticalCenter
                    anchors.left: previousMonth.right
                    anchors.leftMargin: 2
                    anchors.right: nextMonth.left
                    anchors.rightMargin: 2
                }
                HoverButton {
                    id: nextMonth
                    width: parent.height
                    height: width
                    anchors.verticalCenter: parent.verticalCenter
                    anchors.right: parent.right
                    source: "./assets/rightanglearrow.png"
                    onClicked: control.showNextMonth()
                }
                }
            }
        }
    }
    onUpdateCalendarEvents: {
        //this is a hack to update the calendar view whenever the model is updated
        // there must be a better way
        console.log("refreshing calendar. Hack.")
        calendar.showNextMonth()
        calendar.showPreviousMonth()

    }
}