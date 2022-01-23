package fpl

import (
	"io"
	"net/http"
	"net/url"
)

func Request(method string, url string, body io.Reader) (*http.Response, error) {
	client := http.Client{}

	req, err := http.NewRequest(method, url, body)

	if err != nil {
		//Handle Error
		return nil, err
	}

	req.Header = http.Header{
		"Content-Type": []string{"application/json"},
		"User-Agent":   []string{"Golang_Lambda/1.0"},
	}

	return client.Do(req)
}

func CreateGetAPI(uri string) func(map[string]string) (*http.Response, error) {
	return func(queryParams map[string]string) (*http.Response, error) {

		urlA, err := url.Parse(uri)
		if err != nil {
			return nil, err
		}

		values := urlA.Query()
		if queryParams != nil {
			for key, value := range queryParams {
				values.Add(key, value)
			}
		}

		urlA.RawQuery = values.Encode()

		return Request("GET", urlA.String(), nil)
	}
}

var BootstrapStaticURL = "https://fantasy.premierleague.com/api/bootstrap-static/"

var FixturesURL = "https://fantasy.premierleague.com/api/fixtures"

var GetBootstrapStatic = CreateGetAPI(BootstrapStaticURL)

var GetFixtures = CreateGetAPI(FixturesURL)
