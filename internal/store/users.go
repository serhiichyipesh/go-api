package store

import (
	"context"
	"database/sql"
	"errors"
	"github.com/lib/pq"
)

type User struct {
	ID        int64  `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"-"`
	CreatedAt string `json:"created_at"`
}

type UsersStore struct {
	db *sql.DB
}

func (s *UsersStore) Create(ctx context.Context, user *User) error {
	query := `
        INSERT INTO users (username, email, password)
        VALUES ($1, $2, $3)
        RETURNING id, created_at
    `

	row := s.db.QueryRowContext(ctx, query,
		user.Username,
		user.Email,
		user.Password,
	)

	err := row.Scan(&user.ID, &user.CreatedAt)

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			return errors.New("username or email already exists")
		}
		return err
	}

	return nil
}

func (s *UsersStore) GetAll(ctx context.Context) ([]User, error) {
	query := `
        SELECT id, username, email, created_at 
        FROM users 
        ORDER BY created_at DESC
    `

	rows, err := s.db.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []User
	
	for rows.Next() {
		var user User
		err := rows.Scan(
			&user.ID,
			&user.Username,
			&user.Email,
			&user.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
