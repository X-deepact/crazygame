package main

import (
	"log"

	"crazygames.io/config"
	"crazygames.io/handler"
	"crazygames.io/repositories"
	routes "crazygames.io/route"
	"crazygames.io/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()

	db := config.ConnectDatabase()
	redisClient := config.ConnectRedis()
	defer redisClient.Close()
	minioClient := config.ConnectMinIO()
	minioService := services.NewMinIOService(minioClient)

	r := gin.Default()

	r.Use(gin.Logger(), gin.Recovery())

	corsConf := cors.DefaultConfig()
	corsConf.AllowHeaders = append(corsConf.AllowHeaders, "Authorization")
	corsConf.AllowOrigins = config.AppConfig.ALLOW_ORIGINS
	r.Use(cors.New(corsConf))

	categoryRepo := repositories.NewCategoryRepository(db)
	categoryService := services.NewCategoryService(categoryRepo, minioService)
	categoryHandler := handler.NewCategoryHandler(categoryService)

	userRepo := repositories.NewUserRepository(db)
	passwordResetTokenRepo := repositories.NewPasswordResetTokenRepository(db)
	userService := services.NewUserService(userRepo, passwordResetTokenRepo)
	userHandler := handler.NewUserHandler(userService)

	adsRepo := repositories.NewAdsRepository(db)
	adsService := services.NewAdsService(adsRepo, minioClient)
	adsHandler := handler.NewAdsHandler(adsService)

	OAuthService := services.NewOAuthService(userRepo)
	OAuthHandler := handler.NewOAuthHandler(OAuthService)

	gameRepo := repositories.NewGameRepository(db)
	gameService := services.NewGameService(gameRepo, minioClient)
	gameHandler := handler.NewGameHandler(gameService)

	authService := services.NewAuthService(userRepo)
	authHandler := handler.NewAuthHandler(authService, userService)

	router := routes.NewRouter(categoryHandler, userHandler, adsHandler, gameHandler, OAuthHandler, authHandler)

	router.RegisterRoutes(r)

	log.Println("Server is running on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
