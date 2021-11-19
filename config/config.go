package config

import (
	"errors"
	"fmt"
)

var (
	_driver IConfigDriver
)

// Init Init
// path: path of config file
// t: type of config file, it can be json\toml\yaml
// obj optinal，ptr type object, like &obj
func Init(t FileType, path string, obj interface{}) {
	if _driver != nil {
		panic(errors.New("you can only have one config"))
	}
	_driver = NewDriver(t, path)
	if err := _driver.Read(obj); err != nil {
		panic(fmt.Errorf("read config failed, error: %w", err))
	}
}

// Default default config
func Default() (DefaultConfig, error) {
	if _driver == nil {
		return DefaultConfig{}, errors.New("please init config first")
	}
	return _driver.Default(), nil
}
