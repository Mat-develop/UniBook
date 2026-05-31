import axios from "axios";
import React, { useEffect, useState } from "react";
import { type Post } from "../../utils/api";
import PostContainer from "../PostContainer";

const Feed: React.FC = () => {
  const [posts, setPosts] = useState<Post[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const fetchPosts = async () => {
      try {
        const response = await axios.get(`${import.meta.env.VITE_API_URL}/post/c/1`);
        setPosts(response.data);
        setLoading(false);
      } catch (err) {
        setError('Failed to fetch posts');
        setLoading(false);
        console.error('Error fetching posts:', err);
      }
    };

    fetchPosts();
  }, []);

  if (loading) return <div>Loading...</div>;
  if (error) return <div>{error}</div>;

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
