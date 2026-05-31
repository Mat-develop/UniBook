package repository

import (
	"database/sql"
	"fmt"
	"strings"
	"v1/tag/model"
)

const (
	insertTagQuery    = "INSERT IGNORE INTO tags (name, created_by) VALUES (?, ?)"
	selectTagByName   = "SELECT id FROM tags WHERE name = ?"
	findAllTagsQuery  = "SELECT id, name, created_at FROM tags ORDER BY name"
	searchTagsQuery   = "SELECT id, name, created_at FROM tags WHERE name LIKE ? ORDER BY name"
)

type TagRepository interface {
	Create(name string, userID uint64) (uint64, error)
	FindAll() ([]model.Tag, error)
	Search(name string) ([]model.Tag, error)
	FindOrCreate(name string, userID uint64) (uint64, error)
}

type tagRepository struct {
	db *sql.DB
}

func NewTagRepository(db *sql.DB) TagRepository {
	return &tagRepository{db: db}
}

func (r *tagRepository) Create(name string, userID uint64) (uint64, error) {
	name = strings.ToLower(strings.TrimSpace(name))
	res, err := r.db.Exec(insertTagQuery, name, userID)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	if id == 0 {
		return 0, fmt.Errorf("tag %q already exists", name)
	}
	return uint64(id), nil
}

func (r *tagRepository) FindAll() ([]model.Tag, error) {
	rows, err := r.db.Query(findAllTagsQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanTags(rows)
}

func (r *tagRepository) Search(name string) ([]model.Tag, error) {
	rows, err := r.db.Query(searchTagsQuery, "%"+name+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanTags(rows)
}

func (r *tagRepository) FindOrCreate(name string, userID uint64) (uint64, error) {
	name = strings.ToLower(strings.TrimSpace(name))
	var id uint64
	if err := r.db.QueryRow(selectTagByName, name).Scan(&id); err == nil {
		return id, nil
	}
	if _, err := r.db.Exec(insertTagQuery, name, userID); err != nil {
		return 0, err
	}
	if err := r.db.QueryRow(selectTagByName, name).Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func scanTags(rows *sql.Rows) ([]model.Tag, error) {
	tags := []model.Tag{}
	for rows.Next() {
		var t model.Tag
		if err := rows.Scan(&t.ID, &t.Name, &t.CreatedAt); err != nil {
			return nil, err
		}
		tags = append(tags, t)
	}
	return tags, nil
}
