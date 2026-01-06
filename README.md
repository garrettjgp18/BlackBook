# BlackBook - Penetration Testing Toolkit

> An AI-enhanced penetration testing toolkit with interactive terminal and automated reporting

BlackBook is an ambitious project designed to guide penetration testers through the full process of planning, conducting, and documenting penetration tests. Built with GoLang and Wails v2, it combines a powerful terminal interface with AI-powered assistance and automatic report generation.

## ✨ Key Features

- **🖥️ Interactive Terminal**: Full PTY-based terminal with command history tracking
- **🤖 AI Assistant**: Powered by Ollama (local LLM) for command analysis and suggestions
- **📊 Auto Reports**: Generate professional penetration test reports automatically
- **🎨 Split-View UI**: Terminal and AI assistant side-by-side
- **⚡ Serverless**: No external servers required - everything runs locally
- **🔒 Privacy First**: All data stays on your machine

## 🚀 Quick Start

See [SETUP.md](SETUP.md) for detailed installation and usage instructions.

```bash
# Clone repository
git clone https://github.com/garrettjgp18/BlackBook.git
cd BlackBook

# Install dependencies
go mod download
cd frontend && npm install && cd ..

# Run in development mode
wails dev
```

## 🛠️ Technology Stack

- **Backend**: GoLang with Wails v2
- **Frontend**: React with TypeScript
- **Terminal**: xterm.js with PTY backend
- **AI**: Ollama (local LLM - completely free and offline)
- **Styling**: Custom dark theme

## 📖 Documentation

- [Setup Guide](SETUP.md) - Installation and configuration
- [Architecture Overview](SETUP.md#architecture) - Technical details
- [Usage Guide](SETUP.md#usage) - How to use BlackBook

## 🎯 Roadmap

This is the MVP version. Future enhancements include:
- [ ] Tool templates for common pentesting workflows
- [ ] VM integration
- [ ] VPN support
- [ ] Enhanced report customization
- [ ] Session saving and loading
- [ ] Command output filtering and searching

## 📝 License

For educational and authorized security testing purposes only. Use responsibly.

## 🙏 Acknowledgments

- [Wails](https://wails.io/) - Go desktop framework
- [Ollama](https://ollama.ai/) - Local LLM runtime
- [xterm.js](https://xtermjs.org/) - Terminal emulation
