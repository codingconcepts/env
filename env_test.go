package env

import (
	"errors"
	"fmt"
	"math"
	"os"
	"strings"
	"testing"
	"time"
)

func TestEnvBool(t *testing.T) {
	os.Setenv("PROP", "true")

	config := struct {
		Prop bool `env:"PROP"`
	}{}

	ErrorNil(t, Set(&config))
	Equals(t, true, config.Prop)
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

	ErrorNil(t, Set(&config))
	Equals(t, int(123), config.PropInt)
	Equals(t, int8(123), config.PropInt8)
	Equals(t, int16(123), config.PropInt16)
	Equals(t, int32(123), config.PropInt32)
	Equals(t, int64(123), config.PropInt64)
}

func TestIntegerRanges(t *testing.T) {
	testCases := []struct {
		Prop8  int8
		Prop16 int16
		Prop32 int32
		Prop64 int64
	}{
		{Prop8: math.MinInt8, Prop16: math.MinInt16, Prop32: math.MinInt32, Prop64: math.MinInt64},
		{Prop8: math.MaxInt8, Prop16: math.MaxInt16, Prop32: math.MaxInt32, Prop64: math.MaxInt64},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("%v", testCase), func(t *testing.T) {
			os.Setenv("PROP8", fmt.Sprintf("%d", testCase.Prop8))
			os.Setenv("PROP16", fmt.Sprintf("%d", testCase.Prop16))
			os.Setenv("PROP32", fmt.Sprintf("%d", testCase.Prop32))
			os.Setenv("PROP64", fmt.Sprintf("%d", testCase.Prop64))

			config := struct {
				Prop8  int8  `env:"PROP8"`
				Prop16 int16 `env:"PROP16"`
				Prop32 int32 `env:"PROP32"`
				Prop64 int64 `env:"PROP64"`
			}{}

			ErrorNil(t, Set(&config))
			Equals(t, testCase.Prop8, config.Prop8)
			Equals(t, testCase.Prop16, config.Prop16)
			Equals(t, testCase.Prop32, config.Prop32)
			Equals(t, testCase.Prop64, config.Prop64)
		})
	}
}

func TestUnsignedIntegerRanges(t *testing.T) {
	testCases := []struct {
		Prop8  uint8
		Prop16 uint16
		Prop32 uint32
		Prop64 uint64
	}{
		{Prop8: 0, Prop16: 0, Prop32: 0, Prop64: 0},
		{Prop8: math.MaxUint8, Prop16: math.MaxUint16, Prop32: math.MaxUint32, Prop64: math.MaxUint64},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("%v", testCase), func(t *testing.T) {
			os.Setenv("PROP8", fmt.Sprintf("%d", testCase.Prop8))
			os.Setenv("PROP16", fmt.Sprintf("%d", testCase.Prop16))
			os.Setenv("PROP32", fmt.Sprintf("%d", testCase.Prop32))
			os.Setenv("PROP64", fmt.Sprintf("%d", testCase.Prop64))

			config := struct {
				Prop8  uint8  `env:"PROP8"`
				Prop16 uint16 `env:"PROP16"`
				Prop32 uint32 `env:"PROP32"`
				Prop64 uint64 `env:"PROP64"`
			}{}

			ErrorNil(t, Set(&config))
			Equals(t, testCase.Prop8, config.Prop8)
			Equals(t, testCase.Prop16, config.Prop16)
			Equals(t, testCase.Prop32, config.Prop32)
			Equals(t, testCase.Prop64, config.Prop64)
		})
	}
}

func TestUnsignedFloatRanges(t *testing.T) {
	testCases := []struct {
		Prop32 float32
		Prop64 float64
	}{
		{Prop32: math.SmallestNonzeroFloat32, Prop64: math.SmallestNonzeroFloat64},
		{Prop32: math.MaxFloat32, Prop64: math.MaxFloat64},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("%v", testCase), func(t *testing.T) {
			os.Setenv("PROP32", fmt.Sprintf("%g", testCase.Prop32))
			os.Setenv("PROP64", fmt.Sprintf("%g", testCase.Prop64))

			config := struct {
				Prop32 float32 `env:"PROP32"`
				Prop64 float64 `env:"PROP64"`
			}{}

			ErrorNil(t, Set(&config))
			Equals(t, testCase.Prop32, config.Prop32)
			Equals(t, testCase.Prop64, config.Prop64)
		})
	}
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

	ErrorNil(t, Set(&config))
	Equals(t, uint(123), config.PropUint)
	Equals(t, uint8(123), config.PropUint8)
	Equals(t, uint16(123), config.PropUint16)
	Equals(t, uint32(123), config.PropUint32)
	Equals(t, uint64(123), config.PropUint64)
}

