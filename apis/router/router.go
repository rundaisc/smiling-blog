package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetApp() *gin.Engine {
	app := gin.New()
	backend := app.Group("/api/backend")
	//backend.POST("/site/save",controller.SiteController.Save)
	//backend.POST("/site/info",controller.SiteController.Info)
	backend.GET("/test", func(context *gin.Context) {
		context.JSON(http.StatusOK, nil)
		return
	})

	return app
}
