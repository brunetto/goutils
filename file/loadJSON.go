package file

import (
	"io/ioutil"
	"encoding/json"
	"errors"
)

func LoadJsonConf (fileName string, c interface{}) error {
	var (
		jsonData []byte
		err error
	)
	jsonData, err = ioutil.ReadFile(fileName)
	if err != nil {
		return errors.New("LoadJsonConf: error while reading file: " + err.Error())
	}
	err = json.Unmarshal(jsonData, c)
	if err != nil {
		return errors.New("LoadJsonConf: error while unmarshaling JSON file: " + err.Error())
	}
	return nil
}

