package middleware

import (
	"github.com/gin-gonic/gin"
	"smiling-blog/services"
)

func FrontParams() gin.HandlerFunc {
	return func(context *gin.Context) {
		//站点信息
		site := services.NewSiteServices()
		siteInfo := site.GetInfo()

		//导航
		nav := services.NewNavServices()
		navList := nav.GetFrontList()
		//分类
		category := services.NewCategoryServices()
		categories := category.FrontList()

		context.Bind(gin.H{
			"navList":    navList,
			"categories": categories,
			"siteInfo":   siteInfo,
		})

		context.Next()

	}
}
