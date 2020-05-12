package model

import (
	"fmt"

	"github.com/amlwwalker/pickleit/logic"
	"github.com/therecipe/qt/core"
)

var ListModel *listModel

type listModel struct {
	core.QAbstractListModel
	manager *logic.Manager
	_       func() `signal:"update"`
}

func (m *listModel) InitWith(manager *logic.Manager) {
	ListModel = m
	ListModel.manager = manager
	ListModel.ConnectRowCount(ListModel.rowCount)
	ListModel.ConnectData(ListModel.data)
	ListModel.ConnectUpdate(ListModel.update)
}

func (m *listModel) update() {
	m.BeginResetModel()
	m.EndResetModel()
	fmt.Println("listmodel update called")
}
func (m *listModel) rowCount(parent *core.QModelIndex) int {
	return m.manager.GetFileCount()
}

func (m *listModel) data(index *core.QModelIndex, role int) *core.QVariant {
	if role != int(core.Qt__DisplayRole) {
		return core.NewQVariant()
	}
	fmt.Println("index.Row() ", index.Row())
	// if index.Row() == 0 {
	// 	return core.NewQVariant1("All Files")
	// 	// return core.NewQVariant1(m.manager.GetFileForID(index.Row()).Name)
	// }
	fmt.Println("manager returned ", m.manager.GetFileForID(index.Row()+1).Name)
	return core.NewQVariant1(m.manager.GetFileForID(index.Row() + 1).Name)
}
