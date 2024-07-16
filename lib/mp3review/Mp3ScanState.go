// state containing class for mp3 scan bin

package mp3review

import (
	"path/filepath"
)

type Mp3ScanState struct {
    // list of paths of items
    items []string

    // index of current item
    currentItemI int
}

// constructed status view of the state
type Mp3ReviewStatus struct {
	CurrentItem string `json:"currentItem"`
	CurrentItemFolder string `json:"currentItemFolder"`

	TotalItems int `json:"totalItems"`
	CurrentItemIndex int `json:"currentItemIndex"`

	NoMoreItems bool `json:"noMoreItems"`
}

// create new scan state on a target dir
func NewScanState(targetDir string) Mp3ScanState {
    var targetFiles []string=findMp3sShuffled(targetDir)

    return Mp3ScanState{
        items: targetFiles,
        currentItemI: 0,
    }
}

// get current status. returns weird looking one if no more items
func (state *Mp3ScanState) GetStatus() Mp3ReviewStatus {
    if state.currentItemI>=len(state.items) {
        return Mp3ReviewStatus{
            CurrentItem: "",
            CurrentItemFolder: "",
            TotalItems: len(state.items),
            CurrentItemIndex: -1,
            NoMoreItems: true,
        }
    }

    return Mp3ReviewStatus{
		CurrentItem: filepath.Base(state.items[state.currentItemI]),
		CurrentItemFolder: filepath.Base(filepath.Dir(state.items[state.currentItemI])),
		TotalItems: len(state.items),
		CurrentItemIndex: state.currentItemI,
        NoMoreItems: false,
    }
}

// advance to the next item, and return the next state. index will not grow
// larger than the size of the items (being at the size of items means there
// is no more items, though)
func (state *Mp3ScanState) advanceItem() Mp3ReviewStatus {
    state.currentItemI+=1

    if state.currentItemI>len(state.items) {
        state.currentItemI=len(state.items)
    }

    return state.GetStatus()
}

// perform decision on the current item, if there is one, and advance to the next
// item. returns new state. if there is no item, does nothing
func (state *Mp3ScanState) decideItem(decision Mp3Decision) Mp3ReviewStatus {
    if state.currentItemI>=len(state.items) {
        return state.GetStatus()
    }

    var e error=DoItemDecision(
        state.items[state.currentItemI],
        decision,
    )

    if e!=nil {
        panic(e)
    }

    return state.advanceItem()
}