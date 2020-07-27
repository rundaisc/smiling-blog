package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
	"smiling-blog/config"
	"smiling-blog/services"
	"smiling-blog/tools"
	"smiling-blog/tools/slog"
	"strings"
	"time"
)

type Upload struct {
	UpdateServices services.UploadServices
}

func NewUpload() *Upload {
	return &Upload{
		UpdateServices: services.NewUploadServices(),
	}
}

// 上传图片
func (slf *Upload) Images(ctx *gin.Context) {
	file, err := ctx.FormFile("image")
	if err != nil {
		slog.Error(err)
		ctx.JSON(http.StatusOK, tools.BuildFailed(tools.UnKnowError))
		return
	}
	ext := path.Ext(file.Filename)
	f, ferr := file.Open()
	if ferr != nil {
		slog.Error(ferr)
		ctx.JSON(http.StatusOK, tools.BuildFailed(tools.UnKnowError))
		return
	}
	defer f.Close()
	filename := time.Now().Format("20060102150405") + ext
	if uploadName, err := slf.UpdateServices.Qiniu(f, file.Size, filename); err == nil {
		qiniuUrl := fmt.Sprintf("%s/%s", strings.TrimRight(config.Configs.QiNiu.Url, "/"), uploadName)
		ctx.JSON(http.StatusOK, tools.BuildSuccess(map[string]string{
			"file": qiniuUrl,
		}))
	} else {
		slog.Error(err)
		ctx.JSON(http.StatusOK, tools.BuildFailed(tools.UnKnowError))
		return
	}

}
