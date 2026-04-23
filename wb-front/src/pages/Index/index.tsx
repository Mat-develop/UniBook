import { useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import { getStoredToken, isTokenExpired } from '../../utils/auth';


const IndexPage = () => {
  const navigate = useNavigate();

  useEffect(() => {
    const token = getStoredToken();
    
    if (token && !isTokenExpired(token)) {
      navigate('/home', { replace: true });
    } else {
      navigate('/login', { replace: true });
    }
  }, [navigate]);

  return null;
};

export default IndexPage;
