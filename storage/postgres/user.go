package postgres

import (
	"database/sql"

	"github.com/SaidovZohid/blog_db/storage/repo"
	"github.com/jmoiron/sqlx"
)

type userRepo struct {
	db *sqlx.DB
}

func NewUser(db *sqlx.DB) repo.UserStorageI {
	return &userRepo{
		db: db,
	}
}

func (ur *userRepo) Create(user *repo.User) (*repo.User, error) {
	query := `
		INSERT INTO users (
			first_name,
			last_name,
			phone_number,
			email,
			gender,
			password,
			username,
			profile_image_url,
			type
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id, created_at
	`
	err := ur.db.QueryRow(
		query,
		user.FirstName,
		user.LastName,
		user.PhoneNumber,
		user.Email,
		user.Gender,
		user.Password,
		user.UserName,
		user.ProfileImageUrl,
		user.Type,
	).Scan(
		&user.ID,
		&user.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (ur *userRepo) Get(user_id int64) (*repo.User, error) {
	var result repo.User
	
	query := `
		SELECT 
			id,
			first_name,
			last_name,
			phone_number,
			email,
			gender,
			password,
			username,
			profile_image_url,
			type,
			created_at
		FROM users WHERE id = $1
	`
	err := ur.db.QueryRow(
		query,
		user_id,
	).Scan(
		&result.ID,
		&result.FirstName,
		&result.LastName,
		&result.PhoneNumber,
		&result.Email,
		&result.Gender,
		&result.Password,
		&result.UserName,
		&result.ProfileImageUrl,
		&result.Type,
		&result.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (ur *userRepo) Update(user *repo.User) (*repo.User, error) {
	var result repo.User
	
	query := `
		UPDATE users SET
			first_name=$1,
			last_name=$2,
			phone_number=$3,
			email=$4,
			gender=$5,
			password=$6,
			username=$7,
			profile_image_url=$8,
			type=$9
		WHERE id=$10 
		RETURNING 
			id,
			first_name,
			last_name,
			phone_number,
			email,
			gender,
			password,
			username,
			profile_image_url,
			type,
			created_at
	`
	err := ur.db.QueryRow(
		query,
		user.FirstName,
		user.LastName,
		user.PhoneNumber,
		user.Email,
		user.Gender,
		user.Password,
		user.UserName,
		user.ProfileImageUrl,
		user.Type,
		user.ID,
	).Scan(
		&result.ID,
		&result.FirstName,
		&result.LastName,
		&result.PhoneNumber,
		&result.Email,
		&result.Gender,
		&result.Password,
		&result.UserName,
		&result.ProfileImageUrl,
		&result.Type,
		&result.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (ur *userRepo) Delete(user_id int64) error {
	query := `
		DELETE FROM users WHERE id = $1
	`
	res, err := ur.db.Exec(
		query,
		user_id,
	)

	result, err := res.RowsAffected()
	
	if err != nil {
		return err
	}

	if result == 0 {
		return sql.ErrNoRows
	}

	if err != nil {
		return err
	}

	return nil
}