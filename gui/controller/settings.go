package controller

import (
	"fmt"

	"github.com/amlwwalker/pickleit/utilities"
)

//TODO: refactor settings to be part of the manager rather than the frontend.

func (c *Controller) initUXSettings() {
	//configure the settings either from the database, or from defaults
	c.SystemSettings = NewUXSettings(nil)

	c.SystemSettings.ConnectSomeSettingChanged(func() {
		fmt.Println("ConnectSomeSettingChanged",
			c.SystemSettings.IsAutoWatchCheck(),
			c.SystemSettings.IsScreenshotCheck(),
			c.SystemSettings.IsNewFileReadySystemNotifyCheck(),
			c.SystemSettings.IsPatchSystemNotifyCheck(),
			c.SystemSettings.IsOverrideExistingCheck(),
			c.SystemSettings.IsStatisticsCheck(),
			c.SystemSettings.IsWelcomeCheck(),
		)
		tmp := generateSettings(c.SystemSettings)

		if err := c.manager.StoreUXSettings(tmp); err != nil {
			fmt.Println("updating the settings failed ", err)
		}
		c.PushNotification("Settings Changed. Saving")
	})
	if settings, err := c.manager.RetrieveUXSettings(); err != nil {
		fmt.Println("retrieving the settings failed, setting defaults. Error: ", err)
		c.SystemSettings.SetAutoWatchCheck(true)
		// c.SystemSettings.SetScreenshotCheck(true) //removed screenshot check.
		c.SystemSettings.SetNewFileReadySystemNotifyCheck(false)
		c.SystemSettings.SetPatchSystemNotifyCheck(false)
		c.SystemSettings.SetOverrideExistingCheck(true)
		c.SystemSettings.SetStatisticsCheck(true)
		c.SystemSettings.SetWelcomeCheck(true)
	} else {
		fmt.Printf("settings from database %+v\r\n", settings)
		// tmp := UXSettings{} //NewUXSettings(nil)
		c.SystemSettings.SetAutoWatchCheck(settings.AutoWatch)
		// c.SystemSettings.SetScreenshotCheck(settings.Screenshot)
		c.SystemSettings.SetNewFileReadySystemNotifyCheck(settings.NewFileReadySystemNotify)
		c.SystemSettings.SetPatchSystemNotifyCheck(settings.PatchSystemNotify)
		c.SystemSettings.SetOverrideExistingCheck(settings.OverrideExisting)
		c.SystemSettings.SetStatisticsCheck(settings.Statistics)
		c.SystemSettings.SetWelcomeCheck(settings.Welcome)
		fmt.Printf("c.SystemSettings comes out as %+v", *c.SystemSettings)
	}
}
func generateSettings(userSettings *UXSettings) utilities.UXSettings {
	var tmp utilities.UXSettings

	tmp.AutoWatch = userSettings.IsAutoWatchCheck()
	tmp.Statistics = userSettings.IsStatisticsCheck()
	tmp.Screenshot = userSettings.IsScreenshotCheck()
	tmp.NewFileReadySystemNotify = userSettings.IsNewFileReadySystemNotifyCheck()
	tmp.PatchSystemNotify = userSettings.IsPatchSystemNotifyCheck()
	tmp.OverrideExisting = userSettings.IsOverrideExistingCheck()
	tmp.Welcome = userSettings.IsWelcomeCheck()
	return tmp
}

//check whether to show welcome screen
func (c *Controller) welcomeCheck() bool {
	if c.SystemSettings.IsWelcomeCheck() {
		fmt.Println("about to show exanplanation")
		c.ShowExplanation()
		c.SystemSettings.SetWelcomeCheck(false)
		c.SystemSettings.SomeSettingChanged()
		return true
	}
	return false
}

//check whether we should automatically start watching files
func (c *Controller) autoWatchCheck() bool {
	if c.SystemSettings.IsAutoWatchCheck() {
		c.PushNotification("Watching files automatically")
		c.ReloadWatchingRequest()
		c.BeginWatchingRequest()
		return true
	}
	return false
}
func (c *Controller) screenshotCheck() bool {
	if c.SystemSettings.IsScreenshotCheck() {
		//enable screenshots when creating diffs. Requires management to adhere to this
		// c.manager.EnableScreenshots
		return true
	}
	return false
}
func (c *Controller) notifyWhenFileReadyCheck() bool {
	if c.SystemSettings.IsNewFileReadySystemNotifyCheck() {
		// display a notification in the systray if a file is now ready for processing.
		// this should be called whenever something needs to know if it should show a notification
		// for new files
		return true
	}
	return false
}
func (c *Controller) notifyWhenNewPatchCheck() bool {
	if c.SystemSettings.IsPatchSystemNotifyCheck() {
		// display a notification in the systray if a patch has been made.
		// this should be called whenever something needs to know if it should show a notification
		// for new patches
		return true
	}
	return false
}

func (c *Controller) OverrideExistingCheck() bool {
	if c.SystemSettings.IsOverrideExistingCheck() {
		return true
	}
	return false
}

func (c *Controller) SendStatisticsCheck() bool {
	if c.SystemSettings.IsStatisticsCheck() {
		return true
	}
	return false
}
