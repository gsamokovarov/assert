// +build go1.9

package assert

import "testing"

func init() {
	// Use Go 1.9 t.Helper marker.
	markAsHelper = (*testing.T).Helper
}
