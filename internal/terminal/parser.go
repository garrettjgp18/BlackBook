package terminal

import (
	"regexp"
	"strings"
)

// ParseCommandLine attempts to parse a command line into command and arguments
func ParseCommandLine(line string) (string, []string) {
	parts := strings.Fields(line)
	if len(parts) == 0 {
		return "", nil
	}
	return parts[0], parts[1:]
}

// ExtractIPAddresses extracts IP addresses from text
func ExtractIPAddresses(text string) []string {
	ipPattern := regexp.MustCompile(`\b(?:\d{1,3}\.){3}\d{1,3}\b`)
	return ipPattern.FindAllString(text, -1)
}

// ExtractURLs extracts URLs from text
func ExtractURLs(text string) []string {
	urlPattern := regexp.MustCompile(`https?://[^\s]+`)
	return urlPattern.FindAllString(text, -1)
}

// ExtractPorts extracts port numbers from text
func ExtractPorts(text string) []string {
	portPattern := regexp.MustCompile(`(?:port|PORT)\s+(\d+)`)
	matches := portPattern.FindAllStringSubmatch(text, -1)
	ports := make([]string, 0, len(matches))
	for _, match := range matches {
		if len(match) > 1 {
			ports = append(ports, match[1])
		}
	}
	return ports
}

// IsSecurityTool checks if a command is a known security tool
func IsSecurityTool(command string) bool {
	tools := []string{
		"nmap", "masscan", "nikto", "sqlmap", "metasploit",
		"burp", "dirb", "gobuster", "ffuf", "hydra",
		"john", "hashcat", "aircrack-ng", "wireshark", "tcpdump",
		"netcat", "nc", "socat", "crackmapexec", "enum4linux",
		"smbclient", "rpcclient", "ldapsearch", "whatweb", "wpscan",
	}
	
	cmdLower := strings.ToLower(command)
	for _, tool := range tools {
		if strings.Contains(cmdLower, tool) {
			return true
		}
	}
	return false
}

// CleanANSI removes ANSI escape codes from text
func CleanANSI(text string) string {
	ansiPattern := regexp.MustCompile(`\x1b\[[0-9;]*[a-zA-Z]`)
	return ansiPattern.ReplaceAllString(text, "")
}

// TruncateOutput truncates output to a maximum length
func TruncateOutput(output string, maxLines int) string {
	lines := strings.Split(output, "\n")
	if len(lines) <= maxLines {
		return output
	}
	
	truncated := strings.Join(lines[:maxLines], "\n")
	return truncated + "\n... (output truncated)"
}
