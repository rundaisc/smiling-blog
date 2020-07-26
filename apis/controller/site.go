package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"smiling-blog/repo/request"
	"smiling-blog/tools"
)

type SiteController struct {
	Ctx *gin.Context
}

// 保存site信息
func (slf *SiteController) Save() {
	params := request.SiteForm{}
	if err := slf.Ctx.ShouldBind(&params); err != nil {
		slf.Ctx.JSON(http.StatusOK, tools.BuildFailed(tools.ParamsError))
		return
	}
	if err := tools.Validate(params); err != nil {
		slf.Ctx.JSON(http.StatusOK, tools.BuildFailedWithMsg(tools.ValidateError, err.Error()))
		return
	}
}

// 获取site信息
func (slf *SiteController) Info() {

}
