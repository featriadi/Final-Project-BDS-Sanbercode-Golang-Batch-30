package routes

import (
	"Final-Project-BDS-Sanbercode-Golang-Batch-30/controllers"

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

	r.GET("/company", controllers.GetAllCompany)
	r.GET("/company/:id", controllers.GetCompanyById)
	r.POST("/company", controllers.CreateCompany)
	r.PATCH("/company/:id", controllers.UpdateCompany)
	r.DELETE("/company/:id", controllers.DeleteCompany)

	r.GET("/game", controllers.GetAllGame)
	// r.GET("/game/:id", controllers.GetCompanyById)
	// r.POST("/game", controllers.CreateCompany)
	// r.PATCH("/game/:id", controllers.UpdateCompany)
	// r.DELETE("/game/:id", controllers.DeleteCompany)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
