package v1

import (
	"net/http"
	"strconv"

	"github.com/SaidovZohid/blog_db/api/models"
	"github.com/SaidovZohid/blog_db/storage/repo"
	"github.com/gin-gonic/gin"
)

// @Router /posts [post]
// @Summary Create a post
// @Description Create a post
// @Tags post
// @Accept json
// @Produce json
// @Param post body models.CreatePostRequest true "Post"
// @Success 201 {object} models.Post
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) CreatePost(ctx *gin.Context) {
	var (
		req models.CreatePostRequest
	)

	err := ctx.ShouldBindJSON(&req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: err.Error(),
		})
		return
	}

	post, err := h.Storage.Post().Create(&repo.Post{
		Title: req.Title,
		Description: req.Description,
		ImageUrl: req.ImageUrl,
		UserID: req.UserID,
		CategoryID: req.CategoryID,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Post{
		ID: post.ID,
		Title: post.Title,
		Description: post.Description,
		ImageUrl: post.ImageUrl,
		UserID: post.UserID,
		CategoryID: post.CategoryID,
		CreatedAt: post.CreatedAt,
	})
}

// @Router /posts/{id} [get]
// @Summary Get a post with it's id
// @Description Create a post with it's id
// @Tags post
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 201 {object} models.Category
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) GetPost(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: err.Error(),
		})
		return
	}

	post, err := h.Storage.Post().Get(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Post{
		ID: post.ID,
		Title: post.Title,
		Description: post.Description,
		ImageUrl: post.ImageUrl,
		UserID: post.UserID,
		CategoryID: post.CategoryID,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
		ViewsCount: post.ViewsCount,
	})
}

// @Router /posts/update/{id} [put]
// @Summary Update post with it's id as param
// @Description Update post with it's id as param
// @Tags post
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Param post body models.UpdatePostRequest true "Post"
// @Success 201 {object} models.Category
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) UpdatePost(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: err.Error(),
		})
		return
	}

	var (
		req models.UpdatePostRequest
	)

	err = ctx.ShouldBindJSON(&req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: err.Error(),
		})
		return
	}	

	post, err := h.Storage.Post().Update(&repo.Post{
		ID: id,
		Title: req.Title,
		Description: req.Description,
		ImageUrl: req.ImageUrl,
		UserID: req.UserID,
		CategoryID: req.CategoryID,
		ViewsCount: req.ViewsCount,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Post{
		ID: post.ID,
		Title: post.Title,
		Description: post.Description,
		ImageUrl: post.ImageUrl,
		UserID: post.UserID,
		CategoryID: post.CategoryID,
		ViewsCount: post.ViewsCount,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
	})
}

// @Router /posts/delete/{id} [delete]
// @Summary Delete a post
// @Description Create a post
// @Tags post
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 201 {object} models.ResponseSuccess
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) DeletePost(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: err.Error(),
		})
		return
	}

	err = h.Storage.Post().Delete(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: err.Error(),
		})
		return 
	}

	ctx.JSON(http.StatusOK, models.ResponseSuccess{
		Success: "Successfully deleted!",
	})
}