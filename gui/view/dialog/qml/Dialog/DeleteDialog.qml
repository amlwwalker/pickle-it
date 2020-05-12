import QtQuick.Dialogs 1.2  //MessageDialog

import Dialog 1.0

DeleteDialogController {
  property var diffIndex
  MessageDialog {
    id: deleteDialog
    title: "Delete Diff"
    text: ""
    icon: StandardIcon.Question
    standardButtons: StandardButton.Yes | StandardButton.No

    onYes: deleteDiff(diffIndex)
    // onNo:
  }

  onDeleteDiffShowRequest: {
    diffIndex = index
    console.log("requested to delete index ", diffIndex)
    deleteDialog.text = "Are you sure you want to delete this patch?"//'"+diff+"' by '"+name+"'?"
    deleteDialog.visible = true
  }
}
