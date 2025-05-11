package testutils

import (
	"reflect"
	"testing"

	"github.com/castlele/goaccounter/server/utils"
)

type testSuite struct {
	*testing.T
}

func NewSuite(t *testing.T) *testSuite {
	return &testSuite{t}
}

func (ts *testSuite) Parallel(name string, fun func(t *testing.T)) {
	ts.Run(name, func(t *testing.T) {
		t.Parallel()

		fun(t)
	})
}

func (ts *testSuite) AssertStatusCode(expected, actual int) {
	ts.AssertEquals(
		expected == actual,
		"Invalid status code, expected %d, but got: %d",
		expected,
		actual,
	)
}

func (ts *testSuite) AssertError(err error) {
	condition := err != nil

	ts.AssertEquals(condition, "Error is nil, while expecting it")
}

func (ts *testSuite) AssertNil(obj any) {
	if !isNil(obj) {
		ts.Errorf("Expected nil, but got: %v", obj)
	}
}

func (ts *testSuite) AssertNoError(err error) {
	condition := err == nil
	ts.AssertEquals(
		condition,
		"Error isn't nil: %s",
		utils.LazyTernary(
			condition,
			func() string { return "" },
			func() string { return err.Error() },
		),
	)
}

func (ts *testSuite) AssertEquals(condition bool, message string, args ...any) {
	if condition {
		return
	}

	ts.Errorf(message, args...)
}

// Helper function to properly check for nil in all cases
func isNil(obj any) bool {
	if obj == nil {
		return true
	}

	val := reflect.ValueOf(obj)
	switch val.Kind() {
	case reflect.Ptr, reflect.Slice, reflect.Map, reflect.Chan, reflect.Func, reflect.Interface:
		return val.IsNil()
	default:
		return false
	}
}
