package model

type MemorySection struct {
	ADDR_TART   int    `json:"addr_start"`
	ADDR_END    int    `json:"addr_end"`
	PERMISSIONS string `json:"permissions"`
	OFFSET      int    `json:"offset"`
	PATHNAME    string `json:"pathname"`
}

type MemorySections struct {
	PID          int             `json:"pid"`
	MEM_SECTIONS []MemorySection `json:"mem_sections"`
}
