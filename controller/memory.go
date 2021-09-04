package controller

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"dme_service/util"

	"github.com/gin-gonic/gin"
)

func WriteProcMemInt64(c *gin.Context) {
	pid, _ := strconv.Atoi(c.Param("pid"))
	addr, _ := strconv.ParseInt(
		strings.Replace(c.Query("addr"), "0x", "", -1),
		16, 64,
	)
	value, _ := strconv.ParseInt(c.Query("value"), 10, 64)
	log.Printf("pid : %d, addr : %#x, value : %d", pid, addr, value)
	util.WriteProcMemInt64(pid, addr, value)

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func WriteProcMemInt32(c *gin.Context) {
	pid, _ := strconv.Atoi(c.Param("pid"))
	addr, _ := strconv.ParseInt(
		strings.Replace(c.Query("addr"), "0x", "", -1),
		16, 64,
	)
	value, _ := strconv.ParseInt(c.Query("value"), 10, 32)
	log.Printf("pid : %d, addr : %#x, value : %d", pid, addr, value)
	util.WriteProcMemInt32(pid, addr, int32(value))

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func SearchMemInt(c *gin.Context) {
	pid, _ := strconv.Atoi(c.Param("pid"))
	addr_start, _ := strconv.ParseUint(
		strings.Replace(c.Query("addr_start"), "0x", "", -1),
		16, 64,
	)
	addr_end, _ := strconv.ParseUint(
		strings.Replace(c.Query("addr_end"), "0x", "", -1),
		16, 64,
	)
	value, _ := strconv.ParseInt(c.Query("value"), 10, 64)

	if addr_start != 0 && addr_end != 0 {
		foundAddrs := util.SearchMemIntRange(pid, value, addr_start, addr_end)
		c.JSON(http.StatusOK, gin.H{
			"results": foundAddrs,
		})
	} else {
		maxSectionSize, _ := strconv.ParseUint(
			c.DefaultQuery("max_section_size", "10000000"),
			10, 64,
		)
		foundResults := util.SearchMemInt(pid, value, maxSectionSize)
		c.JSON(http.StatusOK, gin.H{
			"results": foundResults,
		})
	}
}

func ReadMemory(c *gin.Context) {
	pid, _ := strconv.Atoi(c.Param("pid"))
	addr_start, _ := strconv.ParseUint(
		strings.Replace(c.Query("addr_start"), "0x", "", -1),
		16, 64,
	)
	addr_end, _ := strconv.ParseUint(
		strings.Replace(c.Query("addr_end"), "0x", "", -1),
		16, 64,
	)
	memoryValue := util.ReadMemRange(pid, addr_start, addr_end)

	c.JSON(http.StatusOK, gin.H{
		"results": memoryValue,
	})
}

func Addr2MemMap(c *gin.Context) {
	pid, _ := strconv.Atoi(c.Param("pid"))
	addr, _ := strconv.ParseUint(
		strings.Replace(c.Query("addr"), "0x", "", -1),
		16, 64,
	)
	pMemMap := util.Addr2MemMap(pid, addr)

	if pMemMap == nil {
		c.JSON(http.StatusOK, gin.H{
			"results": nil,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"results": *pMemMap,
		})
	}
}
