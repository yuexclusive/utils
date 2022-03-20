package config

import (
	"reflect"
	"testing"
)

type DefaultConfig struct {
	Config
	Customers []string `mapstructure:"customers"`
}

func TestDefault(t *testing.T) {
	driver := Init[DefaultConfig]("./config.toml")
	cfg := driver.Get()
	if got, want := cfg.Name, "example.srv"; got != want {
		t.Errorf("got: %s, want: %s", got, want)
	}
	if got, want := cfg.Customers, []string{"aa", "bb", "cc"}; !reflect.DeepEqual(got, want) {
		t.Errorf("got: %s, want: %s", got, want)
	}

	if got, want := cfg.AuthHost, "http://localhost:8080"; got != want {
		t.Errorf("got: %s, want: %s", got, want)
	}
}

type Customer struct {
	A string    `mapstructure:"a"`
	B int       `mapstructure:"b"`
	C []float64 `mapstructure:"c"`
}

func TestCustomer(t *testing.T) {
	driver := Init[Customer]("./customer.toml")

	want := Customer{A: "a", B: 1, C: []float64{1.1, 2.2}}
	got := driver.Get()

	if !reflect.DeepEqual(want, got) {
		t.Errorf("want: %v, got: %v", want, got)
	}
}

func TestCustomerJSON(t *testing.T) {
	driver := Init[Customer]("./customer.json")

	want := Customer{A: "a", B: 1, C: []float64{1.1, 2.2}}
	got := driver.Get()

	if !reflect.DeepEqual(want, got) {
		t.Errorf("want: %v, got: %v", want, got)
	}
}

func TestCustomerYAML(t *testing.T) {
	driver := Init[Customer]("./customer.yaml")

	want := Customer{A: "a", B: 1, C: []float64{1.1, 2.2}}
	got := driver.Get()

	if !reflect.DeepEqual(want, got) {
		t.Errorf("want: %v, got: %v", want, got)
	}
}

func BenchmarkCustomer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		driver := Init[DefaultConfig]("./config.toml")
		_ = driver.Get()
	}
}
