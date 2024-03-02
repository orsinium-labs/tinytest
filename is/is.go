package is

import (
	"fmt"
	"testing"
)

func fail(t *testing.T, msg string, args ...any) {
	t.Helper()
	fargs := make([]any, len(args))
	for i, arg := range args {
		farg := fmt.Sprintf("%#v", arg)
		if len(farg) > 40 {
			farg = fmt.Sprintf("%v", arg)
			if len(farg) > 40 {
				farg = farg[:20] + "<…>" + farg[len(farg)-20:]
			}
		}
		fargs[i] = farg
	}
	t.Fatalf(msg, fargs...)
}

func assert(ok bool, t *testing.T, msg string, args ...any) {
	t.Helper()
	if !ok {
		fail(t, msg, args...)
	}
}

func Equal[V comparable](t *testing.T, a, b V) {
	t.Helper()
	assert(a == b, t, "%s must be equal to %s", a, b)
}

func NotEqual[V comparable](t *testing.T, a, b V) {
	t.Helper()
	assert(a != b, t, "%s must not be equal to %s", a, b)
}

func True(t *testing.T, ok bool) {
	t.Helper()
	assert(ok, t, "must be true")
}

func False(t *testing.T, ok bool) {
	t.Helper()
	assert(!ok, t, "must be false")
}

func Zero[V comparable](t *testing.T, v V) {
	t.Helper()
	assert(v == *new(V), t, "%s must be zero", v)
}

func Panics(t *testing.T, f func()) {
	t.Helper()
	defer func() {
		_ = recover()
	}()
	f()
	fail(t, "%s must panic", f)
}

func NotPanics(t *testing.T, f func()) {
	t.Helper()
	defer func() {
		t.Helper()
		p := recover()
		assert(p == nil, t, "%s must not panic", f)
	}()
	f()
}
