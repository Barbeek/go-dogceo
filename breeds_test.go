package dogceo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListBreeds(t *testing.T) {
	tests := []struct {
		name string
		want *Message
		err  error
	}{
		{
			name: "success",
			want: &Message{
				Status: "success",
			},
			err: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Log(test.name)

			message, err := New().ListBreeds()

			assert.Equal(t, test.err, err, "Error should be : %+v but got : %+v", test.err, err)
			assert.NotNil(t, message.Message, "image should not be nil but got %+v", message.Message)
			assert.Equal(t, test.want.Status, message.Status, "status should be %s but got %s", test.want.Status, message.Status)
		})
	}
}
