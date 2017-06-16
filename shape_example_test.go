package shape

import (
	"fmt"
	"os"
)

func ExampleEnv() {
	os.Setenv("PROP_A", "value a")
	os.Setenv("PROP_B", "42")

	config := struct {
		PropA string `env:"PROP_A" required:"true"`
		PropB int16  `env:"PROP_B"`
	}{}

	Env(&config)

	fmt.Println(config)
	// OUTPUT: {value a 42}
}
