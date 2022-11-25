package postgres

import (
	"fmt"
	"time"

	"github.com/SaidovZohid/blog_db/storage/repo"
	"github.com/jmoiron/sqlx"
)

type commentRepo struct {
	db *sqlx.DB
}

func NewComment(db *sqlx.DB) repo.CommentStorageI {
	return &commentRepo{
		db: db,
	}
}

func (pr *commentRepo) Create(c *repo.Comment) (*repo.Comment, error) {
	query := `
		INSERT INTO comments (
			user_id,
			post_id,
			description
		) VALUES ($1, $2, $3)
		RETURNING 
		id, 
		user_id,
		post_id,
		description,
		created_at
	`

	err := pr.db.QueryRow(
		query,
		c.UserID,
		c.PostID,
		c.Description,
	).Scan(
		&c.ID,
		&c.UserID,
		&c.PostID,
		&c.Description,
		&c.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return c, nil
}

func (pr *commentRepo) Get(comment_id int64) (*repo.Comment, error) {
	var (
		res repo.Comment
	)
	query := `
		SELECT
			id,
			post_id,
			user_id,
			description,
			created_at,
			updated_at,
		FROM comments WHERE id = $1
	`

	err := pr.db.QueryRow(
		query,
		comment_id,
	).Scan(
		&res.ID,
		&res.PostID,
		&res.UserID,
		&res.Description,
		&res.CreatedAt,
		&res.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (cr *commentRepo) Update(c *repo.UpdateComment) (*repo.UpdateComment, error) {
	var (
		res repo.UpdateComment
	)
	query := `
		UPDATE comments SET
			description = $1,
			updated_at = $2
	    WHERE id = $3
		RETURNING 
			id,
			description,
			post_id,
			user_id,
			created_at,
			updated_at
	`

	err := cr.db.QueryRow(
		query,
		c.Description,
		time.Now(),
		c.ID,
	).Scan(
		&res.ID,
		&res.Description,
		&res.PostID,
		&res.UserID,
		&res.CreatedAt,
		&res.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (cr *commentRepo) Delete(comment_id int64) error {
	query := `
		DELETE FROM comments WHERE id = $1
	`

	_, err := cr.db.Exec(
		query,
		comment_id,
	)

	if err != nil {
		return err
	}

	return nil
}

func (pr *commentRepo) GetAll(params *repo.GetCommentsParams) (*repo.GetAllCommentsResult, error) {
	result := repo.GetAllCommentsResult{
		Comments: make([]*repo.Comment, 0),
	}

	offset := (params.Page - 1) * params.Limit

	
	limit := fmt.Sprintf(" LIMIT %d OFFSET %d", params.Limit, offset)
	
	filter := " WHERE true "
	if params.UserID != 0 {
		filter += fmt.Sprintf(" AND c.user_id = %d", params.UserID)
	}

	if params.PostID != 0 {
		filter += fmt.Sprintf(" AND c.post_id = %d", params.PostID)
	}

	query := `
		SELECT
			c.id,
			c.post_id,
			c.user_id,
			c.description,
			c.created_at,
			c.updated_at,
			u.first_name,
			u.last_name,
			u.email,
			u.profile_image_url
		FROM comments c 
		INNER JOIN users u 	ON c.user_id = u.id 
	` + filter  + `
	ORDER BY c.created_at DESC` + limit

	rows, err := pr.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var comment repo.Comment
		err := rows.Scan(
			&comment.ID,
			&comment.PostID,
			&comment.UserID,
			&comment.Description,
			&comment.CreatedAt,
			&comment.UpdatedAt,
			&comment.User.FirstName,
			&comment.User.LastName,
			&comment.User.Email,
			&comment.User.ProfileImageUrl,
		)
		if err != nil {
			return nil, err
		}

		result.Comments = append(result.Comments, &comment)
	}

	queryCount := "SELECT count(1) FROM comments c INNER JOIN users u ON u.id = c.user_id" + filter

	err = pr.db.QueryRow(queryCount).Scan(&result.Count)
	if err != nil {
		return nil, err
	}

	return &result, nil
}