package postgres

import (
	"github.com/SaidovZohid/blog_db/storage/repo"
	"github.com/jmoiron/sqlx"
)

type categoryRepo struct {
	db *sqlx.DB
}

func NewCategory(db *sqlx.DB) repo.CategoryStorageI {
	return &categoryRepo{
		db: db,
	}
}

func (cr *categoryRepo) Create(category *repo.Category) (*repo.Category, error) {
	query := `
		INSERT INTO categories(title) VALUES ($1) RETURNING id, created_at
	`

	err := cr.db.QueryRow(
		query,
		category.Title,
	).Scan(
		&category.ID,
		category.CreatedAt,
	)
	
	if err != nil {
		return nil, err
	}

	return category, nil
}

func (cr *categoryRepo) Get(category_id int64) (*repo.Category, error) {
	query := `
		SELECT 
			id,
			title,
			created_at
		FROM categories WHERE id = $1
	`

	var result repo.Category

	err := cr.db.QueryRow(
		query,
		category_id,
	).Scan(
		&result.ID,
		&result.Title,
		&result.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (cr *categoryRepo) Update(category *repo.Category) (*repo.Category, error) {
	query := `
		UPDATE categories SET 
			title = $1
		WHERE id = $2 
		RETURNING id, title, created_at
	`

	var result repo.Category

	err := cr.db.QueryRow(
		query,
		category.Title,
		category.ID,
	).Scan(
		&result.ID,
		&result.Title,
		&result.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (cr *categoryRepo) Delete(category_id int64) error {
	query := `
		DELETE FROM categories WHERE id = $1
	`

	_, err := cr.db.Exec(
		query,
		category_id,
	)

	if err != nil {
		return err
	}

	return nil
}