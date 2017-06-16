# shape
Tag-based environment configuration for structs.

[![Godoc](https://godoc.org/github.com/codingconcepts/shape?status.svg)](https://godoc.org/github.com/codingconcepts/shape)
[![Build Status](https://travis-ci.org/codingconcepts/shape.svg?branch=master)](https://travis-ci.org/codingconcepts/shape)
[![Exago](https://api.exago.io:443/badge/cov/github.com/codingconcepts/shape)](https://exago.io/project/github.com/codingconcepts/shape)

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