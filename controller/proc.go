package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"dme_service/memory_util"
	"dme_service/model"
)

func ProcList(c *gin.Context) {
	dummy_procs := []model.Proc{
		{
			PID:          1111,
			PACKAGE_NAME: "com.aaa",
		},
		{
			PID:          2222,
			PACKAGE_NAME: "com.bbb",
		},
	}
	c.JSON(http.StatusOK, gin.H{
		"procs": dummy_procs,
	})
}

func ProcMemMaps(c *gin.Context) {
	pid, _ := strconv.Atoi(c.Query("pid"))
	memoryMaps := memory_util.ParseMemMaps(pid)
	c.JSON(http.StatusOK, gin.H{
		"mem_maps": memoryMaps,
	})
}
