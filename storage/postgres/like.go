package postgres

import (
	"github.com/SaidovZohid/blog_db/storage/repo"
	"github.com/jmoiron/sqlx"
)

type likeRepo struct {
	db *sqlx.DB
}

func NewLike(db *sqlx.DB) repo.LikeStorageI {
	return &likeRepo{
		db: db,
	}
}

func (ld *likeRepo) Create(like *repo.Like) error {
	query := `
		INSERT INTO likes (
			post_id,
			user_id,
			status
		) VALUES ($1, $2, $3)
	`

	_, err := ld.db.Exec(
		query,
		like.PostID,
		like.UserID,
		like.Status,
	)
	if err != nil {
		return err
	}

	return nil
}

func (ld *likeRepo) Update(like *repo.Like) error {
	query := `
		UPDATE likes SET 
			status = $1
		WHERE id=$2
	`

	_, err := ld.db.Exec(
		query,
		like.Status,
		like.ID,
	)
	if err != nil {
		return err
	}

	return nil
}

func (ld *likeRepo) Delete(like_id int64) error {
	query := `
		DELETE FROM likes WHERE id=$1
	`

	_, err := ld.db.Exec(
		query,
		like_id,
	)
	if err != nil {
		return err
	}

	return nil
}

func (ld *likeRepo) GetAll(params *repo.GetAllLikesParams) (*repo.GetAllLikes, error) {
	var res repo.GetAllLikes
	query := "SELECT count(status) WHERE status=$1"
	err := ld.db.QueryRow(query, true).Scan(&res.Likes)
	if err != nil {
		return nil, err
	}
	err = ld.db.QueryRow(query, false).Scan(&res.Dislikes)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
