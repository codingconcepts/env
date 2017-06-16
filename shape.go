package shape

import (
	"fmt"
	"os"
	"reflect"
	"strconv"

	"github.com/pkg/errors"
)

type configType string

const (
	configTypeEnvironment configType = "environment"
)

// Env sets the fields of a struct from environment config.
func Env(i interface{}) (err error) {
	t := reflect.TypeOf(i).Elem()
	v := reflect.ValueOf(i).Elem()

	for i := 0; i < t.NumField(); i++ {
		if err = processEnvField(t.Field(i), v.Field(i)); err != nil {
			return
		}
	}

	return
}

func processEnvField(t reflect.StructField, v reflect.Value) (err error) {
	envTag, ok := t.Tag.Lookup("env")
	if !ok {
		return
	}

	env, ok := os.LookupEnv(envTag)
	if !ok {
		return processMissing(t, envTag, configTypeEnvironment)
	}

	if err = setField(v, env); err != nil {
		return errors.Wrapf(err, "error setting %s", t.Name)
	}

	return
}

func processMissing(t reflect.StructField, envTag string, ct configType) (err error) {
	reqTag, ok := t.Tag.Lookup("required")
	if !ok {
		return nil
	}

	var b bool
	if b, err = strconv.ParseBool(reqTag); err != nil {
		return errors.Wrapf(err, fmt.Sprintf("invalid required tag '%s'", reqTag))
	}

	if b {
		return fmt.Errorf("%s %s configuration was missing", envTag, ct)
	}

	return
}

func setField(fieldValue reflect.Value, value string) error {
	if !fieldValue.CanSet() {
		return fmt.Errorf("field is unexported")
	}

	switch fieldValue.Kind() {
	case reflect.Bool:
		b, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		fieldValue.SetBool(b)

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		i, err := strconv.ParseInt(value, 0, 64)
		if err != nil {
			return err
		}
		fieldValue.SetInt(int64(i))

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		i, err := strconv.ParseUint(value, 0, 64)
		if err != nil {
			return err
		}
		fieldValue.SetUint(uint64(i))

	case reflect.Float32, reflect.Float64:
		f, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return err
		}
		fieldValue.SetFloat(f)

	case reflect.String:
		fieldValue.SetString(value)
	}

	return nil
}
