package routes

import (
	"github.com/ApexPlayground/Linkkit/controller"
	"github.com/ApexPlayground/Linkkit/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ShortnerSetupRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()
	apiIndex := "/api/v1"

	//service
	clickService := service.NewClickService(db)

	router.POST(apiIndex+"/shorten", controller.ShortnerController)

	router.GET("/:shortcode", func(c *gin.Context) {
		controller.Redirect(db, clickService, c)
	})

	return router
}
