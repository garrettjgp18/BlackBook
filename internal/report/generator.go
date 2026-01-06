package report

import (
	"BlackBook/internal/terminal"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

// Generator creates reports from session data
type Generator struct {
	session  *terminal.Session
	findings []Finding
}

// NewGenerator creates a new report generator
func NewGenerator(session *terminal.Session) *Generator {
	return &Generator{
		session:  session,
		findings: make([]Finding, 0),
	}
}

// AddFinding adds a finding to the report
func (g *Generator) AddFinding(finding Finding) {
	if finding.ID == "" {
		finding.ID = uuid.New().String()
	}
	g.findings = append(g.findings, finding)
}

// GenerateReport generates a complete report
func (g *Generator) GenerateReport(title string) *Report {
	now := time.Now()
	
	report := &Report{
		ID:              uuid.New().String(),
		Title:           title,
		StartTime:       g.session.StartTime,
		EndTime:         now,
		Findings:        g.findings,
		CommandTimeline: g.session.History,
		Summary:         g.generateSummary(),
		Methodology:     g.generateMethodology(),
		Recommendations: g.generateRecommendations(),
	}
	
	return report
}

// generateSummary generates an executive summary
func (g *Generator) generateSummary() string {
	var sb strings.Builder
	
	sb.WriteString("## Executive Summary\n\n")
	sb.WriteString(fmt.Sprintf("This penetration test was conducted from %s to %s.\n\n",
		g.session.StartTime.Format("2006-01-02 15:04:05"),
		time.Now().Format("2006-01-02 15:04:05")))
	
	successful := 0
	for _, cmd := range g.session.History {
		if cmd.ExitCode == 0 {
			successful++
		}
	}
	
	sb.WriteString(fmt.Sprintf("**Total Commands Executed:** %d\n", len(g.session.History)))
	sb.WriteString(fmt.Sprintf("**Successful Commands:** %d\n", successful))
	sb.WriteString(fmt.Sprintf("**Total Findings:** %d\n\n", len(g.findings)))
	
	if len(g.findings) > 0 {
		criticalCount := 0
		highCount := 0
		mediumCount := 0
		lowCount := 0
		
		for _, f := range g.findings {
			switch strings.ToLower(f.Severity) {
			case "critical":
				criticalCount++
			case "high":
				highCount++
			case "medium":
				mediumCount++
			case "low":
				lowCount++
			}
		}
		
		sb.WriteString("**Findings by Severity:**\n")
		if criticalCount > 0 {
			sb.WriteString(fmt.Sprintf("- Critical: %d\n", criticalCount))
		}
		if highCount > 0 {
			sb.WriteString(fmt.Sprintf("- High: %d\n", highCount))
		}
		if mediumCount > 0 {
			sb.WriteString(fmt.Sprintf("- Medium: %d\n", mediumCount))
		}
		if lowCount > 0 {
			sb.WriteString(fmt.Sprintf("- Low: %d\n", lowCount))
		}
	}
	
	return sb.String()
}

// generateMethodology generates the methodology section
func (g *Generator) generateMethodology() string {
	var sb strings.Builder
	
	sb.WriteString("## Methodology\n\n")
	sb.WriteString("The following approach was taken during this penetration test:\n\n")
	
	// Identify phases based on commands
	phases := make(map[string]bool)
	for _, cmd := range g.session.History {
		cmdLower := strings.ToLower(cmd.Command)
		
		if strings.Contains(cmdLower, "nmap") || strings.Contains(cmdLower, "masscan") {
			phases["Network Scanning"] = true
		}
		if strings.Contains(cmdLower, "gobuster") || strings.Contains(cmdLower, "dirb") || strings.Contains(cmdLower, "ffuf") {
			phases["Web Enumeration"] = true
		}
		if strings.Contains(cmdLower, "hydra") || strings.Contains(cmdLower, "medusa") {
			phases["Password Attacks"] = true
		}
		if strings.Contains(cmdLower, "sqlmap") {
			phases["SQL Injection Testing"] = true
		}
	}
	
	phaseNum := 1
	for phase := range phases {
		sb.WriteString(fmt.Sprintf("%d. %s\n", phaseNum, phase))
		phaseNum++
	}
	
	if len(phases) == 0 {
		sb.WriteString("1. Manual Testing and Exploration\n")
	}
	
	return sb.String()
}

// generateRecommendations generates recommendations
func (g *Generator) generateRecommendations() []string {
	recommendations := make([]string, 0)
	
	// Generate recommendations based on findings
	hasCritical := false
	hasHigh := false
	
	for _, f := range g.findings {
		switch strings.ToLower(f.Severity) {
		case "critical":
			hasCritical = true
		case "high":
			hasHigh = true
		}
	}
	
	if hasCritical {
		recommendations = append(recommendations,
			"Address all Critical severity findings immediately as they pose significant security risks")
	}
	
	if hasHigh {
		recommendations = append(recommendations,
			"Remediate High severity findings within the next security update cycle")
	}
	
	if len(g.findings) > 0 {
		recommendations = append(recommendations,
			"Conduct regular security assessments to identify new vulnerabilities",
			"Implement security best practices and secure coding guidelines",
			"Provide security awareness training to development and operations teams")
	} else {
		recommendations = append(recommendations,
			"Continue maintaining current security posture",
			"Schedule regular security assessments",
			"Stay updated on emerging security threats and vulnerabilities")
	}
	
	return recommendations
}

// GetMetadata returns report metadata
func (g *Generator) GetMetadata() ReportMetadata {
	successful := 0
	failed := 0
	
	for _, cmd := range g.session.History {
		if cmd.ExitCode == 0 {
			successful++
		} else if cmd.ExitCode != -1 {
			failed++
		}
	}
	
	var duration int64
	if len(g.session.History) > 0 {
		lastCmd := g.session.History[len(g.session.History)-1]
		duration = lastCmd.EndTime.Sub(g.session.StartTime).Milliseconds()
	}
	
	return ReportMetadata{
		TotalCommands:      len(g.session.History),
		SuccessfulCommands: successful,
		FailedCommands:     failed,
		TotalFindings:      len(g.findings),
		Duration:           duration,
		GeneratedAt:        time.Now(),
	}
}
