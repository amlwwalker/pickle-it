# PickleIt

PickleIt is a pet project inspired by seeing people organise versions of documents/files with extensions like

```
pres.v1.pptx
pres.v2.pptx
pres.v3FINAL.pptx
pres.v3FINALF**K.pptx
```

I wanted to see if I could build an application that would make versioning of files, something along the lines of git
, for code versioning.

## Basics
At this stage, without going into the technical details (possibly I will here at a later date), PickleIt allows the
 user to drop any file, they are working on, into PickleIt. From that moment on, PickleIt will watch for changes on
  the file (saves) and create a version of the file. This way, at any point, the user can go back to a previous
   version by just "patching" it - which is the click of a button in PickleIt
   

## Detail Views

There are two ways that patches are organised within PickleIt, 

### Calendar View

The calendar view shows the amount of patches that were made per day, when a date is selected, a list of the patches
 from that day will appear on the left. [Still in alpha] On OSX only currently, a screenshot will appear of the
  focused window at the time of the save, giving the viewer an idea of what they were looking at when they made that
   patch.
   
  ![](/images/calendar.png)
   
### List View

The list view displays all the patches, filtered by the drop down menu, for a file. Selecting a file will display the
 detailed view for this file on the right.
 
 ![](/images/listview.png)
 
### Drop Down

The drop down allows slightly better filtering, by file. This affects both the calendar and the list view

### Detail View

The detail view shows you the name of the file, the creation data, and (on OSX and only in alpha), a screenshot of
 the app when it was saved. Clicking the picture will enlarge it.
 
### Options

* There are two things that are possible to do to a patch at the moment, deleting a patch, and patching a patch. 
* You can turn watching off at any time, by clicking the radio button half way down on the right

![](/images/settings.png)


### Note

You can close the window and PickleIt will run in the background. There is an icon in the system tray to bring it
 back to the front.
