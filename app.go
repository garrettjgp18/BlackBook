package main

import (
	"BlackBook/internal/ai"
	"BlackBook/internal/report"
	"BlackBook/internal/terminal"
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"sync"
	"time"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx            context.Context
	terminal       *terminal.Terminal
	aiClient       *ai.OllamaClient
	reportGen      *report.Generator
	mu             sync.RWMutex
	outputBuffer   []byte
	outputChan     chan []byte
	stopOutputChan chan bool
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{
		outputChan:     make(chan []byte, 1000),
		stopOutputChan: make(chan bool),
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	
	// Initialize terminal
	term, err := terminal.NewTerminal()
	if err != nil {
		runtime.LogError(ctx, fmt.Sprintf("Failed to create terminal: %v", err))
		return
	}
	a.terminal = term
	
	// Initialize AI client
	a.aiClient = ai.NewOllamaClient("http://localhost:11434", "llama3")
	
	// Start output streaming
	go a.streamTerminalOutput()
	
	runtime.LogInfo(ctx, "BlackBook started successfully")
}

// streamTerminalOutput streams terminal output to the frontend
func (a *App) streamTerminalOutput() {
	buffer := make([]byte, 4096)
	
	for {
		select {
		case <-a.stopOutputChan:
			return
		default:
			n, err := a.terminal.Read(buffer)
			if err != nil {
				if err != io.EOF {
					runtime.LogError(a.ctx, fmt.Sprintf("Error reading terminal: %v", err))
				}
				time.Sleep(100 * time.Millisecond)
				continue
			}
			
			if n > 0 {
				data := make([]byte, n)
				copy(data, buffer[:n])
				
				// Store in buffer
				a.mu.Lock()
				a.outputBuffer = append(a.outputBuffer, data...)
				a.mu.Unlock()
				
				// Emit to frontend
				encoded := base64.StdEncoding.EncodeToString(data)
				runtime.EventsEmit(a.ctx, "terminal:output", encoded)
			}
		}
	}
}

// Shutdown is called when the app is closing
func (a *App) Shutdown() {
	if a.terminal != nil {
		a.terminal.Close()
	}
	close(a.stopOutputChan)
}

// WriteToTerminal writes input to the terminal
func (a *App) WriteToTerminal(input string) error {
	if a.terminal == nil {
		return fmt.Errorf("terminal not initialized")
	}
	
	return a.terminal.WriteString(input)
}

// ResizeTerminal resizes the terminal
func (a *App) ResizeTerminal(rows, cols int) error {
	if a.terminal == nil {
		return fmt.Errorf("terminal not initialized")
	}
	
	return a.terminal.Resize(uint16(rows), uint16(cols))
}

// GetCommandHistory returns the command history
func (a *App) GetCommandHistory() []terminal.CommandEntry {
	if a.terminal == nil {
		return []terminal.CommandEntry{}
	}
	
	return a.terminal.GetHistory()
}

// AddCommand adds a command to history
func (a *App) AddCommand(command string) string {
	if a.terminal == nil {
		return ""
	}
	
	return a.terminal.AddCommand(command)
}

// UpdateCommandOutput updates command output
func (a *App) UpdateCommandOutput(commandID string, output string, exitCode int) {
	if a.terminal == nil {
		return
	}
	
	a.terminal.UpdateCommandOutput(commandID, output, exitCode)
}

// GetTerminalOutput returns accumulated terminal output
func (a *App) GetTerminalOutput() string {
	a.mu.RLock()
	defer a.mu.RUnlock()
	
	return string(a.outputBuffer)
}

// AnalyzeCommand analyzes a specific command using AI
func (a *App) AnalyzeCommand(commandID string) (string, error) {
	if a.terminal == nil {
		return "", fmt.Errorf("terminal not initialized")
	}
	
	// Check if AI is available
	if !a.aiClient.IsAvailable() {
		return "AI analysis is unavailable. Please ensure Ollama is running at http://localhost:11434\n\nTo install Ollama:\n1. Visit https://ollama.ai\n2. Download and install Ollama\n3. Run: ollama pull llama3\n4. Start Ollama service", nil
	}
	
	// Find the command
	history := a.terminal.GetHistory()
	var targetCmd *terminal.CommandEntry
	
	for i := range history {
		if history[i].ID == commandID {
			targetCmd = &history[i]
			break
		}
	}
	
	if targetCmd == nil {
		return "", fmt.Errorf("command not found")
	}
	
	// Get context (previous commands)
	contextSize := 3
	startIdx := 0
	for i := range history {
		if history[i].ID == commandID {
			startIdx = i - contextSize
			if startIdx < 0 {
				startIdx = 0
			}
			break
		}
	}
	
	var context []terminal.CommandEntry
	for i := range history {
		if history[i].ID == commandID {
			context = history[startIdx:i]
			break
		}
	}
	
	return a.aiClient.AnalyzeCommand(*targetCmd, context)
}

// GetAISuggestion gets AI suggestion for next step
func (a *App) GetAISuggestion() (string, error) {
	if a.terminal == nil {
		return "", fmt.Errorf("terminal not initialized")
	}
	
	// Check if AI is available
	if !a.aiClient.IsAvailable() {
		return "AI suggestions are unavailable. Please ensure Ollama is running at http://localhost:11434", nil
	}
	
	history := a.terminal.GetHistory()
	return a.aiClient.SuggestNextStep(history)
}

// GenerateScript generates a script based on request
func (a *App) GenerateScript(request string) (string, error) {
	if a.terminal == nil {
		return "", fmt.Errorf("terminal not initialized")
	}
	
	// Check if AI is available
	if !a.aiClient.IsAvailable() {
		return "Script generation is unavailable. Please ensure Ollama is running.", nil
	}
	
	history := a.terminal.GetHistory()
	return a.aiClient.GenerateScript(request, history)
}

// GetCurrentReport returns the current report
func (a *App) GetCurrentReport() (*report.Report, error) {
	if a.terminal == nil {
		return nil, fmt.Errorf("terminal not initialized")
	}
	
	session := a.terminal.GetSession()
	
	if a.reportGen == nil {
		a.reportGen = report.NewGenerator(session)
	}
	
	rpt := a.reportGen.GenerateReport("Penetration Test Report - " + time.Now().Format("2006-01-02"))
	return rpt, nil
}

// ExportReport exports the report in the specified format
func (a *App) ExportReport(format string) (string, error) {
	rpt, err := a.GetCurrentReport()
	if err != nil {
		return "", err
	}
	
	if a.reportGen == nil {
		session := a.terminal.GetSession()
		a.reportGen = report.NewGenerator(session)
	}
	
	switch format {
	case "markdown", "md":
		return a.reportGen.ExportMarkdown(rpt), nil
	case "html":
		return a.reportGen.ExportHTML(rpt), nil
	default:
		return "", fmt.Errorf("unsupported format: %s", format)
	}
}

// CheckOllamaStatus checks if Ollama is available
func (a *App) CheckOllamaStatus() bool {
	if a.aiClient == nil {
		return false
	}
	return a.aiClient.IsAvailable()
}
