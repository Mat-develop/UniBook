package repository

import (
	"database/sql"
	"encoding/json"
	"errors"
	"strings"
	"time"
	"v1/post/model"
)

const (
	CreateQuery = "INSERT INTO posts (title, body, community_id, user_id, links) VALUES (?, ?, ?, ?, ?)"

	insertTagStmt    = "INSERT IGNORE INTO tags (name, created_by) VALUES (?, ?)"
	selectTagIDStmt  = "SELECT id FROM tags WHERE name = ?"
	insertPostTagStmt = "INSERT IGNORE INTO post_tags (post_id, tag_id) VALUES (?, ?)"

	baseCols = `
		p.id,
		p.user_id,
		p.community_id,
		c.name AS communityName,
		p.title,
		p.body,
		IFNULL(p.image_url, ''),
		IFNULL(p.likes, 0),
		p.created_at,
		u.nick,
		IFNULL(GROUP_CONCAT(t.name ORDER BY t.name SEPARATOR ','), '') AS tags,
		IF(COUNT(pl.post_id) > 0, 1, 0) AS liked,
		(SELECT COUNT(*) FROM post_comments pc WHERE pc.post_id = p.id) AS comment_count,
		IFNULL(p.links, '[]') AS links`

	baseJoins = `
		FROM posts p
		INNER JOIN users u ON u.id = p.user_id
		INNER JOIN community c ON c.id = p.community_id
		LEFT JOIN post_tags pt ON pt.post_id = p.id
		LEFT JOIN tags t ON t.id = pt.tag_id
		LEFT JOIN post_likes pl ON pl.post_id = p.id AND pl.user_id = ?`

	baseGroup = `
		GROUP BY p.id, p.user_id, p.community_id, c.name, p.title, p.body, p.image_url, p.likes, p.created_at, u.nick`

	feedJoins = `
		FROM posts p
		INNER JOIN users u ON u.id = p.user_id
		INNER JOIN community c ON c.id = p.community_id
		INNER JOIN community_followers cf ON cf.community_id = p.community_id AND cf.user_id = ?
		LEFT JOIN post_tags pt ON pt.post_id = p.id
		LEFT JOIN tags t ON t.id = pt.tag_id
		LEFT JOIN post_likes pl ON pl.post_id = p.id AND pl.user_id = ?`

	FindByCommunityIdQuery = "SELECT" + baseCols + baseJoins + " WHERE p.community_id = ?" + baseGroup
	FindByUserIdQuery      = "SELECT" + baseCols + baseJoins + " WHERE p.user_id = ?" + baseGroup
	SearchByTitleQuery     = "SELECT" + baseCols + baseJoins + " WHERE p.title LIKE ?" + baseGroup
	FindFeedQuery          = "SELECT" + baseCols + feedJoins + baseGroup +
		" ORDER BY ((p.likes + 1) / POW(TIMESTAMPDIFF(HOUR, p.created_at, NOW()) + 2, 1.8)) DESC LIMIT 50"

	UpdateQuery = "UPDATE posts SET title = ?, body = ? WHERE id = ? AND user_id = ?"
	DeleteQuery = "DELETE FROM posts WHERE id = ? AND user_id = ?"
)

type PostRepository interface {
	FindCommunityPosts(viewerID, communityId uint64) ([]model.Post, error)
	FindUserPosts(viewerID, userId uint64) ([]model.Post, error)
	FindFeed(viewerID uint64) ([]model.Post, error)
	FindPostByName() ([]model.Post, error)
	SearchByTitle(viewerID uint64, q string) ([]model.Post, error)
	Create(userId uint64, postBody model.PostDTO) error
	Update(postID uint64, userID uint64, postBody model.PostDTO) error
	Delete(postID uint64, userID uint64) error
	LikePost(postID, userID uint64) (int32, error)
	UnlikePost(postID, userID uint64) (int32, error)
}

type postRepository struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) PostRepository {
	return &postRepository{db: db}
}

