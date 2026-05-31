import React, { useEffect, useMemo, useState } from 'react';
import { useParams, useSearchParams } from 'react-router-dom';
import { Button, Spin, Tag, Typography } from 'antd';
import { PlusOutlined, TeamOutlined, UserAddOutlined, UserDeleteOutlined } from '@ant-design/icons';
import { toast } from 'react-toastify';
import { getCommunity, getCommunityPosts, getJoinedCommunities, followCommunity, unfollowCommunity, type Community, type Post } from '../../utils/api';
import PostContainer from '../../components/PostContainer';
import CreatePostModal from '../../components/CreatePostModal';
import styles from './communityPage.module.scss';

const CommunityPage: React.FC = () => {
  const { communityId } = useParams<{ communityId: string }>();
  const [searchParams, setSearchParams] = useSearchParams();
  const [modalOpen, setModalOpen] = useState(false);
  const activeTag = searchParams.get('tag') ?? '';

  const id = Number(communityId);

  const [community, setCommunity] = useState<Community | null>(null);
  const [posts, setPosts] = useState<Post[]>([]);
  const [isFollowing, setIsFollowing] = useState(false);
  const [memberCount, setMemberCount] = useState(0);
  const [followLoading, setFollowLoading] = useState(false);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    if (!id) return;

    const load = async () => {
      try {
        const [communityData, postsData, joinedIds] = await Promise.all([
          getCommunity(id),
          getCommunityPosts(id),
          getJoinedCommunities(),
        ]);
        setCommunity(communityData);
        setMemberCount(communityData.members);
        setPosts(postsData ?? []);
        setIsFollowing(joinedIds.includes(id));
      } catch {
        setError('Failed to load community');
      } finally {
        setLoading(false);
      }
    };

    load();
  }, [id]);

  const allTags = useMemo(() => {
    const set = new Set<string>();
    posts.forEach((p) => p.tags?.forEach((t) => set.add(t)));
    return Array.from(set).sort();
  }, [posts]);

  const filteredPosts = useMemo(
    () => (activeTag ? posts.filter((p) => p.tags?.includes(activeTag)) : posts),
    [posts, activeTag]
  );

  const handleTagClick = (tag: string) => {
    setSearchParams(tag === activeTag ? {} : { tag });
  };

  const refreshPosts = () =>
    getCommunityPosts(id).then((data) => setPosts(data ?? [])).catch(() => {});

  const handleFollowToggle = async () => {
    setFollowLoading(true);
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
      setFollowLoading(false);
    }
  };

  if (loading) return <div className={styles.center}><Spin size="large" /></div>;
  if (error || !community) return <div className={styles.center}>{error ?? 'Community not found'}</div>;

  return (
    <div className={styles.page}>
      <div className={styles.header}>
        <div className={styles.headerInfo}>
          <Typography.Title level={2}>c/{community.name}</Typography.Title>
          <Typography.Paragraph type="secondary">{community.description}</Typography.Paragraph>
          <div className={styles.meta}>
            <TeamOutlined />
            <span>{memberCount} members</span>
            <Button
              icon={isFollowing ? <UserDeleteOutlined /> : <UserAddOutlined />}
              type={isFollowing ? 'default' : 'primary'}
              loading={followLoading}
              onClick={handleFollowToggle}
              size="small"
            >
              {isFollowing ? 'Leave' : 'Join'}
            </Button>
            <Button
              icon={<PlusOutlined />}
              type="primary"
              size="small"
              onClick={() => setModalOpen(true)}
            >
              Create Post
            </Button>
          </div>
        </div>
      </div>

      <CreatePostModal
        communityId={id}
        open={modalOpen}
        onClose={() => setModalOpen(false)}
        onSuccess={() => { setModalOpen(false); refreshPosts(); }}
      />

      {allTags.length > 0 && (
        <div className={styles.tagBar}>
          <span className={styles.tagLabel}>Filter by tag:</span>
          {allTags.map((tag) => (
            <Tag
              key={tag}
              color={tag === activeTag ? 'blue' : 'default'}
              className={styles.tagChip}
              onClick={() => handleTagClick(tag)}
            >
              #{tag}
            </Tag>
          ))}
          {activeTag && (
            <Button size="small" type="link" onClick={() => setSearchParams({})}>
              Clear
            </Button>
          )}
        </div>
      )}

      <div className={styles.posts}>
        {filteredPosts.length === 0 ? (
          <Typography.Text type="secondary">No posts {activeTag ? `tagged #${activeTag}` : 'yet'}.</Typography.Text>
        ) : (
          filteredPosts.map((post) => (
            <PostContainer
              key={post.id}
              postId={post.id}
              communityId={post.communityId}
              title={post.title}
              body={post.body}
              community={post.communityName}
              img={post.imageUrl}
              likes={post.likes}
              initialLiked={post.liked}
              commentCount={post.commentCount}
              tags={post.tags}
              links={post.links}
              onTagClick={handleTagClick}
            />
          ))
        )}
      </div>
    </div>
  );
};

export default CommunityPage;
