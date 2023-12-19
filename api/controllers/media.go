package controllers

import (
	"strconv"

	"srating/bootstrap"
	"srating/domain"
	"srating/x/rest"

	"github.com/gin-gonic/gin"
)

type MediaController struct {
	MediaService domain.MediaService
	Env          *bootstrap.Env
	*rest.JSONRender
}

// Upload
// @Router /admin/media [post]
// @Tags media
// @Summary Upload media
// @Security ApiKeyAuth
// @Success 200 {object} string
func (t *MediaController) Upload(c *gin.Context) {
	input := []*domain.UploadFileInput{}
	for i := 0; ; i++ {
		index := strconv.Itoa(i)
		file, err := c.FormFile("file[" + index + "]")
		if err != nil {
			break
		}
		input = append(input, &domain.UploadFileInput{
			FileHeader: file,
		})
	}
	result, err := t.MediaService.Upload(c, input)
	rest.AssertNil(err)
	t.SendData(c, result)
}

// GetAll
// @Router /admin/media [get]
// @Tags media
// @Summary Get all media
// @Security ApiKeyAuth
// @Success 200 {object} string
func (t *MediaController) GetAll(c *gin.Context) {
	result, err := t.MediaService.GetAll(c)
	rest.AssertNil(err)
	t.SendData(c, result)
}
