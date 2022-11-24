package v1

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/SaidovZohid/blog_db/api/models"
	emailPkg "github.com/SaidovZohid/blog_db/pkg/email"
	"github.com/SaidovZohid/blog_db/pkg/utils"
	"github.com/SaidovZohid/blog_db/storage/repo"
	"github.com/gin-gonic/gin"
)

var (
	ErrWrongEmailOrPassword = errors.New("wrong email or password")
	ErrUserNotVerifid       = errors.New("user not verified")
	ErrEmailExists          = errors.New("email is already exists")
	ErrIncorrectCode        = errors.New("incorrect verification code")
	ErrCodeExpired          = errors.New("verification is expired")
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

	_, err = h.Storage.User().GetByEmail(req.Email)
	if !errors.Is(err, sql.ErrNoRows) {
		ctx.JSON(http.StatusBadRequest, errResponse(ErrEmailExists))
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	user := repo.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Type:      repo.UserTypeUser,
		Password:  hashedPassword,
	}

	userData, err := json.Marshal(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	err = h.inMemory.Set("user_" + req.Email, string(userData), 10 * time.Minute)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
	}
	
	go func(){
		err = h.sendVereficationCode(req.Email)
		if err != nil {
			fmt.Println("failed to send code")
			fmt.Printf("failed to send verification code: %v", err)
		}
	}()

	ctx.JSON(http.StatusCreated, models.ResponseSuccess{
		Success: "Verification code has been sent!",
	})
}

func (h *handlerV1) sendVereficationCode(email string) error {
	code, err := utils.GenerateRandomCode(6)
	if err != nil {
		return err
	}

	err = h.inMemory.Set("code_" + email, code, time.Minute) 
	if err != nil {
		return err
	}

	err = emailPkg.SendEmail(h.cfg, &emailPkg.SendEmailRequest{
		To:      []string{email},
		Subject: "Verification Email",
		Body: map[string]string{
			"code": code,
		},
		Type: emailPkg.VerificationEmail,
	})
	if err != nil {
		fmt.Println("Failed to sent code to email")
	}
	
	return nil
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

	userData, err := h.inMemory.Get("user_" + req.Email)
	if err != nil {
		ctx.JSON(http.StatusNotFound, errResponse(err))
		return
	}
	// TODO: check verification code
	var user repo.User
	err = json.Unmarshal([]byte(userData), &user)
	if err != nil {
		ctx.JSON(http.StatusNotFound, errResponse(err))
		return
	}

	code, err := h.inMemory.Get("code_" + user.Email)
	if err != nil {
		ctx.JSON(http.StatusForbidden, errResponse(ErrCodeExpired))
		return
	}
	if code != req.Code {
		ctx.JSON(http.StatusForbidden, errResponse(ErrIncorrectCode))
		return
	}

	result, err := h.Storage.User().Create(&user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	token, _, err := utils.CreateToken(h.cfg, &utils.TokenParams{
		UserID:   user.ID,
		Email:    user.Email,
		Duration: time.Hour * 24 * 360,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, models.AuthResponse{
		Id:          result.ID,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		Type:        user.Type,
		CreatedAt:   result.CreatedAt,
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
		Type:        user.Type,
		CreatedAt:   user.CreatedAt,
		AccessToken: token,
	})

}
