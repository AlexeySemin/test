package response

import "time"

type Response struct {
	Message string      `json:"message"`
	Body    interface{} `json:"body"`
}

type Log struct {
	Duration float64 `json:"duration"`
}

func NewLog(start time.Time, end time.Time) *Log {
	duration := end.Sub(start).Seconds()
	return &Log{duration}
}

type LogOnly struct {
	Message string
	Body    Log
}

type MinMaxAvg struct {
	Min int     `json:"min"`
	Max int     `json:"max"`
	Avg float64 `json:"avg"`
}

type MinMaxAvgRating struct {
	Message string
	Body    struct {
		MinMaxAvg
		Log
	}
}

type PerMonthJSON struct {
	Data string `json:"data"`
}

type PerMonthJSONData struct {
	Message string
	Body    struct {
		PerMonthJSON
		Log
	}
}
