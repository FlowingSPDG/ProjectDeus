package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/FlowingSPDG/ProjectDeus/server/cmd/di"
	cs2loghttp "github.com/FlowingSPDG/cs2-log-http"
)

func main() {
	r := gin.Default()
	logController := di.InitializeLogController()
	logHandler := cs2loghttp.NewLogHandler(logController.Handle)
	r.POST("/servers", di.InitializeRegisterServerController().Handle)
	r.POST("/servers/:id/log", logHandler.Handle())
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello!"})
	})
	log.Panicf("Failed to listen port 3090 : %v\n", r.Run("0.0.0.0:3090"))
}
