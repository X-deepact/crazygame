package routes

import (
	docs "crazygames.io/docs"
	"crazygames.io/handler"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Router struct {
	CategoryHandler *handler.CategoryHandler
	UserHandler     *handler.UserHandler
	AdsHander       *handler.AdsHandler
	OAuthHandler    *handler.OAuthHandler
	AuthHandler     *handler.AuthHandler
	GameHander      *handler.GameHandler
}

func NewRouter(category *handler.CategoryHandler, user *handler.UserHandler, ads *handler.AdsHandler, game *handler.GameHandler, Oauth *handler.OAuthHandler, auth *handler.AuthHandler) *Router {
	return &Router{
		CategoryHandler: category,
		UserHandler:     user,
		AdsHander:       ads,
		GameHander:      game,
		OAuthHandler:    Oauth,
		AuthHandler:     auth,
	}
}

func (ro *Router) RegisterRoutes(r *gin.Engine) {
	apiGroup := r.Group("/api")
	{
		categoryApi := apiGroup.Group("/category")
		categoryApi.GET("/", ro.CategoryHandler.GetAll)
		categoryApi.GET("/menu", ro.CategoryHandler.GetMenu)
		categoryApi.GET("/:id", ro.CategoryHandler.GetByID)
		categoryApi.POST("", ro.CategoryHandler.Create)
		categoryApi.PUT("/:id", ro.CategoryHandler.Update)
		categoryApi.DELETE("/:id", ro.CategoryHandler.Delete)

		userApi := apiGroup.Group("/user")
		userApi.GET("/", ro.UserHandler.GetAll)
		userApi.GET("/:id", ro.UserHandler.GetByID)
		userApi.POST("", ro.UserHandler.Create)
		userApi.PUT("/:id", ro.UserHandler.Update)
		userApi.DELETE("/:id", ro.UserHandler.Delete)
		userApi.POST("/forgot-password", ro.UserHandler.ForgotPassword)
		userApi.POST("/reset-password", ro.UserHandler.ResetPassword)

		adsApi := apiGroup.Group("/ads")
		adsApi.GET("/", ro.AdsHander.GetAll)
		adsApi.GET("/:id", ro.AdsHander.GetByID)
		adsApi.POST("/", ro.AdsHander.Create)
		adsApi.PUT("/:id", ro.AdsHander.Update)
		adsApi.DELETE("/:id", ro.AdsHander.Delete)

		gameApi := apiGroup.Group("/game")
		gameApi.GET("/", ro.GameHander.GetAll)
		gameApi.GET("/:id", ro.GameHander.GetByID)
		gameApi.GET("/category/:id", ro.GameHander.GetByCategoryID)
		gameApi.POST("/", ro.GameHander.Create)
		gameApi.PUT("/:id", ro.GameHander.Update)
		gameApi.DELETE("/:id", ro.GameHander.Delete)

		OAuthApi := apiGroup.Group("/Oauth")
		OAuthApi.GET("/google/login", ro.OAuthHandler.GoogleLogin)
		OAuthApi.GET("/google/callback", ro.OAuthHandler.GoogleCallback)

		authApi := apiGroup.Group("/auth")
		authApi.POST("/login", ro.AuthHandler.Login)
		authApi.POST("/register", ro.AuthHandler.Register)
		authApi.POST("/check-email", ro.AuthHandler.CheckEmail)
	}

	{
		docs.SwaggerInfo.BasePath = "/api"
		r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	}
}
