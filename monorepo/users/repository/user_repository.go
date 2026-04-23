package repository

import (
	"database/sql"
	"fmt"
	"v1/users/model"
)

const (
	columns             = "(name, nick, email, password, image_url)"
	columnsNoPassword   = "id ,name, nick, created_at"
	createQuery         = "INSERT INTO users" + columns + "VALUES (?, ?, ?, ?, ?)"
	findByNameQuery     = "SELECT " + columnsNoPassword + " FROM users WHERE name LIKE ? OR nick LIKE ?"
	findAllQuery        = "SELECT " + columnsNoPassword + " FROM users"
	findPasswordById    = "SELECT password from users where id = ?"
	updateQuery         = "UPDATE users SET name = ?, nick = ?, email = ? WHERE id = ?;"
	updatePasswordQuery = "UPDATE users SET password = ? where id = ?;"
	deleteQuery         = "DELETE FROM users WHERE id = ?;"
	findByEmailQuery    = "SELECT id, password FROM users WHERE email = ?"
	followQuery         = "INSERT IGNORE INTO followers (user_id, follower_id) values(?, ?)"
	unfollowQuery       = "DELETE FROM followers WHERE user_id = ? AND follower_id = ?"
	findFollowersQuery = `SELECT u.id, u.name, u.nick, u.image_url, u.created_at
	FROM users u INNER JOIN followers f ON u.id = f.follower_id WHERE f.user_id = ?`
	findFollowingQuery = `SELECT u.id, u.name, u.nick, u.image_url, u.created_at
	FROM users u INNER JOIN followers f ON u.id = f.user_id WHERE f.follower_id = ?`
)

type UserRepository interface {
	Create(user model.User) (uint64, error)
	FindUserByName(userNameOrNick string) ([]model.User, error) // mudar para find by Name
	FindUsers() ([]model.User, error)
	Update(ID uint64, user model.User) error
	UpdatePassword(ID uint64, password string) error
	Delete(ID uint64) error
	FindUserByEmail(email string) (model.User, error)
	FindPasswordById(ID uint64) (string, error)
	Follow(userId, followerID uint64) error
	Unfollow(userId, followerID uint64) error
	FindFollowers(userId uint64) ([]model.User, error)
	FindFollowing(userId uint64) ([]model.User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

// Insert inside the DB
func (u *userRepository) Create(user model.User) (uint64, error) {
	stm, err := u.db.Prepare(
		createQuery,
	)
	if err != nil {
		return 0, err
	}

	defer stm.Close()

	result, err := stm.Exec(user.Name, user.Nick, user.Email, user.Password, user.ImageUrl)
	if err != nil {
		return 0, err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastID), nil
}

func (u *userRepository) FindUserByName(userNameOrNick string) ([]model.User, error) {
	userNameOrNick = fmt.Sprintf("%%%s", userNameOrNick) // to use like %name%

	rows, err := u.db.Query(
		findByNameQuery,
		userNameOrNick,
		userNameOrNick,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var users []model.User

	for rows.Next() {
		var user model.User

		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (u *userRepository) FindUsers() ([]model.User, error) {
	rows, err := u.db.Query(
		findAllQuery,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var users []model.User

	for rows.Next() {
		var user model.User

		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}
	return users, nil
}

func (u *userRepository) FindPasswordById(ID uint64) (string, error) {
	rows, err := u.db.Query(findAllQuery)
	if err != nil {
		return "", err
	}

	defer rows.Close()
	var user model.User

	for rows.Next() {
		if err = rows.Scan(&user.Password); err != nil {
			return "", err
		}
	}

	return user.Password, nil
}

func (u *userRepository) Update(ID uint64, user model.User) error {
	stm, err := u.db.Prepare(
		updateQuery,
	)
	if err != nil {
		return err
	}
	defer stm.Close()
	if _, err = stm.Exec(user.Name, user.Nick, user.Email, ID); err != nil {
		return err
	}
	return nil
}

func (u *userRepository) Delete(ID uint64) error {
	stm, err := u.db.Prepare(
		deleteQuery,
	)
	if err != nil {
		return err
	}
	defer stm.Close()

	if _, err = stm.Exec(ID); err != nil {
		return err
	}
	return nil
}

func (u *userRepository) FindUserByEmail(email string) (model.User, error) {

	row, err := u.db.Query(
		findByEmailQuery,
		email,
	)

	if err != nil {
		return model.User{}, err
	}

	defer row.Close()

	var user model.User

	for row.Next() {
		if err = row.Scan(&user.ID, &user.Password); err != nil {
			return model.User{}, err
		}
	}

	return user, nil

}

func (u *userRepository) Follow(userId, followerID uint64) error {

	stm, err := u.db.Prepare(followQuery)
	if err != nil {
		return err
	}
	defer stm.Close()

	if _, err = stm.Exec(userId, followerID); err != nil {
		return err
	}

	return nil
}

func (u *userRepository) Unfollow(userId, followerID uint64) error {

	stm, err := u.db.Prepare(unfollowQuery)
	if err != nil {
		return err
	}
	defer stm.Close()

	if _, err = stm.Exec(userId, followerID); err != nil {
		return err
	}

	return nil
}

func (u *userRepository) FindFollowers(userId uint64) ([]model.User, error) {
	rows, err := u.db.Query(
		findFollowersQuery,
		userId,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var users []model.User

	for rows.Next() {
		var user model.User

		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.ImageUrl,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (u *userRepository) FindFollowing(userId uint64) ([]model.User, error) {
	rows, err := u.db.Query(
		findFollowingQuery,
		userId,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var users []model.User

	for rows.Next() {
		var user model.User

		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.ImageUrl,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (u *userRepository) UpdatePassword(ID uint64, password string) error {
	stm, err := u.db.Prepare(updatePasswordQuery)
	if err != nil {
		return err
	}
	defer stm.Close()

	if _, err = stm.Exec(password, ID); err != nil {
		return err
	}

	return nil
}
