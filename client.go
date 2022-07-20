package dogceo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// HTTPClient interface
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type apiStatus struct {
	Code    int    `json:"status_code"`
	Message string `json:"status_message"`
}

var (
	httpClient HTTPClient = &http.Client{
		Transport: &http.Transport{},
	}
)

func request(url, method string, body, result interface{}) error {
	data, err := json.Marshal(body)
	if err != nil {
		log.Fatal(err)
	}
	bodyReader := bytes.NewReader(data)
	req, err := http.NewRequest(method, url, bodyReader)
	if err != nil {
		log.Printf("error during request call of %s %s\n error : %s", method, url, err.Error())
		return err
	}

	res, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if res.StatusCode >= http.StatusOK && res.StatusCode < http.StatusMultipleChoices {
		if result != nil {
			if err := json.Unmarshal(resBody, &result); err != nil {
				return err
			}
		}
		return nil
	}

	// Handle failure modes
	var status apiStatus
	if err = json.Unmarshal(resBody, &status); err != nil {
		return err
	}
	return fmt.Errorf("Code (%d): %s", status.Code, status.Message)
}

func New() DogCEO {
	client := &api{
		baseURL: baseURL,
	}
	return client
}
