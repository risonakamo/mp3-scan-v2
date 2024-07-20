// mp3 review bin configuration

package mp3review

import "mp3s-reviewer/lib/utils"

type Mp3ReviewConfig struct {
    // target dir where mp3s reside
    Mp3Dir string `yaml:"mp3Dir"`

    // include "maybe" items or not
    IncludeMaybe bool `yaml:"includeMaybe"`
}

// load mp3 review config yml
func LoadMp3ReviewConfig(path string) Mp3ReviewConfig {
    var result Mp3ReviewConfig
    var e error
    result,e=utils.ReadYaml[Mp3ReviewConfig](path)

    if e!=nil {
        panic(e)
    }

    return result
}