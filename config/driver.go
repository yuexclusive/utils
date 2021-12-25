package config

import (
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

// IDriver IDriver
type IDriver[T any] interface {
	GetType() FileType
	GetPath() string
	Read() error
	GetConfig() T
}

// Driver Driver
type Driver[T any] struct {
	fileType FileType
	path     string
	config   T
}

func (d *Driver[T]) GetType() FileType { return d.fileType }
func (d *Driver[T]) GetPath() string   { return d.path }
func (d *Driver[T]) Read() error {
	var cfg T
	viper.SetConfigType(string(d.fileType))

	viper.SetConfigFile(d.path)

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		return err
	}
	d.config = cfg
	return nil
}

func (d *Driver[T]) GetConfig() T {
	return d.config
}

func NewDriver[T any](fileType FileType, path string) IDriver[T] {
	return &Driver[T]{fileType: fileType, path: path}
}
