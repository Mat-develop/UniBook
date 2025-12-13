import axios from "axios";
import React, { useEffect, useState } from "react";
import PostContainer from "../PostContainer";

interface Post {
  id: string;
  title: string;
  body: string;
  communityName: string;
  imageUrl?: string;
  likes: number;
  comments: number;
}

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
          title={post.title}
          body={post.body}
          community={post.communityName}
          img={post.imageUrl}
          likes={post.likes}
          comments={post.comments}
        />
      ))}
    </div>
  );
};


export default Feed;