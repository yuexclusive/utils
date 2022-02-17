package config

import (
	_ "embed"
)

var _driver interface{}

// Init Init
// path: path of config file
// t: type of config file, it can be json\toml\yaml
// obj optinal，ptr type object, like &obj
func Init[T any](path string) IDriver[T] {
	_driver = NewDriver[T](path)
	driver, ok := _driver.(IDriver[T])
	if !ok {
		panic("invalided driver type")
	}
	if err := driver.Read(); err != nil {
		panic(err)
	}
	return driver
}

// Get T Config, you must init config with t type first
func Get[T any]() T {
	if _driver == nil {
		panic("please init config first")
	}
	driver, ok := _driver.(IDriver[T])
	if !ok {
		panic("invalided driver type")
	}
	return driver.Get()
}

//go:embed config.toml
var defaultConfigContent string

// DefaultConfigContent default config content
func DefaultConfigContent() string {
	return defaultConfigContent
}
