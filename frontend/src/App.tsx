import { useState } from 'react';
import './App.css';
import Terminal from './components/Terminal';
import AIPanel from './components/AIPanel';

function App() {
  const [analysisText] = useState<string>('');
  const [isAnalyzing] = useState(false);

  return (
    <div className="app">
      <div className="app-header">
        <h1 className="app-title">BlackBook</h1>
        <p className="app-subtitle">Penetration Testing Toolkit</p>
      </div>
      
      <div className="app-container">
        <div className="terminal-panel">
          <Terminal />
        </div>
        
        <div className="ai-panel-container">
          <AIPanel analysisText={analysisText} isAnalyzing={isAnalyzing} />
        </div>
      </div>
    </div>
  );
}

export default App;
