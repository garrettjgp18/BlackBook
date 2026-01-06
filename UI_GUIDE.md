# BlackBook UI Guide

This document describes the user interface and user experience of the BlackBook Penetration Testing Toolkit.

## Application Layout

```
┌─────────────────────────────────────────────────────────────────────────┐
│  BlackBook                                                              │
│  Penetration Testing Toolkit                                           │
├─────────────────────────────────────────────────────────────────────────┤
│                                 │                                       │
│   Interactive Terminal (60%)    │   AI Assistant Panel (40%)           │
│                                 │                                       │
│  ┌──────────────────────────┐  │  ┌───────────────────────────────┐  │
│  │ Terminal Header          │  │  │ AI Assistant        [●] Online │  │
│  │ ● ● ●                    │  │  ├───────────────────────────────┤  │
│  ├──────────────────────────┤  │  │ [Analysis] [Suggest] [Script] │  │
│  │                          │  │  ├───────────────────────────────┤  │
│  │ $ ls -la                 │  │  │                               │  │
│  │ drwxr-xr-x  5 user group│  │  │  Command Analysis             │  │
│  │ -rw-r--r--  1 user group│  │  │                               │  │
│  │                          │  │  │  Click on any command in      │  │
│  │ $ nmap -sV 192.168.1.1  │  │  │  the terminal to analyze it   │  │
│  │ Starting Nmap...         │  │  │  with AI assistance.          │  │
│  │ PORT    STATE SERVICE    │  │  │                               │  │
│  │ 22/tcp  open  ssh        │  │  │                               │  │
│  │ 80/tcp  open  http       │  │  │                               │  │
│  │                          │  │  │                               │  │
│  │ $▊                       │  │  │                               │  │
│  │                          │  │  │                               │  │
│  └──────────────────────────┘  │  └───────────────────────────────┘  │
│                                 │                                       │
└─────────────────────────────────┴───────────────────────────────────────┘
```

## Components Overview

### 1. Application Header

**Location**: Top of the window
**Purpose**: Branding and application title

