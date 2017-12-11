package assert

import "testing"

// Diff is a function that stops the execution of a test with a message that is
// supposed to help to differentiate the expected and the actually got values
// of an assertion. It can be swapped for different implementations.
func Diff(t *testing.T, positive bool, expected, actual interface{}) {
	if positive {
		t.Fatalf("\nExpected:\n%v\ngot:\n%v", expected, actual)
	} else {
		t.Fatalf("\nExpected:\n%v\nto not equal:\n%v", expected, actual)
	}
}
