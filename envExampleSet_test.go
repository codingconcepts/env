package env

import (
	"fmt"
	"os"
	"time"
)

func ExampleSet() {
	os.Setenv("HOSTS", "NEHOST1:1234, NEHOST2:1234")
	os.Setenv("PORT", "1234")
	os.Setenv("PEER_TIMEOUT", "2m500ms")

	config := struct {
		Hosts              []string      `env:"HOSTS" required:"true"`
		Port               int16         `env:"PORT" required:"true"`
		PeerConnectTimeout time.Duration `env:"PEER_TIMEOUT" default:"1s500ms"`
	}{}

	err := Set(&config)

	fmt.Println(config, err)
	// OUTPUT: {[NEHOST1:1234 NEHOST2:1234] 1234 2m0.5s} <nil>
}