func TestEnvFloats(t *testing.T) {
	os.Setenv("PROP", "1.23")

	config := struct {
		PropFloat32 float32 `env:"PROP"`
		PropFloat64 float64 `env:"PROP"`
	}{}

	ErrorNil(t, Set(&config))
	Equals(t, float32(1.23), config.PropFloat32)
	Equals(t, float64(1.23), config.PropFloat64)
}

func TestEnvString(t *testing.T) {
	os.Setenv("PROP", "}D-Z2P£T!E*#zE=.gc@")

	config := struct {
		Prop string `env:"PROP"`
	}{}

	ErrorNil(t, Set(&config))
	Equals(t, "}D-Z2P£T!E*#zE=.gc@", config.Prop)
}

func TestEnvDuration(t *testing.T) {
	os.Setenv("PROP", "1m30s")

	config := struct {
		Prop time.Duration `env:"PROP"`
	}{}

	ErrorNil(t, Set(&config))
	Equals(t, "1m30s", config.Prop.String())
}

func TestEnvUnsupportedType(t *testing.T) {
	os.Setenv("PROP", "1")

	config := struct {
		Prop chan int `env:"PROP"`
	}{}

	err := Set(&config)
	ErrorNotNil(t, err)
	Equals(t, `error setting "Prop": chan is not supported`, err.Error())
}

func TestByteSlice(t *testing.T) {
	os.Setenv("PROP", "hello")

	config := struct {
		Prop []byte `env:"PROP"`
	}{}

	err := Set(&config)
	ErrorNil(t, err)
	Equals(t, "hello", string(config.Prop))
}

func TestEnvBoolSlice(t *testing.T) {
	os.Setenv("PROPS", "true, false, true")
	config := struct {
		Items []bool `env:"PROPS"`
	}{}

	ErrorNil(t, Set(&config))
	Equals(t, []bool{true, false, true}, config.Items)
}

func TestEnvStringSlice(t *testing.T) {
	os.Setenv("PROPS", "a, b, c")
	config := struct {
		Items []string `env:"PROPS"`
	}{}

	ErrorNil(t, Set(&config))
	Equals(t, []string{"a", "b", "c"}, config.Items)
}

func TestEnvEmptyStringSlice(t *testing.T) {
	os.Setenv("PROPS", "")
	config := struct {
		Items []string `env:"PROPS"`
	}{}

	ErrorNil(t, Set(&config))
	Equals(t, 0, len(config.Items))
}

func TestEnvIntegerSlices(t *testing.T) {
	os.Setenv("PROP", "1, 2, 3")

	config := struct {
		PropInt   []int   `env:"PROP"`
		PropInt8  []int8  `env:"PROP"`
		PropInt16 []int16 `env:"PROP"`
		PropInt32 []int32 `env:"PROP"`
		PropInt64 []int64 `env:"PROP"`
	}{}

	ErrorNil(t, Set(&config))
	Equals(t, []int{1, 2, 3}, config.PropInt)
	Equals(t, []int8{1, 2, 3}, config.PropInt8)
	Equals(t, []int16{1, 2, 3}, config.PropInt16)
	Equals(t, []int32{1, 2, 3}, config.PropInt32)
	Equals(t, []int64{1, 2, 3}, config.PropInt64)
}

func TestEnvUnsignedIntegerSlices(t *testing.T) {
	os.Setenv("PROP", "1, 2, 3")

	config := struct {
		PropUint   []uint   `env:"PROP"`
		PropUint16 []uint16 `env:"PROP"`
		PropUint32 []uint32 `env:"PROP"`
		PropUint64 []uint64 `env:"PROP"`
	}{}

	ErrorNil(t, Set(&config))
	Equals(t, []uint{1, 2, 3}, config.PropUint)
	Equals(t, []uint16{1, 2, 3}, config.PropUint16)
	Equals(t, []uint32{1, 2, 3}, config.PropUint32)
	Equals(t, []uint64{1, 2, 3}, config.PropUint64)
}

