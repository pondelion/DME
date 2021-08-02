package main

import (
	"dme_service/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	engine.GET("/procs", controller.ProcList)
	memoryEngine := engine.Group("/memory")
	{
		memoryEngine.POST("/write_int", controller.WriteProcMemInt)
	}
	engine.Run()
}
