package postgres_test

import (
	"testing"

	"github.com/SaidovZohid/blog_db/storage/repo"
	"github.com/stretchr/testify/require"
)

func createCategory(t *testing.T) *repo.Category {
	c, err := dbManager.Category().Create(&repo.Category{
		Title: "Entertainment",
	})
	require.NoError(t, err)
	return c
}

func deleteCategory(t *testing.T, id int64) {
	err := dbManager.Category().Delete(id)
	require.NoError(t, err)
}

func TestCreateCategory(t *testing.T) {
	c := createCategory(t)
	require.NotEmpty(t, c)
	deleteCategory(t, c.ID)
}

func TestGetCategory(t *testing.T) {
	c := createCategory(t)
	require.NotEmpty(t, c)
	category, err := dbManager.Category().Get(c.ID)
	require.NoError(t, err)
	require.NotEmpty(t, category)
	deleteCategory(t, c.ID)
}

func TestUpdateCategory(t *testing.T) {
	c := createCategory(t)
	require.NotEmpty(t, c)
	ca, err := dbManager.Category().Update(&repo.Category{
		ID: c.ID,
		Title: "Hello",
	})
	require.NoError(t, err)
	require.NotEmpty(t, ca)
	deleteCategory(t, c.ID)
}

func TestDeleteCategory(t *testing.T) {
	c := createCategory(t)
	require.NotEmpty(t, c)
	deleteCategory(t, c.ID)
}

func TestGetAllCategory(t *testing.T) {
	c := createCategory(t)
	require.NotEmpty(t, c)
	cs, err := dbManager.Category().GetAll(&repo.GetAllCategoryParams{
		Limit: 10,
		Page: 1,
	})
	require.GreaterOrEqual(t, len(cs.Categories), 1)
	require.NoError(t, err)
}