func TestEnvFloatSlices(t *testing.T) {
	os.Setenv("PROP", "1.23, 2.34, 3.45")

	config := struct {
		PropFloat32 []float32 `env:"PROP"`
		PropFloat64 []float64 `env:"PROP"`
	}{}

	ErrorNil(t, Set(&config))
	Equals(t, []float32{1.23, 2.34, 3.45}, config.PropFloat32)
	Equals(t, []float64{1.23, 2.34, 3.45}, config.PropFloat64)
}

func TestEnvDurationSlice(t *testing.T) {
	os.Setenv("PROP", "1s, 2s, 4s")

	config := struct {
		Prop []time.Duration `env:"PROP"`
	}{}

	ErrorNil(t, Set(&config))
	Equals(t, []time.Duration{time.Second, time.Second * 2, time.Second * 4}, config.Prop)
}

func TestEnvUnsupportedBoolSlice(t *testing.T) {
	os.Setenv("PROPS", "true, false, true")
	config := struct {
		Items []chan int `env:"PROPS"`
	}{}

	err := Set(&config)
	ErrorNotNil(t, err)
	Assert(t, "[]chan int is not supported" == err.Error())
}

func TestEnvSetUnexportedProperty(t *testing.T) {
	os.Setenv("PROP", "hello")

	config := struct {
		prop string `env:"PROP"`
	}{}

	err := Set(&config)
	ErrorNotNil(t, err)
	Assert(t, strings.Contains(err.Error(), "field 'prop' cannot be set"))
}

func TestEnvUnsupportedPropertyWithoutTag(t *testing.T) {
	os.Setenv("PROP", "hello")

	config := struct {
		Prop  string `env:"PROP"`
		Prop2 chan int
	}{}

	ErrorNil(t, Set(&config))
	Equals(t, "hello", config.Prop)
}

func TestInvalidValueForRequiredTag(t *testing.T) {
	os.Unsetenv("PROP")

	config := struct {
		Prop int `env:"PROP" required:"invalid"`
	}{}

	err := Set(&config)
	ErrorNotNil(t, err)
	Equals(t, `invalid required tag "invalid": strconv.ParseBool: parsing "invalid": invalid syntax`, err.Error())
}

func TestEnvNoEnvTag(t *testing.T) {
	config := struct {
		Prop string
	}{}

	ErrorNil(t, Set(&config))
}

func TestEnvRequiredWhenProvided(t *testing.T) {
	os.Setenv("PROP", "hello")

	config := struct {
		Prop string `env:"PROP" required:"true"`
	}{}

	ErrorNil(t, Set(&config))
	Equals(t, "hello", config.Prop)
}

func TestEnvRequiredWhenMissing(t *testing.T) {
	config := struct {
		Prop string `env:"MISSING_PROP" required:"true"`
	}{}

	err := Set(&config)
	ErrorNotNil(t, err)
}

func TestEnvWithDefaultWhenProvided(t *testing.T) {
	os.Setenv("PROP", "goodbye")

	config := struct {
		Prop string `env:"PROP" default:"hello"`
	}{}

	ErrorNil(t, Set(&config))
	Equals(t, "goodbye", config.Prop)
}

