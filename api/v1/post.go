package v1

import (
	"net/http"
	"strconv"

	"github.com/SaidovZohid/blog_db/api/models"
	"github.com/SaidovZohid/blog_db/storage/repo"
	"github.com/gin-gonic/gin"
)

// @Security ApiKeyAuth
// @Router /posts [post]
// @Summary Create a post
// @Description Create a post
// @Tags post
// @Accept json
// @Produce json
// @Param post body models.CreatePostRequest true "Post"
// @Success 201 {object} models.Post
// @Failure 500 {object} models.ResponseError
// @Failure 400 {object} models.ResponseError
func (h *handlerV1) CreatePost(ctx *gin.Context) {
	var (
		req models.CreatePostRequest
	)

	err := ctx.ShouldBindJSON(&req)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	post, err := h.Storage.Post().Create(&repo.Post{
		Title:       req.Title,
		Description: req.Description,
		ImageUrl:    req.ImageUrl,
		UserID:      req.UserID,
		CategoryID:  req.CategoryID,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, models.Post{
		ID:          post.ID,
		Title:       post.Title,
		Description: post.Description,
		ImageUrl:    post.ImageUrl,
		UserID:      post.UserID,
		CategoryID:  post.CategoryID,
		CreatedAt:   post.CreatedAt,
	})
}

// @Security ApiKeyAuth
// @Router /posts/{id} [get]
// @Summary Get a post with it's id
// @Description Create a post with it's id
// @Tags post
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 201 {object} models.Post
// @Failure 500 {object} models.ResponseError
// @Failure 400 {object} models.ResponseError
func (h *handlerV1) GetPost(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	res, err := h.Storage.Post().Get(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	post := parsePostModel(res)
	
	likesInfo, err := h.Storage.Like().GetLikesDislikesCount(post.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	post.PostLikeInfo = &models.PostLikeInfo{
		LikesCount: likesInfo.Likes,
		DislikesCount: likesInfo.Dislikes,
	}
	ctx.JSON(http.StatusOK, post)
}

// @Security ApiKeyAuth
// @Router /posts/{id} [put]
// @Summary Update post with it's id as param
// @Description Update post with it's id as param
// @Tags post
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Param post body models.UpdatePostRequest true "Post"
// @Success 201 {object} models.Post
// @Failure 500 {object} models.ResponseError
// @Failure 400 {object} models.ResponseError
func (h *handlerV1) UpdatePost(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	var (
		req models.UpdatePostRequest
	)

	err = ctx.ShouldBindJSON(&req)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	post, err := h.Storage.Post().Update(&repo.Post{
		ID:          id,
		Title:       req.Title,
		Description: req.Description,
		ImageUrl:    req.ImageUrl,
		UserID:      req.UserID,
		CategoryID:  req.CategoryID,
		ViewsCount:  req.ViewsCount,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, models.Post{
		ID:          post.ID,
		Title:       post.Title,
		Description: post.Description,
		ImageUrl:    post.ImageUrl,
		UserID:      post.UserID,
		CategoryID:  post.CategoryID,
		ViewsCount:  post.ViewsCount,
		CreatedAt:   post.CreatedAt,
		UpdatedAt:   post.UpdatedAt,
	})
}

// @Security ApiKeyAuth
// @Router /posts/{id} [delete]
// @Summary Delete a post
// @Description Create a post
// @Tags post
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 201 {object} models.ResponseSuccess
// @Failure 500 {object} models.ResponseError
// @Failure 400 {object} models.ResponseError
func (h *handlerV1) DeletePost(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	err = h.Storage.Post().Delete(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, models.ResponseSuccess{
		Success: "Successfully deleted!",
	})
}

// @Router /posts [get]
// @Summary Get posts by giving limit, page and search for something.
// @Description Get posts by giving limit, page and search for something.
// @Tags post
// @Accept json
// @Produce json
// @Param filter query models.GetAllPostsParams false "Filter"
// @Success 201 {object} models.GetAllPostsResponse
// @Failure 500 {object} models.ResponseError
// @Failure 400 {object} models.ResponseError
func (h *handlerV1) GetAllPosts(c *gin.Context) {
	params, err := validateGetAllPostsParams(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	result, err := h.Storage.Post().GetAll(&repo.GetPostsParams{
		Limit:      params.Limit,
		Page:       params.Page,
		Search:     params.Search,
		UserID:     params.UserID,
		CategoryID: params.CategoryID,
		SortByDate: params.SortByDate,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	c.JSON(http.StatusOK, getPostsResponse(result))
}

func getPostsResponse(data *repo.GetAllPostResult) *models.GetAllPostsResponse {
	response := models.GetAllPostsResponse{
		Posts: make([]*models.Post, 0),
		Count: data.Count,
	}

	for _, post := range data.Posts {
		u := parsePostModel(post)
		response.Posts = append(response.Posts, &u)
	}

	return &response
}

func parsePostModel(post *repo.Post) models.Post {
	return models.Post{
		ID:          post.ID,
		Title:       post.Title,
		Description: post.Title,
		ImageUrl:    post.ImageUrl,
		ViewsCount:  post.ViewsCount,
		UserID:      post.UserID,
		CategoryID:  post.CategoryID,
		CreatedAt:   post.CreatedAt,
		UpdatedAt:   post.UpdatedAt,
	}
}
