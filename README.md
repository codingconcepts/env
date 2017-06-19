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
	ID     string `env:"AWS_ACCESS_KEY_ID" required:"true"`
	Secret string `env:"AWS_SECRET_ACCESS_KEY" required:"true"`
	Region string `env:"AWS_REGION"`
}

func main() {
	config := awsConfig{}
	if err := env.Set(&config); err != nil {
		log.Fatal(err)
	}

	...
}
```