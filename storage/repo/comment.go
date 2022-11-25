package repo

import "time"

type Comment struct {
	ID          int64
	PostID      int64
	UserID      int64
	Description string
	CreatedAt   time.Time
	UpdatedAt   *time.Time
	User        CommentUser
}

type UpdateComment struct {
	ID          int64
	PostID      int64
	UserID      int64
	Description string
	CreatedAt   time.Time
	UpdatedAt   *time.Time
	User        CommentUser
}

type CommentUser struct {
	ID              int64
	FirstName       string
	LastName        string
	Email           string
	ProfileImageUrl *string
}

type CommentStorageI interface {
	Create(u *Comment) (*Comment, error)
	Update(u *UpdateComment) (*UpdateComment, error)
	Delete(comment_id int64) error
	GetAll(params *GetCommentsParams) (*GetAllCommentsResult, error)
}

type GetAllCommentsResult struct {
	Comments []*Comment
	Count    int64
}

type GetCommentsParams struct {
	Limit  int64
	Page   int64
	UserID int64
	PostID int64
}
