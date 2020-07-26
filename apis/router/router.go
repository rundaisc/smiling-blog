package router

import (
	"github.com/gin-gonic/gin"
	"smiling-blog/controller"
)

func GetApp() *gin.Engine {
	app := gin.New()

	backend := app.Group("/api/backend")
	{
		site := controller.NewSite()
		backend.POST("/site/save", site.Save)
		backend.GET("/site/info", site.Info)

		admin := controller.NewAdmin()
		backend.POST("/login", admin.Login)

	}

	return app
}
