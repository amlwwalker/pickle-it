# PickleIt

PickleIt is a project inspired by seeing people organise versions of documents/files with extensions like

```
pres.v1.pptx
pres.v2.pptx
pres.v3FINAL.pptx
pres.v3FINALF**K.pptx
```

I wanted to see if I could build an application that would make versioning of files, something along the lines of git
, for code versioning.


## Download

Currently I have only compiled it for OSX, and as I only have my dev computer available to test it, I can't be sure
 everything will work for you, however the [OSX download link is here](https://drive.google.com/drive/folders/1549Q1h66PTeGEKdlq4bKQKt6osJSSCNM). It's stored on Google Drive, so just click the
  big Download All button, top right. Unzip it, and double click it to run it.

## Compilation

Although there is a command line interface version, this is really about qt and the gui. Unfortunately the Qt library doesn't allow for go modules (or didn't at the time of writing the application) so you will need to install that directly. See https://github.com/therecipe/qt.
* After that, inside the gui directory is a makefile, and something like `make local` or `make desktop` will build you the compiled application. There may end up (again due to the go modules situation) other libraries that you will need to `go get`

## Basics
At this stage, without going into the technical details (possibly I will here at a later date), PickleIt allows the
 user to drop any file, they are working on, into PickleIt. From that moment on, PickleIt will watch for changes on
  the file (saves) and create a version of the file. This way, at any point, the user can go back to a previous
   version by just "patching" it - which is the click of a button in PickleIt
   

1. Create a file that you are going to work on
2. Drop the file into the Dropzone. As long as pickleit is open in the background it will watch for changes to this file
3. When a save occurs on the file, pickleit will take a binary diff of the file and add the diff to the list. You can either delete or patch any diff at any time. Note, you may need to close the original to see the diff take effect when patching.
4. Diffs can have screenshots (if the app has permission) and descriptions to help you remember what was going on at that moment in time
5. There is also a calendar view of diffs so you can go back in time based on the calendar changes.

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
 the app when it was saved. Clicking the picture will enlarge it. There is also a description box so you can add some
  detail for when you come back, to each patch. Enter some text, and hit enter, then click just outside the input
 
### Options

* There are two things that are possible to do to a patch at the moment, deleting a patch, and patching a patch. 
* You can turn watching off at any time, by clicking the radio button half way down on the right

![](/images/settings.png)


### Note

You can close the window and PickleIt will run in the background. There is an icon in the system tray to bring it
 back to the front.


