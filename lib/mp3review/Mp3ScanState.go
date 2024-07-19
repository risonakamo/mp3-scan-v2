// state containing class for mp3 scan bin

package mp3review

import (
	"errors"
	"mp3s-reviewer/lib/utils"
	"path/filepath"

	"github.com/rs/zerolog/log"
)

// mp3 scanner bin state holding class
type Mp3ScanState struct {
    // list of paths of items
    items []string

    // index of current item
    currentItemI int
}

// constructed status view of the state
type Mp3ReviewStatus struct {
	CurrentItem string `json:"currentItem"`
    // folder of the item. if the item is in one of the special folders, gives the parent of that folder
    // instead
	CurrentItemFolder string `json:"currentItemFolder"`

	TotalItems int `json:"totalItems"`
	CurrentItemIndex int `json:"currentItemIndex"`

	NoMoreItems bool `json:"noMoreItems"`
}

type Mp3ScanStateError error
var (
    Mp3ScanStateError_noitems Mp3ScanStateError=errors.New("no more items")
    Mp3ScanStateError_failedToMove Mp3ScanStateError=errors.New("failed to move item")
)

// create new scan state on a target dir
func NewScanState(targetDir string,includeMaybes bool) Mp3ScanState {
    log.Info().Msgf("scanning: %s",targetDir)

    var targetFiles []string=findMp3sShuffled(targetDir,includeMaybes)

    log.Info().Msgf("initialised: tracking %d items",len(targetFiles))

    return Mp3ScanState{
        items: targetFiles,
        currentItemI: 0,
    }
}

// get current status. returns weird looking one if no more items.
// if the item is inside a special dir, the folder name includes the one above the
// special dir
func (state *Mp3ScanState) GetStatus() Mp3ReviewStatus {
    if state.NoMoreItems() {
        return Mp3ReviewStatus{
            CurrentItem: "",
            CurrentItemFolder: "",
            TotalItems: len(state.items),
            CurrentItemIndex: -1,
            NoMoreItems: true,
        }
    }

    var itemFolderPath string=filepath.Dir(state.items[state.currentItemI])
    var itemFolderSingle string=filepath.Base(itemFolderPath)

    // if the item is in a special dir, find the parent of the item's parent, and
    // add it to the item's parent path
    if isSpecialDir(Mp3SpecialDir(itemFolderSingle)) {
        var itemFolderFodler string=filepath.Base(filepath.Dir(itemFolderPath))
        itemFolderSingle=filepath.Join(itemFolderFodler,itemFolderSingle)
    }

    return Mp3ReviewStatus{
		CurrentItem: filepath.Base(state.items[state.currentItemI]),
		CurrentItemFolder: itemFolderSingle,
		TotalItems: len(state.items),
		CurrentItemIndex: state.currentItemI,
        NoMoreItems: false,
    }
}

// advance to the next item, and return the next state. index will not grow
// larger than the size of the items (being at the size of items means there
// is no more items, though)
func (state *Mp3ScanState) AdvanceItem() Mp3ReviewStatus {
    state.currentItemI+=1

    if state.currentItemI>len(state.items) {
        state.currentItemI=len(state.items)
    }

    return state.GetStatus()
}

// perform decision on the current item, if there is one, and advance to the next
// item. returns new state. if there is no item, does nothing.
// if failed to move item, does not advance
func (state *Mp3ScanState) DecideItem(decision Mp3SpecialDir) (Mp3ReviewStatus,error) {
    if state.NoMoreItems() {
        return state.GetStatus(),Mp3ScanStateError_noitems
    }

    log.Info().Msgf("moving item: %s",state.items[state.currentItemI])

    var e error=DoItemDecision(
        state.items[state.currentItemI],
        decision,
    )

    if e!=nil {
        log.Err(e).Msg("failed to move item")
        return state.GetStatus(),Mp3ScanStateError_failedToMove
    }

    return state.AdvanceItem(),nil
}

// return if no more items
func (state *Mp3ScanState) NoMoreItems() bool {
    return state.currentItemI>=len(state.items)
}

// try to open the current item, if any
func (state *Mp3ScanState) OpenItem() {
    if state.NoMoreItems() {
        return
    }

    log.Info().Msgf("opening item: %s",state.items[state.currentItemI])
    utils.OpenTargetWithDefaultProgram(state.items[state.currentItemI])
}