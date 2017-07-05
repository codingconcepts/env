package env

import (
	"reflect"
	"strconv"
	"time"
)

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
