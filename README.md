# Assert

Assert is a minimal assertion library build on top of Go-lang builtin `testing`
package.

## Usage

The library exposes the following assertions:

```go
func Equal(t *testing.T, expected, actual interface{})
func NotEqual(t *testing.T, expected, actual interface{})
func True(t *testing.T, assertion bool)
func False(t *testing.T, assertion bool)
func Nil(t *testing.T, v interface{})
func NotNil(t *testing.T, v interface{})
func Len(t *testing.T, length int, v interface{})
```

This is how they look in action:

```go
func TestFindObject(t *testing.T) {
	obj, err := dbtest.CreateObject(store, nil, nil)
	assert.Nil(t, err)

	object, err := store.FindObject(a.ID, nil)
	assert.Nil(t, err)
	assert.NotNil(t, object)

	assert.Equal(t, a.ID, object.ID)
}
```

## Why

While the builtin `testing` package provides is minimal, it provides everything
needed to define and run tests. However repetitive tasks like checking for
errors and value presence get boring and tedious to write. This leads to noisy
tests with inconsistent error messages.

A simple assertion library like the one above will let you write concise tests
with less copy-pasted noise.
