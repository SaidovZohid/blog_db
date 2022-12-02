package postgres_test

import (
	"testing"

	"github.com/SaidovZohid/blog_db/storage/repo"
	"github.com/bxcodec/faker/v4"
	"github.com/stretchr/testify/require"
)

func createPost(t *testing.T) *repo.Post {
	user := createUser(t)
	catefory := createCategory(t)
	post, err := dbManager.Post().Create(&repo.Post{
		Title:       "Facebook",
		Description: "Facebook is stopped working on Meta Project",
		UserID:      user.ID,
		CategoryID:  catefory.ID,
	})
	require.NoError(t, err)
	require.NotEmpty(t, post)
	deleteUser(t, user.ID)
	deleteCategory(t, catefory.ID)
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
	user := createUser(t)
	catefory := createCategory(t)
	p, err := dbManager.Post().Update(&repo.Post{
		ID:          post.ID,
		Title:       faker.Sentence(),
		Description: faker.Sentence(),
		UserID:      user.ID,
		CategoryID:  catefory.ID,
	})
	deletePost(t, p.ID)
	deleteUser(t, user.ID)
	deleteCategory(t, catefory.ID)
	require.NoError(t, err)
	require.NotEmpty(t, p)
}

func TestDeletePost(t *testing.T) {
	post := createPost(t)
	err := dbManager.Post().Delete(post.ID)
	require.NoError(t, err)
	require.NotEmpty(t, post)
}

func TestGetPost(t *testing.T) {
	post := createPost(t)
	require.NotEmpty(t, post)
	p, err := dbManager.Post().Get(post.ID)
	deletePost(t, post.ID)
	require.NoError(t, err)
	require.NotEmpty(t, p)
}

func TestGetAllPosts(t *testing.T) {
	post := createPost(t)
	require.NotEmpty(t, post)
	posts, err := dbManager.Post().GetAll(&repo.GetPostsParams{
		Limit:      10,
		Page:       1,
		SortByDate: "ASC",
	})
	require.GreaterOrEqual(t, len(posts.Posts), 1)
	require.NoError(t, err)
	deletePost(t, post.ID)
}
