import QtQuick 2.4
import QtQuick.Controls 1.4 //TableView

import Theme 1.0
Rectangle {
    id: rowDel
    property var internalModel
    property TableView tView
    readonly property int modelRow: styleData.row ? styleData.row : 0
    color: styleData.selected ? Theme.dark : modelRow % 2 == 0 ? Theme.white : Theme.gray
    height: 40
    
    Menu { id: contextMenu
        MenuItem {
          text: qsTr('Patch this diff')
          onTriggered: {
            console.log("item id " + modelRow)
            console.log(styleData)
            patchDiff(tableView.model.index(styleData.row, 0))
          }
        }
        MenuItem {
          text: qsTr('Delete this diff')
          onTriggered: {
            console.log("item id " + modelRow)
            console.log(styleData)
            deleteDiff(tableView.model.index(styleData.row, 0))
          }
        }
    }
    MouseArea {
        acceptedButtons: Qt.LeftButton | Qt.RightButton
        anchors.fill: parent
        onClicked: {
            tView.selection.clear()
            tView.selection.select(modelRow)
            // showDiffDetails(internalModel.index(modelRow, 0))
            if (mouse.button == Qt.LeftButton)
            {
                console.log("[LEFT] log: " + modelRow, tableView.model.data(tableView.model.index(styleData.row, 0)));
                showDiffDetails(tableView.model.index(styleData.row, 0))
            } else if (mouse.button == Qt.RightButton) {
                console.log("[RIGHT] log: " + modelRow, tableView.model.data(tableView.model.index(styleData.row, 0)));
                contextMenu.popup()
            }
        }
    }       
}