package controller

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

type Tray struct {
	obj      *QSystemTrayIconWithCustomSlot
	rootmenu *widgets.QMenu
	first    *widgets.QAction
	quit     *widgets.QAction
}

type QSystemTrayIconWithCustomSlot struct {
	widgets.QSystemTrayIcon
	_ func() `slot:"triggerSlot"`

	_ func()                                     `constructor:"init"`
	_ func(text string, action *widgets.QAction) `signal:"setTextForAction"`
}

type UXSettings struct {
	core.QObject

	_ bool   `property:"autoWatchCheck"`
	_ bool   `property:"statisticsCheck"`
	_ bool   `property:"screenshotCheck"`
	_ bool   `property:"newFileReadySystemNotifyCheck"`
	_ bool   `property:"patchSystemNotifyCheck"`
	_ bool   `property:"overrideExistingCheck"`
	_ bool   `property:"welcomeCheck"`
	_ func() `signal:"someSettingChanged"`
}

type UXVersion struct {
	core.QObject

	_ string `property:"tag"`
	_ string `property:"flavour"`
	_ string `property:"version"`
	_ string `property:"hash"`
	_ string `property:"date"`
}
