package memory_util

import (
	"errors"

	"dme_service/model"
)

func ParseMemMaps(pid int) (model.MemorySections, error) {
	var dummy_mem_secs model.MemorySections
	return dummy_mem_secs, errors.New("Not implemented")
}
