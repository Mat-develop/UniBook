import React, { useState } from 'react';
import { Drawer, Layout, Menu } from 'antd';
import {
  HomeOutlined, FireOutlined, ClockCircleOutlined, MergeOutlined,
} from '@ant-design/icons';
import { useNavigate, useLocation, Outlet } from 'react-router-dom';
import styles from './layout.module.scss';
import Header from '../Header/Header';
import SideMenu from '../SideMenu/SideMenu';

const { Content } = Layout;

const menuItems = [
  { key: '/home',        icon: <HomeOutlined />,         label: 'Home' },
  { key: '/communities', icon: <MergeOutlined />,        label: 'Comunidades' },
  { key: '/popular',     icon: <FireOutlined />,         label: 'Popular' },
  { key: '/new',         icon: <ClockCircleOutlined />,  label: 'Novo' },
];

const MainLayout: React.FC = () => {
  const [collapsed, setCollapsed]     = useState(false);
  const [mobileOpen, setMobileOpen]   = useState(false);
  const navigate   = useNavigate();
  const location   = useLocation();

  const handleMobileNav = (key: string) => {
    navigate(key);
    setMobileOpen(false);
  };

  return (
    <Layout className={styles.layout}>
      <Header onMobileMenu={() => setMobileOpen(true)} />
      <Layout className={styles.mainLayout}>
        <SideMenu collapsed={collapsed} onCollapse={setCollapsed} />

        {/* Mobile navigation drawer */}
        <Drawer
          open={mobileOpen}
          onClose={() => setMobileOpen(false)}
          placement="left"
          width={240}
          styles={{
            body:   { padding: 0, background: '#001529' },
            header: { background: '#001529', borderBottom: '1px solid rgba(255,255,255,0.08)' },
          }}
          title={<span style={{ color: 'rgba(255,255,255,0.85)', fontWeight: 600, fontSize: 16 }}>UniBook</span>}
          closeIcon={<span style={{ color: 'rgba(255,255,255,0.5)', fontSize: 16 }}>✕</span>}
          className={styles.mobileDrawer}
        >
          <Menu
            mode="inline"
            theme="dark"
            selectedKeys={[location.pathname]}
            items={menuItems.map(item => ({
              ...item,
              onClick: () => handleMobileNav(item.key),
            }))}
            style={{ border: 'none', fontSize: 15 }}
          />
        </Drawer>

        <Content className={`${styles.content} ${collapsed ? styles.contentCollapsed : ''}`}>
          <div className={styles.contentWrapper}>
            <Outlet />
          </div>
        </Content>
      </Layout>
    </Layout>
  );
};

export default MainLayout;
