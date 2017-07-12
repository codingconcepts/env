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

	"github.com/codingconcepts/env"
)

type awsConfig struct {
	ID                string        `env:"AWS_ACCESS_KEY_ID" required:"true"`
	Secret            string        `env:"AWS_SECRET_ACCESS_KEY" required:"true"`
	Region            string        `env:"AWS_REGION"`
	ConnectionTimeout time.Duration `env:"CONNECTION_TIMEOUT"`
}

func main() {
	config := awsConfig{}
	if err := env.Set(&config); err != nil {
		log.Fatal(err)
	}

	...
}
```

Env currently supports the following data types.  If you'd like to have more, please get in touch or feel free to create a pull request:

- bool and []bool
- string and []string
- int, int8, int16, int32, int64 and all slice equivalents
- uint, uint8, uint16, uint32, uint64 and all slice equivalents
- float32, float64 and all slice equivalents
- time.Duration and []time.Duration

## Todo

- [ ] Allow user to provide custom delimiter for slice types (environment config might be unchangeable)