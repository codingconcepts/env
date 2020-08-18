# env
Tag-based environment configuration for structs.

[![Godoc](https://godoc.org/github.com/codingconcepts/env?status.svg)](https://godoc.org/github.com/codingconcepts/env)
[![Build Status](https://travis-ci.org/codingconcepts/env.svg?branch=master)](https://travis-ci.org/codingconcepts/env)
[![Go Report Card](https://goreportcard.com/badge/github.com/codingconcepts/env)](https://goreportcard.com/report/github.com/codingconcepts/env)

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

type config struct {
	Secret            []byte        `env:"SECRET" required:"true"`
	Region            string        `env:"REGION"`
	Port              int           `env:"PORT" required:"true"`
	Peers             []string      `env:"PEERS"` // you can use `delimiter` tag to specify separator, for example `delimiter:" "` 
	ConnectionTimeout time.Duration `env:"TIMEOUT" default:"10s"`
}

func main() {
	c := config{}
	if err := env.Set(&c); err != nil {
		log.Fatal(err)
	}

	...
}
```

``` bash
$ ID=1 SECRET=shh PORT=1234 PEERS=localhost:1235,localhost:1236 TIMEOUT=5s go run main.go
```

## Supported field types

- `bool` and `[]bool`
- `string` and `[]string`
- `[]byte`
- `int`, `int8`, `int16`, `int32`, `int64`, `[]int`, `[]int8`, `[]int16`, `[]int32`, and `[]int64`
- `uint`, `uint8`, `uint16`, `uint32`, `uint64`, `[]uint`, `[]uint8`, `[]uint16`, `[]uint32`, and `[]uint64`
- `float32`, `float64`, `[]float32`, and `[]float64`
- `time.Duration` and `[]time.Duration`
