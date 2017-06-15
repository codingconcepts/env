package shape

import (
	"fmt"
	"os"
)

func ExampleEnv() {
	os.Setenv("PROP_A", "value a")
	os.Setenv("PROP_B", "value b")

	config := struct {
		PropA string `env:"PROP_A"`
		PropB string `env:"PROP_B"`
	}{}

	Env(&config)

	fmt.Println(config)

	// OUTPUT: {value a value b}
}
