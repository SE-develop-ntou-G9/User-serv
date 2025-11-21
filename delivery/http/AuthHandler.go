package http

import (
	"golangAPI/entity"
	"golangAPI/usecase"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"google.golang.org/api/idtoken"
)

type AuthHandler struct {
	authUC *usecase.AuthUsecase
}

func RegisterAuthRoutes(r gin.IRoutes, authUC *usecase.AuthUsecase) {
	h := &AuthHandler{authUC: authUC}

	r.POST("/auth/google", h.googleCredentialHandler)

	// r.GET("/auth/:provider", h.beginAuthHandler)

	// r.GET("/auth/:provider/callback", h.callbackAuthHandler)
}

type googleLoginRequest struct {
	Credential string `json:"credential" binding:"required"`
}

func (h *AuthHandler) googleCredentialHandler(c *gin.Context) {
	var req googleLoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	// 用 Google 官方套件驗證 ID token
	ctx := c.Request.Context()
	clientID := os.Getenv("GoogleClientID") // 要跟前端使用的 client_id 一樣

	payload, err := idtoken.Validate(ctx, req.Credential, clientID)
	if err != nil {
		// 這裡代表 token 不合法、aud 不對、過期…等
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid id token"})
		return
	}

	// payload.Subject = Google 的 user id (sub)
	sub := payload.Subject

	// 其他資訊在 payload.Claims 裡
	email, _ := payload.Claims["email"].(string)
	name, _ := payload.Claims["name"].(string)

	newUUID, _ := uuid.NewRandom()
	userID := newUUID.String()

	// 你原本的 usecase 是吃 goth.User，我們自己組一個
	oUser := entity.User{
		ID:             userID,
		Provider:       "google",
		ProviderUserID: sub,
		Email:          email,
		Name:           name,
	}

	u, token, err := h.authUC.LoginWithOAuth("google", oUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 回給前端的結構跟你原本一樣
	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"user": gin.H{
			"id":    u.ID,
			"name":  u.Name,
			"email": u.Email,
		},
	})
}

// func (h *AuthHandler) beginAuthHandler(c *gin.Context) {
// 	provider := c.Param("provider")

// 	q := c.Request.URL.Query()
// 	q.Set("provider", provider)
// 	c.Request.URL.RawQuery = q.Encode()

// 	gothic.BeginAuthHandler(c.Writer, c.Request)
// }

// func (h *AuthHandler) callbackAuthHandler(c *gin.Context) {
// 	provider := c.Param("provider")

// 	q := c.Request.URL.Query()
// 	q.Set("provider", provider)
// 	c.Request.URL.RawQuery = q.Encode()

// 	user, err := gothic.CompleteUserAuth(c.Writer, c.Request)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	u, token, err := h.authUC.LoginWithOAuth(provider, user)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"token": token,
// 		"user": gin.H{
// 			"id":    u.ID,
// 			"name":  u.Name,
// 			"email": u.Email,
// 		},
// 	})
// }
