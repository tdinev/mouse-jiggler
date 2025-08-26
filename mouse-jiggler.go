package main

import (
    "fmt"
    "syscall"
    "time"
    "unsafe"
)

var (
    user32          = syscall.NewLazyDLL("user32.dll")
    getCursorPos    = user32.NewProc("GetCursorPos")
    setCursorPos    = user32.NewProc("SetCursorPos")
)

type POINT struct {
    X int32
    Y int32
}

func main() {
    fmt.Println("Mouse jiggler started. Moving cursor by 1 pixel every minute...")

    for {
        var pt POINT

        ret, _, _ := getCursorPos.Call(uintptr(unsafe.Pointer(&pt)))
        if ret == 0 {
            fmt.Println("Failed to get cursor position")
            continue
        }

        setCursorPos.Call(uintptr(pt.X+1), uintptr(pt.Y+1))

        time.Sleep(60 * time.Second)
    }
}
