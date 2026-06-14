import React, { useState } from 'react';
import { Layout } from 'antd';
import styles from './layout.module.scss';
import Header from '../Header/Header';
import SideMenu from '../SideMenu/SideMenu';
import { Outlet } from 'react-router-dom';

const { Content } = Layout;

const MainLayout: React.FC = () => {
  const [collapsed, setCollapsed] = useState(false);

  return (
    <Layout className={styles.layout}>
      <Header />
      <Layout className={styles.mainLayout}>
        <SideMenu collapsed={collapsed} onCollapse={setCollapsed} />
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
