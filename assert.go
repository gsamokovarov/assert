// Package assert introduces a bunch of helpers to simplify the tests writing,
// assuming you are using the default testing package.
package assert

import (
	"math"
	"reflect"
	"testing"
)

// Equal tests two objects for equality.
func Equal(t *testing.T, expected, actual interface{}) {
	if expected == nil {
		if actual == nil {
			return
		}

		val := reflect.ValueOf(actual)
		if val.IsValid() && !val.IsNil() {
			t.Fatalf("\nExpected:\n%v\ngot:\n%v", expected, actual)
		}
	}

	val := reflect.ValueOf(expected)
	typ := reflect.TypeOf(actual)

	if val.IsValid() && val.Type().ConvertibleTo(typ) {
		eval := val.Convert(typ).Interface()

		// Check for NaN. NaN is the only value that is not equal to itself.
		// That's why all the drama.
		if eval, ok := eval.(float64); ok {
			if actual := actual.(float64); ok {
				if math.IsNaN(eval) && math.IsNaN(actual) {
					return
				}
			}
		}

		if !reflect.DeepEqual(eval, actual) {
			t.Fatalf("\nExpected:\n%v\ngot:\n%v", eval, actual)
		}
	}
}

// NotEqual tests two objects for inequality.
func NotEqual(t *testing.T, expected, actual interface{}) {
	// Shortcut the nil check by abusing Go's == nil. This will catch early any
	// nil assertion early. Be it the literal nil value or the zero value of a
	// referential type.
	if expected == nil && actual == nil {
		t.Fatalf("\nExpected:\n%v\nto _not_ equal:\n%v", expected, actual)
	}

	typ := reflect.TypeOf(actual)
	val := reflect.ValueOf(expected)

	if val.IsValid() && val.Type().ConvertibleTo(typ) {
		eval := val.Convert(typ).Interface()

		// Check for NaN. NaN is the only value that is not equal to itself.
		// That's why all the drama.
		if eval, ok := eval.(float64); ok {
			if actual := actual.(float64); ok {
				if math.IsNaN(eval) && math.IsNaN(actual) {
					t.Fatalf("\nExpected:\n%v\nto not equal:\n%v", eval, actual)
				}
			}
		}

		if reflect.DeepEqual(eval, actual) {
			t.Fatalf("\nExpected:\n%v\nto not equal: %v", eval, actual)
		}
	}
}

// True stops the current test if the assertion is false.
func True(t *testing.T, assertion bool) {
	Equal(t, true, assertion)
}

// False stops the current test if the assertion is true.
func False(t *testing.T, assertion bool) {
	Equal(t, false, assertion)
}

// Nil stops the current test if the values is not nil.
func Nil(t *testing.T, v interface{}) {
	Equal(t, nil, v)
}

// NotNil stops the current test if the values is nil.
func NotNil(t *testing.T, v interface{}) {
	NotEqual(t, nil, v)
}

// Len stops the current test if the value doesn't have the expected length.
// Only arrays, chans, maps, slices and strings can have length calculated.
func Len(t *testing.T, length int, v interface{}) {
	val := reflect.Indirect(reflect.ValueOf(v))

	switch val.Kind() {
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Slice, reflect.String:
		Equal(t, length, val.Len())
	default:
		t.Fatalf("Cannot get the length of %v", val)
	}
}
