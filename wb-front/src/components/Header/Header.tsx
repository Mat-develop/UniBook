import React, { useEffect, useState } from "react";
import styles from "./header.module.scss";
import { Input, Avatar, Dropdown, type MenuProps } from "antd";
import { MenuOutlined, KeyOutlined, LogoutOutlined, SkinOutlined, UserOutlined } from '@ant-design/icons';
import logo from "../../assets/logo.svg";
import { useNavigate } from "react-router-dom";
import { setAuthToken, getUserIdFromToken } from "../../utils/auth";
import { getProfile } from "../../utils/api";

interface HeaderProps {
  onMobileMenu?: () => void;
}

const Header: React.FC<HeaderProps> = ({ onMobileMenu }) => {
  const navigate = useNavigate();
  const [avatarSrc, setAvatarSrc] = useState<string | undefined>(undefined);

  useEffect(() => {
    const userId = getUserIdFromToken();
    if (userId) {
      getProfile(userId)
        .then((p) => { if (p.imageUrl) setAvatarSrc(p.imageUrl); })
        .catch(() => {});
    }
  }, []);

  const handleChangePassword = () => {
    navigate("/change-password");
  };

  const handleProfile = () => {
    const userId = getUserIdFromToken();
    if (userId) navigate(`/profile/${userId}`);
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

  const handleSearch = (value: string) => {
    const q = value.trim();
    if (q) navigate(`/search?q=${encodeURIComponent(q)}`);
  };

  return (
    <header className={styles.header}>
      <button className={styles.hamburger} onClick={onMobileMenu} aria-label="Abrir menu">
        <MenuOutlined />
      </button>
      <div className={styles.logoArea}>
        <img src={logo} alt="UniBook" className={styles.logo} />
      </div>
      <div className={styles.searchWrapper}>
        <Input.Search
          placeholder="Buscar comunidades ou posts…"
          className={styles.search}
          onSearch={handleSearch}
          allowClear
        />
      </div>
      <div className={styles.right}>
        <Dropdown menu={{ items }}>
          <Avatar
            size={38}
            src={avatarSrc}
            icon={<UserOutlined />}
            style={{ cursor: 'pointer', backgroundColor: '#1677ff' }}
          />
        </Dropdown>
      </div>
    </header>
  );
};

export default Header;
