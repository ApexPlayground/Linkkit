package routes

import (
	"github.com/ApexPlayground/Linkkit/controller"
	"github.com/gin-gonic/gin"
)

func ShortenerSetupRouter(router *gin.Engine) {
	apiIndex := "/api/v1"

	router.POST(apiIndex+"/shorten", controller.ShortnerController)
	router.GET("/:shortcode", controller.Redirect)
}
