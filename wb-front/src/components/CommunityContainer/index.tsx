import React, { useState } from 'react';
import { Card, Image, Button, Space, Typography } from 'antd';
import { BookOutlined, BlockOutlined, TeamOutlined, UserAddOutlined, UserDeleteOutlined } from '@ant-design/icons';
import { toast } from 'react-toastify';
import { followCommunity, unfollowCommunity } from '../../utils/api';
import styles from './communityContainer.module.scss';

interface CommunityProps {
  id: number;
  name: string;
  about: string;
  img?: string;
  members: number;
}

const CommunityContainer: React.FC<CommunityProps> = ({ id, name, about, img, members = 0 }) => {
  const [isFollowing, setIsFollowing] = useState(false);
  const [memberCount, setMemberCount] = useState(members);
  const [loading, setLoading] = useState(false);

  const handleFollowToggle = async () => {
    setLoading(true);
    try {
      if (isFollowing) {
        await unfollowCommunity(id);
        setIsFollowing(false);
        setMemberCount((c) => c - 1);
      } else {
        await followCommunity(id);
        setIsFollowing(true);
        setMemberCount((c) => c + 1);
      }
    } catch {
      toast.error(`Failed to ${isFollowing ? 'leave' : 'join'} community`);
    } finally {
      setLoading(false);
    }
  };

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
          <Button
            icon={isFollowing ? <UserDeleteOutlined /> : <UserAddOutlined />}
            type={isFollowing ? 'default' : 'primary'}
            loading={loading}
            onClick={handleFollowToggle}
          >
            {isFollowing ? 'Leave' : 'Join'}
          </Button>
          <Button icon={<TeamOutlined />}>
            {memberCount} members
          </Button>
          <Button icon={<BlockOutlined />}>
            Share
          </Button>
        </Space>
      </div>
    </Card>
  );
};

export default CommunityContainer;
