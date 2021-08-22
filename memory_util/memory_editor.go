package memory_util

import (
	"encoding/binary"
	"fmt"
	"os"
	"syscall"
)

func WriteProcMemInt64(pid int, addr int64, value int64) {
	filepath := fmt.Sprintf("/proc/%d/mem", pid)
	file, err := os.OpenFile(filepath, os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()
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
		_ = syscall.PtraceDetach(pid)
		panic(err)
	}

	byte_data := make([]byte, 8)
	binary.LittleEndian.PutUint64(byte_data, uint64(value))
	// log.Printf("%d", int64(binary.LittleEndian.Uint64(byte_data)))
	_, err = file.Write(byte_data)
	if err != nil {
		_ = syscall.PtraceDetach(pid)
		panic(err)
	}
	err = syscall.PtraceDetach(pid)
	if err != nil {
		panic(err)
	}
}

func WriteProcMemInt32(pid int, addr int64, value int32) {
	filepath := fmt.Sprintf("/proc/%d/mem", pid)
	file, err := os.OpenFile(filepath, os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()
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
		_ = syscall.PtraceDetach(pid)
		panic(err)
	}

	byte_data := make([]byte, 4)
	binary.LittleEndian.PutUint32(byte_data, uint32(value))
	_, err = file.Write(byte_data)
	if err != nil {
		_ = syscall.PtraceDetach(pid)
		panic(err)
	}
	err = syscall.PtraceDetach(pid)
	if err != nil {
		panic(err)
	}
}
