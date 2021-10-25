import React from 'react';
import logo from './logo.svg';
import './App.css';

function App() {
  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>
          Edit <code>src/App.tsx</code> and save to reload.
        </p>
        <a className="App-link" href="https://reactjs.org" target="_blank" rel="noopener noreferrer">
          Learn React
        </a>
        <div className="rounded-lg bg-green-300 font-bold p-4 m-2 text-green-900 border-white border-2 shadow-2xl shadow-inner cursor-pointer hover:bg-green-400">
          This will be a button
        </div>
      </header>
    </div>
  );
}

export default App;