func TestEnvWithDefaultWhenMissing(t *testing.T) {
	unsetEnvironment()

	config := struct {
		BoolProp      bool            `env:"MISSING" default:"true"`
		BoolsProp     []bool          `env:"MISSING" default:"true,false,true"`
		StringProp    string          `env:"MISSING" default:"abc"`
		StringsProp   []string        `env:"MISSING" default:"a,b,c"`
		BytesProp     []byte          `env:"MISSING" default:"abc"`
		IntProp       int             `env:"MISSING" default:"1"`
		IntsProp      []int           `env:"MISSING" default:"1,2,3"`
		Int8Prop      int8            `env:"MISSING" default:"1"`
		Int8sProp     []int8          `env:"MISSING" default:"1,2,3"`
		Int16Prop     int16           `env:"MISSING" default:"1"`
		Int16sProp    []int16         `env:"MISSING" default:"1,2,3"`
		Int32Prop     int32           `env:"MISSING" default:"1"`
		Int32sProp    []int32         `env:"MISSING" default:"1,2,3"`
		Int64Prop     int64           `env:"MISSING" default:"1"`
		Int64sProp    []int64         `env:"MISSING" default:"1,2,3"`
		UIntProp      uint            `env:"MISSING" default:"1"`
		UIntsProp     []uint          `env:"MISSING" default:"1,2,3"`
		UInt16Prop    uint16          `env:"MISSING" default:"1"`
		UInt16sProp   []uint16        `env:"MISSING" default:"1,2,3"`
		UInt32Prop    uint32          `env:"MISSING" default:"1"`
		UInt32sProp   []uint32        `env:"MISSING" default:"1,2,3"`
		UInt64Prop    uint64          `env:"MISSING" default:"1"`
		UInt64sProp   []uint64        `env:"MISSING" default:"1,2,3"`
		Float32Prop   float32         `env:"MISSING" default:"1.1"`
		Float32sProp  []float32       `env:"MISSING" default:"1.1,2.2,3.3"`
		Float64Prop   float64         `env:"MISSING" default:"1.1"`
		Float64sProp  []float64       `env:"MISSING" default:"1.1,2.2,3.3"`
		DurationProp  time.Duration   `env:"MISSING" default:"1h2m3s"`
		DurationsProp []time.Duration `env:"MISING" default:"1h,2m,3s"`
	}{}

	ErrorNil(t, Set(&config))

	Equals(t, true, config.BoolProp)
	Equals(t, []bool{true, false, true}, config.BoolsProp)
	Equals(t, "abc", config.StringProp)
	Equals(t, []string{"a", "b", "c"}, config.StringsProp)
	Equals(t, []byte("abc"), config.BytesProp)
	Equals(t, 1, config.IntProp)
	Equals(t, []int{1, 2, 3}, config.IntsProp)
	Equals(t, int8(1), config.Int8Prop)
	Equals(t, []int8{1, 2, 3}, config.Int8sProp)
	Equals(t, int16(1), config.Int16Prop)
	Equals(t, []int16{1, 2, 3}, config.Int16sProp)
	Equals(t, int32(1), config.Int32Prop)
	Equals(t, []int32{1, 2, 3}, config.Int32sProp)
	Equals(t, int64(1), config.Int64Prop)
	Equals(t, []int64{1, 2, 3}, config.Int64sProp)
	Equals(t, 1, config.IntProp)
	Equals(t, []uint{1, 2, 3}, config.UIntsProp)
	Equals(t, uint16(1), config.UInt16Prop)
	Equals(t, []uint16{1, 2, 3}, config.UInt16sProp)
	Equals(t, uint32(1), config.UInt32Prop)
	Equals(t, []uint32{1, 2, 3}, config.UInt32sProp)
	Equals(t, uint64(1), config.UInt64Prop)
	Equals(t, []uint64{1, 2, 3}, config.UInt64sProp)
	Equals(t, float32(1.1), config.Float32Prop)
	Equals(t, []float32{1.1, 2.2, 3.3}, config.Float32sProp)
	Equals(t, float64(1.1), config.Float64Prop)
	Equals(t, []float64{1.1, 2.2, 3.3}, config.Float64sProp)
	Equals(t, time.Hour+time.Minute*2+time.Second*3, config.DurationProp)
	Equals(t, []time.Duration{time.Hour, time.Minute * 2, time.Second * 3}, config.DurationsProp)
}

