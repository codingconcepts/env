package env

import (
	"reflect"
	"testing"
)

// Assert fails the test if the condition is false.
func Assert(tb testing.TB, condition bool) {
	tb.Helper()
	if !condition {
		tb.Fatal("assertion failed")
	}
}

// ErrorNil fails the test if an err is not nil.
func ErrorNil(tb testing.TB, err error) {
	tb.Helper()
	if err != nil {
		tb.Fatalf("unexpected error: %s", err.Error())
	}
}

// ErrorNotNil fails the test if an err is not nil.
func ErrorNotNil(tb testing.TB, err error) {
	tb.Helper()
	if err == nil {
		tb.Fatalf("\nexpected error but got none")
	}
}

// Equals fails the test if expected is not equal to actual.
func Equals(tb testing.TB, exp, act interface{}) {
	tb.Helper()
	if !reflect.DeepEqual(exp, act) {
		tb.Fatalf("\nexp:\t%[1]v (%[1]T)\ngot:\t%[2]v (%[2]T)", exp, act)
	}
}
