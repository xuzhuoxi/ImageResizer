package lib

import (
	"gopkg.in/yaml.v3"
	"os"
)

func UnmarshalYamlData(dataPath string, dataRef interface{}) error {
	bs, err := os.ReadFile(dataPath)
	if nil != err {
		return err
	}
	err = yaml.Unmarshal(bs, dataRef)
	if nil != err {
		return err
	}
	return nil
}
