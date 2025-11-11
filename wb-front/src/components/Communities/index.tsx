import axios from "axios";
import React, { useEffect, useState } from "react";
import CommunityContainer from "../CommunityContainer";
import styles from './communities.module.scss';
interface Community {
  id: number;
  name: string;
  description: string;
  img?: string;
}

const CommunityFeed: React.FC = () => {
  const [community, setCommunity] = useState<Community[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const fetchCommunities = async () => {
      try {
        const response = await axios.get(`${import.meta.env.VITE_API_URL}/c/all`);
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
  
  // TODO: CREATE COMMUNITY MEMBERS COUNT
 return (
    <div className={styles.communitiesPage}>
      {community.map((comu) => (
        <CommunityContainer
          id={comu.id}
          name={comu.name}
          about={comu.description}
          img={comu.img}
          members={0}
        />
      ))}
    </div>
  );
};


export default CommunityFeed;