package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"smiling-blog/repo/request"
	"smiling-blog/services"
	"smiling-blog/tools"
	"smiling-blog/tools/slog"
	"strconv"
)

type article struct {
	ArticleServices services.ArticleServices
}

func NewArticle() *article {
	return &article{
		ArticleServices: services.NewArticleServices(),
	}
}

// 后台文章列表
func (slf *article) List(ctx *gin.Context) {
	params := request.ArticleSearchForm{}
	if err := ctx.BindQuery(&params); err != nil {
		slog.Error(err)
		ctx.JSON(http.StatusOK, tools.BuildFailed(tools.ParamsError))
		return
	}
	if err := tools.Validate(params); err != nil {
		slog.Error(err)
		ctx.JSON(http.StatusOK, tools.BuildFailedWithMsg(tools.ValidateError, err.Error()))
		return
	}
	total, list := slf.ArticleServices.GetAdminList(&params)
	ctx.JSON(http.StatusOK, tools.BuildSuccess(map[string]interface{}{"total": total, "list": list}))
}

// 保存
func (slf *article) Save(c *gin.Context) {
	params := request.ArticleForm{}
	if err := c.BindJSON(&params); err != nil {
		slog.Error(err)
		c.JSON(http.StatusOK, tools.BuildFailed(tools.ParamsError))
		return
	}
	if err := tools.Validate(params); err != nil {
		slog.Error(err)
		c.JSON(http.StatusOK, tools.BuildFailedWithMsg(tools.ValidateError, err.Error()))
		return
	}
	code := slf.ArticleServices.Save(&params)
	if code != tools.OK {
		c.JSON(http.StatusOK, tools.BuildFailed(code))
		return
	}
	c.JSON(http.StatusOK, tools.BuildSuccess(nil))
}

// 删除
func (slf *article) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id == 0 {
		slog.Error(err)
		c.JSON(http.StatusOK, tools.BuildFailedWithMsg(tools.ValidateError, err.Error()))
		return
	}
	code := slf.CategoryServices.Delete(id)
	if code != tools.OK {
		c.JSON(http.StatusOK, tools.BuildFailed(code))
		return
	}
	c.JSON(http.StatusOK, tools.BuildSuccess(nil))
}
