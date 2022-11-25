package repo

import "time"

const (
	UserTypeSuperadmin = "superadmin"
	UserTypeUser       = "user"
)

type User struct {
	ID              int64
	FirstName       string
	LastName        string
	PhoneNumber     *string
	Email           string
	Gender          *string
	Password        string
	UserName        *string
	ProfileImageUrl *string
	Type            string
	CreatedAt       time.Time
}

type UserStorageI interface {
	Create(u *User) (*User, error)
	UpdatePassword(u *UpdatePassword) error
	Get(user_id int64) (*User, error)
	Update(u *User) (*User, error)
	Delete(user_id int64) error
	GetAll(params *GetAllUserParams) (*GetAllUsersResult, error)
	GetByEmail(user_email string) (*User, error)
}

type UpdatePassword struct {
	UserID int64
	Password string 
}

type GetAllUserParams struct {
	Limit  int32
	Page   int32
	Search string
}

type GetAllUsersResult struct {
	Users []*User
	Count int32
}
