package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"smiling-blog/repo/request"
	"smiling-blog/services"
	"smiling-blog/tools"
	"smiling-blog/tools/slog"
)

type Site struct {
	SiteServices services.SiteServices
}

func NewSite() *Site {
	return &Site{
		SiteServices: services.NewSiteServices(),
	}
}

// 保存site信息
func (slf *Site) Save(c *gin.Context) {
	params := request.SiteForm{}
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
	err := slf.SiteServices.Save(&params)
	if err != nil {
		c.JSON(http.StatusOK, tools.BuildFailed(tools.UnKnowError))
		return
	}
	c.JSON(http.StatusOK, tools.BuildSuccess(nil))
}

// 获取site信息
func (slf *Site) Info(c *gin.Context) {
	siteInfo := slf.SiteServices.GetInfo()
	c.JSON(http.StatusOK, tools.BuildSuccess(siteInfo))

}