func TestEnvRequiredWithDefaultWhenMissing(t *testing.T) {
	unsetEnvironment()

	config := struct {
		BoolProp      bool            `env:"MISSING" required:"true" default:"true"`
		BoolsProp     []bool          `env:"MISSING" required:"true" default:"true,false,true"`
		StringProp    string          `env:"MISSING" required:"true" default:"a"`
		StringsProp   []string        `env:"MISSING" required:"true" default:"a,b,c"`
		BytesProp     []byte          `env:"MISSING" required:"true" default:"a"`
		IntProp       int             `env:"MISSING" required:"true" default:"1"`
		IntsProp      []int           `env:"MISSING" required:"true" default:"1,2,3"`
		Int8Prop      int8            `env:"MISSING" required:"true" default:"1"`
		Int8sProp     []int8          `env:"MISSING" required:"true" default:"1,2,3"`
		Int16Prop     int16           `env:"MISSING" required:"true" default:"1"`
		Int16sProp    []int16         `env:"MISSING" required:"true" default:"1,2,3"`
		Int32Prop     int32           `env:"MISSING" required:"true" default:"1"`
		Int32sProp    []int32         `env:"MISSING" required:"true" default:"1,2,3"`
		Int64Prop     int64           `env:"MISSING" required:"true" default:"1"`
		Int64sProp    []int64         `env:"MISSING" required:"true" default:"1,2,3"`
		UIntProp      uint            `env:"MISSING" required:"true" default:"1"`
		UIntsProp     []uint          `env:"MISSING" required:"true" default:"1,2,3"`
		UInt16Prop    uint16          `env:"MISSING" required:"true" default:"1"`
		UInt16sProp   []uint16        `env:"MISSING" required:"true" default:"1,2,3"`
		UInt32Prop    uint32          `env:"MISSING" required:"true" default:"1"`
		UInt32sProp   []uint32        `env:"MISSING" required:"true" default:"1,2,3"`
		UInt64Prop    uint64          `env:"MISSING" required:"true" default:"1"`
		UInt64sProp   []uint64        `env:"MISSING" required:"true" default:"1,2,3"`
		Float32Prop   float32         `env:"MISSING" required:"true" default:"1.1"`
		Float32sProp  []float32       `env:"MISSING" required:"true" default:"1.1,2.2,3.3"`
		Float64Prop   float64         `env:"MISSING" required:"true" default:"1.1"`
		Float64sProp  []float64       `env:"MISSING" required:"true" default:"1.1,2.2,3.3"`
		DurationProp  time.Duration   `env:"MISSING" required:"true" default:"1h2m3s"`
		DurationsProp []time.Duration `env:"MISSING" required:"true" default:"1h,2m,3s"`
	}{}

	ErrorNil(t, Set(&config))

	Equals(t, true, config.BoolProp)
	Equals(t, []bool{true, false, true}, config.BoolsProp)
	Equals(t, "a", config.StringProp)
	Equals(t, []string{"a", "b", "c"}, config.StringsProp)
	Equals(t, []byte("a"), config.BytesProp)
	Equals(t, 1, config.IntProp)
	Equals(t, []int{1, 2, 3}, config.IntsProp)
	Equals(t, int8(1), config.Int8Prop)
	Equals(t, []int8{1, 2, 3}, config.Int8sProp)
	Equals(t, int16(1), config.Int16Prop)
	Equals(t, []int16{1, 2, 3}, config.Int16sProp)
	Equals(t, int32(1), config.Int32Prop)
	Equals(t, []int32{1, 2, 3}, config.Int32sProp)
	Equals(t, int64(1), config.Int64Prop)
	Equals(t, []int64{1, 2, 3}, config.Int64sProp)
	Equals(t, 1, config.IntProp)
	Equals(t, []uint{1, 2, 3}, config.UIntsProp)
	Equals(t, uint16(1), config.UInt16Prop)
	Equals(t, []uint16{1, 2, 3}, config.UInt16sProp)
	Equals(t, uint32(1), config.UInt32Prop)
	Equals(t, []uint32{1, 2, 3}, config.UInt32sProp)
	Equals(t, uint64(1), config.UInt64Prop)
	Equals(t, []uint64{1, 2, 3}, config.UInt64sProp)
	Equals(t, float32(1.1), config.Float32Prop)
	Equals(t, []float32{1.1, 2.2, 3.3}, config.Float32sProp)
	Equals(t, float64(1.1), config.Float64Prop)
	Equals(t, []float64{1.1, 2.2, 3.3}, config.Float64sProp)
	Equals(t, time.Hour+time.Minute*2+time.Second*3, config.DurationProp)
	Equals(t, []time.Duration{time.Hour, time.Minute * 2, time.Second * 3}, config.DurationsProp)
}

func TestEnvCustomDelimiter(t *testing.T) {
	os.Setenv("PROP", "a b c")

	config := struct {
		Prop []string `env:"PROP" delimiter:" "`
	}{}

	ErrorNil(t, Set(&config))
	Equals(t, []string{"a", "b", "c"}, config.Prop)
}

