package repo

type EmailVer struct {
	ID       int64
	UserName string
	Email    string
	Code     int
}

type EmailVerI interface {
	CreateEmailVer(email_ver *EmailVer) error
	GetEmailVer(email string) (*EmailVer, error)
	DeleteEmailVer(email string) error
}


