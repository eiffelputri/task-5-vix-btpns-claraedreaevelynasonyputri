package router

import (
	"task-5-vix-btpns-ClaraEdreaEvelynaSonyPutri/controllers"
	"task-5-vix-btpns-ClaraEdreaEvelynaSonyPutri/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// InitializeRouter initializes the routes for the application
func InitializeRouter(r *gin.Engine, db *gorm.DB) {
	// Public routes
	r.POST("/users/register", func(c *gin.Context) {
		controllers.RegisterUser(c, db)
	})
	r.POST("/users/login", func(c *gin.Context) {
		controllers.LoginUser(c, db)
	})

	// Authenticated routes
	auth := r.Group("/auth")
	auth.Use(middlewares.AuthMiddleware()) // Middleware for authentication
	{
		auth.PUT("/users/:userID", func(c *gin.Context) {
			controllers.UpdateUser(c, db)
		})
		auth.DELETE("/users/:userID", func(c *gin.Context) {
			controllers.DeleteUser(c, db)
		})

		auth.POST("/photos", func(c *gin.Context) {
			controllers.UploadPhoto(c, db)
		})
		auth.PUT("/photos/:photoID", func(c *gin.Context) {
			controllers.UpdatePhoto(c, db)
		})
		auth.DELETE("/photos/:photoID", func(c *gin.Context) {
			controllers.DeletePhoto(c, db)
		})
	}
}
