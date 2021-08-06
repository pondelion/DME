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

func SearchMemInt(c *gin.Context) {
	pid, _ := strconv.Atoi(c.Query("pid"))
	addr_start, err1 := strconv.ParseUint(
		strings.Replace(c.Query("addr_start"), "0x", "", -1),
		16, 64,
	)
	addr_end, err2 := strconv.ParseUint(
		strings.Replace(c.Query("addr_end"), "0x", "", -1),
		16, 64,
	)
	value, _ := strconv.ParseInt(c.Query("value"), 10, 64)

	var foundAddrs []uint64
	if err1 != nil && err2 != nil {
		foundAddrs = memory_util.SearchMemIntRange(pid, value, addr_start, addr_end)
	} else {
		foundAddrs = memory_util.SearchMemInt(pid, value)
	}
	c.JSON(http.StatusOK, gin.H{
		"addrs": foundAddrs,
	})
}
