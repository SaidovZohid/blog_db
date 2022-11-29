package postgres_test

import (
	"testing"

	"github.com/SaidovZohid/blog_db/pkg/utils"
	"github.com/SaidovZohid/blog_db/storage/repo"
	"github.com/stretchr/testify/require"
)

func createUser(t *testing.T) *repo.User {
	hashedPassword, err := utils.HashPassword("1234567890")
	require.NoError(t, err)
	user, err := dbManager.User().Create(&repo.User{
		FirstName: "Zohid",
		LastName:  "Saidov",
		Email:     "zohidsaidov17+1@gmail.com",
		Password:  hashedPassword,
		Type:      "user",
	})
	require.NoError(t, err)
	require.NotEmpty(t, user)
	return user
}

func deleteUser(t *testing.T, user_id int64) {
	err := dbManager.User().Delete(user_id)
	require.NoError(t, err)
}

func TestCreateUser(t *testing.T) {
	user := createUser(t)
	deleteUser(t, user.ID)
}

func TestUpdateUser(t *testing.T) {
	user := createUser(t)
	u, err := dbManager.User().Update(&repo.User{
		ID:        user.ID,
		FirstName: "Zufar",
		LastName:  "Saidov",
		Email:     "zufarsaidov22@gmail.com",
		Type:      "user",
	})
	deleteUser(t, u.ID)
	require.NoError(t, err)
	require.NotEmpty(t, u)
}

func TestDeleteUser(t *testing.T) {
	user := createUser(t)
	deleteUser(t, user.ID)
}

func TestGetUser(t *testing.T) {
	user := createUser(t)
	u, err := dbManager.User().Get(user.ID)
	deleteUser(t, u.ID)
	require.NoError(t, err)
	require.NotEmpty(t, u)
}

func TestGetAll(t *testing.T) {
	user := createUser(t)

	users, err := dbManager.User().GetAll(&repo.GetAllUserParams{
		Limit: 10,
		Page:  1,
	})
	deleteUser(t, user.ID)
	require.GreaterOrEqual(t, len(users.Users), 1)
	require.NoError(t, err)
	require.NotEmpty(t, user)
}
