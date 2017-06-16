package shape

import (
	"os"
	"strings"
	"testing"

	"github.com/codingconcepts/shape/test"
)

func TestEnvBool(t *testing.T) {
	os.Setenv("PROP", "true")

	config := struct {
		Prop bool `env:"PROP"`
	}{}

	test.ErrorNil(t, Env(&config))
	test.Equals(t, true, config.Prop)
}

func TestEnvIntegers(t *testing.T) {
	os.Setenv("PROP", "123")

	config := struct {
		PropInt   int   `env:"PROP"`
		PropInt8  int8  `env:"PROP"`
		PropInt16 int16 `env:"PROP"`
		PropInt32 int32 `env:"PROP"`
		PropInt64 int64 `env:"PROP"`
	}{}

	test.ErrorNil(t, Env(&config))
	test.Equals(t, int(123), config.PropInt)
	test.Equals(t, int8(123), config.PropInt8)
	test.Equals(t, int16(123), config.PropInt16)
	test.Equals(t, int32(123), config.PropInt32)
	test.Equals(t, int64(123), config.PropInt64)
}

func TestEnvUnsignedIntegers(t *testing.T) {
	os.Setenv("PROP", "123")

	config := struct {
		PropUint   uint   `env:"PROP"`
		PropUint8  uint8  `env:"PROP"`
		PropUint16 uint16 `env:"PROP"`
		PropUint32 uint32 `env:"PROP"`
		PropUint64 uint64 `env:"PROP"`
	}{}

	test.ErrorNil(t, Env(&config))
	test.Equals(t, uint(123), config.PropUint)
	test.Equals(t, uint8(123), config.PropUint8)
	test.Equals(t, uint16(123), config.PropUint16)
	test.Equals(t, uint32(123), config.PropUint32)
	test.Equals(t, uint64(123), config.PropUint64)
}

func TestEnvFloats(t *testing.T) {
	os.Setenv("PROP", "1.23")

	config := struct {
		PropFloat32 float32 `env:"PROP"`
		PropFloat64 float64 `env:"PROP"`
	}{}

	test.ErrorNil(t, Env(&config))
	test.Equals(t, float32(1.23), config.PropFloat32)
	test.Equals(t, float64(1.23), config.PropFloat64)
}

func TestEnvString(t *testing.T) {
	os.Setenv("PROP", "}D-Z2P£T!E*#zE=.gc@")

	config := struct {
		Prop string `env:"PROP"`
	}{}

	test.ErrorNil(t, Env(&config))
	test.Equals(t, "}D-Z2P£T!E*#zE=.gc@", config.Prop)
}

func TestInvalidValueForRequiredTag(t *testing.T) {
	os.Unsetenv("PROP")

	config := struct {
		Prop int `env:"PROP" required:"invalid"`
	}{}

	err := Env(&config)
	test.ErrorNotNil(t, err)
	test.Assert(t, strings.HasPrefix(err.Error(), "invalid required tag 'invalid'"))
}

func TestEnvNoEnvTag(t *testing.T) {
	config := struct {
		Prop string
	}{}

	test.ErrorNil(t, Env(&config))
}

func TestEnvRequiredWhenProvided(t *testing.T) {
	os.Setenv("PROP", "hello")

	config := struct {
		Prop string `env:"PROP" required:"true"`
	}{}

	test.ErrorNil(t, Env(&config))
	test.Equals(t, "hello", config.Prop)
}

func TestEnvRequiredWhenMissing(t *testing.T) {
	config := struct {
		Prop string `env:"MISSING_PROP" required:"true"`
	}{}

	err := Env(&config)
	test.ErrorNotNil(t, err)
}

func TestEnvNotRequiredImplicitWhenMissing(t *testing.T) {
	os.Unsetenv("PROP")

	config := struct {
		Prop string `env:"PROP"`
	}{}

	err := Env(&config)
	test.ErrorNil(t, err)
}

func TestEnvNotRequiredExplicitWhenMissing(t *testing.T) {
	os.Unsetenv("PROP")

	config := struct {
		Prop string `env:"PROP" required:"false"`
	}{}

	err := Env(&config)
	test.ErrorNil(t, err)
}

func TestInvalidConfigurationForBoolType(t *testing.T) {
	os.Setenv("PROP", "hello")

	config := struct {
		Prop bool `env:"PROP"`
	}{}

	err := Env(&config)
	test.ErrorNotNil(t, err)
	test.Assert(t, strings.HasPrefix(err.Error(), "error setting Prop"))
}

func TestInvalidConfigurationForIntType(t *testing.T) {
	os.Setenv("PROP", "hello")

	config := struct {
		Prop int `env:"PROP"`
	}{}

	err := Env(&config)
	test.ErrorNotNil(t, err)
	test.Assert(t, strings.HasPrefix(err.Error(), "error setting Prop"))
}

func TestInvalidConfigurationForUintType(t *testing.T) {
	os.Setenv("PROP", "hello")

	config := struct {
		Prop uint `env:"PROP"`
	}{}

	err := Env(&config)
	test.ErrorNotNil(t, err)
	test.Assert(t, strings.HasPrefix(err.Error(), "error setting Prop"))
}

func TestInvalidConfigurationForFloatType(t *testing.T) {
	os.Setenv("PROP", "hello")

	config := struct {
		Prop float32 `env:"PROP"`
	}{}

	err := Env(&config)
	test.ErrorNotNil(t, err)
	test.Assert(t, strings.HasPrefix(err.Error(), "error setting Prop"))
}
