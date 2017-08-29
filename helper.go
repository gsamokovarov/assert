package assert

import "testing"

// Marks a function as a testing helper. Works only on Go 1.9 and above. See
// https://golang.org/pkg/testing/#T.Helper for more information.
var markAsHelper func(t *testing.T)

func init() {
	if markAsHelper == nil {
		markAsHelper = func(t *testing.T) {
			// Noop	for Go 1.8 and below.
		}
	}
}
