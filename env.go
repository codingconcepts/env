package env

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

// Setter is called for any complex struct field with an
// implementation, allowing developers to override Set
// behaviour.
type Setter interface {
	Set(string) error
}

// Set sets the fields of a struct from environment config.
// If a field is unexported or required configuration is not
// found, an error will be returned.
func Set(i interface{}) (err error) {
	v := reflect.ValueOf(i)

	// don't try to process a non-pointer value
	if v.Kind() != reflect.Ptr || v.IsNil() {
		return fmt.Errorf("%s is not a pointer", v.Kind())
	}

	v = v.Elem()
	t := reflect.TypeOf(i).Elem()

	for i := 0; i < t.NumField(); i++ {
		if err = processField(t.Field(i), v.Field(i)); err != nil {
			return
		}
	}

	return
}

// processField will lookup the "env" tag for the property
// and attempt to set it.  If not found, another check for the
// "required" tag will be performed to decided whether an error
// needs to be returned.
func processField(t reflect.StructField, v reflect.Value) (err error) {
	envTag, ok := t.Tag.Lookup("env")
	if !ok {
		// if the env tag isn't found, don't attempt to set a
		// value for the field.
		return
	}

	env, ok := os.LookupEnv(envTag)
	if !ok {
		// an env tag has been provided but a matching environment
		// variable cannot be found, determine if we should return
		// an error or if a missing variable is ok/expected.
		return processMissing(t, envTag, configTypeEnvironment)
	}

	// if the field is unexported or just not settable, bail at
	// this point because all subsequent operations will fail.
	if !v.CanSet() {
		return fmt.Errorf("field cannot be set")
	}

	// if field implements the Setter interface, invoke it now and
	// don't continue attempting to set the primative values.
	if _, ok := v.Interface().(Setter); ok {
		instance := reflect.New(t.Type.Elem())
		v.Set(instance)

		// re-assert the type with the newed-up instance and call.
		setter := v.Interface().(Setter)
		if err = setter.Set(env); err != nil {
			return errors.Wrapf(err, "error in custom setter")
		}
	}

	if err = setField(v, env); err != nil {
		return errors.Wrapf(err, "error setting %s", t.Name)
	}

	return
}

// processMissing returns an error if a required tag is found
// and is set to true.  A different error will be returned if
// the required tag was present but the value could not be parsed
// to a Boolean value.
func processMissing(t reflect.StructField, envTag string, ct configType) (err error) {
	reqTag, ok := t.Tag.Lookup("required")
	if !ok {
		// no required tag was found, this field doesn't expect
		// an env tag to be provided.
		return nil
	}

	var b bool
	if b, err = strconv.ParseBool(reqTag); err != nil {
		// the value provided for the required tag is not a valid
		// Boolean, so inform the user.
		return errors.Wrapf(err, fmt.Sprintf("invalid required tag '%s'", reqTag))
	}

	if b {
		// the value provided for the required tag is valid and is
		// set to true, so the user needs to know that a required
		// environment variable could not be found.
		return fmt.Errorf("%s %s configuration was missing", envTag, ct)
	}

	return
}

// setField determines a field's type and parses the given value
// accordingly.  An error will be returned if the field is unexported.
func setField(fieldValue reflect.Value, value string) (err error) {
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
