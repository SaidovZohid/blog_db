package v1

import (
	"net/http"

	"github.com/SaidovZohid/blog_db/api/models"
	"github.com/SaidovZohid/blog_db/storage/repo"
	"github.com/gin-gonic/gin"
)

// @Security ApiKeyAuth
// @Router /users [post]
// @Summary Create a user
// @Description Create a user
// @Tags user
// @Accept json
// @Produce json
// @Param user body models.CreateUserRequest true "User"
// @Success 201 {object} models.User
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) CreateUser(c *gin.Context) {
	var (
		req models.CreateUserRequest	
	)
	err := c.ShouldBindJSON(&req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: err.Error(),
		})
		return
	}
	resp, err := h.Storage.User().Create(&repo.User{
		FirstName: req.FirstName,
		LastName: req.LastName,
		PhoneNumber: req.PhoneNumber,
		Email: req.Email,
		Gender: req.Gender,
		UserName: req.UserName,
		ProfileImageUrl: req.ProfileImageUrl,
		Type: req.Type,
		Password: req.Password,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.User{
		ID: resp.ID,
		FirstName: resp.FirstName,
		LastName: resp.LastName,
		PhoneNumber: resp.PhoneNumber,
		Email: resp.Email,
		Gender: resp.Gender,
		UserName: resp.UserName, 
		ProfileImageUrl: resp.ProfileImageUrl,
		Type: resp.Type,
		CreatedAt: resp.CreatedAt,
	})
}

// @Security ApiKeyAuth
// @Router /users/user-info [get]
// @Summary Get user
// @Description Get user
// @Tags user
// @Accept json
// @Produce json
// @Success 201 {object} models.User
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) GetUser(c *gin.Context) {
	payload, err := h.GetAuthPayload(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	resp, err := h.Storage.User().Get(payload.UserID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.User{
		ID: resp.ID,
		FirstName: resp.FirstName,
		LastName: resp.LastName,
		PhoneNumber: resp.PhoneNumber,
		Email: resp.Email,
		Gender: resp.Gender,
		UserName: resp.UserName, 
		ProfileImageUrl: resp.ProfileImageUrl,
		Type: resp.Type,
		CreatedAt: resp.CreatedAt,
	})
}

// @Security ApiKeyAuth
// @Router /users/update [put]
// @Summary Update user by taking user id from token
// @Description Update user by taking user id from token
// @Tags user
// @Accept json
// @Produce json
// @Param User body models.CreateUserRequest true "User"
// @Success 201 {object} models.User
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) UpdateUser(c *gin.Context) {
	var (
		req models.CreateUserRequest
	)
	payload, err := h.GetAuthPayload(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	err = c.ShouldBindJSON(&req)

	if err != nil {
		c.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	result, err := h.Storage.User().Update(&repo.User{
		ID: payload.UserID,
		FirstName: req.FirstName,
		LastName: req.LastName,
		PhoneNumber: req.PhoneNumber,
		Email: req.Email,
		Gender: req.Gender,
		UserName: req.UserName, 
		ProfileImageUrl: req.ProfileImageUrl,
		Type: req.Type,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError,errResponse(err))
		return
	}

	c.JSON(http.StatusOK, models.User{
		ID: result.ID,
		FirstName: result.FirstName,
		LastName: result.LastName,
		PhoneNumber: result.PhoneNumber,
		Email: result.Email,
		Gender: result.Gender,
		UserName: result.UserName, 
		ProfileImageUrl: result.ProfileImageUrl,
		Type: result.Type,
	})
}

// @Security ApiKeyAuth
// @Router /users/delete [delete]
// @Summary Delete user by taking user_id from token
// @Description Delete user by taking user_id from token
// @Tags user
// @Accept json
// @Produce json
// @Success 201 {object} models.ResponseSuccess
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) DeleteUser(c *gin.Context) {
	payload, err := h.GetAuthPayload(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	err = h.Storage.User().Delete(payload.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.ResponseSuccess{
		Success: "Successfully deleted!",
	})
}

// @Router /users [get]
// @Summary Get user by giving limit, page and search for something. 
// @Description Get user by giving limit, page and search for something. 
// @Tags user
// @Accept json
// @Produce json
// @Param filter query models.GetAllParams false "Filter"
// @Success 201 {object} models.GetAllUsersResponse
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) GetAllUsers(c *gin.Context) {
	params, err := validateGetAllParams(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	result, err := h.Storage.User().GetAll(&repo.GetAllUserParams{
		Limit: int32(params.Limit),
		Page: int32(params.Page),
		Search: params.Search,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	c.JSON(http.StatusOK, getUsersResponse(result))
}

func getUsersResponse(data *repo.GetAllUsersResult) *models.GetAllUsersResponse {
	response := models.GetAllUsersResponse{
		Users: make([]*models.User, 0),
		Count: data.Count,
	}

	for _, user := range data.Users {
		u := parseUserModel(user)
		response.Users = append(response.Users, &u)
	}

	return &response
}


func parseUserModel(user *repo.User) models.User {
	return models.User{
		ID:              user.ID,
		FirstName:       user.FirstName,
		LastName:        user.LastName,
		PhoneNumber:     user.PhoneNumber,
		Email:           user.Email,
		Gender:          user.Gender,
		UserName:        user.UserName,
		ProfileImageUrl: user.ProfileImageUrl,
		Type:            user.Type,
		CreatedAt:       user.CreatedAt,
	}
}
