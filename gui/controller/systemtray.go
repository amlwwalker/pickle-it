package controller

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)


func (t *QSystemTrayIconWithCustomSlot) init() {
	t.ConnectSetTextForAction(func(text string, action *widgets.QAction) { action.SetText(text) })

}

func NewTray() *Tray {

	var systray = NewQSystemTrayIconWithCustomSlot(nil)
	var systrayMenu = widgets.NewQMenu(nil) //https://doc.qt.io/qt-5/qmenu.html
	var icon *gui.QIcon
	icon = gui.NewQIcon()
	icon.AddPixmap(gui.NewQPixmap3(":/assets/icons/systemtray-grayscale.png", "", core.Qt__AutoColor), gui.QIcon__Normal, gui.QIcon__Off)
	systray.SetIcon(icon)
	systray.SetToolTip("Pickle It")
	systray.SetContextMenu(systrayMenu)
	t := &Tray{
		obj:      systray,
		rootmenu: systrayMenu,
	}

	return t
}

func (t *Tray) Build(showUI func(bool)) {

	t.addAction("Bring to front", showUI)
	t.quit = t.addAction("quit", func(bool) {
		log.Println("exit")
		os.Exit(0)
	})

	t.obj.Show()
}

func (t *Tray) ShowMessage(title, message string) {
	t.obj.ShowMessage(title, message, widgets.QSystemTrayIcon__Information, 2000)
}
func (t *Tray) addAction(str string, fn func(bool)) *widgets.QAction {
	a := widgets.NewQAction2(str, nil)
	a.ConnectTriggered(fn)
	t.rootmenu.AddActions([]*widgets.QAction{a})
	return a
}
func (t *Tray) update() {
	count := 0
	time.Sleep(2 * time.Second)
	for {
		time.Sleep(5 * time.Second)
		t.obj.SetTextForAction(fmt.Sprintf("updated %d", count), t.first)
		log.Println("Set", t.first.Text())
		count++
	}
}
