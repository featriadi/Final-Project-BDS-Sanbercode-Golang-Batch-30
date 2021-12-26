package routes

import (
	"Final-Project-BDS-Sanbercode-Golang-Batch-30/controllers"
	"Final-Project-BDS-Sanbercode-Golang-Batch-30/middlewares"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// set db to gin context
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	r.GET("/companies", controllers.GetAllCompany)
	r.GET("/companies/:id", controllers.GetCompanyById)

	r.GET("/games", controllers.GetAllGame)
	r.GET("/games/:id", controllers.GetGameById)

	r.GET("/tags", controllers.GetAllTag)
	r.GET("/tags/:id", controllers.GetTagById)

	r.GET("/game-tags", controllers.GetAllGameTag)
	r.GET("/game-tags/:id", controllers.GetGameTagById)

	r.GET("/game-images", controllers.GetAllGameImage)
	r.GET("/game-images/:id", controllers.GetGameImageById)

	r.GET("/reviews", controllers.GetAllReview)
	r.GET("/reviews/:id", controllers.GetReviewById)

	companiesMiddlewareRoute := r.Group("/companies")
	companiesMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
	companiesMiddlewareRoute.POST("/", controllers.CreateCompany)
	companiesMiddlewareRoute.PATCH("/:id", controllers.UpdateCompany)
	companiesMiddlewareRoute.DELETE("/:id", controllers.DeleteCompany)

	gamesMiddlewareRoute := r.Group("/games")
	gamesMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
	gamesMiddlewareRoute.POST("/", controllers.CreateGame)
	gamesMiddlewareRoute.PATCH("/:id", controllers.UpdateGame)
	gamesMiddlewareRoute.DELETE("/:id", controllers.DeleteGame)

	tagsMiddlewareRoute := r.Group("/tags")
	tagsMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
	tagsMiddlewareRoute.POST("/", controllers.CreateTag)
	tagsMiddlewareRoute.DELETE("/:id", controllers.DeleteTag)

	gameTagsMiddlewareRoute := r.Group("/game-tags")
	gameTagsMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
	gameTagsMiddlewareRoute.POST("/", controllers.CreateGameTag)
	gameTagsMiddlewareRoute.DELETE("/:id", controllers.DeleteGameTag)

	gameImagesMiddlewareRoute := r.Group("/game-images")
	gameImagesMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
	gameImagesMiddlewareRoute.POST("/", controllers.CreateGameImage)
	gameImagesMiddlewareRoute.DELETE("/:id", controllers.DeleteGameImage)

	reviewsMiddlewareRoute := r.Group("/reviews")
	reviewsMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
	reviewsMiddlewareRoute.POST("/", controllers.CreateReview)
	reviewsMiddlewareRoute.PATCH("/:id", controllers.UpdateReview)
	reviewsMiddlewareRoute.DELETE("/:id", controllers.DeleteReview)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
