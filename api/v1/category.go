package v1

import (
	"net/http"
	"strconv"

	"github.com/SaidovZohid/blog_db/api/models"
	"github.com/SaidovZohid/blog_db/storage/repo"
	"github.com/gin-gonic/gin"
)

// @Router /categories [post]
// @Summary Create a category
// @Description Create a category
// @Tags category
// @Accept json
// @Produce json
// @Param category body models.CreateCategoryRequest true "Category"
// @Success 201 {object} models.Category
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) CreateCategory(ctx *gin.Context) {
	var (
		req models.CreateCategoryRequest
	)

	err := ctx.ShouldBindJSON(&req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: err.Error(),
		})
		return
	}

	category, err := h.Storage.Category().Create(&repo.Category{
		Title: req.Title,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, repo.Category{
		ID: category.ID,
		Title: category.Title,
		CreatedAt: category.CreatedAt,
	})
}

// @Router /categories/{id} [get]
// @Summary Get category by it's id
// @Description Get category by it's id
// @Tags category
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 201 {object} models.Category
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) GetCategory(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: err.Error(),
		})
		return
	}

	category, err := h.Storage.Category().Get(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, repo.Category{
		ID: category.ID,
		Title: category.Title,
		CreatedAt: category.CreatedAt,
	})
}

// @Router /categories/update/{id} [put]
// @Summary Update category by it's id
// @Description Update category by it's id
// @Tags category
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Param category body models.CreateCategoryRequest true "Category"
// @Success 201 {object} models.Category
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) UpdateCategory(ctx *gin.Context) {
	var (
		req models.CreateCategoryRequest
		err error
	)

	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: err.Error(),
		})
		return
	}

	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: err.Error(),
		})
		return
	}

	category, err := h.Storage.Category().Update(&repo.Category{
		ID: id,
		Title: req.Title,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Category{
		ID: category.ID,
		Title: category.Title,
		CreatedAt: category.CreatedAt,
	})
}

// @Router /categories/delete/{id} [delete]
// @Summary Update category by it's id
// @Description Update category by it's id
// @Tags category
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 201 {object} models.ResponseSuccess
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) DeleteCategory(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: err.Error(),
		})
		return
	}

	err = h.Storage.Category().Delete(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.ResponseSuccess{
		Success: "Succesfully deleted!",
	})
}