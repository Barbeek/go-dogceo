package dogceo

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListSubBreeds(t *testing.T) {
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
				Status: "success",
			},
			err: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Log(test.name)

			message, err := New().ListSubBreeds(test.arg)

			if !assert.Equal(t, test.err, err, "Error should be : %+v but got : %+v", test.err, err) {
				return
			}
			if !assert.NotNil(t, message.Message, "image should not be nil but got %+v", message.Message) {
				return
			}
			if !assert.Equal(t, test.want.Status, message.Status) {
				return
			}
		})
	}
}

func TestImagesBySubBreed(t *testing.T) {
	type args struct {
		breed    string
		subbreed string
	}
	tests := []struct {
		name string
		args args
		want *MessageArray
		err  error
	}{
		{
			name: "success",
			args: args{
				breed:    "hound",
				subbreed: "afghan",
			},
			want: &MessageArray{
				Status: "success",
			},
			err: nil,
		},
		{
			name: "error - bad breed",
			args: args{
				breed:    "test",
				subbreed: "afghan",
			},
			want: nil,
			err:  fmt.Errorf("Code (%d): {%s}", http.StatusNotFound, "\"status\":\"error\",\"message\":\"Breed not found (master breed does not exist)\",\"code\":404"),
		},
		{
			name: "error - bad subbreed",
			args: args{
				breed:    "hound",
				subbreed: "test",
			},
			want: nil,
			err:  fmt.Errorf("Code (%d): {%s}", http.StatusNotFound, "\"status\":\"error\",\"message\":\"Breed not found (sub breed does not exist)\",\"code\":404"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Log(test.name)

			messages, err := New().ImagesBySubBreed(test.args.breed, test.args.subbreed)

			if !assert.Equal(t, test.err, err, "Error should be : %+v but got : %+v", test.err, err) {
				return
			}
			if messages != nil && !assert.NotEqual(t, "", messages.Message, "messages should not be empty %+v", messages.Message) {
				return
			}
			if test.want != nil && !assert.Equal(t, test.want.Status, messages.Status) {
				return
			}
		})
	}
}

func TestRandomImageBySubBreed(t *testing.T) {
	type args struct {
		breed    string
		subbreed string
	}
	tests := []struct {
		name string
		args args
		want *Message
		err  error
	}{
		{
			name: "success",
			args: args{
				breed:    "hound",
				subbreed: "afghan",
			},
			want: &Message{
				Status: "success",
			},
			err: nil,
		},
		{
			name: "error - bad breed",
			args: args{
				breed:    "test",
				subbreed: "afghan",
			},
			want: nil,
			err:  fmt.Errorf("Code (%d): {%s}", http.StatusNotFound, "\"status\":\"error\",\"message\":\"Breed not found (master breed does not exist)\",\"code\":404"),
		},
		{
			name: "error - bad subbreed",
			args: args{
				breed:    "hound",
				subbreed: "test",
			},
			want: nil,
			err:  fmt.Errorf("Code (%d): {%s}", http.StatusNotFound, "\"status\":\"error\",\"message\":\"Breed not found (sub breed does not exist)\",\"code\":404"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Log(test.name)

			messages, err := New().RandomImageBySubBreed(test.args.breed, test.args.subbreed)

			if !assert.Equal(t, test.err, err, "Error should be : %+v but got : %+v", test.err, err) {
				return
			}
			if messages != nil && !assert.NotEqual(t, "", messages.Message, "messages should not be empty %+v", messages.Message) {
				return
			}
			if test.want != nil && !assert.Equal(t, test.want.Status, messages.Status) {
				return
			}
		})
	}
}

func TestRandomImagesBySubBreed(t *testing.T) {
	type args struct {
		breed          string
		subbreed       string
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
				subbreed:       "afghan",
				numberOfImages: "3",
			},
			arrayLen: 3,
			want: &MessageArray{
				Status: "success",
			},
			err: nil,
		},
		{
			name: "error - bad breed",
			args: args{
				breed:          "test",
				subbreed:       "afghan",
				numberOfImages: "3",
			},
			want: nil,
			err:  fmt.Errorf("Code (%d): {%s}", http.StatusNotFound, "\"status\":\"error\",\"message\":\"Breed not found (master breed does not exist)\",\"code\":404"),
		},
		{
			name: "error - bad subbreed",
			args: args{
				breed:          "hound",
				subbreed:       "test",
				numberOfImages: "3",
			},
			want: nil,
			err:  fmt.Errorf("Code (%d): {%s}", http.StatusNotFound, "\"status\":\"error\",\"message\":\"Breed not found (sub breed does not exist)\",\"code\":404"),
		},
		{
			name: "bad numberOfImages",
			args: args{
				breed:          "hound",
				subbreed:       "afghan",
				numberOfImages: "test",
			},
			arrayLen: 1,
			want: &MessageArray{
				Status: "success",
			},
			err: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Log(test.name)

			messages, err := New().RandomImagesBySubBreed(test.args.breed, test.args.subbreed, test.args.numberOfImages)

			if !assert.Equal(t, test.err, err, "Error should be : %+v but got : %+v", test.err, err) {
				return
			}
			if messages != nil && !assert.Len(t, messages.Message, test.arrayLen, "messages len should be %d but is %d", test.arrayLen, len(messages.Message)) {
				return
			}
			if test.want != nil && !assert.Equal(t, test.want.Status, messages.Status) {
				return
			}
		})
	}
}
