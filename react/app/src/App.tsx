import React from 'react';
import './static/css//App.css';
import { Routes, Route } from 'react-router-dom';

import { Home } from './Home';
import { Room } from './Room';
import NotFound from './NotFound';

function App() {
  return (
    <>
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/rooms/:id" element={<Room />} />
        <Route path="*" element={<NotFound />} />
      </Routes>
    </>
  );
}

export default App;
