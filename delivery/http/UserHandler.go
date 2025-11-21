package http

import (
	"golangAPI/entity"
	"golangAPI/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userUC *usecase.UserUsecase
}

func RegisterDriverRoutes(r gin.IRoutes, userUC *usecase.UserUsecase) {
	h := &UserHandler{userUC: userUC}

	r.POST("/users/driver", h.PostDriver)
}

func (h *UserHandler) PostUser(c *gin.Context) {
	var body entity.User
	if err := c.ShouldBindBodyWithJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	created, err := h.userUC.CreateUser(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, created)
}

func (h *UserHandler) PostDriver(c *gin.Context) {
	var body entity.Driver
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	created, err := h.userUC.CreateDriver(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, created)
}
