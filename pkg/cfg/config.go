package cfg

import (
	"github.com/jinzhu/configor"
)

func GetConfig(path string, conf interface{}) error {
	err := configor.Load(conf, path)
	if err != nil {
		return err
	}
	return nil
}
