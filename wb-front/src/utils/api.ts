import axios from 'axios';

const BASE = import.meta.env.VITE_API_URL;

export interface Tag {
  id: number;
  name: string;
}

export interface Community {
  id: number;
  name: string;
  description: string;
  imageUrl?: string;
  members: number;
  createdAt: string;
}

export const getCommunity = (id: number): Promise<Community> =>
  axios.get(`${BASE}/c/${id}`).then((r) => r.data);

export interface Post {
  id: number;
  communityId: number;
  title: string;
  body: string;
  communityName: string;
  imageUrl?: string;
  likes: number;
  liked: boolean;
  commentCount: number;
  tags: string[];
  links: string[];
}

export interface SearchResult {
  communities: Community[];
  posts: Post[];
}

export const getCommunityPosts = (communityId: number) =>
  axios.get(`${BASE}/post/c/${communityId}`).then((r) => r.data);

export const search = (q: string): Promise<SearchResult> =>
  axios.get(`${BASE}/search`, { params: { q } }).then((r) => r.data);

export interface Comment {
  id: number;
  postId: number;
  userId: number;
  userNick: string;
  body: string;
  createdAt: string;
}

export const likePost = (postId: number): Promise<{ likes: number }> =>
  axios.post(`${BASE}/post/${postId}/like`).then((r) => r.data);

export const unlikePost = (postId: number): Promise<{ likes: number }> =>
  axios.delete(`${BASE}/post/${postId}/like`).then((r) => r.data);

export const getComments = (postId: number): Promise<Comment[]> =>
  axios.get(`${BASE}/post/${postId}/comments`).then((r) => r.data ?? []);

export const createComment = (postId: number, body: string): Promise<void> =>
  axios.post(`${BASE}/post/${postId}/comments`, { body }).then(() => undefined);

export interface CreatePostDTO {
  communityId: number;
  title: string;
  body: string;
  tags: string[];
  links: string[];
}

export const createPost = (data: CreatePostDTO): Promise<void> =>
  axios.post(`${BASE}/post`, data).then(() => undefined);

export const getTags = (search?: string): Promise<Tag[]> =>
  axios.get(`${BASE}/tags`, { params: search ? { name: search } : {} }).then((r) => r.data ?? []);

export const createTag = (name: string): Promise<number> =>
  axios.post(`${BASE}/tags`, { name }).then((r) => r.data);

export const getJoinedCommunities = (): Promise<number[]> =>
  axios.get(`${BASE}/c/joined`).then((r) => r.data ?? []);

export const followCommunity = (communityId: number) =>
  axios.post(`${BASE}/c/${communityId}/follow`);

export const unfollowCommunity = (communityId: number) =>
  axios.post(`${BASE}/c/${communityId}/unfollow`);

export const followUser = (userId: number) =>
  axios.post(`${BASE}/users/${userId}/follow`);

export const unfollowUser = (userId: number) =>
  axios.post(`${BASE}/users/${userId}/unfollow`);
