package postgres_test

import (
	"testing"

	"github.com/SaidovZohid/blog_db/storage/repo"
	"github.com/stretchr/testify/require"
)

func createPost(t *testing.T) (*repo.Post) {
	post, err := dbManager.Post().Create(&repo.Post{
		Title: "Facebook",
		Description: "Facebook is stopped working on Meta Project",
		UserID: 20,
		CategoryID: 1,
	})
	require.NoError(t, err)
	require.NotEmpty(t, post)
	return post
}

func deletePost(t *testing.T, post_id int64) {
	err := dbManager.Post().Delete(post_id)
	require.NoError(t, err)
} 

func TestCreatePost(t *testing.T) {
	post := createPost(t)
	deletePost(t, post.ID)
	require.NotEmpty(t, post)
}

func TestUpdatePost(t *testing.T) {
	post := createPost(t)
	p, err := dbManager.Post().Update(&repo.Post{
		ID: post.ID,
		Title: "Instagramm",
		Description: "Is a big Company in The World",
		UserID: 20,
		CategoryID: 1,
	})
	deletePost(t, p.ID)
	require.NoError(t, err)
	require.NotEmpty(t, p)
}

func TestDeletePost(t *testing.T) {
	post := createPost(t)
	err := dbManager.Post().Delete(post.ID)
	require.NoError(t, err)
	require.NotEmpty(t, post)
}