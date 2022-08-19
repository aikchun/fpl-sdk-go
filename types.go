package fpl

type BootstrapStatic struct {
	Events []Event `json:"events"`
	Teams  []Team  `json:"teams"`
}

type Event struct {
	ID                int    `json:"id"`
	Name              string `json:"name"`
	DeadlineTimeEpoch int64  `json:"deadline_time_epoch"`
	IsPrevious        bool   `json:"is_previous"`
	IsCurrent         bool   `json:"is_current"`
	IsNext            bool   `json:"is_next"`
}

type Team struct {
	ID        int    `json:"id"`
	Code      int    `json:"code"`
	Name      string `json:"name"`
	ShortName string `json:"short_name"`
}

type Fixture struct {
	ID         int  `json:"id"`
	Event      int  `json:"event"`
	Started    bool `json:"started"`
	Finished   bool `json:"finished"`
	TeamA      int  `json:"team_a"`
	TeamH      int  `json:"team_h"`
	TeamAScore int  `json:"team_a_score"`
	TeamHScore int  `json:"team_h_score"`
	Minutes    int  `json:"minutes"`
}
