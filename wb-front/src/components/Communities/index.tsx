import axios from "axios";
import React, { useEffect, useState } from "react";
import CommunityContainer from "../CommunityContainer";
import CreateCommunityModal from "../CreateCommunityModal";
import { getJoinedCommunities, type Community } from "../../utils/api";
import styles from './communities.module.scss';

const CommunityFeed: React.FC = () => {
  const [communities, setCommunities] = useState<Community[]>([]);
  const [joinedIds, setJoinedIds] = useState<Set<number>>(new Set());
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const [communitiesRes, joined] = await Promise.all([
          axios.get<Community[]>(`${import.meta.env.VITE_API_URL}/c/all`),
          getJoinedCommunities(),
        ]);
        setCommunities(communitiesRes.data);
        setJoinedIds(new Set(joined));
      } catch (err) {
        setError('Failed to fetch communities');
        console.error('Error fetching communities:', err);
      } finally {
        setLoading(false);
      }
    };

    fetchData();
  }, []);

  if (loading) return <div>Loading...</div>;
  if (error) return <div>{error}</div>;

  const reload = async () => {
    const [commRes, joined] = await Promise.all([
      axios.get<Community[]>(`${import.meta.env.VITE_API_URL}/c/all`),
      getJoinedCommunities(),
    ]);
    setCommunities(commRes.data);
    setJoinedIds(new Set(joined));
  };

  return (
    <div className={styles.communitiesPage}>
      <div style={{ display: 'flex', justifyContent: 'flex-end', marginBottom: 16 }}>
        <CreateCommunityModal onCreated={reload} />
      </div>
      {communities.map((community) => (
        <CommunityContainer
          key={community.id}
          id={community.id}
          name={community.name}
          about={community.description}
          img={community.imageUrl}
          members={community.members ?? 0}
          initialFollowing={joinedIds.has(community.id)}
        />
      ))}
    </div>
  );
};

export default CommunityFeed;
