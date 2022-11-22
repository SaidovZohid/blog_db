package postgres

import (
	"database/sql"
	"fmt"

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
			type,
			is_active
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
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
		user.IsActive,
	).Scan(
		&user.ID,
		&user.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (ur *userRepo) Activate(user_id int64) error {
	query := "UPDATE users SET is_active=true WHERE id=$1"
	_, err := ur.db.Exec(query, user_id)
	if err != nil {
		return err
	}

	return nil 
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

func (ur *userRepo) GetAll(params *repo.GetAllUserParams) (*repo.GetAllUsersResult, error) {
	result := repo.GetAllUsersResult{
		Users: make([]*repo.User, 0),
	}

	offset := (params.Page - 1) * params.Limit

	limit := fmt.Sprintf(" LIMIT %d OFFSET %d", params.Limit, offset)

	filter := " WHERE is_active=true "

	if params.Search != "" {
		str := "%" + params.Search + "%"
		filter = fmt.Sprintf(`
			AND first_name ILIKE '%s' OR last_name ILIKE '%s' OR phone_number ILIKE '%s' OR email ILIKE '%s' OR username ILIKE '%s'
		`, str, str, str, str, str)
	}

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
		FROM users
	` + filter  + `
		ORDER BY created_at DESC
	` + limit

	rows, err := ur.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user repo.User
		err := rows.Scan(
			&user.ID,
			&user.FirstName,
			&user.LastName,
			&user.PhoneNumber,
			&user.Email,
			&user.Gender,
			&user.Password,
			&user.UserName,
			&user.ProfileImageUrl,
			&user.Type,
			&user.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		result.Users = append(result.Users, &user)
	}

	queryCount := "SELECT count(1) FROM users " + filter

	err = ur.db.QueryRow(queryCount).Scan(&result.Count)
	
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (ur *userRepo) GetByEmail(user_email string) (*repo.User, error) {
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
		FROM users WHERE email = $1
	`
	err := ur.db.QueryRow(
		query,
		user_email,
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