Features:
- Application name "BlackBook" with gradient styling
- Subtitle "Penetration Testing Toolkit"
- Dark background (#252526) with subtle border

### 2. Interactive Terminal (Left Panel - 60%)

**Purpose**: Full-featured terminal for penetration testing commands

#### Features:
- **Real PTY Integration**: Connects to bash/zsh/PowerShell based on OS
- **Full Terminal Emulation**: Uses xterm.js for authentic terminal experience
- **Command Tracking**: Every command is automatically logged
- **Live Output**: Real-time streaming of command output
- **Terminal Controls**: MacOS-style traffic light buttons (● ● ●)

#### Visual Design:
- Background: Dark (#1e1e1e)
- Text: Light gray (#d4d4d4)
- Cursor: Green blinking cursor
- Font: Monospace (Consolas, Monaco, Courier New)
- Header: Slightly lighter (#2d2d30) with rounded top corners

#### Command History:
All commands are tracked with:
- Timestamp
- Full command text
- Complete output (stdout/stderr)
- Exit code
- Execution duration

### 3. AI Assistant Panel (Right Panel - 40%)

**Purpose**: AI-powered analysis and assistance for penetration testing

#### Status Indicator:
- Green dot (●) + "Connected" when Ollama is available
- Gray dot + "Offline" when Ollama is unavailable

#### Three Tabs:

##### **Analysis Tab**
- Shows detailed analysis of selected commands
- Click any command in terminal history to analyze
- AI provides:
  - Command purpose explanation
  - Effectiveness assessment
  - Security findings from output
  - Recommended next steps

##### **Suggestions Tab**
- "Get Suggestion" button
- AI analyzes entire command history
- Provides next logical pentesting step
- Context-aware recommendations

##### **Script Gen Tab**
- Text input for describing needed script
- "Generate Script" button
- AI creates custom bash/PowerShell scripts
- Includes comments and error handling

#### Visual Design:
- Background: Slightly lighter dark (#252526)
- Tabs: Active tab highlighted with blue underline (#007acc)
- Buttons: Blue (#0e639c) with hover effects
- Results: Displayed in code-style boxes with dark background

#### Ollama Warning:
When Ollama is not available, displays helpful setup instructions:
- Visit ollama.ai
- Installation steps
- How to pull models
- How to start the service

### 4. Color Scheme

**Dark Theme (Hacker Aesthetic)**

Primary Colors:
- Background: `#1e1e1e` (Dark charcoal)
- Panel Background: `#252526` (Slightly lighter)
- Header Background: `#2d2d30`
- Border: `#3e3e42`

Text Colors:
- Primary Text: `#d4d4d4` (Light gray)
- Secondary Text: `#999999` (Medium gray)
- Title/Headers: `#cccccc` to `#ffffff`

Accent Colors:
- Primary Accent: `#007acc` (Blue)
- Success: `#27c93f` (Green)
- Error: `#f14c4c` (Red)
- Warning: `#f5f543` (Yellow)
- Info: `#4fc3f7` (Light blue)

Gradient:
- Title Gradient: Purple to blue (`#667eea` to `#764ba2`)

### 5. Typography

- **Headings**: -apple-system, BlinkMacSystemFont, Segoe UI, Roboto
- **Terminal**: Consolas, Monaco, Courier New (monospace)
- **Code Blocks**: Same as terminal
- **Body Text**: System font stack

Font Sizes:
- App Title: 24px
- Section Headers: 16-18px
- Body Text: 13-14px
- Small Text: 11-12px

### 6. Interactive Elements

#### Buttons:
- Primary: Blue background (#0e639c)
- Hover: Lighter blue (#1177bb)
- Disabled: Gray (#3e3e42) with reduced opacity
- Border Radius: 4px
- Padding: 8-12px

#### Input Fields:
- Background: Dark (#1e1e1e)
- Border: Gray (#3e3e42)
- Focus: Blue border (#007acc)
- Text: Light gray (#d4d4d4)

#### Terminal Dots (Traffic Lights):
- Red: `#ff5f56`
- Yellow: `#ffbd2e`
- Green: `#27c93f`
- Size: 12px diameter

### 7. Command Item Component

Displayed in command history (when implemented):

```
┌────────────────────────────────────┐
│ 15:23:45                        ✓  │  <- Time and status icon
│ nmap -sV 192.168.1.1               │  <- Command text
│ 2.5s                    Exit: 0    │  <- Duration and exit code
└────────────────────────────────────┘
```

Status Icons:
- ✓ (Green) = Success (exit code 0)
- ✗ (Red) = Failed (exit code != 0)
- ⋯ (Yellow) = Pending/Running (exit code -1)

Border Accent:
- Left border (3px) colored by status
- Clickable with hover effects
- Selected state with blue outline

### 8. Report Viewer Component

Located within AI Panel (additional tab or modal):

```
┌────────────────────────────────────┐
│ Report Preview   [Refresh] [Export]│
├────────────────────────────────────┤
│ Penetration Test Report            │
│ Generated: 2024-01-06              │
│                                    │
│ ┌─────────┬─────────┬──────────┐  │
│ │Commands │Findings │  Start   │  │
│ │   45    │    3    │ 14:30:00 │  │
│ └─────────┴─────────┴──────────┘  │
│                                    │
│ Findings:                          │
│ ┌──────────────────────────────┐  │
│ │ 1. Open SSH Port - High      │  │
│ │ 2. HTTP Service - Medium     │  │
│ │ 3. Default Credentials - Low │  │
│ └──────────────────────────────┘  │
└────────────────────────────────────┘
```

Finding Severity Colors:
- Critical: Red border (#d32f2f)
- High: Orange border (#f57c00)
- Medium: Yellow border (#fbc02d)
- Low: Green border (#388e3c)
- Info: Blue border (#1976d2)

### 9. Responsive Behavior

**Window Resize:**
- Terminal panel maintains 60% width
- AI panel maintains 40% width
- Minimum window size: 1024x768

**Terminal Resize:**
- Automatically adjusts PTY size
- xterm.js FitAddon handles terminal fitting
- Proper handling of window resize events

**Mobile/Tablet** (Future consideration):
- Stack panels vertically
- Terminal on top, AI panel below
- Full width for both

### 10. User Interactions

**Terminal:**
- Type commands naturally
- All standard terminal keyboard shortcuts work
- Ctrl+C, Ctrl+D, arrows, etc.
- Command history with up/down arrows
- Tab completion (if shell supports)

**AI Panel:**
- Click tabs to switch between features
- Click "Get Suggestion" for AI analysis
- Type in text area for script generation
- Click "Generate Script" to create code

**Command History** (Future enhancement):
- Click any command to select
- AI automatically analyzes selected command
- Visual highlight on selected command

### 11. Accessibility

**Keyboard Navigation:**
- Tab through interactive elements
- Enter to activate buttons
- Escape to close modals/dialogs

**Screen Readers:**
- Proper ARIA labels on components
- Semantic HTML structure
- Status announcements for AI updates

**Visual:**
- High contrast dark theme
- Color is not the only indicator (icons + text)
- Readable font sizes (minimum 11px)

### 12. Loading States

**AI Processing:**
- "Analyzing command..." text
- Subtle animation or spinner
- Button disabled during processing

**Terminal:**
- Cursor blinks to show activity
- Output streams in real-time
- No blocking operations

**Report Generation:**
- "Loading..." text
- Disabled buttons during generation
- Progress indication

### 13. Error States

**AI Unavailable:**
- Clear warning message with orange/red styling
- Step-by-step setup instructions
- Link to Ollama documentation

**Terminal Error:**
- Error messages in red
- Non-blocking - terminal remains functional
- Error logged to console

**Network Issues:**
- Graceful degradation
- Offline mode for terminal
- Clear indication of what's unavailable

## User Workflows

### Basic Pentesting Session

1. User launches BlackBook
2. Terminal opens with default shell
3. User types reconnaissance command: `nmap -sV 192.168.1.1`
4. Output streams to terminal in real-time
5. Command is tracked in history
6. User switches to AI "Suggestions" tab
7. Clicks "Get Suggestion"
8. AI analyzes the nmap results and suggests next steps
9. User continues with suggested commands
10. At any time, can export report with all findings

### Script Generation Workflow

1. User needs a custom script
2. Switches to "Script Gen" tab
3. Types: "Create a script to scan ports 1-1000 and save results"
4. Clicks "Generate Script"
5. AI creates bash script with comments
6. User copies script to terminal
7. Executes script
8. Results are captured in report

### Report Export Workflow

1. Throughout testing session, commands are tracked
2. User wants to generate report
3. Clicks on Report viewer
4. Clicks "Refresh" to generate current report
5. Reviews summary and findings
6. Clicks "Export MD" to download Markdown file
7. Report includes timeline, methodology, and recommendations

## Technical Implementation Notes

### Performance
- Terminal output is streamed (not buffered)
- AI requests are async (non-blocking)
- Frontend updates are React-managed
- WebGL acceleration for terminal rendering

### Data Flow
1. User types in terminal → Sent to Go backend via Wails
2. Go executes command in PTY → Output streamed back
3. Go tracks command → Stored in session history
4. User requests AI analysis → Go sends context to Ollama
5. Ollama responds → Go returns to frontend
6. Frontend displays result → User sees analysis

### Security Considerations
- PTY runs with user permissions (no privilege escalation)
- AI runs locally (no data sent to cloud)
- Command history stored in memory (not persisted by default)
- Report export asks for save location

## Future UI Enhancements

- Draggable panel divider for custom split ratio
- Command history sidebar with search/filter
- Tool templates as quick-launch buttons
- Settings panel for configuration
- Session save/load functionality
- Dark/Light theme toggle
- Customizable color schemes
- Terminal tabs for multiple sessions

---

**Note**: This UI guide describes the implemented MVP. The actual rendered application may have slight variations based on system fonts, OS theme, and WebView rendering engine.
