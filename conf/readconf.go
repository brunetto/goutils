package conf

import (
	"io/ioutil"
	"encoding/json"
)

func LoadJsonConf (fileName string, c interface{}) error {
	var (
		jsonMsg []byte
		err error
	)
	jsonMsg, err = ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}
	err = json.Unmarshal(jsonMsg, c)
	if err != nil {
		return err
	}
	return nil
}

