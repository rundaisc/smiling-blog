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

type nav struct {
	NavServices services.NavServices
}

func NewNav() *nav {
	return &nav{
		NavServices: services.NewNavServices(),
	}
}

// 列表
func (slf *nav) List(ctx *gin.Context) {
	list := slf.NavServices.GetList()
	ctx.JSON(http.StatusOK, tools.BuildSuccess(list))
}

// 保存
func (slf *nav) Save(c *gin.Context) {
	params := request.NavSave{}
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
	code := slf.NavServices.Save(&params)
	if code != tools.OK {
		c.JSON(http.StatusOK, tools.BuildFailed(code))
		return
	}
	c.JSON(http.StatusOK, tools.BuildSuccess(nil))
}

// 删除
func (slf *nav) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id == 0 {
		slog.Error(err)
		c.JSON(http.StatusOK, tools.BuildFailedWithMsg(tools.ValidateError, err.Error()))
		return
	}
	code := slf.NavServices.Delete(id)
	if code != tools.OK {
		c.JSON(http.StatusOK, tools.BuildFailed(code))
		return
	}
	c.JSON(http.StatusOK, tools.BuildSuccess(nil))
}
