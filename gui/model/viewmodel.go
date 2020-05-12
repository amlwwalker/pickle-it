package model

import (
	"github.com/amlwwalker/pickleit/logic"
	humanize "github.com/dustin/go-humanize"
	"github.com/therecipe/qt/core"
)

type ViewModel struct {
	*core.QAbstractTableModel
	manager *logic.Manager
}

func NewViewModel(manager *logic.Manager) *core.QAbstractItemModel {
	model := &ViewModel{core.NewQAbstractTableModel(nil), manager}

	model.ConnectHeaderData(model.headerData)
	model.ConnectRowCount(model.rowCount)
	model.ConnectColumnCount(model.columnCount)
	model.ConnectData(model.data)
	model.ConnectRoleNames(model.roleNames)

	return model.QAbstractItemModel_PTR()
}

func (m *ViewModel) headerData(section int, orientation core.Qt__Orientation, role int) *core.QVariant {
	if role != int(core.Qt__DisplayRole) && role < int(core.Qt__UserRole) || orientation != core.Qt__Horizontal {
		return core.NewQVariant()
	}
	switch section {
	case 0:
		return core.NewQVariant1("ID")
	case 1:
		return core.NewQVariant1("Subject")
	case 2:
		return core.NewQVariant1("Object")
	case 3:
		return core.NewQVariant1("SubjectHash")
	case 4:
		return core.NewQVariant1("ObjectHash")
	case 5:
		return core.NewQVariant1("Watching")
	case 6:
		return core.NewQVariant1("DiffPath")
	case 7:
		return core.NewQVariant1("Label")
	case 8:
		return core.NewQVariant1("Screenshot")
	case 9:
		return core.NewQVariant1("Fs")
	case 10:
		return core.NewQVariant1("Direction")
	case 11:
		return core.NewQVariant1("Description")
	case 12:
		return core.NewQVariant1("E")
	case 13:
		return core.NewQVariant1("DiffSize")
	case 14:
		return core.NewQVariant1("StartTime")
	case 15:
		return core.NewQVariant1("Message")
	}

	return core.NewQVariant()
}

func (m *ViewModel) rowCount(parent *core.QModelIndex) int {
	// fmt.Println("viewmodel getDiffCount ", getDiffCount())
	return m.manager.GetDiffCount()
}

func (m *ViewModel) columnCount(parent *core.QModelIndex) int {
	return 4 //	ID | Object | Name | DiffPath | Message
}

