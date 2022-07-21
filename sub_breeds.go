package dogceo

import (
	"fmt"
	"net/http"
	"strconv"
)

func (c *api) ListSubBreeds(breed string) (*MessageArray, error) {
	subBreeds := &MessageArray{}
	uri := fmt.Sprintf("%s/breed/%s/list", c.baseURL, breed)
	if err := request(uri, http.MethodGet, nil, &subBreeds); err != nil {
		return nil, err
	}
	return subBreeds, nil
}

func (c *api) ImagesBySubBreed(breed, subbreed string) (*MessageArray, error) {
	images := &MessageArray{}
	uri := fmt.Sprintf("%s/breed/%s/%s/images", c.baseURL, breed, subbreed)
	if err := request(uri, http.MethodGet, nil, &images); err != nil {
		return nil, err
	}
	return images, nil
}

func (c *api) RandomImageBySubBreed(breed, subbreed string) (*Message, error) {
	images := &Message{}
	uri := fmt.Sprintf("%s/breed/%s/%s/images/random", c.baseURL, breed, subbreed)
	if err := request(uri, http.MethodGet, nil, &images); err != nil {
		return nil, err
	}
	return images, nil
}

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
