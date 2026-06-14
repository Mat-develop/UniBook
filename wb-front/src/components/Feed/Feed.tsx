import React, { useEffect, useState } from "react";
import { Typography, Spin } from "antd";
import { getFeed, type Post } from "../../utils/api";
import PostContainer from "../PostContainer";
import styles from "./feed.module.scss";

const Feed: React.FC = () => {
  const [posts, setPosts] = useState<Post[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(false);

  useEffect(() => {
    getFeed()
      .then(setPosts)
      .catch(() => setError(true))
      .finally(() => setLoading(false));
  }, []);

  if (loading) return <div className={styles.center}><Spin size="large" /></div>;
  if (error) return <div className={styles.center}><Typography.Text type="danger">Erro ao carregar o feed.</Typography.Text></div>;
  if (posts.length === 0) return (
    <div className={styles.center}>
      <Typography.Text type="secondary">Nenhum post ainda. Siga comunidades para ver posts aqui.</Typography.Text>
    </div>
  );

  return (
    <div>
      {posts.map((post) => (
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
      ))}
    </div>
  );
};

export default Feed;
