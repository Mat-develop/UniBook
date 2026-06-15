import { Button, Checkbox, Input } from "antd";
import axios from "axios";
import logo from "../../assets/logo.svg"
import { useState } from "react";
import { toast } from "react-toastify";
import styles from "./register.module.scss";
import { LockOutlined, UserOutlined } from '@ant-design/icons';
import { useNavigate } from "react-router-dom";
const Register: React.FC = () => {
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");
    const [email, setEmail] = useState("");
    const [name, setName] = useState("");
    const [loading, setLoading] = useState(false); 
    const navigate = useNavigate();
    const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
   
    if (!username || !password) {
      toast.error("Please fill in all fields");
      return;
    }
   setLoading(true);

    try {
      await axios.post(`${import.meta.env.VITE_API_URL}/users`,{
        name: name,
        nick: username,
        email: email,
        password: password,
      },  {
    headers: {
      "Content-Type": "application/json",
    },});

      toast.success("Cadastrado com Sucesso!");
      navigate("/login");
    } catch (error: any) {
    toast.error(error.response?.data?.erro || "Erro ao Cadastrar! :(");
    } finally {
      setLoading(false);
    } 
  };
  

  return(
    <div className={styles.mainContainer}>
    <div className={styles.registerContainer}>
        <div>
            <img src={logo} alt="Company Logo" style={{ width: "200px", height: "auto" }} />
        </div>
        <form className={styles.form} onSubmit={handleSubmit}>
            <h2>Juntar-se</h2>
            <Input
            placeholder="Nome"
            prefix={<UserOutlined />}
            value={name}
            onChange={(e) => setName(e.target.value)}
            />
            <Input
            placeholder="Username / Nome de Guerra"
            prefix={<UserOutlined />}
            value={username}
            onChange={(e) => setUsername(e.target.value)}
            />
            <Input
            placeholder="Email"
            prefix={<UserOutlined />}
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            />
            <Input.Password
            placeholder="Senha"
                prefix={<LockOutlined />}
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            />
            <Checkbox>
                <p className={styles.confirm}>Ao confirmar estou ciente que é proibido gente chata! </p>
            </Checkbox>
            <Button type="primary" htmlType="submit" className={styles.submit} loading={loading}>
            Submit
            </Button>
        </form>
     </div>
    </div>
  )
}

export default Register;