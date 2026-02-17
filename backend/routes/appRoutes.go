package routes

import (
	"github.com/ApexPlayground/Linkkit/controller"
	"github.com/ApexPlayground/Linkkit/middleware"
	"github.com/gin-gonic/gin"
)

func AppSetupRouter(router *gin.Engine) {
	apiIndex := "/api/v1"

	router.GET("/:shortcode", controller.Redirect)
	router.GET("/qr/:id", controller.QRRedirect)

	protected := router.Group(apiIndex)
	protected.Use(middleware.AuthMiddleware)
	{

		protected.POST("/shorten", controller.ShortnerController)

		protected.POST("/qr/generate", controller.CreateQRCode)
		protected.GET("/qr", controller.QRListController)
		protected.DELETE("/qr/:id", controller.QRDeleteController)
	}
}
