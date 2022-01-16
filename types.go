package fpl

type BootstrapStatic struct {
	Events []Event `json:"events"`
}

type Event struct {
	Id                int64  `json:"id"`
	Name              string `json:"name"`
	DeadlineTimeEpoch int64  `json:"deadline_time_epoch"`
}
