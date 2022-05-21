package is

import (
	"testing"
)

func TestEqual(t *testing.T) {
	Equal(t, 1, 1)
	Equal(t, 10.10, 10.10)
	Equal(t, "hello", "hello")

	t.Run("Sub"+t.Name(), func(sub *testing.T) {
		Equal(sub, 1, 1)
	})
}

func TestTrue(t *testing.T) {
	True(t, 10 > 1)

	t.Run("Sub"+t.Name(), func(sub *testing.T) {
		True(sub, 10 > 1)
	})
}

func TestFalse(t *testing.T) {
	False(t, 10 < 1)

	t.Run("Sub"+t.Name(), func(sub *testing.T) {
		False(sub, 10 < 1)
	})
}
