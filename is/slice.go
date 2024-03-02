package is

import "testing"

func SliceEmpty[S ~[]I, I any](t *testing.T, s S) {
	t.Helper()
	assert(len(s) == 0, t, "%s must be empty", s)
}

func SliceNotEmpty[S ~[]I, I any](t *testing.T, s S) {
	t.Helper()
	assert(len(s) != 0, t, "%s must not be empty", s)
}
