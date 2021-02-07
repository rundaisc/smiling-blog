package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"smiling-blog/repo/request"
	"smiling-blog/services"
	"smiling-blog/tools"
	"smiling-blog/tools/slog"
)

type Home struct {
	ArticleServices services.ArticleServices
}

func NewHome() *Home {
	return &Home{
		ArticleServices: services.NewArticleServices(),
	}
}

//首页以及列表页
func (slf *Home) Index(c *gin.Context) {
	params := request.HomeForm{}
	if err := c.BindQuery(&params); err != nil {
		slog.Error(err)
		c.JSON(http.StatusOK, tools.BuildFailed(tools.ParamsError))
		return
	}
	if err := tools.Validate(params); err != nil {
		slog.Error(err)
		c.JSON(http.StatusOK, tools.BuildFailedWithMsg(tools.ValidateError, err.Error()))
		return
	}
	params.PageSize = 10
	if params.Page == 0 {
		params.Page = 1
	}
	params.CategoryId = c.GetInt("categoryId")
	params.Tag = c.GetString("tag")
	total, list := slf.ArticleServices.GetFrontList(&params)
	pub, _ := c.Get("public")
	data := gin.H{
		"articleNumber": total,
		"articeList":    list,
		"currentPage":   params.Page,
		"pageSize":      params.PageSize,
		"public":        pub,
	}
	c.HTML(http.StatusOK, "posts/index.html", data)

}
