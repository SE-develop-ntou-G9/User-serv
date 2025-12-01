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
	r.PUT("/users/mod", h.EditUser)
	r.GET("/users/:id", h.GetUserByID)
	r.GET("/drivers/user/:user_id", h.GetDriverByUserID)
	r.DELETE("/users/deleteAll", h.DeleteAllUser)
	r.DELETE("/users/delete/:id", h.DeleteUserByID)
	r.DELETE("/drivers/delete/:user_id", h.DeleteDriverByUserID)
	r.PUT("/drivers/mod", h.EditDriver)
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

func (h *UserHandler) EditUser(c *gin.Context) {
	var body entity.User
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	edited, err := h.userUC.EditUser(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, edited)
}

func (h *UserHandler) GetUserByID(c *gin.Context) {
	id := c.Param("id")
	user, err := h.userUC.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) GetDriverByUserID(c *gin.Context) {
	userID := c.Param("user_id")
	driver, err := h.userUC.GetDriverByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, driver)
}

func (h *UserHandler) DeleteAllUser(c *gin.Context) {
	err := h.userUC.DeleteAllUser()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "All users deleted successfully"})
}

func (h *UserHandler) DeleteUserByID(c *gin.Context) {
	id := c.Param("id")
	err := h.userUC.DeleteUserByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

func (h *UserHandler) DeleteDriverByUserID(c *gin.Context) {
	userID := c.Param("user_id")
	err := h.userUC.DeleteDriverByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Driver deleted successfully"})
}

func (h *UserHandler) EditDriver(c *gin.Context) {
	var body entity.Driver
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	edited, err := h.userUC.EditDriver(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, edited)
}
