import { useState, useEffect } from 'react';
import './AIPanel.css';
import {
  GetAISuggestion,
  GenerateScript,
  CheckOllamaStatus,
} from '../../wailsjs/go/main/App';

interface AIPanelProps {
  analysisText?: string;
  isAnalyzing?: boolean;
}

export default function AIPanel({ analysisText, isAnalyzing }: AIPanelProps) {
  const [activeTab, setActiveTab] = useState<'analysis' | 'suggest' | 'script'>('analysis');
  const [suggestion, setSuggestion] = useState<string>('');
  const [scriptRequest, setScriptRequest] = useState<string>('');
  const [generatedScript, setGeneratedScript] = useState<string>('');
  const [isLoading, setIsLoading] = useState(false);
  const [ollamaAvailable, setOllamaAvailable] = useState<boolean>(false);

  useEffect(() => {
    // Check Ollama status on mount
    checkOllamaStatus();
  }, []);

  const checkOllamaStatus = async () => {
    try {
      const available = await CheckOllamaStatus();
      setOllamaAvailable(available);
    } catch (err) {
      console.error('Failed to check Ollama status:', err);
      setOllamaAvailable(false);
    }
  };

  const handleGetSuggestion = async () => {
    setIsLoading(true);
    try {
      const result = await GetAISuggestion();
      setSuggestion(result);
    } catch (err) {
      setSuggestion(`Error: ${err}`);
    } finally {
      setIsLoading(false);
    }
  };

  const handleGenerateScript = async () => {
    if (!scriptRequest.trim()) return;

    setIsLoading(true);
    try {
      const result = await GenerateScript(scriptRequest);
      setGeneratedScript(result);
    } catch (err) {
      setGeneratedScript(`Error: ${err}`);
    } finally {
      setIsLoading(false);
    }
  };

  const renderOllamaWarning = () => {
    if (ollamaAvailable) return null;

    return (
      <div className="ollama-warning">
        <h4>⚠️ Ollama Not Available</h4>
        <p>AI features require Ollama to be running locally.</p>
        <ol>
          <li>Visit <a href="https://ollama.ai" target="_blank" rel="noopener noreferrer">ollama.ai</a></li>
          <li>Install Ollama for your platform</li>
          <li>Run: <code>ollama pull llama3</code></li>
          <li>Start Ollama service</li>
          <li>Refresh this application</li>
        </ol>
      </div>
    );
  };

  return (
    <div className="ai-panel">
      <div className="ai-panel-header">
        <h3>AI Assistant</h3>
        <div className="ai-status">
          <span className={`status-dot ${ollamaAvailable ? 'active' : 'inactive'}`}></span>
          <span className="status-text">{ollamaAvailable ? 'Connected' : 'Offline'}</span>
        </div>
      </div>

      <div className="ai-tabs">
        <button
          className={`ai-tab ${activeTab === 'analysis' ? 'active' : ''}`}
          onClick={() => setActiveTab('analysis')}
        >
          Analysis
        </button>
        <button
          className={`ai-tab ${activeTab === 'suggest' ? 'active' : ''}`}
          onClick={() => setActiveTab('suggest')}
        >
          Suggestions
        </button>
        <button
          className={`ai-tab ${activeTab === 'script' ? 'active' : ''}`}
          onClick={() => setActiveTab('script')}
        >
          Script Gen
        </button>
      </div>

      <div className="ai-content">
        {renderOllamaWarning()}

        {activeTab === 'analysis' && (
          <div className="ai-section">
            <h4>Command Analysis</h4>
            {isAnalyzing && <div className="loading">Analyzing command...</div>}
            {analysisText ? (
              <div className="analysis-result">
                <pre>{analysisText}</pre>
              </div>
            ) : (
              <p className="empty-state">
                Click on any command in the terminal history to analyze it.
              </p>
            )}
          </div>
        )}

        {activeTab === 'suggest' && (
          <div className="ai-section">
            <h4>Next Step Suggestions</h4>
            <button
              className="ai-button"
              onClick={handleGetSuggestion}
              disabled={isLoading || !ollamaAvailable}
            >
              {isLoading ? 'Generating...' : 'Get Suggestion'}
            </button>
            {suggestion && (
              <div className="suggestion-result">
                <pre>{suggestion}</pre>
              </div>
            )}
          </div>
        )}

        {activeTab === 'script' && (
          <div className="ai-section">
            <h4>Script Generator</h4>
            <textarea
              className="script-input"
              placeholder="Describe the script you need (e.g., 'Network scan for open ports')"
              value={scriptRequest}
              onChange={(e) => setScriptRequest(e.target.value)}
              rows={3}
            />
            <button
              className="ai-button"
              onClick={handleGenerateScript}
              disabled={isLoading || !scriptRequest.trim() || !ollamaAvailable}
            >
              {isLoading ? 'Generating...' : 'Generate Script'}
            </button>
            {generatedScript && (
              <div className="script-result">
                <pre>{generatedScript}</pre>
              </div>
            )}
          </div>
        )}
      </div>
    </div>
  );
}
