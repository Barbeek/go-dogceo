package dogceo

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

func (c *api) RandomImage() (*Message, error) {
	random := &Message{}
	uri := fmt.Sprintf("/%s/image/random", c.baseURL)
	if err := request(uri, http.MethodGet, nil, &random); err != nil {
		return nil, err
	}
	return random, nil
}

func (c *api) RandomImages(number string) (*MessageArray, error) {
	random := &MessageArray{}
	nb, err := strconv.Atoi(number)
	if err != nil {
		return nil, err
	}
	if nb > 50 {
		return nil, errors.New("number can not exceed 50")
	}
	uri := fmt.Sprintf("/%s/image/random/%s", c.baseURL, number)
	if err := request(uri, http.MethodGet, nil, &random); err != nil {
		return nil, err
	}
	return random, nil
}

func (c *api) ImagesByBread(breed string) (*MessageArray, error) {
	images := &MessageArray{}
	uri := fmt.Sprintf("/%s/%s/images", c.baseURL, breed)
	if err := request(uri, http.MethodGet, nil, &images); err != nil {
		return nil, err
	}
	return images, nil
}

func (c *api) RandomImageByBread(breed string) (*Message, error) {
	images := &Message{}
	uri := fmt.Sprintf("/%s/%s/images/random", c.baseURL, breed)
	if err := request(uri, http.MethodGet, nil, &images); err != nil {
		return nil, err
	}
	return images, nil
}

func (c *api) RandomImagesByBread(breed, number string) (*MessageArray, error) {
	images := &MessageArray{}
	uri := fmt.Sprintf("/%s/%s/images/random/%s", c.baseURL, breed, number)
	if err := request(uri, http.MethodGet, nil, &images); err != nil {
		return nil, err
	}
	return images, nil
}
