package memory_util

import (
	"encoding/binary"
	"fmt"
	"os"
	"syscall"
)

func WriteProcMemInt(pid int, addr int64, value int64) {
	filepath := fmt.Sprintf("/proc/%d/mem", pid)
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	err = syscall.PtraceAttach(pid)
	if err != nil {
		fmt.Println(err)
		return
	}
	var wsstatus syscall.WaitStatus
	syscall.Wait4(pid, &wsstatus, 0, nil)
	var whence int = 0
	_, err = file.Seek(addr, whence)
	if err != nil {
		err = syscall.PtraceDetach(pid)
		panic(err)
	}
	byte_data := make([]byte, binary.MaxVarintLen64)
	binary.PutVarint(byte_data, value)
	_, err = file.Write(byte_data)
	if err != nil {
		err = syscall.PtraceDetach(pid)
		panic(err)
	}
	err = syscall.PtraceDetach(pid)
	if err != nil {
		panic(err)
	}
}
