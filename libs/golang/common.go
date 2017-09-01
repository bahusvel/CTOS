package ctos

import "time"

type LogLine struct {
	Time       time.Time
	Line       string
	Host       string
	Origin     string
	Attributes map[string]string
}
