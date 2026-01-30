package routes

import (
	"github.com/ApexPlayground/Linkkit/controller"
	"github.com/ApexPlayground/Linkkit/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ShortenerSetupRouter(router *gin.Engine, db *gorm.DB) {
	apiIndex := "/api/v1"

	clickService := service.NewClickService(db)

	router.POST(apiIndex+"/shorten", controller.ShortnerController)

	// Redirect route
	router.GET("/:shortcode", func(c *gin.Context) {
		controller.Redirect(db, clickService, c)
	})
}
