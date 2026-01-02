package routes

import (
	"github.com/ApexPlayground/Linkkit/controller"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ShortnerSetupRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()
	apiIndex := "/api/v1"

	router.POST(apiIndex+"/shorten", controller.ShortnerController)

	router.GET("/:shortcode", func(c *gin.Context) {
		controller.Redirect(db, c)
	})

	return router
}
