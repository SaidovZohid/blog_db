package models

import "time"

type Post struct {
	ID           int64         `json:"id"`
	Title        string        `json:"title"`
	Description  string        `json:"description"`
	ImageUrl     *string       `json:"image_url"`
	UserID       int64         `json:"user_id"`
	CategoryID   int64         `json:"category_id"`
	UpdatedAt    *time.Time    `json:"updated_at"`
	ViewsCount   int32         `json:"views_count"`
	CreatedAt    time.Time     `json:"created_at"`
	PostLikeInfo *PostLikeInfo `json:"like_info"`
}

type PostLikeInfo struct {
	LikesCount    int64 `json:"likes_count"`
	DislikesCount int64 `json:"dislikes_count"`
}

type CreatePostRequest struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	ImageUrl    *string `json:"image_url"`
	UserID      int64   `json:"user_id"`
	CategoryID  int64   `json:"category_id"`
}

type UpdatePostRequest struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	ImageUrl    *string `json:"image_url"`
	UserID      int64   `json:"user_id"`
	CategoryID  int64   `json:"category_id"`
	ViewsCount  int32   `json:"views_count"`
}

type GetAllPostsParams struct {
	Limit      int64  `json:"limit" binding:"required" default:"10"`
	Page       int64  `json:"page" binding:"required" default:"1"`
	Search     string `json:"search"`
	UserID     int64  `json:"user_id"`
	CategoryID int64  `json:"category_id"`
	SortByDate string `json:"sort" enums:"desc,asc" default:"desc"`
}

type GetAllPostsResponse struct {
	Posts []*Post `json:"posts"`
	Count int64   `json:"count"`
}
