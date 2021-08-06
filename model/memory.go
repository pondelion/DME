package model

type MemoryMap struct {
	ADDR_START  uint64 `json:"addr_start"`
	ADDR_END    uint64 `json:"addr_end"`
	PERMISSIONS string `json:"permissions"`
	OFFSET      uint64 `json:"offset"`
	PATHNAME    string `json:"pathname"`
}

type MemoryMaps struct {
	PID      int         `json:"pid"`
	MEM_MAPS []MemoryMap `json:"mem_maps"`
}
