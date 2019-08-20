package response

import "time"

type LogResponse struct {
	Start    time.Time
	End      time.Time
	Duration float64
}

func NewLog(start time.Time, end time.Time) *LogResponse {
	duration := end.Sub(start).Seconds()
	return &LogResponse{start, end, duration}
}

type MinMaxAvgRating struct {
	Min int
	Max int
	Avg float64
}
