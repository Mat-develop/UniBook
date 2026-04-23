import React, { useState } from "react";
import { Input, Button, Checkbox } from "antd";
import styles from "./login.module.scss";
import logo from "../../assets/logo.svg"
import { LockOutlined, UserOutlined } from '@ant-design/icons';

import { toast } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';
import axios from "axios";
import { useNavigate } from "react-router-dom";
import { setAuthToken } from "../../utils/auth";

const Login: React.FC = () => {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [remember, setRemember] = useState(false);
  const [loading, setLoading] = useState(false); 
  const navigate = useNavigate();

  const handleRegister = () =>{
     navigate("/register");
    }
    
  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
   
    if (!username || !password) {
      toast.error("Please fill in all fields");
      return;
    }
   setLoading(true);

    try {
      const response = await axios.post(`${import.meta.env.VITE_API_URL}/login`,{
        email: username,
        password: password
      },  {
        headers: {
          "Content-Type": "application/json",
        },
      });

      const token = response.data.token;

      setAuthToken(token, remember);
      toast.success("Login successful!");

      navigate("/home");
    } catch (error: any) {
    toast.error(error.response?.data?.message || "Login failed");
    } finally {
      setLoading(false);
    } 
  };
  

 

  return (
   <div className={styles.mainContainer}>
    <div className={styles.aboutContainer}>
      <img src={logo} alt="Company Logo" style={{ width: "200px", height: "auto" }} />
      <p> Conecte-se, compartilhe ideias, faça networks e curta os melhores momentos, tudo na mesma rede social!</p>
      
    </div>
    <div className={styles.loginContainer}>
      <form className={styles.form} onSubmit={handleSubmit}>
        <h2>Login</h2>
        <Input
          placeholder="Username"
          prefix={<UserOutlined />}
          value={username}
          onChange={(e) => setUsername(e.target.value)}
        />
        <Input.Password
          placeholder="Password"
            prefix={<LockOutlined />}
          value={password}
          onChange={(e) => setPassword(e.target.value)}
        />
        <div className={styles.remember}>
          <Checkbox
            checked={remember}
            onChange={(e) => setRemember(e.target.checked)}
          >
            Remember me
          </Checkbox>
        </div>
        <Button type="primary" htmlType="submit" className={styles.submit} loading={loading}>
          Submit
        </Button>
      </form>

      <div className={styles.register}>
         <Button type="primary" className={styles.goToReg} loading={loading} onClick={handleRegister}>
          Cadastrar
        </Button>
      </div>
    </div>
  </div>
  );
};

export default Login;