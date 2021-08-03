package memory_util

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"dme_service/model"
)

func ParseMemMaps(pid int) model.MemoryMaps {
	filepath := fmt.Sprintf("/proc/%d/maps", pid)
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	memoryMaps := []model.MemoryMap{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// line format example => 70791000-70a18000 rw-p 00000000 fd:00 1342                               /system/framework/arm64/boot.art
		fmt.Println(line)
		items := strings.Split(line, " ")
		addrRanges := strings.Split(items[0], "-")
		addrStart, _ := strconv.ParseInt(addrRanges[0], 16, 64)
		addrEnd, _ := strconv.ParseInt(addrRanges[1], 16, 64)
		permission := items[1]
		offset, _ := strconv.ParseInt(items[2], 16, 64)
		// device := items[3]
		// inode := items[4]
		pathname := ""
		if len(items) > 5 {
			items2 := strings.Split(line, "   ")
			if len(items) >= 2 {
				pathname = items2[len(items2)-1]
			} else {
				pathname = ""
			}
		}
		memoryMap := model.MemoryMap{
			addrStart,
			addrEnd,
			permission,
			offset,
			pathname,
		}
		memoryMaps = append(memoryMaps, memoryMap)
	}
	return model.MemoryMaps{pid, memoryMaps}
}
