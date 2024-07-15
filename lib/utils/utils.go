package utils

import (
	"fmt"
	"math/rand/v2"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// set zerolog global logger default options
func ConfigureDefaultZeroLogger() {
    log.Logger=log.Output(zerolog.ConsoleWriter{
        Out:os.Stdout,
        TimeFormat: "2006/01/02 15:04:05",
    })
}

// shuffle an array (in place)
func ShuffleArray[T any](array []T) {
    rand.Shuffle(len(array),func (i int,j int) {
        (array)[i],(array)[j]=(array)[j],(array)[i]
    })
}

// try to open web url or file with default program.
// essentially runs program like it was double clicked
func OpenTargetWithDefaultProgram(url string) error {
    fmt.Println("huh",fmt.Sprintf("\"%s\"",url))
    var cmd *exec.Cmd=exec.Command(
        "cmd","/c","start",
        "",
        url,
    )
    var e error=cmd.Run()

    if e!=nil {
        return e
    }

    return nil
}

// give folder location of the exe that calls this func
func GetHereDirExe() string {
    var exePath string
    var e error
    exePath,e=os.Executable()

    if e!=nil {
        panic(e)
    }

    return filepath.Dir(exePath)
}