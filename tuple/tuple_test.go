package tuple

import (
	"testing"
)

func TestTuple_IsVector(t *testing.T) {

	t.Run("is vector", func(t *testing.T) {
		testVector := NewVector(4.3, -4.2, 3.1)
		got := testVector.IsVector()
		want := true

		if got != want {
			t.Errorf(" this is not a vector")
		}
	})
}
