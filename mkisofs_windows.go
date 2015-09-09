// +build windows

package main

import (
//    "fmt"
    "syscall"
//    "unsafe"
)

var (
    kernel32, _        = syscall.LoadLibrary("kernel32.dll")
    getModuleHandle, _ = syscall.GetProcAddress(kernel32, "GetModuleHandleW")
)

func abort(funcname string, err error) {
    //panic(fmt.Sprintf("%s failed: %v", funcname, err))
    return
}
func mkisofs(dir string, isofile string){
    return
}


func GetModuleHandle() (handle uintptr) {
    defer syscall.FreeLibrary(kernel32)
    var nargs uintptr = 0
    if ret, _, callErr := syscall.Syscall(uintptr(getModuleHandle), nargs, 0, 0, 0); callErr != 0 {
        abort("Call GetModuleHandle", callErr)
    } else {
        handle = ret
    }
    return
}