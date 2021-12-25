package config

import (
	_ "embed"
)

// Init Init
// path: path of config file
// t: type of config file, it can be json\toml\yaml
// obj optinal，ptr type object, like &obj
func Init[T any](fileType FileType, path string) IDriver[T] {
	driver := NewDriver[T](fileType, path)

	if err := driver.Read(); err != nil {
		panic(err)
	}
	return driver
}

//go:embed config.toml
var defaultConfig string

func DefaultConfigFile() string {
	return defaultConfig
}
