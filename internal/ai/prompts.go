package ai

import (
	"BlackBook/internal/terminal"
	"fmt"
	"strings"
)

// BuildCommandAnalysisPrompt builds a prompt for command analysis
func BuildCommandAnalysisPrompt(cmd terminal.CommandEntry, context []terminal.CommandEntry) string {
	var sb strings.Builder
	
	sb.WriteString("You are a penetration testing expert assistant. Analyze the following command and provide insights.\n\n")
	
	// Add context
	if len(context) > 0 {
		sb.WriteString("Previous commands for context:\n")
		for i, c := range context {
			if i >= 3 { // Limit context to last 3 commands
				break
			}
			sb.WriteString(fmt.Sprintf("- %s (Exit: %d)\n", c.Command, c.ExitCode))
		}
		sb.WriteString("\n")
	}
	
	// Add the command being analyzed
	sb.WriteString(fmt.Sprintf("Command to analyze: %s\n", cmd.Command))
	sb.WriteString(fmt.Sprintf("Exit code: %d\n", cmd.ExitCode))
	
	if cmd.Output != "" {
		outputPreview := cmd.Output
		if len(outputPreview) > 500 {
			outputPreview = outputPreview[:500] + "... (truncated)"
		}
		sb.WriteString(fmt.Sprintf("Output:\n%s\n\n", outputPreview))
	}
	
	sb.WriteString("Please provide:\n")
	sb.WriteString("1. Purpose: What this command does\n")
	sb.WriteString("2. Effectiveness: Was this command successful? What does the output tell us?\n")
	sb.WriteString("3. Findings: Any vulnerabilities or interesting information discovered\n")
	sb.WriteString("4. Next Steps: What should be done next based on these results\n")
	sb.WriteString("\nKeep your response concise and actionable (max 300 words).")
	
	return sb.String()
}

// BuildNextStepPrompt builds a prompt for next step suggestion
func BuildNextStepPrompt(context []terminal.CommandEntry) string {
	var sb strings.Builder
	
	sb.WriteString("You are a penetration testing expert. Based on the command history, suggest the next logical step.\n\n")
	
	if len(context) == 0 {
		sb.WriteString("No commands have been executed yet.\n")
		sb.WriteString("Suggest initial reconnaissance commands to start a penetration test.\n")
	} else {
		sb.WriteString("Recent command history:\n")
		start := len(context) - 5
		if start < 0 {
			start = 0
		}
		
		for _, cmd := range context[start:] {
			sb.WriteString(fmt.Sprintf("- %s (Exit: %d)\n", cmd.Command, cmd.ExitCode))
		}
		
		// Add last successful command output if available
		for i := len(context) - 1; i >= 0; i-- {
			if context[i].ExitCode == 0 && context[i].Output != "" {
				outputPreview := context[i].Output
				if len(outputPreview) > 300 {
					outputPreview = outputPreview[:300] + "..."
				}
				sb.WriteString(fmt.Sprintf("\nLast successful output:\n%s\n", outputPreview))
				break
			}
		}
	}
	
	sb.WriteString("\nSuggest:\n")
	sb.WriteString("1. The next most logical command to run\n")
	sb.WriteString("2. Why this command is recommended\n")
	sb.WriteString("3. What information we're looking for\n")
	sb.WriteString("\nKeep response under 200 words. Format as actionable advice.")
	
	return sb.String()
}

// BuildScriptGenerationPrompt builds a prompt for script generation
func BuildScriptGenerationPrompt(request string, context []terminal.CommandEntry) string {
	var sb strings.Builder
	
	sb.WriteString("You are a penetration testing expert. Generate a bash script based on the request.\n\n")
	sb.WriteString(fmt.Sprintf("Request: %s\n\n", request))
	
	if len(context) > 0 {
		sb.WriteString("Context from current session:\n")
		for i := len(context) - 3; i < len(context) && i >= 0; i++ {
			sb.WriteString(fmt.Sprintf("- %s\n", context[i].Command))
		}
		sb.WriteString("\n")
	}
	
	sb.WriteString("Generate a complete, working bash script that:\n")
	sb.WriteString("1. Is safe to execute (includes error handling)\n")
	sb.WriteString("2. Has clear comments explaining each step\n")
	sb.WriteString("3. Uses best practices for penetration testing\n")
	sb.WriteString("4. Outputs results in a readable format\n")
	sb.WriteString("\nProvide only the script code with comments. Keep it under 50 lines if possible.")
	
	return sb.String()
}

// BuildOutputExplanationPrompt builds a prompt for output explanation
func BuildOutputExplanationPrompt(output string) string {
	var sb strings.Builder
	
	sb.WriteString("You are a penetration testing expert. Explain what this command output means.\n\n")
	
	outputPreview := output
	if len(outputPreview) > 1000 {
		outputPreview = outputPreview[:1000] + "... (truncated)"
	}
	
	sb.WriteString(fmt.Sprintf("Output:\n%s\n\n", outputPreview))
	sb.WriteString("Explain:\n")
	sb.WriteString("1. What this output indicates\n")
	sb.WriteString("2. Any security findings or vulnerabilities\n")
	sb.WriteString("3. What actions to take based on this output\n")
	sb.WriteString("\nKeep response under 250 words.")
	
	return sb.String()
}

// GetPentestPrompts returns common pentest prompts
func GetPentestPrompts() map[string]string {
	return map[string]string{
		"network_scan":    "Generate a comprehensive network scanning script using nmap",
		"web_enum":        "Create a web application enumeration script",
		"password_attack": "Generate a password attack script with best practices",
		"privilege_esc":   "Suggest privilege escalation checks for Linux",
		"exploit_dev":     "Create a basic buffer overflow exploit template",
	}
}
