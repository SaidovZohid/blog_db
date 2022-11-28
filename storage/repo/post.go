package repo

import "time"

type Post struct {
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

type PostStorageI interface {
	Create(u *Post) (*Post, error)
	Get(post_id int64) (*Post, error)
	Update(u *Post) (*Post, error)
	Delete(post_id int64) error
	GetAll(params *GetPostsParams) (*GetAllPostResult, error)
}

type GetAllPostResult struct {
	Posts []*Post
	Count int64
}

type GetPostsParams struct {
	Limit      int64
	Page       int64
	Search     string
	UserID     int64
	CategoryID int64
	SortByDate string
}
