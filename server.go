package main

import (
	"io"
	"net/http"
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
		c.JSON(http.StatusOK, videoController.FindAll())
	})

	server.POST("/videos", func(c *gin.Context) {
		err := videoController.Save(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"message": "Video saved successfully"})
		}
	})

	server.Run(":8080")
}
