package util

import (
	"bufio"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/cheggaaa/pb/v3"

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
		addrStart, _ := strconv.ParseUint(addrRanges[0], 16, 64)
		addrEnd, _ := strconv.ParseUint(addrRanges[1], 16, 64)
		permission := items[1]
		offset, _ := strconv.ParseUint(items[2], 16, 64)
		// device := items[3]
		// inode := items[4]
		pathname := ""
		if len(items) > 6 {
			items2 := strings.Split(line, "   ")
			if len(items) >= 2 {
				pathname = items2[len(items2)-1]
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

func ReadMemRange(pid int, addrStart uint64, addrEnd uint64) model.MemoryValue {
	filepath := fmt.Sprintf("/proc/%d/mem", pid)
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	buf := make([]byte, addrEnd-addrStart+1)
	file.ReadAt(buf, int64(addrStart))

	// hexStr := hex.Dump(buf)

	hexStr := hex.EncodeToString(buf)

	memoryValue := model.MemoryValue{
		PID:        pid,
		ADDR_START: addrStart,
		ADDR_END:   addrEnd,
		// MEMORY_VALUE: buf,
		MEMORY_VALUE: hexStr,
	}
	return memoryValue
}

func SearchMemIntRange(pid int, value int64, addrStart uint64, addrEnd uint64) []uint64 {
	return searchMemIntRange(pid, value, addrStart, addrEnd)
}

func SearchMemInt(pid int, value int64, max_section_size uint64) []model.MemorySearchResult {
	// Search all memeory sections described in /proc/[pid]/maps.
	memMaps := ParseMemMaps(pid)
	var foundResults []model.MemorySearchResult
	bar := pb.Simple.Start(len(memMaps.MEM_MAPS))
	bar.SetMaxWidth(80)
	SKIP_KEYWORDS := []string{
		".jar", ".ttf", ".ttc", ".so", ".art", ".oat", ".apk", ".dat",
	}
L:
	for _, memMap := range memMaps.MEM_MAPS {
		bar.Increment()
		fmt.Println(memMap.PATHNAME)
		log.Printf("%s : %d", memMap.PATHNAME, memMap.ADDR_END-memMap.ADDR_START)
		// If the memory section does not have read permission, skip searching.
		if !strings.Contains(memMap.PERMISSIONS, "r") {
			continue
		}
		for _, skipKeyword := range SKIP_KEYWORDS {
			if strings.Contains(memMap.PATHNAME, skipKeyword) {
				continue L
			}
		}
		if memMap.ADDR_END-memMap.ADDR_START > max_section_size {
			continue
		}
		foundAddrs := searchMemIntRange(pid, value, memMap.ADDR_START, memMap.ADDR_END)
		var results []model.MemorySearchResult
		for _, addr := range foundAddrs {
			results = append(results, model.MemorySearchResult{
				ADDR:        addr,
				PATHNAME:    memMap.PATHNAME,
				PERMISSIONS: memMap.PERMISSIONS,
			})
		}
		foundResults = append(
			foundResults,
			results...,
		)
	}
	bar.Finish()
	return foundResults
}

func searchMemIntRange(pid int, value int64, addrStart uint64, addrEnd uint64) []uint64 {
	filepath := fmt.Sprintf("/proc/%d/mem", pid)
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var addr uint64 = addrStart
	var foundAddrs []uint64
	buf := make([]byte, 64)
	for addr < addrEnd {
		file.ReadAt(buf, int64(addr))
		readValueInt64 := int64(binary.LittleEndian.Uint64(buf))
		readValueInt32 := int32(binary.LittleEndian.Uint32(buf))
		if readValueInt64 == value || readValueInt32 == int32(value) {
			log.Printf("addr : %#x : %d", addr, readValueInt64)
			foundAddrs = append(foundAddrs, addr)
		}
		addr += 1
	}
	return foundAddrs
}

func Addr2MemMap(pid int, addr uint64) *model.MemoryMap {
	memMaps := ParseMemMaps(pid)
	for _, memMap := range memMaps.MEM_MAPS {
		if addr >= memMap.ADDR_START && addr <= memMap.ADDR_END {
			return &memMap
		}
	}
	return nil
}
