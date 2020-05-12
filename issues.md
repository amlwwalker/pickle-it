# PickleIt

This is a vague outline of what has been completed, and what hasn't 

## v0.1:

**v0.1** Will be a very stripped down, single file monitoring application demonstrating the rawest version of itself. The intent of this version is to configure the architecture of the management to be able to do what PickleIt is intended to do

- [x] flag sets a file to monitor
- [x] flag sets whether to patch
- [x] each time file changes, create a diff
  - [x] github.com/kr/binarydist uses readers/writers so by rights should be able to do this
        with bytes rather than files all the time, however writing out will be necessary -[x] save patches somewhere to the file system
- [x] keep track of files/patches in a database, as references if not in the DB itself
  - [x] look into Storm for BoltDB (https://github.com/asdine/storm)
- [x] create a naming convention for backups and diffs
- [x] patch a single file with a specific patch (currently works if patch name is obvious). TODO: Need to specify the patch name as a flag once the database is implemented

#### Issues

- [-] Diffing is quite slow at the moment, and so a few benchmarks are required, but mainly we need to do some stress tests on larger files
- [] To do the stress tests, we will need to keep some records and so a database is going to be implemented next to keep track of
  the changes and the results.
- [x] Interestingly, it will be worth keeping track of the diffs sizes because it may become obvious when a real file should be kept as a checkpoint, rather than the diffs from the original
- [x] At this point creating a filename structure for the backups and the diffs is a good idea.
- [x] move sync/watch/screenshots to different locations. This is irratating
  - [x] This is slightly irratating as the /sync/path needs to be swapped out of path names

##### Patching

forwards

```
./client -diff=false -file="./testingfiles/newimage.png" -patch="/Users/alex/pickleit/diff/bmV3aW1hZ2UucG5n_1eMFh4OO63y4L2S0f7BhlQ==_1561024469_newimage.png_1561024519_forward_diff.bin"
```

backwards

```
./client -diff=false -file="./testingfiles/newimage.png" -direction=backward -patch="/Users/alex/pickleit/diff/bmV3aW1hZ2UucG5n_1eMFh4OO63y4L2S0f7BhlQ==_1561024469_newimage.png_1561024519_backward_diff.bin"
```

_thursday 20th june, v0.1 tagged as complete_

## v0.2

**v0.2** Will build upon **v0.1** with the intent of cleaning up the _bedrock_ that **v0.1** is built on and giving the author time to reflect on the structure.

- a console UI so that versions can be patched
- Make sure that everything is running on routines and nothing is locking up the application
- Add multiple file/directory management
- improve cancelling diffs when a new diff is kicked off
- make possible to add new and other files to the watcher (not just once file at a time), and while the app is running
- Add ability to 'flick' between versions of a file with ease
- Benchmark compressing the diffs, size and time required to make diffs

#### Issues

- [x] Files with the same path can be added to the database - i.e its not currently unique on path. I think this is because the hash is invovled in the uniqueness when it looks at it. This is not a good system.
- [x] One issue is that it uses the hash of the file to check if its already monitoring a file. As the file is saved/updated the hash invariably changes. The hash should not be used, but the path to the file is the only thing we can go on (and the name). In this case, we should perhaps add the latest hash to the db, rather than use it as the watcher-file-checker.
  - [x] This is actually an issue as it also creates a sync of it, which is not so good....
- [-] Due to the slowness of the suffix sort, in this version, keeping track of a files sort will seriously speed up further diffing in future. This will obviously need to be stored, or regenerated and therefore is either about space, or time.
  - [x] Take the qsufsort and store it with the diff to get to that file.
  - [x] That will mean that we always have the sorted index for a version of a file in the database (need to know size of qsufsort index)
  - [x] Note every time a file changes, it will have a new sort index, so it does need to be calculated, therefore our best option is just storing it in one direction (as we won't have seen the other file yet) - [x] Note. Backward diffs have been disabled and a new structure has been implemented.
    - [x] Save the sorted algorithm to the database, seems easier that way.
    - [0] Clean up when late/out of date (give it a "time since last read" field?)
      - [x] Note if you only diff in one direction this is absolutely possible.
      - [x] to do this you will have to have check points when the diffs are closing in on the size of the file. Changes over 60% are considered enough for a new base

## v0.3

**v0.3** will concentrate on adding a Gui interface to PickleIt.

- [-] Add versions into the app. The version is available just not yet displayed
- [-] Make sure that everything communicates back its state so that progress is known
- [-] Add user concept (who created a patch etc)
  - currently can retrieve the logged in user, but doesn't use/store this information
- [-] Add custom naming conventions/servers/backups/file sharing/right click menu options
  - currently can create a naming convention but this isn't used. This should be displayed in the detailed view
- [-] add purging versions/bases/indexes
- [x] check index size and perhaps remove compression
- [x] make indexing happen in background, and let gui/user know when file is ready for version management
- [x] Mechanism to branch off. Need to set the Current Base back to the last patch applied. This will take affect when the original file is overwritten. Currently concerned about overwriting files. Creating new versions with version name is also an option here
- [x] issue - it tries to create an index for a file that already exists (hashes match). This is testable by moving by a pixel back and forth, where the base image would be identical. Does it matter however, perhaps just the error is a bit aggressive
- [x] calendar clean up, and fix bug when clicking on a date
- [0] fix patch highlight button to display the current patched/versioned checked out file
- [x] filter calendar view
- [x] add settings page, and add settings
- [x] store the screenshot against the patch
- [x] improve detailed view 
- [0] Add feature buttons but send to feedback/information webpage
- [-] setting to feedback usage to us for logging
- [0] display logs in a panel


## v0.4

**v0.4** will concentrate on implementing encryption and user capabilities, ready for backups of diffs and files to s3. File sharing will also begin to form in **v0.4**

- all undefined. Bugs removed from 0.3 first.

- [x] Remove bsdiff and use a simpler alogirthm. Now using fdelta.
- [-] get settings to actually do something. Currently, 

## v0.5

**v0.5** is currently undefined. At the point of **v0.4** testing can begin with a small user base. At this stage it is thought that **v0.5** will include breaking changes to previous versions to fix bugs, issues and user experience problems that make PickleIt less convenient to use

## Roadmap to 1.0

**v1.0** is the goal. At **v1.0** PickleIt will commit to no breaking changes going forward that would cause future versions to be incompatiable with each other.

- Add a plugin mechanism to build out extra capabilities
- Add encryption, i.e assymetric user and secret key encryption and symmetric file encryption
- Add backup and recovery capabilities
- Add custom local version management (how many versions a user keeps locally)
- Adding file sharing with other users will be necessary in v0.4 to fulfil the desires of PickleIt. Link sharing needs to be considered based on feedback, however this must not compromise encryption or convenience in anyway.

## Future

- Refactor out passing functions around and have a channel manager to keep track of everything between routines.
