package repo

import "time"

type User struct {
	ID              int64
	FirstName       string
	LastName        string
	PhoneNumber     *string
	Email           string
	Gender          *string
	Password        string
	UserName        string
	ProfileImageUrl *string
	Type            string
	CreatedAt       time.Time
}

type UserStorageI interface {
	Create(u *User) (*User, error)
	Get(user_id int64) (*User, error)
	Update(u *User) (*User, error)
	Delete(user_id int64) error
}