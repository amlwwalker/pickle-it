# Gui Todo list

- [x] Start/stop watching
- [x] Apply a patch - easy
- [x] Delete a patch
- [x] Set the label on a diff. Need front end changes
- [x] load the index on another thread. Locks up the UI currently
- [x] Add an status if each file (has the base index been created for instance). - On this note, indexes for large files should happen rarely as they take a long time. They should also be created in the background. A large file won't be watched until an index has been created - this will need checking - basically, copy the file, create the index on that file, then compare to the original location - same hash? good. Different hash - make a diff. When the index is ready, the file can be monitored, but not before that.
  - is a quick fix just setting a ready parameter against the file and only displaying it once done? consider progress bar here...
- [] Add progress bars for patch manufacture to the front end. Progress bar in place, need to use from backend
- [] Add progress bar for index manufacture to front end. Progress bar in place, just need to use
- [x] Ask if the user is sure before deleting/patching/uploading etc -> its too dangerous right now

## Patching

- [] If you restore a file to its original, it detects the change. Should have a flag against a file, that even if files are watched, to ignore changes (until reset)

## UX:


- [x] (create a toast manager for the backend). Need to wire up. Notifications can go here. Some notifications can popup on UI if UI not in focus
- [x] Need user settings (default theme/screen etc)
  - [] Add preferences/advanced/notifications area. What notifications to show and telemetry etc. Settings page for now?
  - [-] standard naming convention for files
    - paused for this version
  - [x] user information (which user created this file)
- [-] store user preferences - needs to be stored in the boltdb
  - version format
- [x] store screenshot in database against the diff
- [-] Pass errors to toast/notifier
- [-] Back end informs front end of everything (at the moment front end guesses its suceeded) (via progress bars or toasts)

### Calendar

- [x] for the time being remove the pie chart and add count of diffs on that day
- [x] clear the details when another day is selected that has no elements
- [x] calendar not yet updating properly when new diffs/files are added/form.
- [x] selecting a date doesn't highlight it
  - Seems to be becuase of the SelectedDate property not updating properly at the moment

### detail

- [.] add options/click tools to the items (right click menu)
  - upgraded to not use right click menu, just action buttons
- [-] make sure the detail updates when diffs are created for a selected date
  - currently there is the ability to clear, but not update. Would need to be reloaded and we would have to keep track of 
  - what last updated it. One approach would be to have two arrays, one is the last update and one is the current state
  - if you need to do an update of the last, you update the buffer, otherwise you update both

### file

- [] scrollable file list

## backend


- [x] add close window not app feature
- [x] add system tray - need to implement it, but code is in place
- [x] add notifications. See above comment on toast etc

## on load

- [] on first load (need count of loads), show a pop up explanation gif - use similar approach to big image. Perhaps a setting to re show it from settings page.


## telemetry

- [] send count of loads on startup, fail quietly
- [] send count of patch creation, patching, deletion

## optimisations

- [] compile dyanmically for windows and mac - add installers
- [x] fix compression of index array (just dont compress perhaps)
- [] remove the backward diff code
- [] add auto update feature
- [] add a logging window for alpha/beta testing
- [x] stress test with large files. - issue is creating the index takes time. Need to process in the background. This can all be done on a thread however, and inform the front end via channels
- [] When a file being monitored is deleted, we get the error that its detected. Should it be removed from the database, or should the database be able to retrieve it and replace it if the user still wants?



## Wishlist

- [] Add the status of a diff, for instance which one is the currently applied one, and which ones are based images <-- this is tricky. Put this later down
