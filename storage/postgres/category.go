package postgres

import (
	"fmt"

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
		&category.CreatedAt,
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

func (ur *categoryRepo) GetAll(params *repo.GetAllCategoryParams) (*repo.GetAllCategoryResult, error) {
	result := repo.GetAllCategoryResult{
		Categories: make([]*repo.Category, 0),
	}

	offset := (params.Page - 1) * params.Limit

	
	limit := fmt.Sprintf(" LIMIT %d OFFSET %d", params.Limit, offset)
	
	filter := ""
	if params.Search != "" {
		filter = "WHERE title ILIKE '%s'" + "%" + params.Search + "%"
	}

	query := `
		SELECT 
			id,
			title,
			created_at
		FROM categories
	` + filter  + `
		ORDER BY created_at DESC
	` + limit

	rows, err := ur.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var category repo.Category
		err := rows.Scan(
			&category.ID,
			&category.Title,
			&category.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		result.Categories = append(result.Categories, &category)
	}

	queryCount := "SELECT count(1) FROM categories " + filter

	err = ur.db.QueryRow(queryCount).Scan(&result.Count)
	
	if err != nil {
		return nil, err
	}

	return &result, nil
}