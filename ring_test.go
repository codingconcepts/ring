package ring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	cases := []struct {
		name        string
		size        int
		shouldPanic bool
	}{
		{
			name:        "zero size",
			size:        0,
			shouldPanic: true,
		},
		{
			name:        "non-zero size",
			size:        1,
			shouldPanic: false,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			r, panicked := testNew[int](t, c.size)
			assert.Equal(t, c.shouldPanic, panicked)

			if panicked {
				return
			}

			// Fill the buffer.
			for i := 0; i < c.size; i++ {
				r.Add(i)
			}

			slice := r.Slice()
			assert.Equal(t, c.size, len(slice))
		})
	}
}

func TestAddSlice(t *testing.T) {
	r := New[int](5)

	for i := 0; i < 10; i++ {
		r.Add(i)
	}

	assert.Equal(t, []int{5, 6, 7, 8, 9}, r.Slice())
}

func testNew[T any](t *testing.T, size int) (ring *Ring[T], panics bool) {
	t.Helper()

	defer func() {
		if r := recover(); r != nil {
			panics = true
		}
	}()

	ring = New[T](size)

	return
}