func TestEnvNotRequiredImplicitWhenMissing(t *testing.T) {
	os.Unsetenv("PROP")

	config := struct {
		Prop string `env:"PROP"`
	}{}

	err := Set(&config)
	ErrorNil(t, err)
}

func TestEnvNotRequiredExplicitWhenMissing(t *testing.T) {
	os.Unsetenv("PROP")

	config := struct {
		Prop string `env:"PROP" required:"false"`
	}{}

	err := Set(&config)
	ErrorNil(t, err)
}

func TestInvalidConfigurationForBoolType(t *testing.T) {
	os.Setenv("PROP", "hello")

	config := struct {
		Prop bool `env:"PROP"`
	}{}

	err := Set(&config)
	ErrorNotNil(t, err)
	Equals(t, `error setting "Prop": strconv.ParseBool: parsing "hello": invalid syntax`, err.Error())
}

func TestInvalidConfigurationForIntType(t *testing.T) {
	os.Setenv("PROP", "hello")

	config := struct {
		Prop int `env:"PROP"`
	}{}

	err := Set(&config)
	ErrorNotNil(t, err)
	Equals(t, `error setting "Prop": strconv.ParseInt: parsing "hello": invalid syntax`, err.Error())
}

func TestInvalidConfigurationForUintType(t *testing.T) {
	os.Setenv("PROP", "hello")

	config := struct {
		Prop uint `env:"PROP"`
	}{}

	err := Set(&config)
	ErrorNotNil(t, err)
	Equals(t, `error setting "Prop": strconv.ParseUint: parsing "hello": invalid syntax`, err.Error())
}

func TestInvalidConfigurationForFloatType(t *testing.T) {
	os.Setenv("PROP", "hello")

	config := struct {
		Prop float32 `env:"PROP"`
	}{}

	err := Set(&config)
	ErrorNotNil(t, err)
	Equals(t, `error setting "Prop": strconv.ParseFloat: parsing "hello": invalid syntax`, err.Error())
}

func TestInvalidConfigurationForDuration(t *testing.T) {
	os.Setenv("PROP", "1hh")

	config := struct {
		Prop time.Duration `env:"PROP"`
	}{}

	err := Set(&config)
	ErrorNotNil(t, err)
	Equals(t, `error setting "Prop": time: unknown unit "hh" in duration "1hh"`, err.Error())
}

func TestEnvNonPointer(t *testing.T) {
	config := struct {
		Prop float32 `env:"PROP"`
	}{}

	err := Set(config)
	ErrorNotNil(t, err)
	Equals(t, err.Error(), "struct is not a pointer")
}

func TestEnvCustomTypeAliasedPrimitiveWithoutSetter(t *testing.T) {
	os.Setenv("PROP", "1234")

	config := struct {
		Prop myInt `env:"PROP"`
	}{}

	ErrorNil(t, Set(&config))
	Equals(t, myInt(1234), config.Prop)
}

type myInt int16

func TestEnvCustomTypeStruct(t *testing.T) {
	os.Setenv("PROP", "3h2m1s")

	config := struct {
		Timeout *configDuration `env:"PROP"`
	}{}

	ErrorNil(t, Set(&config))
	Equals(t, time.Hour*3+time.Minute*2+time.Second*1, config.Timeout.Duration)
}

func TestEnvCustomTypeStructWithError(t *testing.T) {
	os.Setenv("PROP", "3h2m1s")

	config := struct {
		Timeout *configDurationError `env:"PROP"`
	}{}

	err := Set(&config)
	ErrorNotNil(t, err)
	Assert(t, strings.HasPrefix(err.Error(), "error in custom setter"))
	Assert(t, strings.Contains(err.Error(), errConfigDurationError.Error()))
}

func TestEnvPrefixed(t *testing.T) {
	os.Setenv("PROP_PROP", "hello")

	config := struct {
		Prop string `env:"PROP"`
	}{}

	ErrorNil(t, SetPrefix(&config, "PROP_"))
	Equals(t, "hello", config.Prop)
}

type configDuration struct {
	Duration time.Duration
}

func (d *configDuration) Set(config string) (err error) {
	d.Duration, err = time.ParseDuration(config)
	return
}

type configDurationError struct {
	Duration time.Duration
}

var errConfigDurationError = errors.New("example error from custom Set code")

func (d *configDurationError) Set(config string) (err error) {
	return errConfigDurationError
}
