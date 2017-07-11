package env

import (
	"fmt"
	"os"
)

func ExampleSet() {
	os.Setenv("PROP_A", "value a")
	os.Setenv("PROP_B", "42")

	config := struct {
		PropA string `env:"PROP_A" required:"true"`
		PropB int16  `env:"PROP_B"`
	}{}

	err := Set(&config)

	fmt.Println(config, err)
	// OUTPUT: {value a 42} <nil>
}
