package main

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Cocoa
#include "b-r-o-w-s-e.h"
*/
import "C"

import (
	"os/exec"
)

var urlListener chan string

func main() {
	urlListener = make(chan string)
	go C.RunApp()
	url := <-urlListener

	browser := "Safari"
	cmd := exec.Command("open", "-a", browser, url)

	cmd.Run()
}

//export HandleURL
func HandleURL(u *C.char) {
	urlListener <- C.GoString(u)
}
