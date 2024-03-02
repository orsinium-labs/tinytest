package is

func Equal[V comparable](c *config, a, b V) {
	c.hide()
	assert(a == b, c, "%s must be equal to %s", a, b)
}

func NotEqual[V comparable](c *config, a, b V) {
	c.hide()
	assert(a != b, c, "%s must not be equal to %s", a, b)
}

func True(c *config, ok bool) {
	c.hide()
	assert(ok, c, "must be true")
}

func False(c *config, ok bool) {
	c.hide()
	assert(!ok, c, "must be false")
}

func Zero[V comparable](c *config, v V) {
	c.hide()
	assert(v == *new(V), c, "%s must be zero", v)
}

func Panic(c *config, f func()) {
	c.hide()
	defer func() {
		_ = recover()
	}()
	f()
	fail(c, "%s must panic", f)
}

func NotPanic(c *config, f func()) {
	c.hide()
	defer func() {
		c.hide()
		p := recover()
		assert(p == nil, c, "%s must not panic", f)
	}()
	f()
}
