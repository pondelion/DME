package main

import (
	"dme_service/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	engine.GET("/procs", controller.ProcList)
	memoryEngine := engine.Group("/proc/:pid")
	{
		memoryEngine.GET("/mem_maps", controller.ProcMemMaps)
		memoryEngine.POST("/write_int64", controller.WriteProcMemInt64)
		memoryEngine.GET("/search_mem_int", controller.SearchMemInt)
		memoryEngine.GET("/read_memory", controller.ReadMemory)
		memoryEngine.GET("/addr2mem_map", controller.Addr2MemMap)

	}
	engine.Run()
}
