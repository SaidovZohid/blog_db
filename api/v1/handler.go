package v1

import (
	"strconv"

	"github.com/SaidovZohid/blog_db/api/models"
	"github.com/SaidovZohid/blog_db/config"
	"github.com/SaidovZohid/blog_db/storage"
	"github.com/gin-gonic/gin"
)

type handlerV1 struct {
	cfg *config.Config
	Storage storage.StorageI
}

type HandlerV1Options struct {
	Cfg *config.Config
	Storage *storage.StorageI
}

func New(options *HandlerV1Options) *handlerV1 {
	return &handlerV1{
		cfg: options.Cfg,
		Storage: *options.Storage,
	}
}

func errResponse(err error) *models.ResponseError {
	return &models.ResponseError{
		Error: err.Error(),
	}
}

func validateGetAllParams(ctx *gin.Context) (*models.GetAllParams, error) {
	var (
		limit int64 = 10
		page int64 = 1
		err error
	)
	if ctx.Query("limit") != "" {
		limit, err = strconv.ParseInt(ctx.Query("limit"), 10, 64)
		if err != nil {
			return nil, err
		}
	}
	
	if ctx.Query("page") != "" {
		page, err = strconv.ParseInt(ctx.Query("page"), 10, 64)
		if err != nil {
			return nil, err
		}
	}

	return &models.GetAllParams{
		Limit: limit,
		Page: page,
		Search: ctx.Query("search"),
	}, nil
}

func validateGetAllPostsParams(ctx *gin.Context) (*models.GetAllPostsParams, error) {
	var (
		limit int64 = 10
		page int64 = 1
		err error
		userId, categoryId int64
		sortByDate string 
	)
	if ctx.Query("limit") != "" {
		limit, err = strconv.ParseInt(ctx.Query("limit"), 10, 64)
		if err != nil {
			return nil, err
		}
	}
	
	if ctx.Query("page") != "" {
		page, err = strconv.ParseInt(ctx.Query("page"), 10, 64)
		if err != nil {
			return nil, err
		}
	}

	if ctx.Query("user_id") != "" {
		userId, err = strconv.ParseInt(ctx.Query("user_id"), 10, 64)
		if err != nil {
			return nil, err
		}
	}
	
	if ctx.Query("category_id") != "" {
		categoryId, err = strconv.ParseInt(ctx.Query("category_id"), 10, 64)
		if err != nil {
			return nil, err
		}
	}
	if ctx.Query("sort") != "" && 
	(ctx.Query("sort") == "desc" || ctx.Query("sort") == "asc"){
		sortByDate = ctx.Query("sort")
	}

	return &models.GetAllPostsParams{
		Limit: limit,
		Page: page,
		Search: ctx.Query("search"),
		UserID: userId,
		CategoryID: categoryId,
		SortByDate: sortByDate,
	}, nil
}