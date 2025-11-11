import React from 'react';
import ReactDOM from 'react-dom/client';
import { BrowserRouter, Routes, Route } from 'react-router-dom';

import './index.css';
import Login from './pages/Login/Login';
import { ToastContainer } from 'react-toastify';
import Register from './pages/Login/Register';
import MainLayout from './components/Layout/Layout';
import Feed from './components/Feed/Feed';
import CommunityFeed from './components/Communities';



function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/login" element={<Login />} />
        <Route path="/register" element={<Register />} />
        <Route path="/" element={<MainLayout />}>
          <Route index element={<Feed />} />
          <Route path="home" element={<Feed />} />
          <Route path="communities" element={<CommunityFeed />} />
          <Route path="popular" element={<Feed />} />
          <Route path="new" element={<Feed />} />
        </Route>
      </Routes>
      <ToastContainer
        position="top-center"
        autoClose={3000}
        hideProgressBar={false}
        newestOnTop={false}
        closeOnClick
        pauseOnFocusLoss
        draggable
        pauseOnHover
        theme="colored" 
      />
    </BrowserRouter>
  );
}

ReactDOM.createRoot(document.getElementById('root') as HTMLElement).render(
  <React.StrictMode>
    <App />
  </React.StrictMode>
);

console.log(import.meta.env.VITE_API_URL);

fetch(`${import.meta.env.VITE_API_URL}/login`)
  .then(res => res.json())
  .then(data => console.log(data));