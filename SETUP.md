# BlackBook - Penetration Testing Toolkit

BlackBook is a robust, serverless penetration testing toolkit built with GoLang and Wails v2, featuring a split-view interface with an interactive terminal and AI-enhanced script generator panel.

## Features

- **Interactive Terminal**: Full PTY-based terminal emulation (bash/zsh/powershell based on OS)
- **AI-Powered Analysis**: Integration with Ollama (local LLM) for command analysis and script generation
- **Command History Tracking**: Track every command with timestamp, output, exit code, and duration
- **Auto Report Generation**: Automatically generate penetration test reports in Markdown/HTML format
- **Split-View Interface**: Terminal on the left, AI assistant on the right
- **Dark Theme**: Professional hacker aesthetic

## Architecture

### Backend (Go)
- `internal/terminal`: PTY management and command history tracking
- `internal/ai`: Ollama client for AI features
- `internal/report`: Report generation and templates
- `app.go`: Main application with Wails bindings

### Frontend (React + TypeScript)
- `Terminal.tsx`: Interactive terminal using xterm.js
- `AIPanel.tsx`: AI assistant interface
- `CommandItem.tsx`: Clickable command history items
- `ReportViewer.tsx`: Report preview and export

## Prerequisites

### Development Dependencies
- Go 1.21 or higher
- Node.js 16+ and npm
- Wails CLI: `go install github.com/wailsapp/wails/v2/cmd/wails@latest`

### Platform-Specific Dependencies

#### Linux
```bash
sudo apt-get install build-essential libgtk-3-dev libwebkit2gtk-4.0-dev
```

#### macOS
```bash
xcode-select --install
```

#### Windows
- Download and install the [WebView2 Runtime](https://developer.microsoft.com/en-us/microsoft-edge/webview2/)
- Install [MinGW-w64](https://www.mingw-w64.org/) or use MSYS2

### Optional: Ollama (for AI features)
```bash
# Install Ollama from https://ollama.ai
# Then pull a model:
ollama pull llama3
```

## Installation

1. Clone the repository:
```bash
git clone https://github.com/garrettjgp18/BlackBook.git
cd BlackBook
```

2. Install Go dependencies:
```bash
go mod download
```

3. Install frontend dependencies:
```bash
cd frontend
npm install
cd ..
```

## Development

### Run in Development Mode
```bash
wails dev
```

This will start the application in development mode with hot-reload for both frontend and backend.

### Build for Production
```bash
wails build
```

The compiled application will be in `build/bin/`.

### Frontend Only Development
```bash
cd frontend
npm run dev
```

## Usage

### Basic Workflow

1. **Launch BlackBook**: The app opens with a split view - terminal on the left, AI panel on the right

2. **Use the Terminal**: Type commands as you would in any terminal
   - All commands are automatically tracked
   - Output is captured and stored
   - Command history is maintained

3. **Get AI Assistance** (requires Ollama):
   - **Analysis Tab**: View automatic command analysis
   - **Suggestions Tab**: Click "Get Suggestion" for next step recommendations
   - **Script Gen Tab**: Describe a script you need and let AI generate it

4. **Generate Reports**:
   - Reports are automatically built in the background
   - Access via the report viewer
   - Export as Markdown or HTML

### AI Features

The AI assistant provides:
- **Command Analysis**: Understanding what each command does and its effectiveness
- **Next Step Suggestions**: Recommendations for the next logical pentesting action
- **Script Generation**: Custom pentesting scripts based on your requirements
- **Context Awareness**: AI understands your session history

### Report Generation

Reports include:
- Executive Summary
- Methodology (auto-detected from commands used)
- Command Timeline
- Findings (if any)
- Recommendations

Export formats:
- Markdown (`.md`)
- HTML (basic formatting)

## Configuration

### Ollama Settings
By default, BlackBook connects to Ollama at `http://localhost:11434` using the `llama3` model.

To use a different model or endpoint, modify `app.go`:
```go
a.aiClient = ai.NewOllamaClient("http://localhost:11434", "mistral")
```

### Terminal Shell
The default shell is automatically detected based on your OS:
- Linux/macOS: `$SHELL` environment variable or `/bin/bash`
- Windows: PowerShell or `cmd.exe`

## Project Structure

```
BlackBook/
├── main.go                     # Wails entry point
├── app.go                      # Core application logic
├── wails.json                  # Wails configuration
├── go.mod / go.sum            # Go dependencies
├── internal/
│   ├── terminal/              # Terminal package
│   │   ├── pty.go            # PTY management
│   │   ├── history.go        # Command history
│   │   ├── parser.go         # Output parsing
│   │   └── models.go         # Data models
│   ├── ai/                    # AI package
│   │   ├── ollama.go         # Ollama client
│   │   ├── prompts.go        # AI prompts
│   │   └── context.go        # Context builder
│   └── report/                # Report package
│       ├── generator.go      # Report generation
│       ├── templates.go      # Report templates
│       └── models.go         # Data models
├── frontend/
│   ├── src/
│   │   ├── App.tsx           # Main app component
│   │   ├── components/       # React components
│   │   ├── types/            # TypeScript types
│   │   └── main.tsx          # React entry
│   ├── package.json
│   └── tsconfig.json
└── build/                     # Build output
```

## Troubleshooting

### Terminal Not Working
- Ensure PTY dependencies are installed on your system
- Check that the application has proper permissions

### AI Features Not Available
1. Verify Ollama is running: `curl http://localhost:11434/api/tags`
2. Check that the model is downloaded: `ollama list`
3. Pull the model if missing: `ollama pull llama3`

### Build Errors
- **Linux**: Install GTK and WebKit dependencies
- **macOS**: Ensure Xcode command line tools are installed
- **Windows**: Install WebView2 Runtime and MinGW-w64

### Frontend Build Issues
```bash
cd frontend
rm -rf node_modules package-lock.json
npm install
npm run build
```

## Known Limitations

This is the MVP (Minimum Viable Product) version. Future enhancements will include:
- Tool templates for common pentesting workflows
- VM integration
- VPN support
- Enhanced report customization
- Command output filtering and searching
- Session saving and loading

## Contributing

This project is primarily for learning GoLang and building penetration testing tools. Contributions, issues, and feature requests are welcome!

## License

This project is for educational and authorized security testing purposes only. Use responsibly and only on systems you have permission to test.

## Acknowledgments

- [Wails](https://wails.io/) - The electron alternative for Go
- [Ollama](https://ollama.ai/) - Local LLM runtime
- [xterm.js](https://xtermjs.org/) - Terminal emulation library
- [creack/pty](https://github.com/creack/pty) - PTY interface for Go

---

**BlackBook** - Making penetration testing more accessible through AI-assisted workflows.
