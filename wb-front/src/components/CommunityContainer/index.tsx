import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { Avatar, Card, Button, Typography } from 'antd';
import { TeamOutlined, UserAddOutlined, UserDeleteOutlined } from '@ant-design/icons';
import { toast } from 'react-toastify';
import { followCommunity, unfollowCommunity } from '../../utils/api';
import styles from './communityContainer.module.scss';

interface CommunityProps {
  id: number;
  name: string;
  about: string;
  img?: string;
  members: number;
  initialFollowing?: boolean;
}

const CommunityContainer: React.FC<CommunityProps> = ({ id, name, about, img, members = 0, initialFollowing = false }) => {
  const navigate = useNavigate();
  const [isFollowing, setIsFollowing] = useState(initialFollowing);
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
      toast.error(`Erro ao ${isFollowing ? 'sair da' : 'entrar na'} comunidade`);
    } finally {
      setLoading(false);
    }
  };

  return (
    <Card className={styles.communityContainer}>
      <div className={styles.header}>
        <Avatar
          size={52}
          src={img || undefined}
          className={styles.avatar}
          onClick={() => navigate(`/c/${id}`)}
        >
          {!img && name.charAt(0).toUpperCase()}
        </Avatar>
        <Typography.Title
          level={4}
          className={styles.communityName}
          onClick={() => navigate(`/c/${id}`)}
        >
          c/{name}
        </Typography.Title>
      </div>

      <div className={styles.content}>
        <Typography.Paragraph>{about}</Typography.Paragraph>
      </div>

      <div className={styles.footer}>
        <Button
          icon={isFollowing ? <UserDeleteOutlined /> : <UserAddOutlined />}
          type={isFollowing ? 'default' : 'primary'}
          loading={loading}
          onClick={handleFollowToggle}
        >
          {isFollowing ? 'Sair' : 'Participar'}
        </Button>
        <Button icon={<TeamOutlined />} disabled>
          {memberCount} membros
        </Button>
      </div>
    </Card>
  );
};

export default CommunityContainer;
