package fpl

type BootstrapStatic struct {
	Events []Event `json:"events"`
	Teams  []Team  `json:"teams"`
}

type Event struct {
	Id                int64  `json:"id"`
	Name              string `json:"name"`
	DeadlineTimeEpoch int64  `json:"deadline_time_epoch"`
}

type Team struct {
	Id        int    `json:"id"`
	Code      int    `json:"code"`
	Name      string `json:"name"`
	ShortName string `json:"short_name"`
}
