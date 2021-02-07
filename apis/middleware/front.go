package middleware

import (
	"github.com/gin-gonic/gin"
	"smiling-blog/repo/response"
	"smiling-blog/services"
	"strconv"
)

func FrontParams(context *gin.Context) {
	//站点信息
	site := services.NewSiteServices()
	siteInfo := site.GetInfo()

	categoryId := context.GetInt("categoryId")
	articleId := context.GetInt("articleId")

	path := ""
	//如果是分类路由
	if categoryId != 0 {
		path = "/category/" + strconv.Itoa(categoryId)
	} else if articleId != 0 {
		// 如果是文章详情
		article := services.NewArticleServices()
		articelDetail, err := article.GetAdminDetail(articleId)
		if err != nil {
			context.Abort()
			return
		} else {
			path = "/category/" + strconv.Itoa(articelDetail.CategoryId)
		}
		context.Set("articelDetail", articelDetail)
	} else {
		path = context.FullPath()
	}

	//导航
	nav := services.NewNavServices()
	navList := nav.GetFrontList(path)
	//分类
	category := services.NewCategoryServices()
	categories := category.FrontList()
	public := response.Public{
		SiteInfo:     siteInfo,
		NavList:      navList,
		CategoryList: categories,
	}
	context.Set("public", public)
	context.Next()
}
