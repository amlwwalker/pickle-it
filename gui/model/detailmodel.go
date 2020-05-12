package model

import (
	"fmt"

	"github.com/amlwwalker/pickleit/logic"
	"github.com/therecipe/qt/core"
)

var DetailModel *detailModel

type detailModel struct {
	core.QAbstractListModel
	manager *logic.Manager
	details []*DiffDetail

	_ func()                                `signal:"clear"`
	_ func(diffs []*DiffDetail, reset bool) `signal:"update"`
}

// type DetailObject struct {
// 	core.QObject

// 	_ int             `property:"id"`
// 	_ string          `property:"name"`
// 	_ string          `property:"description"`
// 	_ string          `property:"startTime"`
// 	_ string          `property:"screenshot"`
// 	_ *core.QDateTime `property:"startDate"`
// 	_ *core.QDateTime `property:"endDate"`
// }

func (m *detailModel) InitWith(manager *logic.Manager) {
	DetailModel = m
	DetailModel.manager = manager
	DetailModel.ConnectRowCount(DetailModel.rowCount)
	DetailModel.ConnectData(DetailModel.data)
	DetailModel.ConnectUpdate(DetailModel.update)
	DetailModel.ConnectClear(DetailModel.clear)
}

// func (m *detailModel) update() {
// 	m.BeginResetModel()
// 	m.EndResetModel()
// 	fmt.Println("detailModel update called")
// }
func (m *detailModel) rowCount(parent *core.QModelIndex) int {
	return len(m.details)
}

/*
so the controller needs to read the events from the calendar and pass them to the update method here. the model
*/
func (m *detailModel) data(index *core.QModelIndex, role int) *core.QVariant {
	//displayRole is the default, and sits at 0, while UserRole is the beginning of custom roles and is 256
	if role != int(core.Qt__DisplayRole) {
		fmt.Println("role is not equal to displayRole ", role)
		return core.NewQVariant()
	}

	//so this needs to collect the current 'details' and be updated accordingly
	//lets start with just the calendar as that will be reasonably straight forward.
	v := m.DetailFromIndex(index)
	return core.NewQVariant1(v)
}

//needed only for qml
func (m *detailModel) roleNames() map[int]*core.QByteArray {
	return map[int]*core.QByteArray{
		int(core.Qt__UserRole) + 1: core.NewQByteArray2("ID", -1),
		int(core.Qt__UserRole) + 2: core.NewQByteArray2("Name", -1),
		int(core.Qt__UserRole) + 3: core.NewQByteArray2("Description", -1),
		int(core.Qt__UserRole) + 4: core.NewQByteArray2("StartTime", -1),
		int(core.Qt__UserRole) + 5: core.NewQByteArray2("Screenshot", -1),
		int(core.Qt__UserRole) + 6: core.NewQByteArray2("StartDate", -1),
		int(core.Qt__UserRole) + 7: core.NewQByteArray2("EndDate", -1),
	}
}

func (m *detailModel) DetailFromIndex(index *core.QModelIndex) *DiffDetail {
	row := index.Row()
	for i, v := range m.details {
		//as details is an array, the indexRow, will equal the index of the current array.
		//on that basis the index.Row will pull out the correct diffDetail by matching index
		// fmt.Printf("vId: %+v -- index %d\r\n", v.Id(), index.Row())
		//THEN you will know what diffDetail was clicked on...
		if i == row { //does it need to be +1?
			// fmt.Printf("found it %+v\r\n", v.Screenshot())
			return v
		}
	}
	return &DiffDetail{}
}

func (m *detailModel) clear() {
	fmt.Println("clering all details")
	m.Update([]*DiffDetail{}, true)
}
func (m *detailModel) update(diffs []*DiffDetail, reset bool) {
	m.BeginResetModel()
	if reset {
		fmt.Println("resetting")
		m.details = diffs
		fmt.Printf("details %+v\r\n", m.details)
	} else {
		m.details = append(m.details, diffs...)
	}
	m.EndResetModel()
}
