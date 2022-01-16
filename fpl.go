package fpl

import (
	"encoding/json"
	"net/http"
	"os"
)

func GetSeason() (*BootstrapStatic, error) {
	client := http.Client{}

	seasonUrl, _ := os.LookupEnv("SEASON_URL")

	req, err := http.NewRequest("GET", seasonUrl, nil)

	if err != nil {
		//Handle Error
		return nil, err
	}

	req.Header = http.Header{
		"Content-Type": []string{"application/json"},
		"User-Agent":   []string{"Golang_Lambda/1.0"},
	}

	r, err := client.Do(req)

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
