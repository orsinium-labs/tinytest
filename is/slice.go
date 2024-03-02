package is

// Assert that the given slice is empty
func SliceEmpty[S ~[]I, I any](c *config, s S) {
	c.hide()
	assert(len(s) == 0, c, "%s must be empty", s)
}

// Assert that the given slice is not empty
func SliceNotEmpty[S ~[]I, I any](c *config, s S) {
	c.hide()
	assert(len(s) != 0, c, "%s must not be empty", s)
}

// Assert that the given slice contains the given element.
func SliceContains[S ~[]I, I comparable](c *config, s S, v I) {
	c.hide()
	for _, el := range s {
		if el == v {
			return
		}
	}
	fail(c, "%s must contain %s", s, v)
}

// Assert that the given slice does not contain the given element.
func SliceNotContains[S ~[]I, I comparable](c *config, s S, v I) {
	c.hide()
	for _, el := range s {
		if el == v {
			fail(c, "%s must not contain %s", s, v)
		}
	}
}

func SliceEqual[S1 ~[]I, S2 ~[]I, I comparable](c *config, s1 S1, s2 S2) {
	c.hide()
	assert(
		len(s1) == len(s2), c,
		"len(%s) (%s) must be equal to len(%s) (%s)",
		s1, s2, len(s1), len(s2),
	)
	for i, e1 := range s1 {
		e2 := s2[i]
		Equal(c, e1, e2)
	}
}
