package is_test

import (
	"testing"

	"github.com/orsinium-labs/tinytest/is"
)

func TestOk(t *testing.T) {
	c := is.NewRelaxed(t)

	is.Equal(c, 1, 1)
	is.Equal(c, true, true)
	is.Equal(c, "ab", "ab")

	is.Equal(is.Not(c), 1, 2)
	is.Equal(is.Not(c), "ac", "ab")
	is.Equal(is.Not(c), "", "ab")

	is.True(c, true)
	is.True(is.Not(c), false)

	is.Zero(c, 0)
	is.Zero(c, false)
	is.Zero(c, "")

	is.SliceEqual(c, []int{1, 2, 3}, []int{1, 2, 3})

	is.Panic(c, func() { panic("oh no") })
	is.Panic(is.Not(c), func() { _ = 1 + 2 })
}
