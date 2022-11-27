package models

import "time"

type Comment struct {
	ID          int64        `json:"id"`
	Description string       `json:"description"`
	UserID      int64        `json:"user_id"`
	PostID      int64        `json:"post_id"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   *time.Time   `json:"updated_at"`
	User        *CommentUser `json:"user"`
}

type UpdateComment struct {
	ID          int64        `json:"id"`
	Description string       `json:"description"`
	UserID      int64        `json:"user_id"`
	PostID      int64        `json:"post_id"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   *time.Time   `json:"updated_at"`
}

type CommentUser struct {
	FirstName       string  `json:"first_name"`
	Lastname        string  `json:"last_name"`
	Email           string  `json:"email"`
	ProfileImageUrl *string `json:"profile_image_url"`
}

type CreateCommentRequest struct {
	Description string `json:"description"`
	PostID      int64  `json:"post_id"`
}

type UpdateCommentRequest struct {
	Description string `json:"description"`
}

type GetAllCommentsParams struct {
	Limit  int64 `json:"limit" binding:"required" default:"10"`
	Page   int64 `json:"page" binding:"required" default:"1"`
	UserID int64 `json:"user_id"`
	PostID int64 `json:"category_id"`
}

type GetAllCommentsResponse struct {
	Comments []*Comment `json:"comments"`
	Count    int64      `json:"count"`
}
