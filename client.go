package dogceo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

// HTTPClient interface
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

const (
	baseURL               string = "https://dog.ceo/api"
	defaultNumberOfImages string = "1"
)

var (
	httpClient HTTPClient = &http.Client{}
)

func request(baseURL, method string, body, result interface{}) error {
	data, err := json.Marshal(body)
	if err != nil {
		log.Fatal(err)
	}
	uri, err := url.Parse(baseURL)
	if err != nil {
		return err
	}
	bodyReader := bytes.NewReader(data)
	req, err := http.NewRequest(method, uri.String(), bodyReader)
	if err != nil {
		log.Printf("error during request call of %s %s\n error : %s", method, uri.String(), err.Error())
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

	return fmt.Errorf("Code (%d): %s", res.StatusCode, string(resBody))
}

func New() DogCEO {
	client := &api{
		baseURL: baseURL,
	}
	return client
}
