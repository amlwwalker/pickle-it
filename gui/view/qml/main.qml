import QtQuick 2.11
import QtQuick.Window 2.2
import QtQuick.Controls 2.2

import "."
View {}
// Loader {
//     id: loader
//     Component {
//         id: splash
//         Splashscreen {}
//     }

//     Component {
//         id: root
//         View {}
//     }

//     sourceComponent: splash
//     active: true
//     visible: true
//     onStatusChanged: {
//         if (loader.status === Loader.Ready)
//             item.show();
//     }

//     Connections {
//         id: connection
//         target: loader.item
//         onTimeout: {
//             connection.target = null;
//             loader.sourceComponent = root;
//         }
//     }
// }