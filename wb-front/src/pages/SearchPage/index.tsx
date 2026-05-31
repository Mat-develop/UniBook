import React, { useEffect, useState } from 'react';
import { Link, useSearchParams } from 'react-router-dom';
import { Spin, Typography } from 'antd';
import { TeamOutlined } from '@ant-design/icons';
import { search, type SearchResult } from '../../utils/api';
import PostContainer from '../../components/PostContainer';
import styles from './searchPage.module.scss';

const SearchPage: React.FC = () => {
  const [searchParams] = useSearchParams();
  const q = searchParams.get('q') ?? '';

  const [results, setResults] = useState<SearchResult | null>(null);
  const [loading, setLoading] = useState(false);

  useEffect(() => {
    if (!q.trim()) { setResults(null); return; }
    setLoading(true);
    search(q)
      .then(setResults)
      .catch(() => {})
      .finally(() => setLoading(false));
  }, [q]);

  if (!q.trim()) {
    return <div className={styles.empty}><Typography.Text type="secondary">Type something to search.</Typography.Text></div>;
  }

  if (loading) {
    return <div className={styles.empty}><Spin size="large" /></div>;
  }

  const communityCount = results?.communities.length ?? 0;
  const postCount = results?.posts.length ?? 0;

  return (
    <div className={styles.page}>
      <Typography.Title level={3} className={styles.heading}>
        Results for <em>"{q}"</em>
      </Typography.Title>

      <section className={styles.section}>
        <Typography.Title level={4}>
          Communities <span className={styles.count}>{communityCount}</span>
        </Typography.Title>
        {communityCount === 0 ? (
          <Typography.Text type="secondary">No communities found.</Typography.Text>
        ) : (
          <div className={styles.communityList}>
            {results!.communities.map((c) => (
              <Link key={c.id} to={`/c/${c.id}`} className={styles.communityCard}>
                <div className={styles.communityInfo}>
                  <Typography.Text strong>c/{c.name}</Typography.Text>
                  {c.description && (
                    <Typography.Text type="secondary" className={styles.communityDesc}>
                      {c.description}
                    </Typography.Text>
                  )}
                </div>
                <div className={styles.communityMeta}>
                  <TeamOutlined />
                  <Typography.Text type="secondary">{c.members} members</Typography.Text>
                </div>
              </Link>
            ))}
          </div>
        )}
      </section>

      <section className={styles.section}>
        <Typography.Title level={4}>
          Posts <span className={styles.count}>{postCount}</span>
        </Typography.Title>
        {postCount === 0 ? (
          <Typography.Text type="secondary">No posts found.</Typography.Text>
        ) : (
          results!.posts.map((post) => (
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
            />
          ))
        )}
      </section>
    </div>
  );
};

export default SearchPage;
