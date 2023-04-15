import React from 'react';
import './App.css';
import { Routes, Route } from 'react-router-dom';

import Home from './Home';
import About from './About';
import NotFound from './NotFound';

function App() {
  return (
    <div className="App">
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/about" element={<About />} />
        <Route path="*" element={<NotFound />} />
      </Routes>
    </div>
  );
}

export default App;
