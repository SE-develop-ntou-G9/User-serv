package main

import (
	delivHTTP "golangAPI/delivery/http"
	"golangAPI/infrastructure"
	"golangAPI/repositories"
	"golangAPI/usecase"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := infrastructure.NewDBInstance()

	if err != nil {
		panic("Failed to connect to database")
	}

	infrastructure.InitAuth()

	userRepo := repositories.NewUserRepository(db)

	router := gin.Default()
	v1 := router.Group("/v1")

	router.Use(cors.Default())

	authUC := usecase.NewAuthUsecase(userRepo) // Assuming you have a userRepo instance
	delivHTTP.RegisterAuthRoutes(v1, authUC)

	router.Run(":8080")
}
