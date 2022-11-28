package v1

import (
	"net/http"
	"strconv"

	"github.com/SaidovZohid/blog_db/api/models"
	"github.com/SaidovZohid/blog_db/storage/repo"
	"github.com/gin-gonic/gin"
)

// @Security ApiKeyAuth
// @Router /likes [post]
// @Summary Create Or Update like
// @Description Create Or Update like
// @Tags like
// @Accept json
// @Produce json
// @Param like body models.CreateOrUpdateLikeRequest true "like"
// @Success 201 {object} models.Like
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) CreateOrUpdateLike(ctx *gin.Context) {
	var (
		req models.CreateOrUpdateLikeRequest
	)

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	payload, err := h.GetAuthPayload(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	resp, err := h.Storage.Like().CreateOrUpdate(&repo.Like{
		UserID: payload.UserID,
		PostID: req.PostID,
		Status: req.Status,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, models.Like{
		ID:     resp.ID,
		PostID: resp.PostID,
		UserID: resp.UserID,
		Status: resp.Status,
	})
}

// @Security ApiKeyAuth
// @Router /likes/user-post [get]
// @Summary Get like by giving to query post_id
// @Description Get like by giving to query post_id
// @Tags like
// @Accept json
// @Produce json
// @Param post_id query int true "post_id"
// @Success 201 {object} models.Like
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) GetLike(ctx *gin.Context) {
	postID, err := strconv.Atoi(ctx.Query("post_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	payload, err := h.GetAuthPayload(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	resp, err := h.Storage.Like().Get(payload.UserID, int64(postID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, models.Like{
		ID:     resp.ID,
		PostID: resp.PostID,
		UserID: resp.UserID,
		Status: resp.Status,
	})
}
