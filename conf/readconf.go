package conf

import (
	"github.com/brunetto/goutils/file"
)

func LoadJsonConf (fileName string, c interface{}) error {
	return file.LoadJsonConf(fileName, c)
}

