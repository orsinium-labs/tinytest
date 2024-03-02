package is

// Assert that the two given comparable values are equal.
//
// Use [SliceEqual] to compare slices.
func Equal[V comparable](c *config, a, b V) {
	c.hide()
	assert(a == b, c, "%s != %s", "%s == %s", a, b)
}

// Assert that the given condition is true.
func True(c *config, ok bool) {
	c.hide()
	assert(ok, c, "is not true", "is true")
}

// Assert that the given value is the default value for the given type.
func Default[V comparable](c *config, v V) {
	c.hide()
	assert(v == *new(V), c, "%s is not zero", "%s is zero", v)
}

// Assert that the given pointer is nil.
func NilPointer[V any](c *config, v *V) {
	assert(v == nil, c, "%s is not nil", "%s is nil", v)
}

// Assert that the given error is not nil.
func Err(c *config, err error) {
	assert(err != nil, c, "%s is nil error", "%s is error", err)
}

// Assert that the given function panics.
func Panic(c *config, f func()) {
	c.hide()
	defer func() {
		c.hide()
		p := recover()
		assert(p != nil, c, "%s does not panic", "%s panics", f)
	}()
	f()
}
