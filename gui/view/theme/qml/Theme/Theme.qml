
pragma Singleton

import QtQuick 2.8

QtObject {
    readonly property color charcoal: "#36454f"
    readonly property color darkGray: "#878787"
    readonly property color gray: "#b2b1b1"
    readonly property color lightGray: "#dddddd"
    readonly property color white: "#ffffff"
    // readonly property color blue: "#2d548b"
    property color mainColor: "#C30694"
    readonly property color dark: "#6B0451"
    readonly property color mainColorDarker: "#A0057A"
    readonly property color mainColorLighter: "#E38DCE"

    property int baseSize: 10
    property real minWidth: 1024
    property real minHeight: 768

    readonly property int smallSize: 10
    readonly property int largeSize: 16

    property font font
    font.bold: false
    font.underline: false
    font.pixelSize: 15
    font.family: "Helvetica"
    // font.capitalization: Font.SmallCaps
}