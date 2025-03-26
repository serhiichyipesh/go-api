package store

import (
	"context"
	"database/sql"
)

type Storage struct {
	Posts interface {
		Create(ctx context.Context, post *Post) error
		GetAll(ctx context.Context) ([]Post, error)
		GetByID(ctx context.Context, id int64) (*Post, error)
	}
	Users interface {
		Create(ctx context.Context, user *User) error
		GetAll(ctx context.Context) ([]User, error)
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Posts: &PostsStore{db},
		Users: &UsersStore{db},
	}
}