func (p *postRepository) scanPosts(rows *sql.Rows) ([]model.Post, error) {
	posts := []model.Post{}
	for rows.Next() {
		var (
			id            int64
			userId        int64
			community     sql.NullInt64
			communityName string
			title         string
			body          string
			imageURL      string
			likes         int64
			createdAt     time.Time
			userNick      string
			tagsRaw       string
			likedInt      int
			commentCount  int64
			linksJSON     string
		)
		if err := rows.Scan(&id, &userId, &community, &communityName, &title, &body, &imageURL, &likes, &createdAt, &userNick, &tagsRaw, &likedInt, &commentCount, &linksJSON); err != nil {
			return nil, err
		}

		tags := []string{}
		if tagsRaw != "" {
			tags = strings.Split(tagsRaw, ",")
		}

		links := []string{}
		if linksJSON != "" && linksJSON != "[]" {
			json.Unmarshal([]byte(linksJSON), &links)
		}

		post := model.Post{
			ID:            uint64(id),
			CommunityName: communityName,
			UserId:        uint64(userId),
			UserNick:      userNick,
			Title:         title,
			Body:          body,
			ImageUrl:      imageURL,
			Likes:         int32(likes),
			Liked:         likedInt != 0,
			CommentCount:  int32(commentCount),
			Tags:          tags,
			Links:         links,
			CreatedAt:     createdAt,
		}
		if community.Valid {
			post.CommunityId = int32(community.Int64)
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func (p *postRepository) FindCommunityPosts(viewerID, communityId uint64) ([]model.Post, error) {
	rows, err := p.db.Query(FindByCommunityIdQuery, viewerID, communityId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return p.scanPosts(rows)
}

func (p *postRepository) FindUserPosts(viewerID, userId uint64) ([]model.Post, error) {
	rows, err := p.db.Query(FindByUserIdQuery, viewerID, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return p.scanPosts(rows)
}

func (p *postRepository) FindFeed(viewerID uint64) ([]model.Post, error) {
	rows, err := p.db.Query(FindFeedQuery, viewerID, viewerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return p.scanPosts(rows)
}

func (p *postRepository) FindPostByName() ([]model.Post, error) {
	return []model.Post{}, errors.New("not implemented")
}

func (p *postRepository) SearchByTitle(viewerID uint64, q string) ([]model.Post, error) {
	rows, err := p.db.Query(SearchByTitleQuery, viewerID, "%"+q+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return p.scanPosts(rows)
}

func (p *postRepository) Create(userId uint64, postBody model.PostDTO) error {
	tx, err := p.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	linksJSON, _ := json.Marshal(postBody.Links)
	res, err := tx.Exec(CreateQuery, postBody.Title, postBody.Body, postBody.CommunityId, userId, string(linksJSON))
	if err != nil {
		return err
	}

	postID, err := res.LastInsertId()
	if err != nil {
		return err
	}

	for _, tagName := range postBody.Tags {
		tagName = strings.ToLower(strings.TrimSpace(tagName))
		if tagName == "" {
			continue
		}
		if _, err = tx.Exec(insertTagStmt, tagName, userId); err != nil {
			return err
		}
		var tagID uint64
		if err = tx.QueryRow(selectTagIDStmt, tagName).Scan(&tagID); err != nil {
			return err
		}
		if _, err = tx.Exec(insertPostTagStmt, postID, tagID); err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (p *postRepository) Update(postID uint64, userID uint64, postBody model.PostDTO) error {
	stmt, err := p.db.Prepare(UpdateQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(postBody.Title, postBody.Body, postID, userID)
	if err != nil {
		return err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if affected == 0 {
		return errors.New("no post updated (not found or not owner)")
	}
	return nil
}

func (p *postRepository) Delete(postID uint64, userID uint64) error {
	stmt, err := p.db.Prepare(DeleteQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(postID, userID)
	if err != nil {
		return err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if affected == 0 {
		return errors.New("no post deleted (not found or not owner)")
	}
	return nil
}

func (p *postRepository) LikePost(postID, userID uint64) (int32, error) {
	tx, err := p.db.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	res, err := tx.Exec("INSERT IGNORE INTO post_likes (post_id, user_id) VALUES (?, ?)", postID, userID)
	if err != nil {
		return 0, err
	}
	if n, _ := res.RowsAffected(); n == 0 {
		return 0, errors.New("already liked")
	}
	if _, err = tx.Exec("UPDATE posts SET likes = likes + 1 WHERE id = ?", postID); err != nil {
		return 0, err
	}
	var count int32
	if err = tx.QueryRow("SELECT likes FROM posts WHERE id = ?", postID).Scan(&count); err != nil {
		return 0, err
	}
	return count, tx.Commit()
}

func (p *postRepository) UnlikePost(postID, userID uint64) (int32, error) {
	tx, err := p.db.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	res, err := tx.Exec("DELETE FROM post_likes WHERE post_id = ? AND user_id = ?", postID, userID)
	if err != nil {
		return 0, err
	}
	if n, _ := res.RowsAffected(); n == 0 {
		return 0, errors.New("not liked")
	}
	if _, err = tx.Exec("UPDATE posts SET likes = GREATEST(likes - 1, 0) WHERE id = ?", postID); err != nil {
		return 0, err
	}
	var count int32
	if err = tx.QueryRow("SELECT likes FROM posts WHERE id = ?", postID).Scan(&count); err != nil {
		return 0, err
	}
	return count, tx.Commit()
}
