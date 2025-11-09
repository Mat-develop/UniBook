import React from 'react';
import { Card, Image, Button, Space, Typography } from 'antd';
import { LikeOutlined, CommentOutlined, ShareAltOutlined } from '@ant-design/icons';
import styles from './postContainer.module.scss';

interface PostProps {
  title: string;
  body: string;
  img?: string;
  community: string;
  likes?: number;
  comments?: number;
}

const PostContainer: React.FC<PostProps> = ({ 
  title, 
  body, 
  img, 
  community,
  likes = 0,
  comments = 0 
}) => {
  return (
    <Card className={styles.post}>
      <div className={styles.header}>
        <Typography.Text type="secondary">c/{community}</Typography.Text>
        <Typography.Title level={3}>{title}</Typography.Title>
      </div>

      <div className={styles.content}>
        <Typography.Paragraph>{body}</Typography.Paragraph>
        {img && (
          <div className={styles.imageContainer}>
            <Image src={img} alt={title} />
          </div>
        )}
      </div>

      <div className={styles.footer}>
        <Space>
          <Button icon={<LikeOutlined />}>
            {likes} Likes
          </Button>
          <Button icon={<CommentOutlined />}>
            {comments} Comments
          </Button>
          <Button icon={<ShareAltOutlined />}>
            Share
          </Button>
        </Space>
      </div>
    </Card>
  );
};

export default PostContainer;