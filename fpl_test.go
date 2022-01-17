package fpl_test

import (
	"testing"

	"github.com/aikchun/fpl-sdk-go"
)

func TestGetSeason(t *testing.T) {
	season, err := fpl.GetSeason()

	if err != nil {
		t.Errorf("shouldn't have error")
	}

	if len(season.Events) == 0 {
		t.Errorf("Shouldn't have 0 events")
	}

	if len(season.Teams) == 0 {
		t.Errorf("Shouldn't have 0 teams")
	}
}

func TestGetFixtures(t *testing.T) {
	fixtures, err := fpl.GetFixtures(22)

	if err != nil {
		t.Errorf("shouldn't have error")
	}

	if len(*fixtures) > 20 {
		t.Errorf("Should have more than 20 fixtures, have %d instead", len(*fixtures))
	}

}
