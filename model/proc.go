package model

type ProcDroid struct {
	PID          int    `json:"pid"`
	PPID         int    `json:"ppid"`
	PACKAGE_NAME string `json:"package_name"`
}

type Proc struct {
	PID          int    `json:"pid"`
	PPID         int    `json:"ppid"`
	PROCESS_NAME string `json:"process_name"`
}
