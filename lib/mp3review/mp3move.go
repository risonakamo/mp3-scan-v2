// funcs implementing the special move procedure

package mp3review

import (
	"errors"
	"os"
	"path/filepath"
)

// possible mp3 decisions
type Mp3Decision string
const (
    Mp3Decision_yes="yes"
    Mp3Decision_no="no"
    Mp3Decision_maybe="maybe"
)

// more scoped version of move item into dir. only takes the valid decision types
func DoItemDecision(targetItem string,decision Mp3Decision) error {
    if !checkMp3Decision(decision) {
        return errors.New("invalid decision type")
    }

    return moveItemIntoDir(targetItem,string(decision))
}

// move a target file into a dir relative to the file. create dir if it
// doesnt exist
func moveItemIntoDir(target string,dirName string) error {
    var e error
    _,e=os.Stat(target)

    if e!=nil {
        return e
    }

    var targetDirLoc string=filepath.Join(
        filepath.Dir(target),
        dirName,
    )

    e=os.MkdirAll(targetDirLoc,0755)

    if e!=nil {
        return e
    }

    e=os.Rename(
        target,
        filepath.Join(targetDirLoc,filepath.Base(target)),
    )

    if e!=nil {
        return e
    }

    return nil
}

// returns true if the decision is valid
func checkMp3Decision(decision Mp3Decision) bool {
    switch decision {
        case Mp3Decision_yes,Mp3Decision_no,Mp3Decision_maybe:
            return true
    }

    return false
}