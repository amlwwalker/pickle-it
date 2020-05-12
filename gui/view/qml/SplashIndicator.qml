import QtQuick 2.13
import QtQuick.Controls 2.3

import Theme 1.0

BusyIndicator {
    id: control

    contentItem: Item {
        implicitWidth: 128
        implicitHeight: 128

        Item {
            id: item
            x: parent.width / 2 - 64
            y: parent.height / 2 - 24// - (128 + 32)
            width: 128
            height: 128
            opacity: control.running ? 1 : 0

            Behavior on opacity {
                OpacityAnimator {
                    duration: 250
                }
            }

            RotationAnimator {
                target: item
                running: control.visible && control.running
                from: 0
                to: 360
                loops: Animation.Infinite
                duration: 3000
            }

            Repeater {
                id: repeater
                model: 8

                Rectangle {
                    x: item.width / 2 - width / 2
                    y: item.height / 2 - height / 2
                    implicitWidth: 16
                    implicitHeight: 16
                    radius: 8
                    color: Theme.mainColor
                    transform: [
                        Translate {
                            y: -Math.min(item.width, item.height) * 0.5 + 5
                        },
                        Rotation {
                            angle: index / repeater.count * 360
                            origin.x: 8
                            origin.y: 8
                        }
                    ]
                }
            }
        }
    }
}