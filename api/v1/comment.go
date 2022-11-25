package v1

import (
	"net/http"
	"strconv"

	"github.com/SaidovZohid/blog_db/api/models"
	"github.com/SaidovZohid/blog_db/storage/repo"
	"github.com/gin-gonic/gin"
)

// @Security ApiKeyAuth
// @Router /comments [post]
// @Summary Create a comment
// @Description Create a comment
// @Tags comment
// @Accept json
// @Produce json
// @Param post body models.CreateCommentRequest true "Post"
// @Success 201 {object} models.Comment
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) CreateComment(ctx *gin.Context) {
	var (
		req models.CreateCommentRequest
	)

	err := ctx.ShouldBindJSON(&req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: err.Error(),
		})
		return
	}

	payload, err := h.GetAuthPayload(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	comment, err := h.Storage.Comment().Create(&repo.Comment{
		Description: req.Description,
		UserID:      payload.UserID,
		PostID:      req.PostID,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}
	c := parseCommentModel(comment)
	ctx.JSON(http.StatusOK, c)
}

// @Security ApiKeyAuth
// @Router /comments/update/{id} [put]
// @Summary Update comment with it's id as param
// @Description Update comment with it's id as param
// @Tags comment
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Param comment body models.UpdateCommentRequest true "Comment"
// @Success 201 {object} models.UpdateComment
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) UpdateComment(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	var (
		req models.UpdateCommentRequest
	)

	if err = ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	comment, err := h.Storage.Comment().Update(&repo.UpdateComment{
		ID:          id,
		Description: req.Description,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, models.UpdateComment{
		ID:          comment.ID,
		Description: comment.Description,
		UserID:      comment.UserID,
		PostID:      comment.PostID,
		CreatedAt:   comment.CreatedAt,
		UpdatedAt:   comment.UpdatedAt,
	})
}

// @Security ApiKeyAuth
// @Router /comments/delete/{id} [delete]
// @Summary Delete a comment
// @Description Delete a comment
// @Tags comment
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 201 {object} models.ResponseSuccess
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) DeleteComment(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	err = h.Storage.Comment().Delete(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, models.ResponseSuccess{
		Success: "Successfully deleted!",
	})
}

// @Router /comments [get]
// @Summary Get comments by giving limit, page and user_id, post_id.
// @Description Get comments by giving limit, page and user_id, post_id.
// @Tags comment
// @Accept json
// @Produce json
// @Param filter query models.GetAllCommentsParams false "Filter"
// @Success 201 {object} models.GetAllCommentsResponse
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) GetAllComments(c *gin.Context) {
	params, err := validateGetAllCommentsParams(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	result, err := h.Storage.Comment().GetAll(&repo.GetCommentsParams{
		Limit:  params.Limit,
		Page:   params.Page,
		UserID: params.UserID,
		PostID: params.PostID,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	c.JSON(http.StatusOK, getCommentsResponse(result))
}

func getCommentsResponse(data *repo.GetAllCommentsResult) *models.GetAllCommentsResponse {
	response := models.GetAllCommentsResponse{
		Comments: make([]*models.Comment, 0),
		Count:    data.Count,
	}

	for _, comment := range data.Comments {
		u := parseCommentModel(comment)
		response.Comments = append(response.Comments, &u)
	}

	return &response
}

func parseCommentModel(comment *repo.Comment) models.Comment {
	return models.Comment{
		ID:          comment.ID,
		Description: comment.Description,
		UserID:      comment.UserID,
		PostID:      comment.PostID,
		CreatedAt:   comment.CreatedAt,
		UpdatedAt:   comment.UpdatedAt,
		User: &models.CommentUser{
			FirstName: comment.User.FirstName,
			Lastname: comment.User.LastName,
			Email: comment.User.Email,
			ProfileImageUrl: comment.User.ProfileImageUrl,
		},
	}
}
