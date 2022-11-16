package postgres

import (
	"time"

	"github.com/SaidovZohid/blog_db/storage/repo"
	"github.com/jmoiron/sqlx"
)

type postRepo struct {
	db *sqlx.DB
}

func NewPost(db *sqlx.DB) repo.PostStorageI {
	return &postRepo{
		db: db,
	}
}

func (pr *postRepo) Create(p *repo.Post) (*repo.Post, error) {
	query := `
		INSERT INTO posts(
			title,
			description,
			image_url,
			user_id,
			category_id
		) VALUES ($1, $2, $3, $4, $5)
		RETURNING id, created_at
	`

	err := pr.db.QueryRow(
		query,
		p.Title,
		p.Description,
		p.ImageUrl,
		p.UserID,
		p.CategoryID,
	).Scan(
		&p.ID,
		&p.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return p, nil
}

func (pr *postRepo) Get(post_id int64) (*repo.Post, error) {
	var (
		res repo.Post
	)
	query := `
		SELECT
			id,
			title,
			description,
			image_url,
			user_id,
			category_id,
			created_at,
			updated_at,
			views_count
		FROM posts WHERE id = $1
	`

	err := pr.db.QueryRow(
		query,
		post_id,
	).Scan(
		&res.ID,
		&res.Title,
		&res.Description,
		&res.ImageUrl,
		&res.UserID,
		&res.CategoryID,
		&res.CreatedAt,
		&res.UpdatedAt,
		&res.ViewsCount,
	)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (pr *postRepo) Update(p *repo.Post) (*repo.Post, error) {
	var (
		res repo.Post
	)
	query := `
		UPDATE posts SET
			title = $1,
			description = $2,
			image_url = $3,
			user_id = $4,
			category_id = $5,
			views_count = $6,
			updated_at = $7
	    WHERE id = $8
		RETURNING 
			id,
			title,
			description,
			image_url,
			user_id,
			category_id,
			created_at,
			updated_at,
			views_count
	`

	err := pr.db.QueryRow(
		query,
		p.Title,
		p.Description,
		p.ImageUrl,
		p.UserID,
		p.CategoryID,
		p.ViewsCount,
		time.Now(),
		p.ID,
	).Scan(
		&res.ID,
		&res.Title,
		&res.Description,
		&res.ImageUrl,
		&res.UserID,
		&res.CategoryID,
		&res.CreatedAt,
		&res.UpdatedAt,
		&res.ViewsCount,
	)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (pr *postRepo) Delete(post_id int64) error {
	query := `
		DELETE FROM posts WHERE id = $1
	`

	_, err := pr.db.Exec(
		query,
		post_id,
	)

	if err != nil {
		return err
	}

	return nil
}