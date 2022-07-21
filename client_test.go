package dogceo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	t.Run("New()", func(t *testing.T) {
		t.Log("New()")
		got := New()
		want := &api{
			baseURL: baseURL,
		}
		assert.Equal(t, want, got, "Expected: %+v\n Got: %+v\n", want, got)
	})
}
