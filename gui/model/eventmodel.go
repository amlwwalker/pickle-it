// +build qml
package model

import (
	"fmt"
	// "time"
	"strings"

	"github.com/amlwwalker/pickleit/logic"
	"github.com/therecipe/qt/core"
)

var EventModel *eventModel

type eventModel struct {
	core.QAbstractListModel
	manager *logic.Manager
	_       *core.QDate `property:"selectedDate"`
	_       string      `property:"filterExpression"`
	_       func()      `signal:"update"`
}

func (e *eventModel) InitWith(manager *logic.Manager) {
	EventModel = e
	EventModel.manager = manager
	EventModel.ConnectRowCount(EventModel.rowCount)
	EventModel.ConnectData(EventModel.data)
	EventModel.ConnectUpdate(EventModel.update)
	EventModel.SetFilterExpression("")
}

func (e *eventModel) update() {
	e.BeginResetModel()
	e.EndResetModel()
}

func (e *eventModel) rowCount(*core.QModelIndex) int {
	if e.SelectedDate() == nil {
		return 0
	}
	fmt.Println("row count selectedDate ", e.SelectedDate())
	return len(e.EventsForDate(e.SelectedDate()))
}

func (e *eventModel) data(index *core.QModelIndex, role int) *core.QVariant {
	fmt.Printf("role %d; index %+v\r\n", role, index)
	if e.SelectedDate() == nil {
		fmt.Println("attempting to convert, but received selected Date is nil???")
		return core.NewQVariant()
	}
	if role != int(core.Qt__DisplayRole) && role < int(core.Qt__UserRole) {
		fmt.Printf("attempting to convert, but role is %+v\r\n", role)
		return core.NewQVariant()
	}
	fmt.Println("index row ", index.Row())
	fmt.Println("data selected ", e.EventsForDate(e.SelectedDate())[index.Row()])
	return e.EventsForDate(e.SelectedDate())[index.Row()].ToVariant()
}

func (e *eventModel) EventsForDate(d *core.QDate) (o []*DiffDetail) {

	diffs, _ := e.manager.RetrieveAllDiffs() //retrieve all the diffs, now compare the date elements

	//so now this will filter the events based on the date passed in
	for _, diff := range diffs {
		startDate := diff.StartTime
		//e is a QDateTime, need to get the Date() object to get the Year/Month/Day
		if strings.Contains(diff.Object, e.FilterExpression()) && (startDate.Year() == d.Year() && int(startDate.Month()) == d.Month() && startDate.Day() == d.Day()) {
			//now create a QDateTime object to append
			ev := NewDiffDetail(nil)
			// ev.SetName(fmt.Sprintf("event (%v) on the %v/%v/%v", diff.Watching, startDate.Day(), startDate.Month(), startDate.Year()))
			ev.SetName(diff.Object)
			ev.SetId(diff.ID)
			ev.SetScreenshot(diff.Screenshot)
			ev.SetDescription(diff.Description)
			// ev.SetStartTime(diff.StartTime.String())
			st := core.NewQDateTime()
			st.SetDate(core.NewQDate3(startDate.Year(), int(startDate.Month()), startDate.Day()))
			st.SetTime(core.NewQTime3(startDate.Hour(), startDate.Minute(), startDate.Second(), 0))
			ev.SetStartDate(st)

			et := core.NewQDateTime()
			et.SetDate(core.NewQDate3(startDate.Year(), int(startDate.Month()), startDate.Day()))
			et.SetTime(core.NewQTime3(startDate.Hour(), startDate.Minute(), startDate.Second(), 0))
			ev.SetEndDate(et)
			o = append(o, ev)
		}
	}
	return o
}
