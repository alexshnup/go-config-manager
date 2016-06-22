##	Golang Config Manager

[![Go Report Card](https://goreportcard.com/badge/github.com/gkiryaziev/go-config-manager)](https://goreportcard.com/report/github.com/gkiryaziev/go-config-manager)


### Installation:
```sh
go get github.com/gkiryaziev/go-config-manager
```

### Usage:
```go
import (
	"fmt"
	yamlCfg "github.com/gkiryaziev/go-config-manager/yaml"
)

// Config struct
type Config struct {
	Name string `yaml:"name"`
	Age  int    `yaml:"age"`
}

	// config object
	var config Config

	// config manager
	err := yamlCfg.NewConfig("config.yaml").Load(&config)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(config.Name, config.Age)
```

```yaml
#config.yaml

name: "John"
age:  30
```