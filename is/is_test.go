package is_test

import (
	"testing"

	"github.com/orsinium-labs/tinytest/is"
)

func TestOk(t *testing.T) {
	is.Equal(t, 1, 1)
	is.Equal(t, 1, 2)
}
