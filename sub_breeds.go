package dogceo

import (
	"fmt"
	"net/http"
)

func (c *api) ListSubBreads(breed string) (*MessageArray, error) {
	subBreeds := &MessageArray{}
	uri := fmt.Sprintf("/%s/%s/list", c.baseURL, breed)
	if err := request(uri, http.MethodGet, nil, &subBreeds); err != nil {
		return nil, err
	}
	return subBreeds, nil
}

func (c *api) ImagesBySubBreed(breed, subbreed string) (*MessageArray, error) {
	images := &MessageArray{}
	uri := fmt.Sprintf("/%s/%s/%s/list", c.baseURL, breed, subbreed)
	if err := request(uri, http.MethodGet, nil, &images); err != nil {
		return nil, err
	}
	return images, nil
}

func (c *api) RandomImageBySubBreed(breed, subbreed string) (*Message, error) {
	images := &Message{}
	uri := fmt.Sprintf("/%s/%s/%s/random", c.baseURL, breed, subbreed)
	if err := request(uri, http.MethodGet, nil, &images); err != nil {
		return nil, err
	}
	return images, nil
}

func (c *api) RandomImagesBySubBreed(breed, subbreed, number string) (*MessageArray, error) {
	images := &MessageArray{}
	uri := fmt.Sprintf("/%s/%s/%s/random/%s", c.baseURL, breed, subbreed, number)
	if err := request(uri, http.MethodGet, nil, &images); err != nil {
		return nil, err
	}
	return images, nil
}
