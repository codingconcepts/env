package env

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// setField determines a field's type and parses the given value
// accordingly.  An error will be returned if the field is unexported.
func setBuiltInField(fieldValue reflect.Value, value string) (err error) {
	switch fieldValue.Kind() {
	case reflect.Bool:
		return setBool(fieldValue, value)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return setInt(fieldValue, value)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return setUint(fieldValue, value)
	case reflect.Float32, reflect.Float64:
		return setFloat(fieldValue, value)
	case reflect.String:
		return setString(fieldValue, value)
	default:
		return fmt.Errorf("%s is not supported", fieldValue.Kind())
	}
}

func setBool(fieldValue reflect.Value, value string) (err error) {
	var b bool
	if b, err = strconv.ParseBool(value); err != nil {
		return err
	}

	fieldValue.SetBool(b)
	return
}

func setInt(fieldValue reflect.Value, value string) (err error) {
	if fieldValue.Type() == reflect.TypeOf((*time.Duration)(nil)).Elem() {
		return setDuration(fieldValue, value)
	}

	var i int64
	if i, err = strconv.ParseInt(value, 0, 64); err != nil {
		return err
	}

	fieldValue.SetInt(int64(i))
	return
}

func setUint(fieldValue reflect.Value, value string) (err error) {
	var i uint64
	if i, err = strconv.ParseUint(value, 0, 64); err != nil {
		return err
	}

	fieldValue.SetUint(uint64(i))
	return
}

func setFloat(fieldValue reflect.Value, value string) (err error) {
	var f float64
	if f, err = strconv.ParseFloat(value, 64); err != nil {
		return err
	}

	fieldValue.SetFloat(f)
	return
}

func setDuration(fieldValue reflect.Value, value string) (err error) {
	var d time.Duration
	if d, err = time.ParseDuration(value); err != nil {
		return err
	}

	fieldValue.SetInt(d.Nanoseconds())
	return
}

func setString(fieldValue reflect.Value, value string) (err error) {
	fieldValue.SetString(value)
	return
}

func setSlice(t reflect.StructField, v reflect.Value, value string) (err error) {
	// allow the user to provide their own delimiter, falling back to a
	// comma if one isn't provided.
	delimiter := getDelimiter(t)
	rawValues := split(value, delimiter)

	sliceValue, err := makeSlice(v, len(rawValues))
	if err != nil {
		return
	}

	populateSlice(sliceValue, rawValues)
	v.Set(sliceValue)

	return
}

func makeSlice(v reflect.Value, n int) (slice reflect.Value, err error) {
	switch v.Type() {
	case reflect.TypeOf([]string{}):
		slice = reflect.MakeSlice(reflect.TypeOf([]string{}), n, n)
	case reflect.TypeOf([]bool{}):
		slice = reflect.MakeSlice(reflect.TypeOf([]bool{}), n, n)
	case reflect.TypeOf([]int{}):
		slice = reflect.MakeSlice(reflect.TypeOf([]int{}), n, n)
	case reflect.TypeOf([]int8{}):
		slice = reflect.MakeSlice(reflect.TypeOf([]int8{}), n, n)
	case reflect.TypeOf([]int16{}):
		slice = reflect.MakeSlice(reflect.TypeOf([]int16{}), n, n)
	case reflect.TypeOf([]int32{}):
		slice = reflect.MakeSlice(reflect.TypeOf([]int32{}), n, n)
	case reflect.TypeOf([]int64{}):
		slice = reflect.MakeSlice(reflect.TypeOf([]int64{}), n, n)
	case reflect.TypeOf([]uint{}):
		slice = reflect.MakeSlice(reflect.TypeOf([]uint{}), n, n)
	case reflect.TypeOf([]uint8{}):
		slice = reflect.MakeSlice(reflect.TypeOf([]uint8{}), n, n)
	case reflect.TypeOf([]uint16{}):
		slice = reflect.MakeSlice(reflect.TypeOf([]uint16{}), n, n)
	case reflect.TypeOf([]uint32{}):
		slice = reflect.MakeSlice(reflect.TypeOf([]uint32{}), n, n)
	case reflect.TypeOf([]uint64{}):
		slice = reflect.MakeSlice(reflect.TypeOf([]uint64{}), n, n)
	case reflect.TypeOf([]float32{}):
		slice = reflect.MakeSlice(reflect.TypeOf([]float32{}), n, n)
	case reflect.TypeOf([]float64{}):
		slice = reflect.MakeSlice(reflect.TypeOf([]float64{}), n, n)
	case reflect.TypeOf([]time.Duration{}):
		slice = reflect.MakeSlice(reflect.TypeOf([]time.Duration{}), n, n)
	default:
		err = fmt.Errorf("%v is not supported", v.Type())
	}
	return
}

func populateSlice(sliceValue reflect.Value, rawItems []string) {
	for i, item := range rawItems {
		setBuiltInField(sliceValue.Index(i), item)
	}
}

func split(value string, delimeter string) (out []string) {
	out = strings.Split(value, delimeter)
	for i, s := range out {
		out[i] = strings.Trim(s, " ")
	}

	return
}

func getDelimiter(t reflect.StructField) (d string) {
	if d, ok := t.Tag.Lookup("delimiter"); ok {
		return d
	}
	return ","
}
