import React, { useState } from 'react';
import { Link } from 'react-router-dom';
import { Card, Image, Button, Typography, Tag, Input } from 'antd';
import { LikeOutlined, LikeFilled, CommentOutlined, SendOutlined, LinkOutlined } from '@ant-design/icons';
import ReactMarkdown from 'react-markdown';
import remarkGfm from 'remark-gfm';
import { toast } from 'react-toastify';
import { likePost, unlikePost, getComments, createComment, type Comment } from '../../utils/api';
import styles from './postContainer.module.scss';

interface PostProps {
  postId: number;
  communityId?: number;
  title: string;
  body: string;
  img?: string;
  community: string;
  likes?: number;
  initialLiked?: boolean;
  commentCount?: number;
  tags?: string[];
  links?: string[];
  onTagClick?: (tag: string) => void;
}

const PostContainer: React.FC<PostProps> = ({
  postId,
  communityId,
  title,
  body,
  img,
  community,
  likes: initialLikes = 0,
  initialLiked = false,
  commentCount = 0,
  tags = [],
  links = [],
  onTagClick,
}) => {
  const [liked, setLiked] = useState(initialLiked);
  const [likeCount, setLikeCount] = useState(initialLikes);
  const [likeLoading, setLikeLoading] = useState(false);

  const [commentsOpen, setCommentsOpen] = useState(false);
  const [comments, setComments] = useState<Comment[]>([]);
  const [commentsLoaded, setCommentsLoaded] = useState(false);
  const [loadingComments, setLoadingComments] = useState(false);
  const [commentInput, setCommentInput] = useState('');
  const [submittingComment, setSubmittingComment] = useState(false);

  const handleLike = async () => {
    setLikeLoading(true);
    try {
      if (liked) {
        const { likes } = await unlikePost(postId);
        setLikeCount(likes);
        setLiked(false);
      } else {
        const { likes } = await likePost(postId);
        setLikeCount(likes);
        setLiked(true);
      }
    } catch (err: unknown) {
      const msg: string = (err as { response?: { data?: { erro?: string } } })?.response?.data?.erro ?? '';
      if (msg.includes('already liked')) setLiked(true);
      else if (msg.includes('not liked')) setLiked(false);
      else toast.error('Failed to update like');
    } finally {
      setLikeLoading(false);
    }
  };

  const loadComments = async () => {
    setLoadingComments(true);
    try {
      const data = await getComments(postId);
      setComments(data);
      setCommentsLoaded(true);
    } catch {
      toast.error('Failed to load comments');
    } finally {
      setLoadingComments(false);
    }
  };

  const handleToggleComments = () => {
    const next = !commentsOpen;
    setCommentsOpen(next);
    if (next && !commentsLoaded) loadComments();
  };

  const handleSubmitComment = async () => {
    if (!commentInput.trim()) return;
    setSubmittingComment(true);
    try {
      await createComment(postId, commentInput.trim());
      setCommentInput('');
      setComments(await getComments(postId));
    } catch {
      toast.error('Failed to post comment');
    } finally {
      setSubmittingComment(false);
    }
  };

  return (
    <Card className={styles.post}>
      <div className={styles.header}>
        {communityId ? (
          <Link to={`/c/${communityId}`} className={styles.communityLink}>c/{community}</Link>
        ) : (
          <Typography.Text type="secondary">c/{community}</Typography.Text>
        )}
        <Typography.Title level={3}>{title}</Typography.Title>
      </div>

      <div className={styles.content}>
        <div className={styles.markdownBody}>
          <ReactMarkdown remarkPlugins={[remarkGfm]}>{body}</ReactMarkdown>
        </div>
        {img && (
          <div className={styles.imageContainer}>
            <Image src={img} alt={title} />
          </div>
        )}
      </div>

      {links.length > 0 && (
        <div className={styles.links}>
          {links.map((url) => {
            let host = url;
            try { host = new URL(url).hostname; } catch {}
            return (
              <a key={url} href={url} target="_blank" rel="noopener noreferrer">
                <Tag icon={<LinkOutlined />} color="geekblue">{host}</Tag>
              </a>
            );
          })}
        </div>
      )}

      {tags.length > 0 && (
        <div className={styles.tags}>
          {tags.map((tag) => (
            <Tag
              key={tag}
              color="blue"
              className={onTagClick ? styles.clickableTag : ''}
              onClick={() => onTagClick?.(tag)}
            >
              #{tag}
            </Tag>
          ))}
        </div>
      )}

      <div className={styles.footer}>
        <div className={styles.likeArea}>
          <Button
            icon={liked ? <LikeFilled /> : <LikeOutlined />}
            type={liked ? 'primary' : 'default'}
            loading={likeLoading}
            onClick={handleLike}
            size="small"
          >
            {liked ? 'Unlike' : 'Like'}
          </Button>
          {likeCount > 0 && (
            <Typography.Text type="secondary" className={styles.likeCount}>
              {likeCount} {likeCount === 1 ? 'like' : 'likes'}
            </Typography.Text>
          )}
        </div>

        <Button
          icon={<CommentOutlined />}
          type={commentsOpen ? 'primary' : 'default'}
          onClick={handleToggleComments}
          size="small"
        >
          {commentsLoaded ? comments.length : commentCount} Comments
        </Button>
      </div>

      {commentsOpen && (
        <div className={styles.commentSection}>
          {loadingComments ? (
            <Typography.Text type="secondary">Loading…</Typography.Text>
          ) : (
            <>
              <div className={styles.commentList}>
                {comments.length === 0 ? (
                  <Typography.Text type="secondary">No comments yet. Be the first!</Typography.Text>
                ) : (
                  comments.map((c) => (
                    <div key={c.id} className={styles.commentItem}>
                      <Typography.Text strong className={styles.commentNick}>{c.userNick}</Typography.Text>
                      <Typography.Text>{c.body}</Typography.Text>
                    </div>
                  ))
                )}
              </div>
              <div className={styles.commentForm}>
                <Input
                  value={commentInput}
                  onChange={(e) => setCommentInput(e.target.value)}
                  onPressEnter={handleSubmitComment}
                  placeholder="Write a comment…"
                />
                <Button
                  icon={<SendOutlined />}
                  type="primary"
                  loading={submittingComment}
                  disabled={!commentInput.trim()}
                  onClick={handleSubmitComment}
                >
                  Post
                </Button>
              </div>
            </>
          )}
        </div>
      )}
    </Card>
  );
};

export default PostContainer;
