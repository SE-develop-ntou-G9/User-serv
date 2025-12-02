package http

import (
	"golangAPI/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ImageHandler struct {
	imageUsecase *usecase.ImageUsecase
}

func NewImageHandler(imageUsecase *usecase.ImageUsecase) *ImageHandler {
	return &ImageHandler{
		imageUsecase: imageUsecase,
	}
}

func RegisterImageRoutes(r gin.IRoutes, imageUsecase *usecase.ImageUsecase) {
	h := NewImageHandler(imageUsecase)
	r.POST("/images/avatar", h.UploadAvatar)
	r.POST("/images/license", h.UploadLicense)
}

func (h *ImageHandler) UploadAvatar(c *gin.Context) {

	fileHeader, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing file"})
		return
	}

	file, err := fileHeader.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to open file"})
		return
	}
	contentType := fileHeader.Header.Get("Content-Type")

	url, err := h.imageUsecase.UploadAvatarToS3(file, fileHeader.Filename, contentType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "upload success",
		"url":     url,
	})
}

func (h *ImageHandler) UploadLicense(c *gin.Context) {

	fileHeader, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing file"})
		return
	}
	file, err := fileHeader.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to open file"})
		return
	}
	contentType := fileHeader.Header.Get("Content-Type")

	url, err := h.imageUsecase.UploadLicenseToS3(file, fileHeader.Filename, contentType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "upload success",
		"url":     url,
	})
}
