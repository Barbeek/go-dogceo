package dogceo

import (
	"fmt"
	"net/http"
)

const baseURL string = "https://dog.ceo/api/breed"

func (c *api) ListBreads() (*MessageMap, error) {
	breeds := &MessageMap{}
	uri := fmt.Sprintf("/%s/list/all", c.baseURL)
	if err := request(uri, http.MethodGet, nil, &breeds); err != nil {
		return nil, err
	}
	return breeds, nil
}
