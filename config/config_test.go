package config

import (
	"reflect"
	"testing"
)

type CustomerConfig struct {
	Config    `mapstructure:"config"`
	Customers []string `mapstructure:"customers"`
}

func TestCustomer(t *testing.T) {
	var c CustomerConfig

	Init(TOML, "./customer_config.toml", &c)

	cfg, err := Default()

	if err != nil {
		t.Error(err)
	}

	if got, want := cfg.Name, "example.srv"; got != want {
		t.Errorf("got: %s, want: %s", got, want)
	}
	if got, want := c.Customers, []string{"aa", "bb", "cc"}; !reflect.DeepEqual(got, want) {
		t.Errorf("got: %s, want: %s", got, want)
	}
}
