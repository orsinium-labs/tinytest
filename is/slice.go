package is

// Assert that the given slice is empty
func SliceEmpty[S ~[]I, I any](c *config, s S) {
	c.hide()
	assert(len(s) == 0, c, "%s is not empty", "%s is empty", s)
}

// Assert that the given slice contains the given element.
func SliceContains[S ~[]I, I comparable](c *config, s S, v I) {
	c.hide()
	found := false
	for _, el := range s {
		if el == v {
			found = true
			break
		}
	}
	assert(found, c, "%s does not contain %s", "%s contains %s", s, v)
}

func SliceEqual[S1 ~[]I, S2 ~[]I, I comparable](c *config, s1 S1, s2 S2) {
	c.hide()
	assert(
		len(s1) == len(s2), c,
		"len(%s) (%s) != len(%s) (%s)",
		"len(%s) (%s) == len(%s) (%s)",
		s1, len(s1), s2, len(s2),
	)
	// In relaxed mode, do not check slice elements if the length is different.
	if c.t.Failed() {
		return
	}
	for i, e1 := range s1 {
		e2 := s2[i]
		Equal(c, e1, e2)
	}
}
