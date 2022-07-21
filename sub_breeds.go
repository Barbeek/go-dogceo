package dogceo

import (
	"fmt"
	"net/http"
	"strconv"
)

// ListSubBreeds returns an array of all the sub-breeds from a breed
// If the breed passed in parameter doesn't exist, it will return 404.
func (c *api) ListSubBreeds(breed string) (*MessageArray, error) {
	subBreeds := &MessageArray{}
	uri := fmt.Sprintf("%s/breed/%s/list", c.baseURL, breed)
	if err := request(uri, http.MethodGet, nil, &subBreeds); err != nil {
		return nil, err
	}
	return subBreeds, nil
}

// ImagesBySubBreed returns an array of all the images from the sub-breed
// If the breed or the sub-breed passed in parameter don't exist, it will return 404.
func (c *api) ImagesBySubBreed(breed, subbreed string) (*MessageArray, error) {
	images := &MessageArray{}
	uri := fmt.Sprintf("%s/breed/%s/%s/images", c.baseURL, breed, subbreed)
	if err := request(uri, http.MethodGet, nil, &images); err != nil {
		return nil, err
	}
	return images, nil
}

// RandomImageBySubBreed returns a random dog image from a breed's sub-bread.
// If the breed or the sub-breed passed in parameter don't exist, it will return 404.
func (c *api) RandomImageBySubBreed(breed, subbreed string) (*Message, error) {
	images := &Message{}
	uri := fmt.Sprintf("%s/breed/%s/%s/images/random", c.baseURL, breed, subbreed)
	if err := request(uri, http.MethodGet, nil, &images); err != nil {
		return nil, err
	}
	return images, nil
}

// RandomImagesBySubBreed returns, randomly, the number of images passed in parameter from a breed's sub-breed
// If the breed or the sub-breed passed in parameter don't exist, it will return 404.
// If the number is not an integer, it will be set by default at 1.
func (c *api) RandomImagesBySubBreed(breed, subbreed, numberOfImages string) (*MessageArray, error) {
	images := &MessageArray{}
	if _, err := strconv.Atoi(numberOfImages); err != nil {
		numberOfImages = defaultNumberOfImages
	}
	uri := fmt.Sprintf("%s/breed/%s/%s/images/random/%s", c.baseURL, breed, subbreed, numberOfImages)
	if err := request(uri, http.MethodGet, nil, &images); err != nil {
		return nil, err
	}
	return images, nil
}
