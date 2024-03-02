package is

import (
	"fmt"
)

type T interface {
	Fail()
	FailNow()
	Helper()
	Logf(format string, args ...any)
}

type config struct {
	t    T
	fail func()
}

func (c config) hide() {
	c.t.Helper()
}

func New(t T) *config {
	return &config{t, t.FailNow}
}

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

func assert(ok bool, c *config, msg string, args ...any) {
	c.hide()
	if !ok {
		fail(c, msg, args...)
	}
}
