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
	markAsHelper(t)

	if isNil(expected) || isNil(actual) {
		if isNil(expected) && isNil(actual) {
			return
		}

		t.Fatalf("\nExpected:\n%v\ngot:\n%v", expected, actual)
	}

	val := reflect.ValueOf(expected)
	typ := reflect.TypeOf(actual)

	if val.Type().ConvertibleTo(typ) {
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
	markAsHelper(t)

	// Shortcut the nil check by abusing Go's == nil. This will catch early any
	// nil assertion early. Be it the literal nil value or the zero value of a
	// referential type.
	if isNil(expected) || isNil(actual) {
		if isNil(expected) && isNil(actual) {
			t.Fatalf("\nExpected:\n%v\nto _not_ equal:\n%v", expected, actual)
		}

		return
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
	markAsHelper(t)

	Equal(t, true, assertion)
}

// False stops the current test if the assertion is true.
func False(t *testing.T, assertion bool) {
	markAsHelper(t)

	Equal(t, false, assertion)
}

// Nil stops the current test if the values is not nil.
func Nil(t *testing.T, v interface{}) {
	markAsHelper(t)

	Equal(t, nil, v)
}

// NotNil stops the current test if the values is nil.
func NotNil(t *testing.T, v interface{}) {
	markAsHelper(t)

	NotEqual(t, nil, v)
}

// Len stops the current test if the value doesn't have the expected length.
// Only arrays, chans, maps, slices and strings can have length calculated.
func Len(t *testing.T, length int, v interface{}) {
	markAsHelper(t)

	val := reflect.Indirect(reflect.ValueOf(v))

	switch val.Kind() {
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Slice, reflect.String:
		Equal(t, length, val.Len())
	default:
		t.Fatalf("Cannot get the length of %v", val)
	}
}

func isNil(v interface{}) bool {
	if v == nil {
		return true
	}

	vv, ok := v.(reflect.Value)
	if !ok {
		vv = reflect.ValueOf(v)
	}

	switch vv.Kind() {
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr, reflect.Interface, reflect.Slice:
		return vv.IsNil()
	}

	return false

}
