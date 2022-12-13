package storage

import (
	"github.com/SaidovZohid/blog_db/storage/postgres"
	"github.com/SaidovZohid/blog_db/storage/repo"
	"github.com/jmoiron/sqlx"
)

type StorageI interface {
	User() repo.UserStorageI
	Category() repo.CategoryStorageI
	Post() repo.PostStorageI
	Comment() repo.CommentStorageI
	Like() repo.LikeStorageI
}

type StoragePg struct {
	userRepo     repo.UserStorageI
	categoryRepo repo.CategoryStorageI
	postRepo     repo.PostStorageI
	commentRepo  repo.CommentStorageI
	likeRepo     repo.LikeStorageI
}

func NewStoragePg(db *sqlx.DB) StorageI {
	return &StoragePg{
		userRepo:     postgres.NewUser(db),
		categoryRepo: postgres.NewCategory(db),
		postRepo:     postgres.NewPost(db),
		commentRepo:  postgres.NewComment(db),
		likeRepo:     postgres.NewLike(db),
	}
}

func (s *StoragePg) User() repo.UserStorageI {
	return s.userRepo
}

func (s *StoragePg) Category() repo.CategoryStorageI {
	return s.categoryRepo
}

func (s *StoragePg) Post() repo.PostStorageI {
	return s.postRepo
}

func (s *StoragePg) Comment() repo.CommentStorageI {
	return s.commentRepo
}

func (s *StoragePg) Like() repo.LikeStorageI {
	return s.likeRepo
}
