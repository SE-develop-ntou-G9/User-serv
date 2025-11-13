package http

import (
	"golangAPI/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
)

type AuthHandler struct {
	authUC *usecase.AuthUsecase
}

func RegisterAuthRoutes(r gin.IRoutes, authUC *usecase.AuthUsecase) {
	h := &AuthHandler{authUC: authUC}

	r.GET("/auth/:provider", h.beginAuthHandler)

	r.GET("/auth/:provider/callback", h.callbackAuthHandler)
}

func (h *AuthHandler) beginAuthHandler(c *gin.Context) {
	provider := c.Param("provider")

	q := c.Request.URL.Query()
	q.Set("provider", provider)
	c.Request.URL.RawQuery = q.Encode()

	gothic.BeginAuthHandler(c.Writer, c.Request)
}

func (h *AuthHandler) callbackAuthHandler(c *gin.Context) {
	provider := c.Param("provider")

	q := c.Request.URL.Query()
	q.Set("provider", provider)
	c.Request.URL.RawQuery = q.Encode()

	user, err := gothic.CompleteUserAuth(c.Writer, c.Request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, token, err := h.authUC.LoginWithOAuth(provider, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"user": gin.H{
			"id":    u.ID,
			"name":  u.Name,
			"email": u.Email,
		},
	})
}
