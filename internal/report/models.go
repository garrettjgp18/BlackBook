package report

import (
	"BlackBook/internal/terminal"
	"time"
)

// Finding represents a security finding
type Finding struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Severity    string    `json:"severity"` // Critical, High, Medium, Low, Info
	Description string    `json:"description"`
	Evidence    string    `json:"evidence"`
	Command     string    `json:"command"`
	Timestamp   time.Time `json:"timestamp"`
	Remediation string    `json:"remediation"`
}

// Report represents a penetration test report
type Report struct {
	ID              string                  `json:"id"`
	Title           string                  `json:"title"`
	StartTime       time.Time               `json:"startTime"`
	EndTime         time.Time               `json:"endTime"`
	Findings        []Finding               `json:"findings"`
	CommandTimeline []terminal.CommandEntry `json:"commandTimeline"`
	Summary         string                  `json:"summary"`
	Methodology     string                  `json:"methodology"`
	Recommendations []string                `json:"recommendations"`
}

// ReportMetadata contains metadata about the report
type ReportMetadata struct {
	TotalCommands      int       `json:"totalCommands"`
	SuccessfulCommands int       `json:"successfulCommands"`
	FailedCommands     int       `json:"failedCommands"`
	TotalFindings      int       `json:"totalFindings"`
	Duration           int64     `json:"duration"` // milliseconds
	GeneratedAt        time.Time `json:"generatedAt"`
}
