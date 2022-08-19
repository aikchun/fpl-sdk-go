package fpl_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aikchun/fpl-sdk-go"
)

func TestGetBootstrapStatic(t *testing.T) {
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(BootstrapStaticPayload))
	}))
	defer svr.Close()
	getSeason := fpl.CreateGetAPI(svr.URL)

	var data fpl.BootstrapStatic
	r, err := getSeason(nil)

	defer r.Body.Close()

	if err != nil {
		t.Errorf("expected err to be nil got %v", err)
	}

	err = json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		t.Errorf("Failed to decode BootstrapStatic API call into BootstrapStatic struct")
	}

	if data.Events[0].Id != 22 {
		t.Errorf("Id should be 22, instead it is %d", data.Events[1].Id)
	}

	if data.Events[0].Name != "Gameweek 22" {
		t.Errorf("Name should be Gameweek 22, instead it is %s", data.Events[0].Name)
	}

	if data.Events[0].DeadlineTimeEpoch != 1642185000 {
		t.Errorf("DeadlineTimeEpoch should be 1642185000, instead it is %d", data.Events[0].DeadlineTimeEpoch)
	}

	if !data.Events[0].IsPrevious {
		t.Errorf("IsPrevious should be true, instead it is %t", data.Events[0].IsPrevious)
	}

	if data.Events[0].IsCurrent {
		t.Errorf("IsCurrent should be false, instead it is %t", data.Events[0].IsCurrent)
	}

	if data.Events[0].IsNext {
		t.Errorf("IsNext should be false, instead it is %t", data.Events[0].IsNext)
	}
}

func TestGetFixtures(t *testing.T) {
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(FixtureResponsePayload))
	}))
	defer svr.Close()

	getFixtures := fpl.CreateGetAPI(svr.URL)

	var data []fpl.Fixture
	r, err := getFixtures(map[string]string{"event": "22"})

	defer r.Body.Close()

	if err != nil {
		t.Errorf("Error shouldn't occur when making request")
	}

	err = json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		t.Errorf("Failed to decode Fixtures API into Fixture struct %s", err)
	}

	if len(data) != 2 {
		t.Errorf("Number of Fixtures record should be 2")
	}

	if data[0].Id != 215 {
		t.Errorf("First fixture Id should be 215, but instead it's %d", data[0].Id)
	}

	if data[0].Event != 22 {
		t.Errorf("First fixture Event should be 22, but instead it's %d", data[0].Event)
	}

	if !data[0].Finished {
		t.Errorf("First fixture Finished should be true, but instead it's %t", data[0].Finished)
	}

	if data[0].TeamA != 6 {
		t.Errorf("First fixture TeamA should be 6, but instead it's %d", data[0].TeamA)
	}

	if data[0].TeamH != 12 {
		t.Errorf("First fixture TeamH should be 12, but instead it's %d", data[0].TeamH)
	}

	if data[0].TeamAScore != 0 {
		t.Errorf("First fixture TeamAScore should be 0, but instead it's %d", data[0].TeamAScore)
	}

	if data[0].TeamHScore != 1 {
		t.Errorf("First fixture TeamHScore should be 1, but instead it's %d", data[0].TeamHScore)
	}

	if data[0].Minutes != 90 {
		t.Errorf("First fixture Minutes should be 90, but instead it's %d", data[0].Minutes)
	}

	if !data[0].Started {
		t.Errorf("First fixture Started should be true, but instead it's %t", data[0].Started)
	}

	if data[1].Id != 216 {
		t.Errorf("First fixture Id should be 216, but instead it's %d", data[1].Id)
	}

	if data[1].Event != 22 {
		t.Errorf("First fixture Event should be 22, but instead it's %d", data[1].Event)
	}

	if data[1].Finished {
		t.Errorf("First fixture Finished should be false, but instead it's %t", data[1].Finished)
	}

	if data[1].TeamA != 18 {
		t.Errorf("First fixture TeamA should be 6, but instead it's %d", data[1].TeamA)
	}

	if data[1].TeamH != 14 {
		t.Errorf("First fixture TeamH should be 14, but instead it's %d", data[1].TeamH)
	}

	if data[1].TeamAScore != 2 {
		t.Errorf("First fixture TeamAScore should be 2, but instead it's %d", data[1].TeamAScore)
	}

	if data[1].TeamHScore != 3 {
		t.Errorf("First fixture TeamHScore should be 3, but instead it's %d", data[1].TeamHScore)
	}

	if data[1].Minutes != 70 {
		t.Errorf("First fixture Minutes should be 70, but instead it's %d", data[1].Minutes)
	}

	if data[1].Started {
		t.Errorf("First fixture Started should be false, but instead it's %t", data[1].Started)
	}
}

func TestCreateGetAPI(t *testing.T) {
	urlString := "dummyurl"
	faultyAPI := fpl.CreateGetAPI("yaabbaadaabaadoo")
	_, err := faultyAPI(nil)

	if err == nil {
		t.Errorf("%s should not be parsable into a URL", urlString)
	}
}
