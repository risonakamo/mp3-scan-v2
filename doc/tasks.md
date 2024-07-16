- [x] link button needs "selected" state. have decision buttons become "selected" on click
- [x] "next item" needs to be in disabled state until decision is locked in
- [x] folder needs to be populated
- [x] implement move item into folder func
- [x] reorganise server to use class for state management
- [x] performs moving
    - [x] api to move the current item and move to the next item
    - [x] either replace current next item or change to skip item. or have the next item able to skip as well?
- [ ] show error somewhere when failed to do something
- [x] on successfully changing item, need to reset the decision

# 2
- [ ] progress meter functionality
- [ ] when out of items, controls should all become disabled, and info zone should accordingly say something about everything being done. maybe keep around the progress indicators and change name to "ALL DONE"?
- [ ] while opening item, disable and set text to "opening item..."
- [ ] keyboard shortcuts
- [ ] skip item
- [ ] maybe mode
- [ ] tscheck in terminal not detecting ts errors, but vscode is detecting it

# 3
- [ ] also detect all the items in y/n/m and add it to the progress number (but they are non-interactable except in maybe mode)
- [ ] flashing on page load due to delay in getting state

# when back online
- [ ] switch to lodash
- [ ] switch submodule url to github
- [ ] update gomod/package json