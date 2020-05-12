import QtQuick 2.2
import QtQuick.Controls 2.5
import QtQuick.Controls.Styles 1.4
import QtQuick.Layouts 1.3
import QtCharts 2.0


//https://github.com/qyvlik/Chart.qml

ChartView {
  id: chartView
  property var data
    // width: parent.width
    // height: parent.height
    margins.top: 0
    margins.bottom: 0
    margins.left: 0
    margins.right: 0
    legend.visible: false
    anchors.fill: parent
    plotArea: Qt.rect(0, 0, parent.width, parent.height)
    ValueAxis {
      labelsVisible: false
      gridVisible:false
      minorGridVisible:false
      titleVisible: false
    }    
    antialiasing: true
    ToolTip {
        id: id_tooltip
        contentItem: Text{
            color: "hotpink"
            text: id_tooltip.text
        }
        background: Rectangle {
            border.color: "hotpink"
        }
    }
    PieSeries {
        id: pieSeries
        PieSlice { label: "eaten"; value: 94.9; color: "hotpink" }
        PieSlice { label: "not yet eaten"; value: 5.1; color: "pink" }
        onHovered: {
          if (!data.visibleMonth) return false
          console.log(slice.label, slice.value, state)
          if (state) {
            pieSeries.find(slice.label).exploded = true;
            // var p = chartView.mapToPosition(point)
            id_tooltip.text = slice.label
            // id_tooltip.x = p.x
            // id_tooltip.y = p.y
            id_tooltip.visible = true
          } else {
            pieSeries.find(slice.label).exploded = false;
            id_tooltip.visible = false
          }
        }
    }
}