package fpl

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

func GetSeason() (*BootstrapStatic, error) {
	r, err := Request("GET", "https://fantasy.premierleague.com/api/bootstrap-static/", nil)

	defer r.Body.Close()

	if err != nil {
		return nil, err
	}

	var data BootstrapStatic

	err = json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		return nil, err
	}

	return &data, err
}

func GetFixtures(eventId int) (*[]Fixture, error) {

	fixturesUrl := "https://fantasy.premierleague.com/api/fixtures/"

	if eventId > 0 {
		fixturesUrl = fmt.Sprintf("https://fantasy.premierleague.com/api/fixtures/?events=%d", eventId)
	}

	r, err := Request("GET", fixturesUrl, nil)

	defer r.Body.Close()

	if err != nil {
		return nil, err
	}

	var data []Fixture

	err = json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		return nil, err
	}

	return &data, err
}
