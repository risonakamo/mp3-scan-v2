// funcs implementing main special scanning function to find mp3s

package mp3review

import (
	"mp3s-reviewer/lib/utils"
	"path/filepath"
	"strings"

	"github.com/bmatcuk/doublestar/v4"
)

// find all mp3s while applying special rules (see should be excluded)
func FindMp3s(topDir string) []string {
    var targetFiles []string
    var e error
    targetFiles,e=doublestar.FilepathGlob(filepath.Join(
        topDir,
        "**/*.mp3",
    ))

    if e!=nil {
        panic(e)
    }

    var goodFiles []string
    var file string
    for _,file = range targetFiles {
        if !shouldBeExcluded(file) {
            goodFiles=append(goodFiles,file)
        }
    }

    return goodFiles
}

// determine if a path should be ignored or not.
// only includes the file components that are present in the path (non-absolute). i.e., if user provides
// path while sitting inside of a folder that would be ignored, this func should have no knowledge
// of that.
// file should be ignored if:
// - not an mp3 file
// - anywhere in the path, it is under a folder called:
//   - done, y, m, n
func shouldBeExcluded(file string) bool {
    if !strings.HasSuffix(strings.ToLower(file),".mp3") {
        return true
    }

    file=filepath.Clean(file)
    file=filepath.ToSlash(file)

    var splitPath []string=strings.Split(file,"/")

    var pathPiece string
    for _,pathPiece = range splitPath {
        switch pathPiece {
            case "y","n","m","done":
            return true
        }
    }

    return false
}

// uses find mp3s to find mp3s. shuffles and returns result
func findMp3sShuffled(targetDir string) []string {
	var foundFiles []string=FindMp3s(targetDir)
	utils.ShuffleArray(foundFiles)

	return foundFiles
}