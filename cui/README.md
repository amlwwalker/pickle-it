# Workhorse Cui

The Cui is the command line user interface - a text based user interface. Mainly this is used to make sure that communciation between the different elements of the Workhorse framework are designed and built to communicate with an interface before a graphical user interface is developed, however this interface will be more of a 'nightly' build of Workhorse and therefore the UI here may have capabilities before the graphical user interface does. It does obviously require understanding of running a terminal application

## To do

- [x] List all the currently watched files (from database)
- [x] Add a file to the watcher from the cui
- [x] List all the patches that are in the database for a file, both forward and backward. (is forward and backward patching strictly necessary? Don't know at this stage)
- [x] Patch a file with a selected patch
  - [x] Forward patch works
    - [] Work out if we want bi-directional patching anyway. Would cause a refactor.
- [x] When Workhorse Cui is told to start watching, it should automatically watch any file that is in the database
- [x] Make sure that the UI is reloading automatically, rather than manually (selecting something else)
- [-] Update the Cui with new patches as they get created in real time
  - [x] This actually is less necessary in the Cui app. It sounds like a cop out, however the UI needs to have the correct file selected, to show the patches. This would create a slightly overcomplex flow. More relevant in the gui, rather than the cui.
- [-] Update the Cui with information about how the patch is doing (completed/running etc)
- [] Recover the original base file. Currently this is not actually possible
  - [] When a new patch is made and a new base is created the patch will produce that base, so that is straight forward to return to
    - [] However there is no way to restore the original back up
    - [] Also is there need to store both the patch and the base if a new base is created, and a base is recoverable? Perhaps base files should be listed aswell and then they are recoverable, and they can be considered as commits?
- [-] make sure cancelling diffing works properly
  - [-] involves going into the sort method and adding cancelation points in there
  - [-] seems to work, need further testing
- [] Mark a file as watching/not currently watching/broken link (file is in db, but doesn't exist anymore)
- [] Delete patches that can't be applied because whichever file they should be applied to no longer exists.
- [] Document all functions

#### Issues

- [x]It is not able to listen to the files, and I wonder if its because of different waitGroups.
  - [x] Perhaps make it so that you can set the manager waitgroup rather than it init automatically (client will need to provide it then in a similar pattern to cui)
- [x] Extreme cpu usage due to the file watcher. Perhaps can optimize out some go routines and profile why its so intensive. Also update fsnotify
  - [x] updating fsnotify didn't work, so going to look into replacing it with watcher.
    - [x] see https://twitter.com/amlwwalker/status/1143255867292364801
  - [x] Backward patch broken (probably subject/object wrong way around or file has been changed since patch was made...?)
    - [-] Is due to the versions changing. Should clean up patches that don't apply, based on the hash they are coming from. This is interesting, patches that cant be applied because the hash has changed should be deleted.
  - [] The cui could benefit from real time patch update display, however it will require keeping track of which file has been selected, and then if its the file that just changed, update the patch list. You could use a channel, or pass through a "udateUI" function, but its then involving the backend quite a lot. Lets see how we go about pushing the patche status (percentage) first and cancellation, and perhaps over the same channel we can send other discussion points which would allow this.
  - [-] Refactor the channels communicating between the patch and the manager about progress and cancelation
