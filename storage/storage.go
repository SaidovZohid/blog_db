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
	EmailVer() repo.EmailVerI
}

type storagePg struct {
	userRepo repo.UserStorageI
	categoryRepo repo.CategoryStorageI
	postRepo repo.PostStorageI
	emailVerRepo repo.EmailVerI
}

func NewStoragePg(db *sqlx.DB) StorageI {
	return &storagePg{
		userRepo: postgres.NewUser(db),
		categoryRepo: postgres.NewCategory(db),
		postRepo: postgres.NewPost(db),
		emailVerRepo: postgres.NewEmailVer(db),
	}
}

func (s *storagePg) User() repo.UserStorageI {
	return s.userRepo
}

func (s *storagePg) Category() repo.CategoryStorageI {
	return s.categoryRepo
}

func (s *storagePg) Post() repo.PostStorageI {
	return s.postRepo
}

func (s *storagePg) EmailVer() repo.EmailVerI {
	return s.emailVerRepo
}