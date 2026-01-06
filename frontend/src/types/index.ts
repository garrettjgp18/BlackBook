// Type definitions for BlackBook

export interface CommandEntry {
  id: string;
  command: string;
  output: string;
  exitCode: number;
  startTime: string;
  endTime: string;
  duration: number; // milliseconds
}

export interface Session {
  id: string;
  startTime: string;
  history: CommandEntry[];
  active: boolean;
}

export interface Finding {
  id: string;
  title: string;
  severity: 'Critical' | 'High' | 'Medium' | 'Low' | 'Info';
  description: string;
  evidence: string;
  command: string;
  timestamp: string;
  remediation: string;
}

export interface Report {
  id: string;
  title: string;
  startTime: string;
  endTime: string;
  findings: Finding[];
  commandTimeline: CommandEntry[];
  summary: string;
  methodology: string;
  recommendations: string[];
}

export interface AIAnalysis {
  purpose?: string;
  effectiveness?: string;
  findings?: string;
  nextSteps?: string;
  fullText: string;
}

export interface PentestPrompt {
  id: string;
  title: string;
  description: string;
  prompt: string;
}
