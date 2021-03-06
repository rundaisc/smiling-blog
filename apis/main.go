package main

import (
	"github.com/gin-gonic/gin"
	"os"
	"path"
	"smiling-blog/config"
	"smiling-blog/router"
	"smiling-blog/tools/slog"
)

func main() {
	configInit()
	logInit()
	app := router.GetApp()
	if config.Configs.App.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	//视图文件
	app.LoadHTMLGlob("./resources/views/**/*")

	app.Static("static", "./resources/static")
	app.Use(gin.Logger())
	app.Use(gin.Recovery())
	app.Run(config.GetAppAddr())
}

func configInit() {
	currentDir, _ := os.Getwd()
	configureName := path.Join(currentDir, "application.yaml")
	config.Init(configureName)
}

func logInit() {
	slog.InitLogger(config.Configs.Log)
}
