package router

import (
	"github.com/gin-gonic/gin"
	"smiling-blog/controller"
	"smiling-blog/middleware"
)

func GetApp() *gin.Engine {
	app := gin.New()

	backend := app.Group("/api/backend")
	{
		//上传
		upload := controller.NewUpload()
		{
			backend.POST("/upload/image", upload.Images)
		}
		// 站点信息
		site := controller.NewSite()
		{
			backend.POST("/site/save", site.Save)
			backend.GET("/site/info", site.Info)
		}
		// 登录
		admin := controller.NewAdmin()
		{
			backend.POST("/login", admin.Login)
		}
		// 友情链接
		flink := controller.NewFlink()
		{
			backend.GET("/flink/list", flink.List)
			backend.POST("/flink/save", flink.Save)
			backend.DELETE("/flink/:id", flink.Delete)
		}

		// 分类管理
		category := controller.NewCategory()
		{
			backend.GET("/category/list", category.List)
			backend.POST("/category/save", category.Save)
			backend.DELETE("/category/:id", category.Delete)
		}

		// 导航管理
		nav := controller.NewNav()
		{
			backend.GET("/nav/list", nav.List)
			backend.POST("/nav/save", nav.Save)
			backend.DELETE("/nav/:id", nav.Delete)
		}

		// 文章管理
		article := controller.NewArticle()
		{
			backend.GET("/article/list", article.List)
			backend.POST("/article/save", article.Save)
			backend.DELETE("/article/:id", article.Delete)
			backend.GET("/article/detail/:id", article.Detail)
		}

	}
	home := controller.NewHome()
	app.GET("/", home.Index).Use(middleware.FrontParams())
	app.GET("/category/:categoryId", home.Index)
	app.GET("/tag/:tag", home.Index)
	app.GET("/posts/:articleId", home.Index)

	return app
}
