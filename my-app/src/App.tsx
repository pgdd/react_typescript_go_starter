import React, { useEffect, useState } from 'react';
import './App.css';

const App: React.FC = () => {
  const [data, setData] = useState<string>('');
  const [error, setError] = useState<string>('');

  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await fetch('http://localhost:8080');
        if (!response.ok) {
          throw new Error(`HTTP error! Status: ${response.status}`);
        }
        const data = await response.text();
        setData(data);
      } catch (error) {
        setError(error instanceof Error ? error.message : 'Unknown error occurred');
      }
    };

    fetchData();
  }, []);

  return (
    <div className="App">
      <header className="App-header">
        {error ? (
          <p>Error fetching data: {error}</p>
        ) : (
          <p>Data from Go: {data}</p>
        )}
      </header>
    </div>
  );
};

export default App;
