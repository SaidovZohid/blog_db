package storage

import (
	"github.com/SaidovZohid/blog_db/storage/postgres"
	"github.com/SaidovZohid/blog_db/storage/repo"
	"github.com/jmoiron/sqlx"
)

type StorageI interface {
	User() repo.UserStorageI
}

type storagePg struct {
	userRepo repo.UserStorageI
}

func NewStoragePg(db *sqlx.DB) StorageI {
	return &storagePg{
		userRepo: postgres.NewUser(db),
	}
}

func (s *storagePg) User() repo.UserStorageI {
	return s.userRepo
}