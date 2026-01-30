package routes

import (
	"github.com/ApexPlayground/Linkkit/controller"
	"github.com/ApexPlayground/Linkkit/middleware"
	"github.com/gin-gonic/gin"
)

func UserSetupRouter(router *gin.Engine) {
	apiIndex := "/api/v1/users"

	// Public routes
	router.POST(apiIndex+"/signup", controller.CreateUser)
	router.POST(apiIndex+"/login", controller.Login)

	// Protected routes (require JWT)
	protected := router.Group(apiIndex)
	protected.Use(middleware.AuthMiddleware)

	protected.GET("/me", controller.GetUser)
	protected.GET("/", controller.ListUsers)
	protected.PUT("/", controller.UpdateUser)
	protected.DELETE("/", controller.DeleteUser)
}
