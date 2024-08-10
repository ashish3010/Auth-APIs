package routes

import (
	"go-auth/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(routes *gin.Engine) {

	routes.POST("user/signup", controllers.Signup())
	routes.POST("user/login", controllers.Login())
}
