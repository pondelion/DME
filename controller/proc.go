package controller

import (
	"net/http"
	"runtime"
	"strconv"

	"github.com/gin-gonic/gin"

	"dme_service/util"
)

func ProcList(c *gin.Context) {
	// dummy_procs := []model.Proc{
	// 	{
	// 		PID:          1111,
	// 		PACKAGE_NAME: "com.aaa",
	// 	},
	// 	{
	// 		PID:          2222,
	// 		PACKAGE_NAME: "com.bbb",
	// 	},
	// }
	if runtime.GOOS == "android" {
		procs := util.ProcessListAndroid()
		c.JSON(http.StatusOK, gin.H{
			"procs": procs,
		})
	} else {
		procs := util.ProcessList()
		c.JSON(http.StatusOK, gin.H{
			"procs": procs,
		})
	}
}

func ProcMemMaps(c *gin.Context) {
	pid, _ := strconv.Atoi(c.Param("pid"))
	memoryMaps := util.ParseMemMaps(pid)
	c.JSON(http.StatusOK, gin.H{
		"mem_maps": memoryMaps,
	})
}
