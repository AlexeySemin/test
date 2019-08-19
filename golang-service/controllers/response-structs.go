package controllers

import "time"

type logResponse struct {
	Start    time.Time
	End      time.Time
	Duration float64
}

func NewLogResponse(start time.Time, end time.Time) *logResponse {
	duration := end.Sub(start).Seconds()
	return &logResponse{start, end, duration}
}
