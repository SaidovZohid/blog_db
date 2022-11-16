package repo

import "time"

type Category struct {
	ID        int64
	Title     string
	CreatedAt time.Time
}

type CategoryStorageI interface {
	Create(c *Category) (*Category, error)
	Get(category_id int64) (*Category, error)
	Update(u *Category) (*Category, error)
	Delete(category_id int64) error
}