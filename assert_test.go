package assert

import (
	"errors"
	"math"
	"testing"
)

func TestEqual(t *testing.T) {
	intVar := 42
	nanVar := math.NaN()

	Equal(t, 42, intVar)
	Equal(t, nanVar, math.NaN())
}

func TestNotEqual(t *testing.T) {
	NotEqual(t, math.NaN(), math.MaxFloat32)
	NotEqual(t, nil, struct{}{})
	NotEqual(t, struct{}{}, nil)
}

func TestNil(t *testing.T) {
	var nilInterface error
	var nilPtr *struct{}
	var nilFunc func()
	var nilSlice []interface{}
	var nilMap map[interface{}]interface{}

	Nil(t, nilInterface)
	Nil(t, nilPtr)
	Nil(t, nilFunc)
	Nil(t, nilSlice)
	Nil(t, nilMap)
	Nil(t, nil)
}

func TestNotNil(t *testing.T) {
	NotNil(t, 42)
	NotNil(t, errors.New("Something Bad"))
	NotNil(t, struct{ A int }{A: 42})
	NotNil(t, &struct{ A int }{A: 42})
}

func TestError(t *testing.T) {
	Len(t, 1, []interface{}{"foo"})
	Len(t, 1, &[]interface{}{"foo"})
	Len(t, 0, []interface{}{})
}

func TestLen(t *testing.T) {
	err := errors.New("assert: error")

	Error(t, err)
	Error(t, err, "assert: error")
	Error(t, err, "assert: ", "error")
}

func TestPanic(t *testing.T) {
	Panic(t, func() {
		panic("pass")
	})
}

func TestChaining(t *testing.T) {
	err := errors.New("assert: error")

	Present(t, err).Error(t, err, "assert: error")
}
