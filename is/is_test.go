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

	is.NotEqual(c, 1, 2)
	is.NotEqual(c, "ac", "ab")
	is.NotEqual(c, "", "ab")

	is.True(c, true)
	is.False(c, false)

	is.Zero(c, 0)
	is.Zero(c, false)
	is.Zero(c, "")

	is.SliceEqual(c, []int{1, 2, 3}, []int{1, 2, 3})

	is.Panic(c, func() { panic("oh no") })

	is.NotPanic(c, func() { _ = 1 + 2 })
}
