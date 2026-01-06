import { CommandEntry } from '../types';
import './CommandItem.css';

interface CommandItemProps {
  command: CommandEntry;
  onClick: (commandId: string) => void;
  isSelected?: boolean;
}

export default function CommandItem({ command, onClick, isSelected }: CommandItemProps) {
  const getStatusColor = (exitCode: number) => {
    if (exitCode === 0) return 'success';
    if (exitCode === -1) return 'pending';
    return 'error';
  };

  const formatDuration = (ms: number) => {
    if (ms < 1000) return `${ms}ms`;
    return `${(ms / 1000).toFixed(2)}s`;
  };

  const formatTime = (timeStr: string) => {
    try {
      const date = new Date(timeStr);
      return date.toLocaleTimeString();
    } catch {
      return timeStr;
    }
  };

  return (
    <div
      className={`command-item ${isSelected ? 'selected' : ''} ${getStatusColor(command.exitCode)}`}
      onClick={() => onClick(command.id)}
    >
      <div className="command-header">
        <span className="command-time">{formatTime(command.startTime)}</span>
        <span className={`command-status ${getStatusColor(command.exitCode)}`}>
          {command.exitCode === 0 ? '✓' : command.exitCode === -1 ? '⋯' : '✗'}
        </span>
      </div>
      <div className="command-text">{command.command}</div>
      <div className="command-footer">
        <span className="command-duration">{formatDuration(command.duration)}</span>
        <span className="command-exit-code">Exit: {command.exitCode}</span>
      </div>
    </div>
  );
}
