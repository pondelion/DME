package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Proc struct {
	PID          int    `json:"pid"`
	PACKAGE_NAME string `json:"package_name"`
}

func main() {
	engine := gin.Default()
	dummy_procs := []Proc{
		{
			PID:          1111,
			PACKAGE_NAME: "com.aaa",
		},
		{
			PID:          2222,
			PACKAGE_NAME: "com.bbb",
		},
	}
	engine.GET("/procs", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"procs": dummy_procs,
		})
	})
	engine.Run()
}
