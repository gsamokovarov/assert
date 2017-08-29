// +build !go1.9

package assert

import "testing"

// Marks a function as a testing helper. Works only on Go 1.9 and above. See
// https://golang.org/pkg/testing/#T.Helper for more information.
func markAsHelper(t *testing.T) {
	// Noop	for Go 1.8 and below.
}
