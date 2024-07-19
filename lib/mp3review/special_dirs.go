// funcs dealing with the special mp3 review dirs

package mp3review

import "path/filepath"

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
    return isSpecialDir(filepath.Base(filepath.Dir(path)))
}

// determines if a string is one of the special dirs
func isSpecialDir(dirname string) bool {
    switch dirname {
        case "y","n","m","done","yes","no","maybe":
        return true
    }

    return false
}