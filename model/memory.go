package model

type MemoryMap struct {
	ADDR_START  int64  `json:"addr_start"`
	ADDR_END    int64  `json:"addr_end"`
	PERMISSIONS string `json:"permissions"`
	OFFSET      int64  `json:"offset"`
	PATHNAME    string `json:"pathname"`
}

type MemoryMaps struct {
	PID      int         `json:"pid"`
	MEM_MAPS []MemoryMap `json:"mem_maps"`
}
