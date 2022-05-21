package is

import "testing"

func Equal[T comparable](t *testing.T, a T, b T) {
	t.Helper()
	if a != b {
		t.Error(a, "!=", b)
	}
}

func True(t *testing.T, expression bool) {
	t.Helper()
	if !expression {
		t.Error(expression, "is not true")
	}
}

func False(t *testing.T, expression bool) {
	t.Helper()
	if expression {
		t.Error(expression, "is not false")
	}
}
