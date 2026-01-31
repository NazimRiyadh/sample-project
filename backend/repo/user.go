package repo

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type User struct {
	ID        int    `json:"id" db:"id"`
	Firstname string `json:"first_name" db:"first_name"`
	Lastname  string `json:"last_name" db:"last_name"`
	Email     string `json:"email" db:"email"`
	Password  string `json:"password" db:"password"`
	IsOwner   bool   `json:"is_owner" db:"is_owner"`
}

type UserRepo interface {
	Create(user User) (*User, error)
	Find(email, pass string) (*User, error)
}

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) *userRepo {
	return &userRepo{db: db}
}

func (r *userRepo) Create(user User) (*User, error) {
	query := `
		INSERT INTO users (
			first_name,
			last_name,
			email,
			password,
			is_owner
		)
		VALUES (
			:first_name,
			:last_name,
			:email,
			:password,
			:is_owner
		)
		RETURNING id
	`

	var userID int
	rows, err := r.db.NamedQuery(query, user)
	if err != nil {
		return nil, err
	}

	if rows.Next() {
		err = rows.Scan(&userID)
		if err != nil {
			return nil, err
		}
	}
	user.ID = userID
	return &user, nil
}

func (r *userRepo) Find(email, password string) (*User, error) {
	query := `
		SELECT
			id,
			first_name,
			last_name,
			email,
			password,
			is_owner
		FROM users
		WHERE email = $1 AND password = $2
		LIMIT 1
	`

	var user User
	err := r.db.Get(&user, query, email, password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // user not found
		}
		return nil, err
	}

	return &user, nil
}
