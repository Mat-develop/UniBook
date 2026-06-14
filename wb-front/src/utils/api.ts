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

export const getFeed = (): Promise<Post[]> =>
  axios.get(`${BASE}/feed`).then((r) => r.data ?? []);

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

export const updateUserImage = (userId: number, imageData: string): Promise<void> =>
  axios.put(`${BASE}/users/${userId}/image`, { imageData }).then(() => undefined);

export const followUser = (userId: number) =>
  axios.post(`${BASE}/users/${userId}/follow`);

export const unfollowUser = (userId: number) =>
  axios.post(`${BASE}/users/${userId}/unfollow`);

// ── Profile ──────────────────────────────────────────────────────────────────

export interface Education {
  id?: number;
  institution: string;
  degree: string;
  fieldOfStudy: string;
  startYear: number;
  endYear?: number | null;
  description?: string | null;
}

export interface Project {
  id?: number;
  title: string;
  description: string;
  url?: string | null;
  year?: number | null;
}

export interface Course {
  id?: number;
  title: string;
  institution: string;
  year?: number | null;
  url?: string | null;
}

export interface UserProfile {
  id: number;
  name: string;
  nick: string;
  imageUrl: string;
  education: Education[];
  projects: Project[];
  courses: Course[];
}

export const getProfile = (userId: number): Promise<UserProfile> =>
  axios.get(`${BASE}/users/${userId}/profile`).then((r) => r.data);

export const addEducation = (data: Education): Promise<{ id: number }> =>
  axios.post(`${BASE}/profile/education`, data).then((r) => r.data);
export const updateEducation = (id: number, data: Education): Promise<void> =>
  axios.put(`${BASE}/profile/education/${id}`, data).then(() => undefined);
export const deleteEducation = (id: number): Promise<void> =>
  axios.delete(`${BASE}/profile/education/${id}`).then(() => undefined);

export const addProject = (data: Project): Promise<{ id: number }> =>
  axios.post(`${BASE}/profile/projects`, data).then((r) => r.data);
export const updateProject = (id: number, data: Project): Promise<void> =>
  axios.put(`${BASE}/profile/projects/${id}`, data).then(() => undefined);
export const deleteProject = (id: number): Promise<void> =>
  axios.delete(`${BASE}/profile/projects/${id}`).then(() => undefined);

export const addCourse = (data: Course): Promise<{ id: number }> =>
  axios.post(`${BASE}/profile/courses`, data).then((r) => r.data);
export const updateCourse = (id: number, data: Course): Promise<void> =>
  axios.put(`${BASE}/profile/courses/${id}`, data).then(() => undefined);
export const deleteCourse = (id: number): Promise<void> =>
  axios.delete(`${BASE}/profile/courses/${id}`).then(() => undefined);
