package dogceo

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

// RandomImage returns a single random image from all dog collection
func (c *api) RandomImage() (*Message, error) {
	random := &Message{}
	uri := fmt.Sprintf("%s/breeds/image/random", c.baseURL)
	if err := request(uri, http.MethodGet, nil, &random); err != nil {
		return nil, err
	}
	return random, nil
}

// RandomImages returns, randomly, the number of images passed in parameter from all dog collection
// If the number is not an integer, it will be set by default at 1.
func (c *api) RandomImages(numberOfImages string) (*MessageArray, error) {
	random := &MessageArray{}
	nb, err := strconv.Atoi(numberOfImages)
	if err != nil {
		numberOfImages = defaultNumberOfImages
	} else if nb > 50 {
		return nil, errors.New("numberOfImages can not exceed 50")
	}
	uri := fmt.Sprintf("%s/breeds/image/random/%s", c.baseURL, numberOfImages)
	if err := request(uri, http.MethodGet, nil, &random); err != nil {
		return nil, err
	}
	return random, nil
}

// ImagesByBreed returns an array of all the images from a breed
func (c *api) ImagesByBreed(breed string) (*MessageArray, error) {
	images := &MessageArray{}
	uri := fmt.Sprintf("%s/breed/%s/images", c.baseURL, breed)
	if err := request(uri, http.MethodGet, nil, &images); err != nil {
		return nil, err
	}
	return images, nil
}

// RandomImageByBreed returns a random dog image from a breed
// If the breed passed in parameter doesn't exist, it will return 404.
func (c *api) RandomImageByBreed(breed string) (*Message, error) {
	images := &Message{}
	uri := fmt.Sprintf("%s/breed/%s/images/random", c.baseURL, breed)
	if err := request(uri, http.MethodGet, nil, &images); err != nil {
		return nil, err
	}
	return images, nil
}

// RandomImagesByBreed returns, randomly, the number of images passed in parameter from a breed
// If the breed passed in parameter doesn't exist, it will return 404.
// If the number is not an integer, it will be set by default at 1.
func (c *api) RandomImagesByBreed(breed, numberOfImages string) (*MessageArray, error) {
	images := &MessageArray{}
	if _, err := strconv.Atoi(numberOfImages); err != nil {
		numberOfImages = defaultNumberOfImages
	}
	uri := fmt.Sprintf("%s/breed/%s/images/random/%s", c.baseURL, breed, numberOfImages)
	if err := request(uri, http.MethodGet, nil, &images); err != nil {
		return nil, err
	}
	return images, nil
}
