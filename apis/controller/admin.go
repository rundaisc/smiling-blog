package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"smiling-blog/repo/request"
	"smiling-blog/services"
	"smiling-blog/tools"
	"smiling-blog/tools/slog"
)

type Admin struct {
	AdminServices services.AdminServices
}

func NewAdmin() *Admin {
	return &Admin{
		AdminServices: services.NewAdminServices(),
	}
}

// 登录
func (slf *Admin) Login(ctx *gin.Context) {
	params := request.LoginForm{}
	if err := ctx.BindJSON(&params); err != nil {
		slog.Error(err)
		ctx.JSON(http.StatusOK, tools.BuildFailed(tools.ParamsError))
		return
	}
	if err := tools.Validate(params); err != nil {
		slog.Error(err)
		ctx.JSON(http.StatusOK, tools.BuildFailedWithMsg(tools.ValidateError, err.Error()))
		return
	}
	request.ClientIp = ctx.ClientIP()
	code, resp := slf.AdminServices.Login(&params)
	if code != tools.OK {
		ctx.JSON(http.StatusOK, tools.BuildFailed(code))
		return
	}
	ctx.JSON(http.StatusOK, tools.BuildSuccess(resp))

}
