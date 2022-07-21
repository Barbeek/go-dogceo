package dogceo

import (
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomImage(t *testing.T) {
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

			message, err := New().RandomImage()

			assert.Equal(t, test.err, err, "Error should be : %+v but got : %+v", test.err, err)
			assert.NotNil(t, message.Message, "image should not be nil but got %+v", message.Message)
			assert.Equal(t, test.want.Status, message.Status, "status should be %s but got %s", test.want.Status, message.Status)
		})
	}
}

func TestRandomImages(t *testing.T) {
	tests := []struct {
		name        string
		arg         string
		arrayLenght int
		want        *MessageArray
		err         error
	}{
		{
			name:        "success",
			arg:         "2",
			arrayLenght: 2,
			want: &MessageArray{
				Message: []string{},
				Status:  "success",
			},
			err: nil,
		},
		{
			name:        "error - wrong argument",
			arg:         "bad",
			arrayLenght: 1,
			want: &MessageArray{
				Message: []string{},
				Status:  "success",
			},
			err: nil,
		},
		{
			name:        "error - argument must not exceed 50",
			arg:         "300",
			arrayLenght: 0,
			want:        nil,
			err:         errors.New("numberOfImages can not exceed 50"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Log(test.name)

			messages, err := New().RandomImages(test.arg)

			assert.Equal(t, test.err, err, "Error should be : %+v but got : %+v", test.err, err)
			if messages != nil {
				assert.Len(t, messages.Message, test.arrayLenght, "messages len should be %d but is %d", test.arrayLenght, len(messages.Message))
			}
			if test.want != nil {
				assert.Equal(t, test.want.Status, messages.Status, "status should be %s but got %s", test.want.Status, messages.Status)
			}
		})
	}
}

func TestImagesByBreed(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want *MessageArray
		err  error
	}{
		{
			name: "success",
			arg:  "hound",
			want: &MessageArray{
				Message: []string{},
				Status:  "success",
			},
			err: nil,
		},
		{
			name: "error - bad argument",
			arg:  "test",
			err:  fmt.Errorf("Code (%d): {%s}", http.StatusNotFound, "\"status\":\"error\",\"message\":\"Breed not found (master breed does not exist)\",\"code\":404"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Log(test.name)

			messages, err := New().ImagesByBreed(test.arg)

			assert.Equal(t, test.err, err, "Error should be : %+v but got : %+v", test.err, err)
			if messages != nil {
				assert.NotEmpty(t, messages.Message, "messages should not be empty %+v", messages.Message)
			}
			if test.want != nil {
				assert.Equal(t, test.want.Status, messages.Status, "status should be %s but got %s", test.want.Status, messages.Status)
			}
		})
	}
}

func TestRandomImageByBreed(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want *Message
		err  error
	}{
		{
			name: "success",
			arg:  "hound",
			want: &Message{
				Message: "",
				Status:  "success",
			},
			err: nil,
		},
		{
			name: "error - bad argument",
			arg:  "test",
			err:  fmt.Errorf("Code (%d): {%s}", http.StatusNotFound, "\"status\":\"error\",\"message\":\"Breed not found (master breed does not exist)\",\"code\":404"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Log(test.name)

			messages, err := New().RandomImageByBreed(test.arg)

			assert.Equal(t, test.err, err, "Error should be : %+v but got : %+v", test.err, err)
			if messages != nil {
				assert.NotEqual(t, "", messages.Message, "messages should not be empty %+v", messages.Message)
			}
			if test.want != nil {
				assert.Equal(t, test.want.Status, messages.Status, "status should be %s but got %s", test.want.Status, messages.Status)
			}
		})
	}
}

func TestRandomImagesByBreed(t *testing.T) {
	type args struct {
		breed          string
		numberOfImages string
	}
	tests := []struct {
		name     string
		args     args
		arrayLen int
		want     *MessageArray
		err      error
	}{
		{
			name: "success",
			args: args{
				breed:          "hound",
				numberOfImages: "2",
			},
			arrayLen: 2,
			want: &MessageArray{
				Message: make([]string, 2),
				Status:  "success",
			},
			err: nil,
		},
		{
			name: "error - bad breed",
			args: args{
				breed:          "test",
				numberOfImages: "2",
			},
			arrayLen: 0,
			want:     nil,
			err:      fmt.Errorf("Code (%d): {%s}", http.StatusNotFound, "\"status\":\"error\",\"message\":\"Breed not found (master breed does not exist)\",\"code\":404"),
		},
		{
			name: "error - bad numberOfImages",
			args: args{
				breed:          "hound",
				numberOfImages: "a",
			},
			arrayLen: 1,
			want: &MessageArray{
				Message: make([]string, 1),
				Status:  "success",
			},
			err: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Log(test.name)

			messages, err := New().RandomImagesByBreed(test.args.breed, test.args.numberOfImages)

			assert.Equal(t, test.err, err, "Error should be : %+v but got : %+v", test.err, err)
			if messages != nil {
				assert.Len(t, messages.Message, test.arrayLen, "messages len should be %d but is %d", test.arrayLen, len(messages.Message))
			}
			if test.want != nil {
				assert.Equal(t, test.want.Status, messages.Status, "status should be %s but got %s", test.want.Status, messages.Status)
			}
		})
	}
}
