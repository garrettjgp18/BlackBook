package terminal

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/creack/pty"
	"github.com/google/uuid"
)

// Terminal manages PTY and command tracking
type Terminal struct {
	ptmx           *os.File
	cmd            *exec.Cmd
	session        *Session
	mu             sync.RWMutex
	outputBuffer   strings.Builder
	currentCommand *CommandEntry
	shellType      string
}

// NewTerminal creates a new terminal instance
func NewTerminal() (*Terminal, error) {
	shell := getDefaultShell()
	
	// Create command
	cmd := exec.Command(shell)
	
	// Set environment
	cmd.Env = append(os.Environ(), "TERM=xterm-256color")
	
	// Start the command with a PTY
	ptmx, err := pty.Start(cmd)
	if err != nil {
		return nil, fmt.Errorf("failed to start pty: %w", err)
	}

	t := &Terminal{
		ptmx:      ptmx,
		cmd:       cmd,
		shellType: shell,
		session: &Session{
			ID:        uuid.New().String(),
			StartTime: time.Now(),
			History:   make([]CommandEntry, 0),
			Active:    true,
		},
	}

	return t, nil
}

// getDefaultShell returns the default shell for the current OS
func getDefaultShell() string {
	switch runtime.GOOS {
	case "windows":
		// Try PowerShell first, fall back to cmd
		if _, err := exec.LookPath("powershell.exe"); err == nil {
			return "powershell.exe"
		}
		return "cmd.exe"
	default:
		// Unix-like systems
		shell := os.Getenv("SHELL")
		if shell == "" {
			shell = "/bin/bash"
		}
		return shell
	}
}

// Read reads output from the PTY
func (t *Terminal) Read(p []byte) (n int, err error) {
	return t.ptmx.Read(p)
}

// Write writes input to the PTY
func (t *Terminal) Write(p []byte) (n int, err error) {
	return t.ptmx.Write(p)
}

// WriteString writes a string to the PTY
func (t *Terminal) WriteString(s string) error {
	_, err := t.ptmx.Write([]byte(s))
	return err
}

// StartOutputCapture starts capturing terminal output
func (t *Terminal) StartOutputCapture() {
	go func() {
		reader := bufio.NewReader(t.ptmx)
		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				if err != io.EOF {
					// Error reading - terminal may be closed
				}
				return
			}
			
			t.mu.Lock()
			t.outputBuffer.WriteString(line)
			t.mu.Unlock()
		}
	}()
}

// GetOutput returns the current output buffer
func (t *Terminal) GetOutput() string {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.outputBuffer.String()
}

// AddCommand adds a command to the history
func (t *Terminal) AddCommand(command string) string {
	t.mu.Lock()
	defer t.mu.Unlock()

	entry := CommandEntry{
		ID:        uuid.New().String(),
		Command:   command,
		StartTime: time.Now(),
		ExitCode:  -1, // Unknown until command completes
	}

	t.currentCommand = &entry
	t.session.History = append(t.session.History, entry)
	
	return entry.ID
}

// UpdateCommandOutput updates the output for the current command
func (t *Terminal) UpdateCommandOutput(commandID string, output string, exitCode int) {
	t.mu.Lock()
	defer t.mu.Unlock()

	for i := range t.session.History {
		if t.session.History[i].ID == commandID {
			t.session.History[i].Output = output
			t.session.History[i].ExitCode = exitCode
			t.session.History[i].EndTime = time.Now()
			t.session.History[i].Duration = t.session.History[i].EndTime.Sub(t.session.History[i].StartTime).Milliseconds()
			break
		}
	}
}

// GetHistory returns the command history
func (t *Terminal) GetHistory() []CommandEntry {
	t.mu.RLock()
	defer t.mu.RUnlock()
	
	history := make([]CommandEntry, len(t.session.History))
	copy(history, t.session.History)
	return history
}

// GetSession returns the current session
func (t *Terminal) GetSession() *Session {
	t.mu.RLock()
	defer t.mu.RUnlock()
	
	sessionCopy := *t.session
	sessionCopy.History = make([]CommandEntry, len(t.session.History))
	copy(sessionCopy.History, t.session.History)
	
	return &sessionCopy
}

// Resize resizes the PTY
func (t *Terminal) Resize(rows, cols uint16) error {
	return pty.Setsize(t.ptmx, &pty.Winsize{
		Rows: rows,
		Cols: cols,
	})
}

// Close closes the terminal
func (t *Terminal) Close() error {
	t.mu.Lock()
	t.session.Active = false
	t.mu.Unlock()

	if t.ptmx != nil {
		t.ptmx.Close()
	}
	
	if t.cmd != nil && t.cmd.Process != nil {
		return t.cmd.Process.Kill()
	}
	
	return nil
}

// GetFD returns the PTY file descriptor for external use
func (t *Terminal) GetFD() *os.File {
	return t.ptmx
}
