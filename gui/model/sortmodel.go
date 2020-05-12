package model

import (
	"github.com/amlwwalker/pickleit/logic"
	"github.com/therecipe/qt/core"
)

var SortFilterModel *sortFilterModel

type sortFilterModel struct {
	core.QSortFilterProxyModel
	viewModel *core.QAbstractItemModel
	_         func() `signal:"update"`
}

func (m *sortFilterModel) InitWith(manager *logic.Manager) {
	SortFilterModel = m
	SortFilterModel.viewModel = NewViewModel(manager)
	SortFilterModel.SetSourceModel(SortFilterModel.viewModel)
	SortFilterModel.ConnectUpdate(SortFilterModel.update)
}

func (m *sortFilterModel) update() {
	m.viewModel.BeginResetModel()
	m.viewModel.EndResetModel()
}
