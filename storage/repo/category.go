package repo

import "time"

type Category struct {
	ID          int64
	Title       string
	Description string
	ImageUrl    *string
	UserID      int64
	CategoryID  int64
	CreatedAt   time.Time
	UpdatedAt   *time.Time
	ViewsCount  int32
}

type CategoryStorageI interface {
	Create(u *Category) (*Category, error)
	Get(category_id int64) (*Category, error)
	Update(u *Category) (*Category, error)
	Delete(category_id int64) error
}