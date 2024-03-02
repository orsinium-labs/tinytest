package is

import (
	"fmt"
)

// T is the current test state manager.
//
// T is implemented by [testing.T], [testing.B], and [testing.F].
type T interface {
	// Fail marks the test as failed but continues execution.
	Fail()

	// FailNow marks the test as failed and immediately interrupts execution.
	FailNow()

	// True if the test is already marked as failed.
	Failed() bool

	// Helper marks a function as a helper function.
	//
	// Helper functions are not included in tracebacks.
	Helper()

	// Logf is like [fmt.Sprintf] and used in tests for error messages and traces.
	Logf(format string, args ...any)
}

// config controls how assetions behave.
type config struct {
	// t is the current test state manager.
	t T
	// fail marks the test as failed an possibly stops its execution.
	fail func()
	// not reverses the checked condition
	not bool
}

// hide the current frame drom the traceback.
//
// Must be called from every assertion function.
func (c config) hide() {
	c.t.Helper()
}

// func New(t T) *config {
// 	return &config{t, t.FailNow}
// }

func NewRelaxed(t T) *config {
	return &config{t, t.Fail, false}
}

func Not(c *config) *config {
	r := *c
	r.not = true
	return &r
}

// fail formats an logs the given error message and possibly interrupts the test.
func fail(c *config, msg string, args ...any) {
	c.hide()
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
	c.t.Logf(msg, fargs...)
	c.fail()
}

// assert marks the test as failed if the given condition is false
func assert(ok bool, c *config, msg, msgNot string, args ...any) {
	c.hide()
	if c.not {
		if ok {
			fail(c, msgNot, args...)
		}
	} else {
		if !ok {
			fail(c, msg, args...)
		}
	}
}
