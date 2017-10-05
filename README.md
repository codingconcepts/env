# env
Tag-based environment configuration for structs.

[![Godoc](https://godoc.org/github.com/codingconcepts/env?status.svg)](https://godoc.org/github.com/codingconcepts/env)
[![Build Status](https://travis-ci.org/codingconcepts/env.svg?branch=master)](https://travis-ci.org/codingconcepts/env)
[![Exago](https://api.exago.io:443/badge/cov/github.com/codingconcepts/env)](https://exago.io/project/github.com/codingconcepts/env)

## Installation

``` bash
$ go get -u github.com/codingconcepts/env
```

## Usage

``` go
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/codingconcepts/env"
)

type awsConfig struct {
	Secret            string        `env:"SECRET" required:"true"`
	Region            string        `env:"REGION"`
	Port              int           `env:"PORT" required:"true"`
	Peers             []string      `env:"PEERS"`
	ConnectionTimeout time.Duration `env:"TIMEOUT"`
}

func main() {
	config := awsConfig{}
	if err := env.Set(&config); err != nil {
		log.Fatal(err)
	}

	...
}
```

``` bash
$ ID=1 SECRET=shh PORT=1234 PEERS=localhost:1235,localhost:1236 TIMEOUT=5s go run main.go
```

Env currently supports the following data types.  If you'd like to have more, please get in touch or feel free to create a pull request:

- bool and []bool
- string and []string
- int, int8, int16, int32, int64 and all slice equivalents
- uint, uint8, uint16, uint32, uint64 and all slice equivalents
- float32, float64 and all slice equivalents
- time.Duration and []time.Duration

### Default Values

If a field isn't required, it's also possible to specify a default value:

``` go
type config struct {
	Address string `env:"ADDRESS" default:"0.0.0.0"`
}
```
