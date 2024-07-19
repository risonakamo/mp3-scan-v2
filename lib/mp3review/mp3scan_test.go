package mp3review

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/bmatcuk/doublestar/v4"
	"github.com/k0kubun/pp/v3"
)

func Test_shouldBeExcluded(t *testing.T) {
    something:=filepath.Join(
        "C:\\Users\\ktkm2\\Desktop\\song jobs\\2024-04-27",
        "**/*.mp3",
    )

    fmt.Println(something)

    targetItems,e:=doublestar.FilepathGlob(
        filepath.Join(
            "E:\\mp3\\new\\2024-04-27",
            "**/*.mp3",
        ),
    )

    if e!=nil {
        panic(e)
    }

    // pp.Println(targetItems)

    var item string
    for _,item = range targetItems {
        // shouldBeExcluded(item)
        fmt.Println(item)
        fmt.Println("->",shouldBeExcluded(item,false))
        fmt.Println()
    }
}

func Test_findMp3s(t *testing.T) {
    result:=FindMp3s("C:/Users/ktkm2/Desktop/song jobs/2024-06-20",false)

    pp.Println(result)
}