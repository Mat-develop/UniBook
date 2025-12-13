import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import MainLayout from './components/Layout/Layout';
import Feed from './components/Feed/Feed';
import Login from './pages/Login/Login';
import Register from './pages/Login/Register';

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/login" element={<Login />} />
        <Route path="/register" element={<Register />} />
        <Route path="/home" element={<MainLayout />}>
          <Route index element={<Feed />} />
          <Route path="popular" element={<Feed />} />
          <Route path="new" element={<Feed />} />
          <Route path="topics" element={<Feed />} />
        </Route>
      </Routes>
    </Router>
  );
}

export default App;