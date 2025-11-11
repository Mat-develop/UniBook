import React from 'react';
import { Card, Image, Button, Space, Typography } from 'antd';
import { LikeOutlined, CommentOutlined, ShareAltOutlined, HeatMapOutlined, BlockOutlined, BookOutlined } from '@ant-design/icons';
import styles from './communityContainer.module.scss';

interface PostProps {
  id: number;
  name: string;
  about: string;
  img?: string;
  members: number;
}

const CommunityContainer: React.FC<PostProps> = ({ 
  id,
  name, 
  about, 
  img, 
  members = 0,
}) => {
    
  return (
    <Card className={styles.communityContainer}>
      <div className={styles.header}>
          {img && (
          <div className={styles.imageContainer}>
            <Image src={img} alt={name} />
          </div>
        )}
        <Typography.Title level={3}>c/{name}</Typography.Title>
      </div>

      <div className={styles.content}>
        <Typography.Paragraph>{about}</Typography.Paragraph>
      </div>

      <div className={styles.footer}>
        <Space>
          <Button icon={<HeatMapOutlined />}>
            Join
          </Button>
          <Button icon={<BookOutlined />}>
            {members} members
          </Button>
          <Button icon={<BlockOutlined /> }>
            Share
          </Button>
        </Space>
      </div>
    </Card>
  );
};

export default CommunityContainer;