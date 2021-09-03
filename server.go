package main

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/llucasreis/go-gin/controller"
	"github.com/llucasreis/go-gin/middlewares"
	"github.com/llucasreis/go-gin/service"
	gindump "github.com/tpkeeper/gin-dump"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	server := gin.New()

	setupLogOutput()

	server.Use(gin.Recovery(), middlewares.Logger(),
		middlewares.BasicAuth(), gindump.Dump())

	server.GET("/videos", func(c *gin.Context) {
		c.JSON(200, videoController.FindAll())
	})

	server.POST("/videos", func(c *gin.Context) {
		c.JSON(200, videoController.Save(c))
	})

	server.Run(":8080")
}
