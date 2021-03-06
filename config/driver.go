package config

import (
	"path/filepath"
	"strings"

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
	Read() error
	GetType() FileType
	GetPath() string
	Get() T
}

// driver driver
type driver[T any] struct {
	fileType FileType
	path     string
	config   T
}

func (d *driver[T]) GetType() FileType { return d.fileType }
func (d *driver[T]) GetPath() string   { return d.path }
func (d *driver[T]) Read() error {
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

func (d *driver[T]) Get() T {
	return d.config
}

func NewDriver[T any](path string) IDriver[T] {
	driver := &driver[T]{path: path}
	ext := strings.ToLower(filepath.Ext(path))
	var fileType FileType = TOML
	switch ext {
	case ".toml":
		fileType = TOML
	case ".yaml", ".yml":
		fileType = YAML
	case ".json":
		fileType = JSON
	}
	driver.fileType = fileType
	return driver
}
