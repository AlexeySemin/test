package response

import "time"

type Log struct {
	Duration float64 `json:"duration"`
}

func NewLog(start time.Time, end time.Time) *Log {
	duration := end.Sub(start).Seconds()
	return &Log{duration}
}

type MinMaxAvgRating struct {
	Min int     `json:"min"`
	Max int     `json:"max"`
	Avg float64 `json:"avg"`
}

type PerMonthJSONData struct {
	Data string `json:"data"`
}
