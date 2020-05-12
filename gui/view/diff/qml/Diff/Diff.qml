import QtQuick.Controls 1.4 //TableView
import QtQuick.Controls.Styles 1.4
import QtQuick 2.2

import "."
import Diff 1.0

import Theme 1.0
DiffController {
    TableView {
      id: tableView

      anchors.fill: parent
      horizontalScrollBarPolicy: Qt.ScrollBarAlwaysOff
      backgroundVisible: false
      sortIndicatorVisible: true
      onSortIndicatorColumnChanged: sortTableView(sortIndicatorColumn, sortIndicatorOrder)
      onSortIndicatorOrderChanged: sortTableView(sortIndicatorColumn, sortIndicatorOrder)

      onActivated: showDiffDetails(model.index(row, 0))
      model: viewModel
      headerDelegate: HeaderDelegate {
      }
      rowDelegate: RowDelegate {
        internalModel: viewModel
        tView: tableView
      } 
      TableColumnDelegate { role: "Object"; title: "File"; }
      TableColumnDelegate { role: "StartTime"; title: "Created"; format: "pretty"}
      TableColumnDelegate { role: "DiffSize"; title: "Patch Size"; }
      TableColumnActionDelegate {
        role: "ACTIONS"
        currentPatch: true
        tableView: tableView
      }
      
    }

  onDiffAdded: {
    tableView.selection.clear()
    tableView.selection.select(tableView.rowCount - 1)
    showDiffDetails(tableView.model.index(tableView.rowCount - 1, 0))
  }

  onDeleteDiffRequest: {
    var index = tableView.model.index(tableView.currentRow, 0)
    var title = tableView.model.data(index, Qt.UserRole + 2)
    var name = tableView.model.data(index, Qt.UserRole + 3)
    deleteDiffShowRequest(title, name)
  }

  onDeleteDiffCommand: {
    deleteDiff(tableView.model.index(tableView.currentRow, 0))
    showImageLabel()
  }
}
