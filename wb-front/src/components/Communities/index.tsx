import axios from "axios";
import React, { useEffect, useState } from "react";

interface Community {
  id: number;
  name: string;
  description: string;
  img?: string;
}

const Feed: React.FC = () => {
  const [community, setCommunity] = useState<Community[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const fetchCommunities = async () => {
      try {
        const response = await axios.get(`${import.meta.env.VITE_API_URL}/post/c/1`);
        setCommunity(response.data);
        setLoading(false);
      } catch (err) {
        setError('Failed to fetch communities');
        setLoading(false);
        console.error('Error fetching communities:', err);
      }
    };

    fetchCommunities();
  }, []);

  if (loading) return <div>Loading...</div>;
  if (error) return <div>{error}</div>;
  
  // TODO: CREATE COMMUNITY CONTAINER
 return (
    <div>
      {community.map((comu) => (
        <CommunityContainer
          key={comu.id}
          title={comu.name}
          body={comu.description}
          img={comu.img}
        />
      ))}
    </div>
  );
};


export default Feed;