func (m *ViewModel) data(index *core.QModelIndex, role int) *core.QVariant {
	if role != int(core.Qt__DisplayRole) && role < int(core.Qt__UserRole) {
		return core.NewQVariant()
	}

	dbArray := getModelArray(m.manager)
	// fmt.Printf("dbArray %+v\r\n", dbArray)
	if index.Row() < 0 || index.Row() >= len(dbArray) {
		return core.NewQVariant()
	}

	dbItem := dbArray[index.Row()]

	switch {
	case
		index.Column() == 0 && role == int(core.Qt__DisplayRole) || //widgets
			role == int(core.Qt__UserRole)+1: //qml
		return core.NewQVariant1(dbItem.Diff.ID)
	case
		index.Column() == 1 && role == int(core.Qt__DisplayRole) || //widgets
			role == int(core.Qt__UserRole)+2: //qml
		return core.NewQVariant1(dbItem.Diff.Subject)
	case
		index.Column() == 2 && role == int(core.Qt__DisplayRole) || //widgets
			role == int(core.Qt__UserRole)+3: //qml
		return core.NewQVariant1(dbItem.Diff.Object)
	case
		index.Column() == 3 && role == int(core.Qt__DisplayRole) || //widgets
			role == int(core.Qt__UserRole)+4: //qml
		return core.NewQVariant1(dbItem.Diff.SubjectHash)
	case
		index.Column() == 4 && role == int(core.Qt__DisplayRole) || //widgets
			role == int(core.Qt__UserRole)+5: //qml
		return core.NewQVariant1(dbItem.Diff.ObjectHash)
	case
		index.Column() == 5 && role == int(core.Qt__DisplayRole) || //widgets
			role == int(core.Qt__UserRole)+6: //qml
		return core.NewQVariant1(dbItem.Diff.Watching)
	case
		index.Column() == 6 && role == int(core.Qt__DisplayRole) || //widgets
			role == int(core.Qt__UserRole)+7: //qml
		return core.NewQVariant1(dbItem.Diff.DiffPath)
	case
		index.Column() == 7 && role == int(core.Qt__DisplayRole) || //widgets
			role == int(core.Qt__UserRole)+8: //qml
		return core.NewQVariant1(dbItem.Diff.Label)
	case
		index.Column() == 8 && role == int(core.Qt__DisplayRole) || //widgets
			role == int(core.Qt__UserRole)+9: //qml
		return core.NewQVariant1(dbItem.Diff.Screenshot)
	case
		index.Column() == 9 && role == int(core.Qt__DisplayRole) || //widgets
			role == int(core.Qt__UserRole)+10: //qml
		return core.NewQVariant1(dbItem.Diff.Fs)
	case
		index.Column() == 10 && role == int(core.Qt__DisplayRole) || //widgets
			role == int(core.Qt__UserRole)+11: //qml
		return core.NewQVariant1(dbItem.Diff.Direction)
	case
		index.Column() == 11 && role == int(core.Qt__DisplayRole) || //widgets
			role == int(core.Qt__UserRole)+12: //qml
		return core.NewQVariant1(dbItem.Diff.Description)
	case
		index.Column() == 12 && role == int(core.Qt__DisplayRole) || //widgets
			role == int(core.Qt__UserRole)+13: //qml
		return core.NewQVariant1(dbItem.Diff.E)
	case
		index.Column() == 13 && role == int(core.Qt__DisplayRole) || //widgets
			role == int(core.Qt__UserRole)+14: //qml
		return core.NewQVariant1(humanize.Bytes(uint64(dbItem.Diff.DiffSize)))
	case
		index.Column() == 14 && role == int(core.Qt__DisplayRole) || //widgets
			role == int(core.Qt__UserRole)+15: //qml
		startDate := dbItem.Diff.StartTime
		st := core.NewQDateTime()
		st.SetDate(core.NewQDate3(startDate.Year(), int(startDate.Month()), startDate.Day()))
		st.SetTime(core.NewQTime3(startDate.Hour(), startDate.Minute(), startDate.Second(), 0))
		return core.NewQVariant1(st.ToString2(core.Qt__ISODate)) //time recorded when diff began
	case
		index.Column() == 15 && role == int(core.Qt__DisplayRole) || //widgets
			role == int(core.Qt__UserRole)+16: //qml
		return core.NewQVariant1(dbItem.Diff.Message)
	}

	return core.NewQVariant()
}

//needed only for qml
func (m *ViewModel) roleNames() map[int]*core.QByteArray {
	return map[int]*core.QByteArray{
		int(core.Qt__UserRole) + 1:  core.NewQByteArray2("ID", -1),
		int(core.Qt__UserRole) + 2:  core.NewQByteArray2("Subject", -1),
		int(core.Qt__UserRole) + 3:  core.NewQByteArray2("Object", -1),
		int(core.Qt__UserRole) + 4:  core.NewQByteArray2("SubjectHash", -1),
		int(core.Qt__UserRole) + 5:  core.NewQByteArray2("ObjectHash", -1),
		int(core.Qt__UserRole) + 6:  core.NewQByteArray2("Watching", -1),
		int(core.Qt__UserRole) + 7:  core.NewQByteArray2("DiffPath", -1),
		int(core.Qt__UserRole) + 8:  core.NewQByteArray2("Label", -1),
		int(core.Qt__UserRole) + 9:  core.NewQByteArray2("Screenshot", -1),
		int(core.Qt__UserRole) + 10: core.NewQByteArray2("Fs", -1),
		int(core.Qt__UserRole) + 11: core.NewQByteArray2("Direction", -1),
		int(core.Qt__UserRole) + 12: core.NewQByteArray2("Description", -1),
		int(core.Qt__UserRole) + 13: core.NewQByteArray2("E", -1),
		int(core.Qt__UserRole) + 14: core.NewQByteArray2("DiffSize", -1),
		int(core.Qt__UserRole) + 15: core.NewQByteArray2("StartTime", -1),
		int(core.Qt__UserRole) + 16: core.NewQByteArray2("Message", -1),
	}
}
