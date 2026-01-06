package terminal

import "time"

// CommandEntry represents a tracked command
type CommandEntry struct {
	ID        string    `json:"id"`
	Command   string    `json:"command"`
	Output    string    `json:"output"`
	ExitCode  int       `json:"exitCode"`
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
	Duration  int64     `json:"duration"` // milliseconds
}

// Session represents a terminal session
type Session struct {
	ID        string         `json:"id"`
	StartTime time.Time      `json:"startTime"`
	History   []CommandEntry `json:"history"`
	Active    bool           `json:"active"`
}
