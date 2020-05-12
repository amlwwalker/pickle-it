import QtQuick 2.0
import QtQuick.Controls 2.0
import QtQuick.Window 2.2

import Theme 1.0
Window {
    id: splashWindow
    signal timeout()
    width: 400
    height: 200
    modality: Qt.ApplicationModal
    flags: Qt.SplashScreen
    color: Theme.white
    SplashIndicator {
        id: busyAnimation
        anchors.horizontalCenter: parent.horizontalCenter
        anchors.bottom: parent.bottom
        anchors.bottomMargin: parent.height / 5
        width: parent.width / 2
        height: width
        running: true
    }
    Text {
        text: "pickling..."
        anchors.centerIn: parent
        font.italic: true
        font.pixelSize: 15
        color: Theme.mainColorDarker
    }
    ProgressBar {
        id: progress
        anchors {
            left: parent.left
            right: parent.right
            bottom: parent.bottom
        }
        value: 0
        to : 100
        from : 0
    }
    Timer {
        id: timer
        interval: 20
        running: true
        repeat: true
        onTriggered: {
            progress.value++;
            if(progress.value >= 100) {
                timer.stop();
                splashWindow.timeout();
            }
        }
    }
}