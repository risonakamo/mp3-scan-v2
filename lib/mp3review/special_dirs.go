// funcs dealing with the special mp3 review dirs

package mp3review

import "path/filepath"

type Mp3SpecialDir string
const (
    SpecialDir_yes Mp3SpecialDir="yes"
    SpecialDir_no Mp3SpecialDir="no"
    SpecialDir_maybe Mp3SpecialDir="maybe"
    SpecialDir_y Mp3SpecialDir="y"
    SpecialDir_n Mp3SpecialDir="n"
    SpecialDir_m Mp3SpecialDir="m"
    SpecialDir_done Mp3SpecialDir="done"
)

// given path to an item and a special dir, determine the new filepath of the item so that
// it resides in a special dir right next to the original file path.
// if the filepath is already in a special dir, moves up 1 level before creating the special
// dir, so all special dirs are right next to each other
func determineNearestSpecialDir(path string,specialDir string) string {
    if !isInSpecialDir(path) {
        return filepath.Join(
            filepath.Dir(path),
            specialDir,
            filepath.Base(path),
        )
    }

    return filepath.Join(
        filepath.Dir(path),
        "..",
        specialDir,
        filepath.Base(path),
    )
}

// given a file path to an item, determine if the item is within a special dir. only
// considers the current dir of the item
func isInSpecialDir(path string) bool {
    return isSpecialDir(Mp3SpecialDir(filepath.Base(filepath.Dir(path))))
}

// determines if a string is one of the special dirs
func isSpecialDir(dirname Mp3SpecialDir) bool {
    switch dirname {
        case SpecialDir_yes,
            SpecialDir_no,
            SpecialDir_maybe,
            SpecialDir_y,
            SpecialDir_n,
            SpecialDir_m,
            SpecialDir_done:
        return true
    }

    return false
}