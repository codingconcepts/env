# shape
Tag-based environment configuration for structs.

## Installation

``` bash
$ go get -u github.com/codingconcepts/shape
```

## Usage

``` go
package main

import (
	"fmt"
	"log"

	"github.com/codingconcepts/shape"
)

type awsConfig struct {
	ID     string `env:"AWS_ACCESS_KEY_ID" required:"true"`
	Secret string `env:"AWS_SECRET_ACCESS_KEY" required:"true"`
	Region string `env:"AWS_REGION"`
}

func main() {
	config := awsConfig{}
	if err := shape.Env(&config); err != nil {
		log.Fatal(err)
	}

	...
}
```

## Todos

- [x] required tag
- [ ] check for unexported