import React from "react";
import styles from "./header.module.scss";
import { Input, Avatar, Dropdown, type MenuProps } from "antd";
import logo from "../../assets/logo.svg";
import { KeyOutlined, LogoutOutlined, SkinOutlined, UserOutlined } from '@ant-design/icons';
import { useNavigate } from "react-router-dom";
import { setAuthToken } from "../../utils/auth";

const Header: React.FC = () => {
  const navigate = useNavigate(); 

  const handleChangePassword = () => {
    navigate("/change-password"); 
  };

  const handleProfile = () => {
    navigate("/profile");
  };

  const handleLogout = () => {
    setAuthToken(null);
    navigate("/login");
  };

  const items: MenuProps['items'] = [
    {
      key: '1',
      label: 'Perfil',
      icon: <SkinOutlined />,
      onClick: handleProfile
    },
    {
      key: '2',
      label: 'Trocar Senha',
      icon: <KeyOutlined />,
      onClick: handleChangePassword
    },
    {
      key: '3',
      label: 'Deslogar',
      icon: <LogoutOutlined />,
      onClick: handleLogout
    }
  ];

  return (
    <header className={styles.header}>
      <img src={logo} alt="Logo" className={styles.logo} />
      <Input.Search placeholder="Search..." className={styles.search} />
      <Dropdown menu={{items}} >
        <Avatar size={25} icon={<UserOutlined />} />
      </Dropdown>
    </header>
  );
};

export default Header;