package middleware

import (
	"github.com/gin-gonic/gin"
	"smiling-blog/dao"
	"smiling-blog/repo/response"
	"smiling-blog/services"
	"strconv"
)

func FrontParams(context *gin.Context) {
	//站点信息
	site := services.NewSiteServices()
	siteInfo := site.GetInfo()
	var categoryId int
	var articleId int
	categoryIds := context.Param("categoryId")
	if categoryIds != "" {
		categoryId, _ = strconv.Atoi(categoryIds)
	}
	articleIds := context.Param("articleId")
	if articleIds != "" {
		articleId, _ = strconv.Atoi(articleIds)
	}
	context.Set("categoryId", categoryId)

	path := ""
	//如果是分类路由
	if categoryId != 0 {
		path = "/category/" + categoryIds
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
	//标签
	tag := dao.NewArticleRelationDao()
	tagList := tag.FrontShow()
	//友情链接
	flink := services.NewFlinkServices()
	flinkList := flink.GetFrontList()

	public := response.Public{
		SiteInfo:     siteInfo,
		NavList:      navList,
		CategoryList: categories,
		TagList:      tagList,
		FlinkList:    flinkList,
	}
	context.Set("public", public)
	context.Next()
}
