package response

import "time"

type Log struct {
	Start    time.Time
	End      time.Time
	Duration float64
}

func NewLog(start time.Time, end time.Time) *Log {
	duration := end.Sub(start).Seconds()
	return &Log{start, end, duration}
}

type MinMaxAvgRating struct {
	Min int
	Max int
	Avg float64
}

type PerMonthJSONData struct {
	Data string
}
