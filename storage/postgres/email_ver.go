package postgres

import (
	"github.com/SaidovZohid/blog_db/storage/repo"
	"github.com/jmoiron/sqlx"
)

type emailVerRepo struct {
	db *sqlx.DB
}

func NewEmailVer(db *sqlx.DB) repo.EmailVerI {
	return &emailVerRepo{
		db: db,
	}
}

func (er *emailVerRepo) CreateEmailVer(email_ver *repo.EmailVer) error {
	query := `
		INSERT INTO email_ver (
			username,
			email,
			ver_code
		) VALUES ($1, $2, $3)
	`
	_, err := er.db.Exec(query, email_ver.UserName, email_ver.Email, email_ver.Code)
	if err != nil {
		return err
	}		

	return nil
}

func (er *emailVerRepo) GetEmailVer(email string) (*repo.EmailVer, error) {
	var email_ver repo.EmailVer
	query := `
		SELECT
			ver_code
		FROM email_ver WHERE email=$1
	`

	err := er.db.QueryRow(
		query, 
		email,
	).Scan(
		&email_ver.Code,
	)
	if err != nil {
		return nil, err
	}		

	return &email_ver, nil
}

func (er *emailVerRepo) DeleteEmailVer(email string) error {
	query := "DELETE FROM email_ver WHERE email=$1"
	_, err := er.db.Exec(query, email)
	if err  != nil {
		return err
	}

	return nil
}