package repository

import (
	"database/sql"
	"time"
	"v1/comment/model"
)

const (
	insertCommentQuery = `
		INSERT INTO post_comments (post_id, user_id, body) VALUES (?, ?, ?)`

	findByPostQuery = `
		SELECT c.id, c.post_id, c.user_id, u.nick, c.body, c.created_at
		FROM post_comments c
		INNER JOIN users u ON u.id = c.user_id
		WHERE c.post_id = ?
		ORDER BY c.created_at ASC`
)

type CommentRepository interface {
	Create(postID, userID uint64, body string) error
	FindByPost(postID uint64) ([]model.Comment, error)
}

type commentRepository struct {
	db *sql.DB
}

func NewCommentRepository(db *sql.DB) CommentRepository {
	return &commentRepository{db: db}
}

func (r *commentRepository) Create(postID, userID uint64, body string) error {
	_, err := r.db.Exec(insertCommentQuery, postID, userID, body)
	return err
}

func (r *commentRepository) FindByPost(postID uint64) ([]model.Comment, error) {
	rows, err := r.db.Query(findByPostQuery, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	comments := []model.Comment{}
	for rows.Next() {
		var (
			id        uint64
			postId    uint64
			userId    uint64
			userNick  string
			body      string
			createdAt time.Time
		)
		if err = rows.Scan(&id, &postId, &userId, &userNick, &body, &createdAt); err != nil {
			return nil, err
		}
		comments = append(comments, model.Comment{
			ID:        id,
			PostID:    postId,
			UserID:    userId,
			UserNick:  userNick,
			Body:      body,
			CreatedAt: createdAt,
		})
	}
	return comments, nil
}
