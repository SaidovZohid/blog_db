package postgres_test

// import (
// 	"math/rand"
	
// 	"testing"
// 	"time"

// 	"github.com/SaidovZohid/blog_db/storage/repo"
// 	"github.com/bxcodec/faker/v4"
// 	"github.com/stretchr/testify/require"
// )

// func getRandomBoolValue() bool {
// 	s := time.Now().UnixMilli()
// 	if s % 2 == 0 {
// 		return true
// 	}
// 	return false
// }

// func createLike(t *testing.T) *repo.Like {
// 	post := createPost(t)
// 	user := createUser(t)
// 	like, err := dbManager.Like().CreateOrUpdate(&repo.Like{
// 		PostID: post.ID,
// 		UserID: user.ID,
// 		Status: getRandomBoolValue(),
// 	})
// 	require.NoError(t, err)
// 	require.NotEmpty(t, like)
// 	deletePost(t, post.ID)
// 	deleteUser(t, user.ID)
// }
