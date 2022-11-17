package v1

import (
	"log"
	"net/http"
	"time"

	"github.com/SaidovZohid/blog_db/api/models"
	"github.com/SaidovZohid/blog_db/pkg/utils"
	"github.com/SaidovZohid/blog_db/storage/repo"
	"github.com/gin-gonic/gin"
)

// @Router /auth/register [post]
// @Summary Create user with token key and get token key.
// @Description Create user with token key and get token key.
// @Tags register
// @Accept json
// @Produce json
// @Param data body models.RegisterRequest true "Data"
// @Success 200 {object} models.RegisterResponse
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) Register(ctx *gin.Context) {
	var (
		req models.RegisterRequest
	)

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errRespone(err))
		return
	}

	hashedPassword, err := utils.HashPassword(req.Password)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, errRespone(err))
		return
	}

	result, err := h.Storage.User().Create(&repo.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		UserName:  req.UserName,
		Type:      repo.UserTypeUser,
		Password:  hashedPassword,
	})

	if err != nil {
		ctx.JSON(http.StatusBadRequest, errRespone(err))
		return
	}

	token, _, err := utils.CreateToken(result.UserName, result.Email, time.Hour*24)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errRespone(err))
		return
	}
	log.Print("Hello")
	ctx.JSON(http.StatusOK, models.RegisterResponse{
		Id:          result.ID,
		FirstName:   result.FirstName,
		LastName:    result.LastName,
		Email:       result.Email,
		UserName:    result.UserName,
		Type:        result.Type,
		CreatedAt:   result.CreatedAt,
		AccessToken: token,
	})
}
