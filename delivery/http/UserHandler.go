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
