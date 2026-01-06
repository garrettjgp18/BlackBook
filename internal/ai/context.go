package ai

import (
	"BlackBook/internal/terminal"
	"fmt"
	"strings"
)

// ContextBuilder builds context for AI prompts
type ContextBuilder struct {
	history []terminal.CommandEntry
}

// NewContextBuilder creates a new context builder
func NewContextBuilder(history []terminal.CommandEntry) *ContextBuilder {
	return &ContextBuilder{
		history: history,
	}
}

// GetRecentCommands returns the N most recent commands
func (cb *ContextBuilder) GetRecentCommands(n int) []terminal.CommandEntry {
	if n >= len(cb.history) {
		return cb.history
	}
	return cb.history[len(cb.history)-n:]
}

// GetSuccessfulCommands returns only successful commands
func (cb *ContextBuilder) GetSuccessfulCommands() []terminal.CommandEntry {
	successful := make([]terminal.CommandEntry, 0)
	for _, cmd := range cb.history {
		if cmd.ExitCode == 0 {
			successful = append(successful, cmd)
		}
	}
	return successful
}

// GetSecurityToolCommands returns commands that used security tools
func (cb *ContextBuilder) GetSecurityToolCommands() []terminal.CommandEntry {
	toolCommands := make([]terminal.CommandEntry, 0)
	for _, cmd := range cb.history {
		if terminal.IsSecurityTool(cmd.Command) {
			toolCommands = append(toolCommands, cmd)
		}
	}
	return toolCommands
}

// GetSummary returns a summary of the session
func (cb *ContextBuilder) GetSummary() string {
	if len(cb.history) == 0 {
		return "No commands executed yet."
	}

	var sb strings.Builder
	sb.WriteString("Session Summary:\n")
	sb.WriteString(fmt.Sprintf("Total commands: %d\n", len(cb.history)))
	
	successful := cb.GetSuccessfulCommands()
	sb.WriteString(fmt.Sprintf("Successful: %d\n", len(successful)))
	
	toolCommands := cb.GetSecurityToolCommands()
	sb.WriteString(fmt.Sprintf("Security tools used: %d\n", len(toolCommands)))
	
	return sb.String()
}

// BuildContext builds a context string for AI prompts
func (cb *ContextBuilder) BuildContext(maxCommands int) string {
	recent := cb.GetRecentCommands(maxCommands)
	if len(recent) == 0 {
		return "No command history available."
	}

	var sb strings.Builder
	for _, cmd := range recent {
		sb.WriteString("Command: " + cmd.Command + "\n")
		sb.WriteString(fmt.Sprintf("Exit Code: %d\n", cmd.ExitCode))
		if cmd.Output != "" {
			output := cmd.Output
			if len(output) > 200 {
				output = output[:200] + "..."
			}
			sb.WriteString("Output: " + output + "\n")
		}
		sb.WriteString("---\n")
	}
	
	return sb.String()
}
