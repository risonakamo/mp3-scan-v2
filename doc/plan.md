# interface
initialising job

- yml file + cli: call program on a yml file. saves having to keep giving the same path when returning to a job, but still requires calling on yml file (maybe can get drag yml file into the exe working?)
- global yml file: scuffed version of the previous, so dont
- cli: need to type in the target dir. easy to remember how to use but requires opening command line, ect. could get annoying.

reviewing

- web ui
- cli

# usage case
once in a while, a folder is created with a bunch of mp3s. the mp3s each need to be reviewed in random order, and a decision made on each item. when decision is made, the mp3 needs to be placed into a folder right next to the file based on the decision.

when randomly selecting an item, the item must already not be in one of the decision folders, unless counting "maybe" decisions.

# possible improvements
## where to place decided items
would it be better to just move all items of a particular decision into the same folder? why move them into a folder next to the original location

answer: the original location is needed as when tagging, all items that are associated with each other (ie, album) need to be right next to each other.