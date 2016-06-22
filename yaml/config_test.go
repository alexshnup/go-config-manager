package yaml

import "testing"

// Config struct
type Config struct {
	Name string `yaml:"name"`
	Data Data   `yaml:"data"`
}

// Data struct
type Data struct {
	Age int `yaml:"age"`
}

func TestLoad(t *testing.T) {
	// config object
	var config Config

	// config manager
	err := NewConfig("config.yaml").Load(&config)
	if err != nil {
		t.Fatal(err)
	}

	// chech parameters
	switch {
	case "John" != config.Name:
		t.Error("Error, John !=", config.Name)
	case 30 != config.Data.Age:
		t.Error("Error, 30 !=", config.Data.Age)
	}
}
