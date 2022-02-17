package config

import (
	"reflect"
	"testing"
)

type CustomerConfig struct {
	Config
	Customers []string `mapstructure:"customers"`
}

func TestCustomer(t *testing.T) {
	driver := Init[CustomerConfig]("./config.toml")
	cfg := driver.Get()
	if got, want := cfg.Name, "example.srv"; got != want {
		t.Errorf("got: %s, want: %s", got, want)
	}
	if got, want := cfg.Customers, []string{"aa", "bb", "cc"}; !reflect.DeepEqual(got, want) {
		t.Errorf("got: %s, want: %s", got, want)
	}

	if got, want := cfg.AuthHost, "http://localhost:8080"; !reflect.DeepEqual(got, want) {
		t.Errorf("got: %s, want: %s", got, want)
	}
}
