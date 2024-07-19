package mp3review

import (
	"fmt"
	"testing"
)

func Test_inSpecialDir(t *testing.T) {
    var testItems []string=[]string{
        "C:/Users/ktkm2/Desktop/song jobs/2024-06-20/mostu Rokudenashi - Eureka (Anime Shuumatsu Train Doko e Iku ED) MP3/01.Eureka.mp3",
        "C:/Users/ktkm2/Desktop/song jobs/2024-06-20/m/geoby SSGIRLS - Follow your arrows MP3/m/01.Follow your arrows",
    }

    var correct []bool=[]bool{
        false,
        true,
    }

    var i int
    var testItem string
    for i,testItem = range testItems {
        var result bool=isInSpecialDir(testItem)

        fmt.Println(
            testItem,
            result,
        )

        if correct[i]!=result {
            fmt.Println("result did not match")
            t.Fail()
        }
    }
}

func Test_determineSpecialDir(t *testing.T) {
    var testItems []string=[]string{
        "C:/Users/ktkm2/Desktop/song jobs/2024-06-20/mostu Rokudenashi - Eureka (Anime Shuumatsu Train Doko e Iku ED) MP3/01.Eureka.mp3",
        "C:/Users/ktkm2/Desktop/song jobs/2024-06-20/m/geoby SSGIRLS - Follow your arrows MP3/m/01.Follow your arrows",
    }

    var testitem string
    for _,testitem = range testItems {
        fmt.Println(determineNearestSpecialDir(
            testitem,
            "m",
        ))
    }
}