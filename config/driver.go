package config

import (
	"fmt"
	"reflect"

	"github.com/spf13/viper"
)

// FileType FileType
type FileType string

const (
	// TOML TOML
	TOML FileType = "toml"
	// YAML YAML
	YAML FileType = "yaml"
	// JSON JSON
	JSON FileType = "json"
)

// IConfigDriver IConfigDriver
type IConfigDriver interface {
	Read(obj interface{}) error
	Default() DefaultConfig
}

// Driver Driver
type Driver struct {
	Type          FileType
	DefaultConfig DefaultConfig
	Path          string
}

// NewDriver NewDriver
func NewDriver(t FileType, path string) IConfigDriver {
	return &Driver{Type: t, Path: path}
}

// Init Init
func (d *Driver) Read(obj interface{}) error {
	viper.SetConfigType(string(d.Type))
	viper.SetConfigFile(d.Path)

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	var defaultConfig DefaultConfig

	if err := viper.Unmarshal(&defaultConfig); err != nil {
		return err
	}

	d.DefaultConfig = defaultConfig

	if obj != nil {
		objType := reflect.TypeOf(obj)

		if field, ok := objType.Elem().FieldByName("Config"); ok {
			if field.Type.String() != "config.Config" {
				return fmt.Errorf("obj must have field config.Config")
			}
		} else {
			return fmt.Errorf("obj must have field config.Config")
		}

		if objType.Kind() != reflect.Ptr {
			return fmt.Errorf("config Init: obj must be ptr type")
		}
		if err := viper.Unmarshal(obj); err != nil {
			return err
		}
	}
	return nil
}

// Default Default
func (d *Driver) Default() DefaultConfig {
	return d.DefaultConfig
}
