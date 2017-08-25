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

	// don't try to process a non-pointer value.
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
		return
	}

	// if the field is unexported or just not settable, bail at
	// this point because subsequent operations will fail.
	if !v.CanSet() {
		return fmt.Errorf("field '%s' cannot be set", t.Name)
	}

	// lookup the environment variable and if found, set and
	// return
	env, ok := os.LookupEnv(envTag)
	if ok {
		return setField(t, v, env)
	}

	// if the value isn't found in the environment, look for a
	// user-defined default value
	d, ok := t.Tag.Lookup("default")
	if ok {
		return setField(t, v, d)
	}

	// an env tag has been provided but a matching environment
	// variable cannot be found, determine if we should return
	// an error or if a missing variable is ok/expected.
	return processMissing(t, envTag, configTypeEnvironment)
}

func setField(t reflect.StructField, v reflect.Value, value string) (err error) {
	// if field implements the Setter interface, invoke it now and
	// don't continue attempting to set the primative values.
	if _, ok := v.Interface().(Setter); ok {
		instance := reflect.New(t.Type.Elem())
		v.Set(instance)

		// re-assert the type with the newed-up instance and call.
		setter := v.Interface().(Setter)
		if err = setter.Set(value); err != nil {
			return errors.Wrapf(err, "error in custom setter")
		}
		return
	}

	// if the given type is a slice, create a slice and return,
	// otherwise, we're dealing with a primitive type
	if v.Kind() == reflect.Slice {
		return setSlice(t, v, value)
	}

	if err = setBuiltInField(v, value); err != nil {
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
