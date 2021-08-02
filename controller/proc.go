package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

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
