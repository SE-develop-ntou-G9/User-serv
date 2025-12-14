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
	s3client := infrastructure.NewS3Client()
	if err != nil {
		panic("Failed to connect to database")
	}

	infrastructure.InitAuth()

	userRepo := repositories.NewUserRepository(db)
	imageRepo := repositories.NewImageRepository(s3client)
	notifyRepo := repositories.NewNotifyRepository(db)

	router := gin.Default()
	v1 := router.Group("/v1")

	router.Use(cors.Default())

	authUC := usecase.NewAuthUsecase(userRepo) // Assuming you have a userRepo instance
	userUC := usecase.NewUserUsecase(userRepo)
	imageUC := usecase.NewImageUsecase(imageRepo)
	notifyUC := usecase.NewNotifyUsecase(notifyRepo)
	delivHTTP.RegisterAuthRoutes(v1, authUC)
	delivHTTP.RegisterUserRoutes(v1, userUC)
	delivHTTP.RegisterImageRoutes(v1, imageUC)
	delivHTTP.RegisterNotifyRoutes(v1, notifyUC)

	router.Run(":8080")
}
