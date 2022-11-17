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
	GetAll(params *GetAllCategoryParams) (*GetAllCategoryResult, error)
}

type GetAllCategoryParams struct {
	Limit int32
	Page int32
	Search string
}

type GetAllCategoryResult struct {
	Categories []*Category
	Count int32
}