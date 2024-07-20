// generic file read/write utils

package utils

import (
	"errors"
	"io/fs"
	"os"

	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v3"
)

// read a yaml file and return result
func ReadYaml[DataT any](filename string) (DataT,error) {
	var data []byte
	var e error
	data,e=os.ReadFile(filename)

	if errors.Is(e,fs.ErrNotExist) {
		log.Info().Msgf("file not found: %s",filename)
		var def DataT
		return def,e
	}

	if e!=nil {
		var def DataT
		return def,e
	}

	var parsedData DataT
	yaml.Unmarshal(data,&parsedData)

	return parsedData,nil
}