package routes

import (
	"go-auth/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(routes *gin.Engine) {

	// routes.Use(middleware.Authenticate())

	routes.GET("users", controllers.GetUsers())
	routes.GET("users/:user_id", controllers.GetUser())

}
