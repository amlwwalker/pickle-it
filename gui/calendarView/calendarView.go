// +build qml
package calendarview

import (
	"fmt"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/quick"

	"github.com/amlwwalker/pickleit/gui/controller"
	"github.com/amlwwalker/pickleit/gui/model"
)

func init() {
	calendarController_QmlRegisterType2("CalendarView", 1, 0, "CalendarController")
}

type calendarController struct {
	quick.QQuickItem

	_ func() `constructor:"init"`

	_ *core.QAbstractListModel `property:"listModel"`

	// events []*model.DiffDetail
	_ func(*core.QDate) []*model.DiffDetail `slot:"eventsForDate"`
	_ func(*core.QDate) []*model.DiffDetail `signal:"displayEventDetails"`

	_ *core.QDate `property:"selectedDate"` //move to model and controller!!!
	_ func()      `signal:"update"`
	_ func()      `signal:"updateCalendarEvents"`
}

// TODO: move as much of the controller's model information to the eventModel.go file (/models)
func (c *calendarController) init() {

	// c.SetListModel(core.NewQAbstractListModel(nil))
	c.SetListModel(controller.Instance().EventModel())
	c.SetSelectedDate(controller.Instance().SelectedDate()) //this needs to come from the event model who owns the date
	// c.ListModel().ConnectRowCount(func(*core.QModelIndex) int {
	// 	if c.SelectedDate() == nil {
	// 		return 0
	// 	}
	// 	return len(c.EventsForDate(c.SelectedDate()))
	// })

	// c.ListModel().ConnectData(func(index *core.QModelIndex, role int) *core.QVariant {
	// 	if c.SelectedDate() == nil || role != int(core.Qt__DisplayRole) {
	// 		return core.NewQVariant()
	// 	}
	// 	return c.EventsForDate(c.SelectedDate())[index.Row()].ToVariant()
	// })

	// c.ConnectSelectedDateChanged(func(*core.QDate) {
	// 	c.ListModel().BeginResetModel()
	// 	c.ListModel().EndResetModel()
	// })
	controller.Instance().ConnectUpdateCalendarEvents(c.UpdateCalendarEvents)
	c.ConnectUpdate(func() {
		fmt.Println("calendar update!")
		controller.Instance().EventUpdate()
	})

	c.ConnectDisplayEventDetails(func(date *core.QDate) {
		events := controller.Instance().EventsForDate(date)
		fmt.Printf("events for date %+v\r\n", len(events))
		controller.Instance().DisplayEventDetails(events)
	})

	c.ConnectEventsForDate(func(date *core.QDate) []*model.DiffDetail {
		events := controller.Instance().EventsForDate(date)
		return events
	})
	//this should be on the controller...
	c.ConnectSelectedDateChanged(func(d *core.QDate) {
		fmt.Println("[calendarview.go] Selected date changed ", d)
		// controller.EventModel().BeginResetModel()
		// controller.EventModel().EndResetModel()
		controller.Instance().SetSelectedDate(d)
	})

}

// func (c *calendarController) update() {
// 	c.ListModel().BeginResetModel()
// 	c.ListModel().EndResetModel()
// }
