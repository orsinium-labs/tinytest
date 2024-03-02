package is_test

import (
	"testing"

	"github.com/orsinium-labs/tinytest/is"
)

func TestOk(t *testing.T) {
	is.Equal(t, 1, 1)
	is.Equal(t, true, true)
	is.Equal(t, "ab", "ab")

	is.NotEqual(t, 1, 2)
	is.NotEqual(t, "ac", "ab")
	is.NotEqual(t, "", "ab")

	is.True(t, true)
	is.False(t, false)

	is.Zero(t, 0)
	is.Zero(t, false)
	is.Zero(t, "")

	is.SliceEqual(t, []int{1, 2, 3}, []int{1, 2, 3})

	is.Panics(t, func() { panic("oh no") })

	is.NotPanics(t, func() { _ = 1 + 2 })
}
