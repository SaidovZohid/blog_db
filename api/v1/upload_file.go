package v1

import (
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/SaidovZohid/blog_db/api/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type File struct {
	File *multipart.FileHeader `form:"file" binding:"required"`
}

// @Security ApiKeyAuth
// @Router /file_upload [post]
// @Summary File upload
// @Description File upload
// @Tags file-upload
// @Accept json
// @Produce json
// @Param file formData file true "File"
// @Success 200 {object} models.ResponseSuccess
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) UploadFile(ctx *gin.Context) {
	var file File

	if err := ctx.ShouldBind(&file); err != nil {
		log.Print("HEllo World 1")
		ctx.JSON(http.StatusBadRequest, models.ResponseError{
			Error: err.Error(),
		})
		return
	}
	id := uuid.New()
	log.Print(filepath.Ext(file.File.Filename))
	fileName := id.String() + filepath.Ext(file.File.Filename)
	dir, _ := os.Getwd()

	if _, err := os.Stat(dir + "/media"); os.IsNotExist(err) {
		log.Print("HEllo World 2")
		os.Mkdir(dir + "/media", os.ModePerm)
	}

	filePath := "/media/" + fileName
	err := ctx.SaveUploadedFile(file.File, dir + filePath)
	if err != nil {
		log.Print("HEllo World 3")
		ctx.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: err.Error(),
		})
		return
	}
	log.Print("HEllo World 4")
	ctx.JSON(http.StatusCreated, models.ResponseSuccess{
		Success: filePath,
	})
}
