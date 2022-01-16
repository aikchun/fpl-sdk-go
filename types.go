package fpl

type BootstrapStatic struct {
	Events []Event `json:"events"`
	Teams  []Team  `json:"teams"`
}

type Event struct {
	Id                int64  `json:"id"`
	Name              string `json:"name"`
	DeadlineTimeEpoch int64  `json:"deadline_time_epoch"`
	IsPrevious        bool   `json:"is_previous"`
	IsCurrent         bool   `json:"is_previous"`
	IsNext            bool   `json:"is_next"`
}

type Team struct {
	Id        int    `json:"id"`
	Code      int    `json:"code"`
	Name      string `json:"name"`
	ShortName string `json:"short_name"`
}

type Fixture struct {
	Id         int  `json:"id"`
	Event      int  `json:"event"`
	Finished   bool `json:"finished"`
	TeamA      int  `json:"team_a"`
	TeamB      int  `json:"team_b"`
	TeamAScore int  `json:"team_a_score"`
	TeamBScore int  `json:"team_b_score"`
	Minutes    int  `json:"minutes"`
}
