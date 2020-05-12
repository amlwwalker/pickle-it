import QtQuick.Dialogs 1.2  //MessageDialog

import Dialog 1.0

PatchDialogController {
  property var patchIndex
  MessageDialog {
    id: patchDialog
    title: "Apply patch"
    text: ""
    icon: StandardIcon.Question
    standardButtons: StandardButton.Yes | StandardButton.No

    onYes: patchFile(patchIndex)
    // onNo:
  }

  onPatchFileShowRequest: {
    patchIndex = index
    console.log("requested to patch index ", patchIndex)
    patchDialog.text = "Are you sure you want to patch the file with this version?"//'"+diff+"' by '"+name+"'?"
    patchDialog.visible = true
  }
}
