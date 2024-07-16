// funcs implementing the special move procedure

package mp3review

import (
	"os"
	"path/filepath"
)

// move a target file into a dir relative to the file. create dir if it
// doesnt exist
func MoveItemIntoDir(target string,dirName string) {
    var e error
    _,e=os.Stat(target)

    if e!=nil {
        panic(e)
    }

    var targetDirLoc string=filepath.Join(
        filepath.Dir(target),
        dirName,
    )

    e=os.MkdirAll(targetDirLoc,0755)

    if e!=nil {
        panic(e)
    }

    e=os.Rename(
        target,
        filepath.Join(targetDirLoc,filepath.Base(target)),
    )

    if e!=nil {
        panic(e)
    }
}