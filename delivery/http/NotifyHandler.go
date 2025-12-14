package http

import (
	"golangAPI/entity"
	"golangAPI/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type NotifyHandler struct {
	notifyUC *usecase.NotifyUsecase
}

func RegisterNotifyRoutes(r gin.IRoutes, notifyUC *usecase.NotifyUsecase) {
	h := &NotifyHandler{notifyUC: notifyUC}
	r.POST("/notifications", h.PostNotification)
	r.GET("/notifications/:reciever_id", h.GetNotificationsByRecieverID)
	r.DELETE("/notifications/all/:reciever_id", h.DeleteNotificationsByRecieverID)
	r.DELETE("/notifications/:id", h.DeleteNotificationByID)
}

func (h *NotifyHandler) PostNotification(c *gin.Context) {
	var body entity.Notify
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	createdNotify, err := h.notifyUC.CreateNotification(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, createdNotify)
}

func (h *NotifyHandler) GetNotificationsByRecieverID(c *gin.Context) {
	recieverID := c.Param("reciever_id")
	notifications, err := h.notifyUC.GetNotificationByRecieverID(recieverID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, notifications)
}

func (h *NotifyHandler) DeleteNotificationsByRecieverID(c *gin.Context) {
	recieverID := c.Param("reciever_id")
	err := h.notifyUC.DeleteNotificationByRecieverID(recieverID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "notifications deleted"})
}

func (h *NotifyHandler) DeleteNotificationByID(c *gin.Context) {
	id := c.Param("id")
	err := h.notifyUC.DeleteNotificationByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "notification deleted"})
}
