import QtQuick 2.7          //Component
import QtQuick.Controls 1.4 //TableViewColumn
import QtQuick.Controls 2.1 //Label

import Theme 1.0            //Theme

TableViewColumn {

  title: role
  resizable: false
  movable: false
  width: parent.width * 0.25
  property var format
  // elideMode: Text.ElideLeft;
  delegate: Component { Item { 
    Label {
    anchors.centerIn: parent

    color: styleData.selected ? Theme.mainColorLighter : Theme.mainColor
    text: {
      const elipse = 20
      if (styleData.value !== "" && format !== undefined && format == "pretty") {
          let date = new Date(styleData.value)
          // console.log("date - >" + styleData.value + "< - ", date.getFullYear())
          date = date.getDate() + "/" + (date.getMonth() + 1) + "/" + date.getFullYear() + ", " + (date.getHours() < 10 ? '0' : '') + date.getHours() + ":" + (date.getMinutes() < 10 ? '0' : '') + date.getMinutes() + ":" + (date.getSeconds()  < 10 ? '0' : '') + date.getSeconds()
          return date
      }
      var text = styleData.value.substring(styleData.value.length - elipse,styleData.value.length)
      if (styleData.value.length > elipse) {
        text = "..." + text
      }
      return text
    }
    // font.bold: styleData.column == 2
    // font.italic: styleData.column == 3
    font: Theme.font
  } } }
  
}

// TableViewColumn { role: "Object"; title: "Object"; width: 200; elideMode: Text.ElideLeft; }
// TableViewColumn { role: "Subject"; title: "Subject"; width: 200; elideMode: Text.ElideLeft; }
// TableViewColumn { role: "SubjectHash"; title: "SubjectHash"; width: 100; visible: false }
// TableViewColumn { role: "ObjectHash"; title: "ObjectHash"; width: 100; visible: false }
// TableViewColumn { role: "Watching"; title: "Watching"; width: 200; elideMode: Text.ElideLeft; visible: false }
// TableViewColumn { role: "DiffPath"; title: "DiffPath"; width: 200; elideMode: Text.ElideLeft; visible: false }
// TableViewColumn { role: "Label"; title: "Label"; width: 100; visible: false }
// TableViewColumn { role: "Screenshot"; title: "Screenshot"; width: 100; visible: false }
// TableViewColumn { role: "Fs"; title: "Fs"; width: 100; visible: false }
// TableViewColumn { role: "Direction"; title: "Direction"; width: 100 ;visible: false }
// TableViewColumn { role: "Description"; title: "Description"; width: 100; visible: false }
// TableViewColumn { role: "E"; title: "E"; width: 100; visible: false }
// TableViewColumn { role: "DiffSize"; title: "DiffSize"; width: 100 }
// TableViewColumn { role: "StartTime"; title: "StartTime"; width: 100; visible: false }
// TableViewColumn { role: "Message"; title: "Message"; width: 200; elideMode: Text.ElideLeft; }