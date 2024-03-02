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

func SliceContains[S ~[]I, I comparable](t *testing.T, s S, v I) {
	t.Helper()
	for _, el := range s {
		if el == v {
			return
		}
	}
	fail(t, "%s must contain %s", s, v)
}

func SliceNotContains[S ~[]I, I comparable](t *testing.T, s S, v I) {
	t.Helper()
	for _, el := range s {
		if el == v {
			fail(t, "%s must not contain %s", s, v)
		}
	}
}

func SliceEqual[S1 ~[]I, S2 ~[]I, I comparable](t *testing.T, s1 S1, s2 S2) {
	t.Helper()
	assert(
		len(s1) == len(s2), t,
		"len(%s) (%s) must be equal to len(%s) (%s)",
		s1, s2, len(s1), len(s2),
	)
	for i, e1 := range s1 {
		e2 := s2[i]
		Equal(t, e1, e2)
	}
}
