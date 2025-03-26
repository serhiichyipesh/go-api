package store

import (
	"context"
	"database/sql"
)

type Post struct {
	ID        int64    `json:"id"`
	Content   string   `json:"content"`
	Title     string   `json:"title"`
	UserID    int64    `json:"user_id"`
	CreatedAt string   `json:"created_at"`
	UpdatedAt string   `json:"updated_at"`
}

type PostsStore struct {
	db *sql.DB
}

func (s *PostsStore) Create(ctx context.Context, post *Post) error {
	query := `
		INSERT INTO posts (content, title, user_id)
		VALUES ($1, $2, $3)
		RETURNING id, created_at, updated_at
	`

	row := s.db.QueryRowContext(ctx, query, post.Content, post.Title, post.UserID)

	err := row.Scan(&post.ID, &post.CreatedAt, &post.UpdatedAt)

	if err != nil {
		return err
	}

	return nil
}

func (s *PostsStore) GetAll(ctx context.Context) ([]Post, error) {
	query := `
        SELECT id, content, title, user_id, created_at, updated_at
        FROM posts
        ORDER BY created_at DESC
    `

	rows, err := s.db.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var posts []Post

	for rows.Next() {
		var post Post
		err := rows.Scan(
			&post.ID,
			&post.Content,
			&post.Title,
			&post.UserID,
			&post.CreatedAt,
			&post.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil

}

func (s *PostsStore) GetByID(ctx context.Context, id int64) (*Post, error) {
	query := `
		SELECT id, content, title, user_id, created_at, updated_at
		FROM posts
		WHERE id = $1
	`

	row := s.db.QueryRowContext(ctx, query, id)

	var post Post
	err := row.Scan(
		&post.ID,
		&post.Content,
		&post.Title,
		&post.UserID,
		&post.CreatedAt,
		&post.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &post, nil
}