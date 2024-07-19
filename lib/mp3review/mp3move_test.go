package mp3review

import "testing"

func Test_move(t *testing.T) {
	var e error=moveItemIntoDir("test/thing/a.txt","../thing")

    if e!=nil {
        panic(e)
    }
}