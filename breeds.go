package dogceo

import (
	"fmt"
	"net/http"
)

func (c *api) ListBreeds() (*MessageMap, error) {
	breeds := &MessageMap{}
	uri := fmt.Sprintf("%s/breeds/list/all", c.baseURL)
	if err := request(uri, http.MethodGet, nil, &breeds); err != nil {
		return nil, err
	}
	return breeds, nil
}
