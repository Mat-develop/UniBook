import axios from 'axios';

const BASE = import.meta.env.VITE_API_URL;

export const followCommunity = (communityId: number) =>
  axios.post(`${BASE}/c/${communityId}/follow`);

export const unfollowCommunity = (communityId: number) =>
  axios.post(`${BASE}/c/${communityId}/unfollow`);

export const followUser = (userId: number) =>
  axios.post(`${BASE}/users/${userId}/follow`);

export const unfollowUser = (userId: number) =>
  axios.post(`${BASE}/users/${userId}/unfollow`);
