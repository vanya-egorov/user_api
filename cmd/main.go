package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"github.com/vanya-egorov/user_api/config"
	"github.com/vanya-egorov/user_api/handler"
	"github.com/vanya-egorov/user_api/repository"
	"github.com/vanya-egorov/user_api/service"
	"log"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("no env file found")
	}

	db := config.InitDB()
	repo := repository.NewUserRepository(db)
	srv := service.NewUserService(repo)
	h := handler.NewHandler(srv)

	r := gin.Default()
	h.RegisterRoutes(r)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL("/docs/swagger.json")))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	log.Println("Server running on port", port)

	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
