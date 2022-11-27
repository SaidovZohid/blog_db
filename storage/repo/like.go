package repo

type Like struct {
	ID     int64
	PostID int64
	UserID int64
	Status bool
}

type LikeStorageI interface {
	Create(like *Like) error
	Update(like *Like) error
	Delete(like_id int64) error
	GetAll(params *GetAllLikesParams) (*GetAllLikes, error)
}

type GetAllLikes struct {
	Likes    int
	Dislikes int
}

type GetAllLikesParams struct {
	UserID int64
	PostID int64
}
