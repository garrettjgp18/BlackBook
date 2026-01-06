import { useState } from 'react';
import './ReportViewer.css';
import { GetCurrentReport, ExportReport } from '../../wailsjs/go/main/App';
import { report } from '../../wailsjs/go/models';

export default function ReportViewer() {
  const [reportData, setReportData] = useState<report.Report | null>(null);
  const [isLoading, setIsLoading] = useState(false);

  const handleLoadReport = async () => {
    setIsLoading(true);
    try {
      const result = await GetCurrentReport();
      setReportData(result);
    } catch (err) {
      console.error('Failed to load report:', err);
    } finally {
      setIsLoading(false);
    }
  };

  const handleExport = async (format: string) => {
    setIsLoading(true);
    try {
      const content = await ExportReport(format);
      
      // Download the file
      const blob = new Blob([content], { type: 'text/plain' });
      const url = URL.createObjectURL(blob);
      const a = document.createElement('a');
      a.href = url;
      a.download = `pentest-report-${new Date().getTime()}.${format === 'html' ? 'html' : 'md'}`;
      document.body.appendChild(a);
      a.click();
      document.body.removeChild(a);
      URL.revokeObjectURL(url);
    } catch (err) {
      console.error('Failed to export report:', err);
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <div className="report-viewer">
      <div className="report-header">
        <h4>Report Preview</h4>
        <div className="report-actions">
          <button onClick={handleLoadReport} disabled={isLoading}>
            {isLoading ? 'Loading...' : 'Refresh'}
          </button>
          <button onClick={() => handleExport('markdown')} disabled={isLoading || !reportData}>
            Export MD
          </button>
        </div>
      </div>

      <div className="report-content">
        {!reportData ? (
          <p className="empty-state">Click "Refresh" to generate a report from your session</p>
        ) : (
          <div className="report-summary">
            <h3>{reportData.title}</h3>
            <div className="report-stats">
              <div className="stat">
                <span className="stat-label">Commands:</span>
                <span className="stat-value">{reportData.commandTimeline.length}</span>
              </div>
              <div className="stat">
                <span className="stat-label">Findings:</span>
                <span className="stat-value">{reportData.findings.length}</span>
              </div>
              <div className="stat">
                <span className="stat-label">Start:</span>
                <span className="stat-value">{new Date(reportData.startTime).toLocaleString()}</span>
              </div>
            </div>
            
            {reportData.findings.length > 0 && (
              <div className="findings-summary">
                <h5>Findings</h5>
                {reportData.findings.map((finding, idx) => (
                  <div key={finding.id} className={`finding-item ${finding.severity.toLowerCase()}`}>
                    <span className="finding-number">{idx + 1}</span>
                    <span className="finding-title">{finding.title}</span>
                    <span className="finding-severity">{finding.severity}</span>
                  </div>
                ))}
              </div>
            )}
          </div>
        )}
      </div>
    </div>
  );
}
