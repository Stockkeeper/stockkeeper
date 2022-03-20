import React, { useCallback, useEffect, useState } from 'react';
import logo from './logo.svg';
import './App.css';
import * as admin from './lib/services/admin';

function App() {
  const [password, setPassword] = useState('');

  const fetchData = useCallback(async () => {
    const data = await admin.getAdminPassword();
    setPassword(data.password);
  }, []);

  const handleSubmit = useCallback(async (e) => {
    const data = await admin.createAdminPassword(password);
    console.log(data);
  }, [password]);

  useEffect(() => {
    fetchData().catch(console.error);
  }, [fetchData]);



  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>
          Edit <code>src/App.tsx</code> and save to reload.
        </p>
        <a
          className="App-link"
          href="https://reactjs.org"
          target="_blank"
          rel="noopener noreferrer"
        >
          Learn React
        </a>
        <div>
          <form onSubmit={handleSubmit}>
            <label>Admin Password</label>
            <input type="text" value={password} onChange={() => setPassword(password)} />
            <button type="submit">Save</button>
          </form>
        </div>
      </header>
    </div>
  );
}

export default App;
