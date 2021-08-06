package main

import (
	"dme_service/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	engine.GET("/procs", controller.ProcList)
	procEngine := engine.Group("/proc")
	{
		procEngine.GET("/mem_maps", controller.ProcMemMaps)
	}
	memoryEngine := engine.Group("/memory")
	{
		memoryEngine.POST("/write_int", controller.WriteProcMemInt)
		memoryEngine.GET("/search_mem_int", controller.SearchMemInt)
	}
	engine.Run()
}
