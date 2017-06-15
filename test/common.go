package test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"

	"github.com/Sirupsen/logrus"
)

// Assert fails the test if the condition is false.
func Assert(tb testing.TB, condition bool, v ...interface{}) {
	if !condition {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("%s:%d:\n\n", append([]interface{}{filepath.Base(file), line}, v...)...)
		tb.FailNow()
	}
}

// ErrorNil fails the test if an err is not nil.
func ErrorNil(tb testing.TB, err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("%s:%d: unexpected error: %s\n\n", filepath.Base(file), line, err.Error())
		tb.FailNow()
	}
}

// ErrorNotNil fails the test if an err is not nil.
func ErrorNotNil(tb testing.TB, err error) {
	if err == nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("%s:%d: expected error but got none\n\n", filepath.Base(file), line)
		tb.FailNow()
	}
}

// Equals fails the test if expected is not equal to actual.
func Equals(tb testing.TB, expected, actual interface{}) {
	if !reflect.DeepEqual(expected, actual) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("%s:%d:\n\n\texp: %#v\n\n\tgot: %#v\n\n", filepath.Base(file), line, expected, actual)
		tb.FailNow()
	}
}

// AnyLogEntryContainsMessage fails if the log entries message
// does not match the expected value.
func AnyLogEntryContainsMessage(tb testing.TB, expectedValue string, logEntries []*logrus.Entry) {
	if len(logEntries) == 0 {
		tb.Fatal("log entry expected")
	}

	for _, logEntry := range logEntries {
		if logEntry.Message == expectedValue {
			return
		}
	}

	// if we get here, none of the log entries contain the
	// the message we were looking for.
	tb.Fatalf("log entry expected: '%s'", expectedValue)
}

// LogEntryContainsField fails if the log entry does not contain a field
// with the given key.
func LogEntryContainsField(tb testing.TB, key string, expectedValue string, logEntry *logrus.Entry) {
	if logEntry == nil {
		tb.Fatal("log entry expected")
	}

	actual := fmt.Sprintf("%v", logEntry.Data[key])
	if actual != expectedValue {
		tb.Fatalf("log entry expected: '%s' but got: '%s'", expectedValue, actual)
	}
}

// AnyLogEntryContainsField fails if no log entry is found to contain a field
// with the given key.
func AnyLogEntryContainsField(tb testing.TB, key string, expectedValue interface{}, logEntries []*logrus.Entry) {
	if len(logEntries) == 0 {
		tb.Fatal("log entries expected")
	}

	for _, logEntry := range logEntries {
		if expectedValue == logEntry.Data[key] {
			return
		}
	}

	// if we get here, none of the given entries contain
	// the value we were looking for
	tb.Fatalf("log entry expected: '%s' but none found", expectedValue)
}

// MakeGetRequest performs a GET request using a given handler.
func MakeGetRequest(handler http.Handler, uri string) (resp *httptest.ResponseRecorder) {
	req, _ := http.NewRequest(http.MethodGet, uri, nil)
	resp = httptest.NewRecorder()
	handler.ServeHTTP(resp, req)

	return
}
