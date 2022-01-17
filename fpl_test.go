package fpl_test

import (
	"testing"

	"github.com/aikchun/fpl-sdk-go"
)

func TestGetSeason(t *testing.T) {
	season, err := fpl.GetSeason()

	if len(season.Events) == 0 {
		t.Errorf("Shouldn't have 0 events")
	}

	if len(season.Teams) == 0 {
		t.Errorf("Shouldn't have 0 teams")
	}
}
