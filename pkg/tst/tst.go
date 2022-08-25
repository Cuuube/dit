package tst

import (
	"reflect"
	"testing"
)

type Utils struct {
	*testing.T
}

func New(t *testing.T) *Utils {
	return &Utils{t}
}

func (t *Utils) MustEqual(a any, b any) {
	if !reflect.DeepEqual(a, b) {
		t.Fail()
	}
}

func (t *Utils) MustNotEqual(a any, b any) {
	if reflect.DeepEqual(a, b) {
		t.Fail()
	}
}

func (t *Utils) MustIsNil(a any) {
	if !reflect.ValueOf(a).IsNil() {
		t.Fail()
	}
}

func (t *Utils) MustNotNil(a any) {
	if reflect.ValueOf(a).IsNil() {
		t.Fail()
	}
}

func (t *Utils) MustIsEmpty(a []any) {
	if len(a) != 0 {
		t.Fail()
	}
}

func (t *Utils) MustNotEmpty(a []any) {
	if len(a) == 0 {
		t.Fail()
	}
}

func MustEqual[T any](t *testing.T, a T, b T) {
	if !reflect.DeepEqual(a, b) {
		t.Fail()
	}
}

func MustNotEqual[T any](t *testing.T, a T, b T) {
	if reflect.DeepEqual(a, b) {
		t.Fail()
	}
}

func MustIsNil[T any](t *testing.T, a T) {
	if !reflect.ValueOf(a).IsNil() {
		t.Fail()
	}
}

func MustNotNil[T any](t *testing.T, a T) {
	if reflect.ValueOf(a).IsNil() {
		t.Fail()
	}
}

func MustIsEmpty[T any](t *testing.T, a []T) {
	if len(a) != 0 {
		t.Fail()
	}
}

func MustNotEmpty[T any](t *testing.T, a []T) {
	if len(a) == 0 {
		t.Fail()
	}
}
