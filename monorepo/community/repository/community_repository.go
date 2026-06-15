package repository

import (
	"database/sql"
	"fmt"

	"v1/community/model"
	userModel "v1/users/model"
)

const (
	createCommunityQuery   = "INSERT INTO community (name, description, image_url, created_by) VALUES (?, ?, ?, ?)"
	findByCreatorQuery     = `SELECT c.id, c.name, c.description, c.image_url, c.created_at, COUNT(cf.user_id) AS members
	FROM community c LEFT JOIN community_followers cf ON c.id = cf.community_id
	WHERE c.created_by = ?
	GROUP BY c.id, c.name, c.description, c.image_url, c.created_at`
	findCommunityByID = `SELECT c.id, c.name, c.description, c.image_url, c.created_at, COUNT(cf.user_id) AS members
	FROM community c LEFT JOIN community_followers cf ON c.id = cf.community_id
	WHERE c.id = ?
	GROUP BY c.id, c.name, c.description, c.image_url, c.created_at`
	findCommunityByName    = "SELECT id, name, description, image_url, created_at FROM community WHERE name = ?"
	findAllCommunities     = `SELECT c.id, c.name, c.description, c.image_url, c.created_at, COUNT(cf.user_id) AS members
	FROM community c LEFT JOIN community_followers cf ON c.id = cf.community_id
	GROUP BY c.id, c.name, c.description, c.image_url, c.created_at`
	searchCommunitiesQuery = `SELECT c.id, c.name, c.description, c.image_url, c.created_at, COUNT(cf.user_id) AS members
	FROM community c LEFT JOIN community_followers cf ON c.id = cf.community_id
	WHERE c.name LIKE ?
	GROUP BY c.id, c.name, c.description, c.image_url, c.created_at`
	deleteCommunityQuery   = "DELETE FROM community WHERE id = ?"
	followCommunityQuery   = "INSERT IGNORE INTO community_followers (user_id, community_id) VALUES (?, ?)"
	unfollowCommunityQuery = "DELETE FROM community_followers WHERE user_id = ? AND community_id = ?"
	findFollowersQuery     = "SELECT u.id, u.name, u.nick, u.image_url, u.created_at FROM users u INNER JOIN community_followers cf ON u.id = cf.user_id WHERE cf.community_id = ?"
	findJoinedByUserQuery  = "SELECT community_id FROM community_followers WHERE user_id = ?"
)

type CommunityRepository interface {
	Create(c model.Community) (uint64, error)
	FindByID(id uint64) (model.Community, error)
	FindByName(name string) (model.Community, error)
	FindAll() ([]model.Community, error)
	FindByCreator(userID uint64) ([]model.Community, error)
	Search(q string) ([]model.Community, error)
	Delete(id uint64) error

	Follow(userID, communityID uint64) error
	Unfollow(userID, communityID uint64) error
	FindFollowers(communityID uint64) ([]userModel.User, error)
	FindJoinedByUser(userID uint64) ([]uint64, error)
}

type communityRepository struct {
	db *sql.DB
}

func NewCommunityRepository(db *sql.DB) CommunityRepository {
	return &communityRepository{db: db}
}

func (r *communityRepository) Create(community model.Community) (uint64, error) {
	stmt, err := r.db.Prepare(createCommunityQuery)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(community.Name, community.Description, community.ImageUrl, community.CreatedBy)
	if err != nil {
		return 0, err
	}
	lastID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return uint64(lastID), nil
}

func (r *communityRepository) FindByCreator(userID uint64) ([]model.Community, error) {
	rows, err := r.db.Query(findByCreatorQuery, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	communities := []model.Community{}
	for rows.Next() {
		var c model.Community
		if err := rows.Scan(&c.Id, &c.Name, &c.Description, &c.ImageUrl, &c.CreatedAt, &c.Members); err != nil {
			return nil, err
		}
		communities = append(communities, c)
	}
	return communities, nil
}

func (r *communityRepository) FindByID(id uint64) (model.Community, error) {
	var community model.Community
	row := r.db.QueryRow(findCommunityByID, id)
	if err := row.Scan(&community.Id, &community.Name, &community.Description, &community.ImageUrl, &community.CreatedAt, &community.Members); err != nil {
		if err == sql.ErrNoRows {
			return community, fmt.Errorf("community %d not found", id)
		}
		return community, err
	}
	return community, nil
}

func (r *communityRepository) FindByName(name string) (model.Community, error) {
	var community model.Community
	row := r.db.QueryRow(findCommunityByName, name)
	if err := row.Scan(&community.Id, &community.Name, &community.Description, &community.ImageUrl, &community.CreatedAt); err != nil {
		if err == sql.ErrNoRows {
			return community, fmt.Errorf("community %s not found", name)
		}
		return community, err
	}
	return community, nil
}

func (r *communityRepository) FindAll() ([]model.Community, error) {
	rows, err := r.db.Query(findAllCommunities)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	communities := []model.Community{}
	for rows.Next() {
		var community model.Community
		if err := rows.Scan(&community.Id, &community.Name, &community.Description, &community.ImageUrl, &community.CreatedAt, &community.Members); err != nil {
			return nil, err
		}
		communities = append(communities, community)
	}
	return communities, nil
}

func (r *communityRepository) Search(q string) ([]model.Community, error) {
	rows, err := r.db.Query(searchCommunitiesQuery, "%"+q+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	communities := []model.Community{}
	for rows.Next() {
		var c model.Community
		if err := rows.Scan(&c.Id, &c.Name, &c.Description, &c.ImageUrl, &c.CreatedAt, &c.Members); err != nil {
			return nil, err
		}
		communities = append(communities, c)
	}
	return communities, nil
}

func (r *communityRepository) Delete(id uint64) error {
	stmt, err := r.db.Prepare(deleteCommunityQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	return err
}

func (r *communityRepository) Follow(userID, communityID uint64) error {
	stmt, err := r.db.Prepare(followCommunityQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(userID, communityID)
	return err
}

func (r *communityRepository) Unfollow(userID, communityID uint64) error {
	stmt, err := r.db.Prepare(unfollowCommunityQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(userID, communityID)
	return err
}

func (r *communityRepository) FindJoinedByUser(userID uint64) ([]uint64, error) {
	rows, err := r.db.Query(findJoinedByUserQuery, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ids := []uint64{}
	for rows.Next() {
		var id uint64
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}
	return ids, nil
}

func (r *communityRepository) FindFollowers(communityID uint64) ([]userModel.User, error) {
	rows, err := r.db.Query(findFollowersQuery, communityID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	followers := []userModel.User{}
	for rows.Next() {
		var user userModel.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Nick, &user.ImageUrl, &user.CreatedAt); err != nil {
			return nil, err
		}
		followers = append(followers, user)
	}
	return followers, nil
}
