package is

func Equal[V comparable](c *config, a, b V) {
	c.hide()
	assert(a == b, c, "%s != %s", "%s == %s", a, b)
}

func True(c *config, ok bool) {
	c.hide()
	assert(ok, c, "is not true", "is true")
}

func Zero[V comparable](c *config, v V) {
	c.hide()
	assert(v == *new(V), c, "%s is not zero", "%s is zero", v)
}

func Panic(c *config, f func()) {
	c.hide()
	defer func() {
		c.hide()
		p := recover()
		assert(p != nil, c, "%s does not panic", "%s panics", f)
	}()
	f()
}
