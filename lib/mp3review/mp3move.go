// funcs implementing the special move procedure

package mp3review

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/rs/zerolog/log"
)

// more scoped version of move item into dir. only takes the valid special dir types.
// if the target item's current dir is one of the special dirs, moves up 2 levels.
func DoItemDecision(targetItem string,decision Mp3SpecialDir) error {
    if !isSpecialDir(decision) {
        log.Error().Msgf("tried to use bad special dir: %s",decision)
        return errors.New("invalid special dir type")
    }

    var destination string=string(decision)

    if isInSpecialDir(targetItem) {
        destination="../"+destination
    }

    return moveItemIntoDir(targetItem,destination)
}

// move a target file into a dir relative to the file. create dir if it
// doesnt exist. if the dir is the same as the current dir, do nothing
func moveItemIntoDir(target string,dirName string) error {
    var result fs.FileInfo
    var e error
    result,e=os.Stat(target)

    if e!=nil {
        return e
    }

    // need to return result to check if actually found?
    result.Name()

    target,e=filepath.Abs(target)

    if e!=nil {
        return e
    }

    var targetDirLoc string=filepath.Join(
        filepath.Dir(target),
        dirName,
    )

    targetDirLoc,e=filepath.Abs(targetDirLoc)

    if e!=nil {
        return e
    }

    if targetDirLoc==filepath.Dir(target) {
        log.Info().Msg("not moving - same directory")
        return nil
    }

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