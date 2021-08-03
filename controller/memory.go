package controller

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"dme_service/memory_util"

	"github.com/gin-gonic/gin"
)

func WriteProcMemInt(c *gin.Context) {
	pid, _ := strconv.Atoi(c.Query("pid"))
	addr, _ := strconv.ParseInt(
		strings.Replace(c.Query("addr"), "0x", "", -1),
		16, 64,
	)
	value, _ := strconv.ParseInt(c.Query("value"), 10, 64)
	log.Printf("pid : %d, addr : %#x, value : %d", pid, addr, value)
	memory_util.WriteProcMemInt(pid, addr, value)

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
