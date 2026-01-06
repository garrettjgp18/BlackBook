package terminal

import (
	"strings"
)

// CommandHistory manages the command history
type CommandHistory struct {
	entries []CommandEntry
}

// NewCommandHistory creates a new command history
func NewCommandHistory() *CommandHistory {
	return &CommandHistory{
		entries: make([]CommandEntry, 0),
	}
}

// Add adds a command entry to the history
func (h *CommandHistory) Add(entry CommandEntry) {
	h.entries = append(h.entries, entry)
}

// Get returns all command entries
func (h *CommandHistory) Get() []CommandEntry {
	return h.entries
}

// GetByID returns a command entry by ID
func (h *CommandHistory) GetByID(id string) *CommandEntry {
	for i := range h.entries {
		if h.entries[i].ID == id {
			return &h.entries[i]
		}
	}
	return nil
}

// GetLast returns the last N command entries
func (h *CommandHistory) GetLast(n int) []CommandEntry {
	if n >= len(h.entries) {
		return h.entries
	}
	return h.entries[len(h.entries)-n:]
}

// Clear clears the command history
func (h *CommandHistory) Clear() {
	h.entries = make([]CommandEntry, 0)
}

// Filter returns command entries matching a filter
func (h *CommandHistory) Filter(filter func(CommandEntry) bool) []CommandEntry {
	filtered := make([]CommandEntry, 0)
	for _, entry := range h.entries {
		if filter(entry) {
			filtered = append(filtered, entry)
		}
	}
	return filtered
}

// Search searches for commands containing the given text
func (h *CommandHistory) Search(query string) []CommandEntry {
	return h.Filter(func(entry CommandEntry) bool {
		return strings.Contains(strings.ToLower(entry.Command), strings.ToLower(query))
	})
}

// GetSuccessful returns only successful commands (exit code 0)
func (h *CommandHistory) GetSuccessful() []CommandEntry {
	return h.Filter(func(entry CommandEntry) bool {
		return entry.ExitCode == 0
	})
}

// GetFailed returns only failed commands (exit code != 0)
func (h *CommandHistory) GetFailed() []CommandEntry {
	return h.Filter(func(entry CommandEntry) bool {
		return entry.ExitCode != 0 && entry.ExitCode != -1
	})
}
