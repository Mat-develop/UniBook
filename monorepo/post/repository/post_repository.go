package repository

import (
	"database/sql"
	"errors"
	"time"
	"v1/post/model"
)

const (
	CreateQuery            = "INSERT INTO posts (title, body, community_id, user_id) VALUES (?, ?, ?, ?)"
	FindByCommunityIdQuery = `
	SELECT 
		p.id, 
		p.user_id, 
		p.community_id, 
		c.name as communityName,
		p.title, 
		p.body, 
		IFNULL(p.image_url,''), 
		IFNULL(p.likes,0), 
		p.created_at, 
		u.nick 
	FROM posts p 
	INNER JOIN users u ON u.id = p.user_id 
	INNER JOIN community c on c.id = p.community_id
	WHERE p.community_id = ?`

	FindByUserIdQuery = `
	SELECT 
		p.id, 
		p.user_id, 
		p.community_id, 
		c.name as communityName,
		p.title, 
		p.body, 
		IFNULL(p.image_url,''), 
		IFNULL(p.likes,0), 
		p.created_at, 
		u.nick 
	FROM posts p 
	INNER JOIN users u ON u.id = p.user_id 
	INNER JOIN community c on c.id = p.community_id
	WHERE p.user_id = ?`

	UpdateQuery = "UPDATE posts SET title = ?, body = ? WHERE id = ? AND user_id = ?"
	DeleteQuery = "DELETE FROM posts WHERE id = ? AND user_id = ?"
)

type PostRepository interface {
	FindCommunityPosts(communityId uint64) ([]model.Post, error)
	FindUserPosts(userId uint64) ([]model.Post, error)
	FindPostByName() ([]model.Post, error)
	Create(userId uint64, postBody model.PostDTO) error
	Update(postID uint64, userID uint64, postBody model.PostDTO) error
	Delete(postID uint64, userID uint64) error
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
		)
		if err := rows.Scan(&id, &userId, &community, &communityName, &title, &body, &imageURL, &likes, &createdAt, &userNick); err != nil {
			return nil, err
		}

		post := model.Post{
			ID:            uint64(id),
			CommunityId:   int32(0),
			CommunityName: communityName,
			UserId:        uint64(userId),
			UserNick:      userNick,
			Title:         title,
			Body:          body,
			ImageUrl:      imageURL,
			Likes:         int32(likes),
			CreatedAt:     createdAt,
		}
		if community.Valid {
			post.CommunityId = int32(community.Int64)
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func (p *postRepository) FindCommunityPosts(communityId uint64) ([]model.Post, error) {
	rows, err := p.db.Query(FindByCommunityIdQuery, communityId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return p.scanPosts(rows)
}

func (p *postRepository) FindUserPosts(userId uint64) ([]model.Post, error) {
	rows, err := p.db.Query(FindByUserIdQuery, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return p.scanPosts(rows)
}

func (p *postRepository) FindPostByName() ([]model.Post, error) {
	return []model.Post{}, errors.New("hello")
}

func (p *postRepository) Create(userId uint64, postBody model.PostDTO) error {
	statement, err := p.db.Prepare(CreateQuery)
	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(postBody.Title, postBody.Body, postBody.CommunityId, userId)
	if err != nil {
		return err
	}

	return nil
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
