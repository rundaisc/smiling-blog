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

type Flink struct {
	FlinkServices services.FlinkServices
}

func NewFlink() *Flink {
	return &Flink{
		FlinkServices: services.NewFlinkServices(),
	}
}

// 列表
func (slf *Flink) List(ctx *gin.Context) {
	list := slf.FlinkServices.GetList()
	ctx.JSON(http.StatusOK, tools.BuildSuccess(list))
}

// 保存
func (slf *Flink) Save(c *gin.Context) {
	params := request.FlinkSave{}
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
	code := slf.FlinkServices.Save(&params)
	if code != tools.OK {
		c.JSON(http.StatusOK, tools.BuildFailed(code))
		return
	}
	c.JSON(http.StatusOK, tools.BuildSuccess(nil))
}

// 删除
func (slf *Flink) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id == 0 {
		slog.Error(err)
		c.JSON(http.StatusOK, tools.BuildFailedWithMsg(tools.ValidateError, err.Error()))
		return
	}
	code := slf.FlinkServices.Delete(id)
	if code != tools.OK {
		c.JSON(http.StatusOK, tools.BuildFailed(code))
		return
	}
	c.JSON(http.StatusOK, tools.BuildSuccess(nil))
}
