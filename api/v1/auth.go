package v1

import (
	"math/rand"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/SaidovZohid/blog_db/api/models"
	"github.com/SaidovZohid/blog_db/pkg/email"
	"github.com/SaidovZohid/blog_db/pkg/utils"
	"github.com/SaidovZohid/blog_db/storage/repo"
	"github.com/gin-gonic/gin"
)

var (
	ErrWrongEmailOrPassword = errors.New("wrong email or password")
	ErrUserNotVerifid       = errors.New("user not verified")
)

// @Router /auth/register [post]
// @Summary Create user with token key and get token key.
// @Description Create user with token key and get token key.
// @Tags register
// @Accept json
// @Produce json
// @Param data body models.RegisterRequest true "Data"
// @Success 200 {object} models.AuthResponse
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) Register(ctx *gin.Context) {
	var (
		req models.RegisterRequest
	)

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	hashedPassword, err := utils.HashPassword(req.Password)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	result, err := h.Storage.User().Create(&repo.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		UserName:  req.UserName,
		Type:      repo.UserTypeUser,
		Password:  hashedPassword,
		IsActive:  false,
	})

	if err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}
	
	rand.Seed(time.Now().UnixNano())
	randNum := rand.Intn((10000 - 9999 + 1) + 9999)

	go func() {
		err = email.SendEmail(h.cfg, &email.SendEmailRequest{
			To:      []string{result.Email},
			Subject: "Verification Email",
			Body: map[string]int{
				"code": randNum,
			},
			Type: email.VerificationEmail,
		})
		if err != nil {
			fmt.Println("Failed to sent code to email")
		}
	}()

	ctx.JSON(http.StatusOK, models.ResponseSuccess{
		Success: "Verification code has been sent!",
	})
}

// @Router /auth/verify [post]
// @Summary Create user with token key and get token key.
// @Description Create user with token key and get token key.
// @Tags register
// @Accept json
// @Produce json
// @Param data body models.VerifyRequest true "Data"
// @Success 200 {object} models.AuthResponse
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) Verify(ctx *gin.Context) {
	var (
		req models.VerifyRequest
	)

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	user, err := h.Storage.User().GetByEmail(req.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, errResponse(err))
		}
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	// TODO: check verification code

	err = h.Storage.User().Activate(user.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	token, _, err := utils.CreateToken(h.cfg, &utils.TokenParams{
		UserID:   user.ID,
		Username: user.UserName,
		Email:    user.Email,
		Duration: time.Hour * 24 * 360,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, models.AuthResponse{
		Id:          user.ID,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		UserName:    user.UserName,
		Type:        user.Type,
		CreatedAt:   user.CreatedAt,
		AccessToken: token,
	})

}

// @Router /auth/login [post]
// @Summary Login User
// @Description Login User
// @Tags register
// @Accept json
// @Produce json
// @Param login body models.LoginRequest true "Login"
// @Success 200 {object} models.AuthResponse
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) Login(ctx *gin.Context) {
	var (
		req models.LoginRequest
	)

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	user, err := h.Storage.User().GetByEmail(req.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusForbidden, errResponse(ErrWrongEmailOrPassword))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	err = utils.CheckPassword(req.Password, user.Password)
	if err != nil {
		ctx.JSON(http.StatusForbidden, errResponse(ErrWrongEmailOrPassword))
		return
	}

	token, _, err := utils.CreateToken(h.cfg, &utils.TokenParams{
		UserID:   user.ID,
		Username: user.UserName,
		Email:    user.Email,
		Duration: time.Hour * 24 * 360,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, models.AuthResponse{
		Id:          user.ID,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		UserName:    user.UserName,
		Type:        user.Type,
		CreatedAt:   user.CreatedAt,
		AccessToken: token,
	})

}
