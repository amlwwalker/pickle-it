import QtQuick 2.7          //Rectangle
import QtQuick.Controls 2.1 //StackView
import QtQuick.Layouts 1.3  //GridLayout

// import Stack 1.0

import Theme 1.0
import Detail 1.0
import CalendarView 1.0

DetailController {
  Rectangle {
    anchors.fill: parent
    StackView {
      id: stackView
      anchors.fill: parent

      initialItem: detail
      Detail {
        id: detail
        visible: true
      }
      CalendarView { id: events
       visible: true }
    }
  }
  onPushNewStackCode: {
    console.log("clicked code ", code)
    var next = nextItem(code)
    if (next != null && next != stackView.currentItem) {
      stackView.currentItem.visible = false
      stackView.replace(next, StackView.Immediate)
      stackView.currentItem.visible = true
    }
  }

  function nextItem(code) {
    switch (code) {
    case "detail":
      return detail
    case "events":
      return events
    // case "hosting":
    //   return hosting
    // case "wallet":
    //   return wallet
    // case "terminal":
    //   return terminal
    default:
      console.log(code + " is not an option")
      return null
    }
  }
}
