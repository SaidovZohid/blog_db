package postgres

import (
	"database/sql"
	"errors"

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

func (ld *likeRepo) CreateOrUpdate(like *repo.Like) (*repo.Like, error) {
	l, err := ld.Get(like.UserID, like.PostID)
	if errors.Is(err, sql.ErrNoRows) {
		query := `
			INSERT INTO likes (post_id, user_id, status) 
			VALUES ($1, $2, $3) RETURNING id
		`

		_, err = ld.db.Exec(
			query,
			like.PostID,
			like.UserID,
			like.Status,
		)
		if err != nil {
			return nil, err
		}
	} else if like != nil {
		if like.Status == l.Status {
			query := "DELETE FROM likes WHERE id = $1"
			_, err = ld.db.Exec(query, l.ID)
			if err != nil {
				return nil, err
			}
		} else {
			query := "UPDATE likes SET status = $1 WHERE id = $2"
			_, err = ld.db.Exec(query, like.Status, l.ID)
			if err != nil {
				return nil, err
			}
		}
	}

	return like, nil
}

func (ld *likeRepo) Get(userID, postID int64) (*repo.Like, error) {
	var res repo.Like

	query := `
		SELECT 
			id,
			user_id,
			post_id,
			status
		FROM likes 
		WHERE user_id = $1 AND post_id = $2
	`

	err := ld.db.QueryRow(query, userID, postID).Scan(
		&res.ID,
		&res.UserID,
		&res.PostID,
		&res.Status,
	)

	if err != nil {
		return nil, err
	}

	return &res, nil
} 

func (ld *likeRepo) GetLikesDislikesCount(postID int64) (*repo.LikesDislikesCountResult, error) {
	var res repo.LikesDislikesCountResult

	query := `
		SELECT 
			count(1) FILTER (WHERE status=true) as likes_count,
			count(1) FILTER (WHERE status=false) as dislikes_count
		FROM likes 
		WHERE post_id=$1
	`

	err := ld.db.QueryRow(query, postID).Scan(
		&res.Likes,
		&res.Dislikes,
	)
	if err != nil {
		return nil, err
	}

	return &res, nil
}