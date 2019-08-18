package main

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Cocoa
#include "b-r-o-w-s-e.h"
*/
import "C"

import (
	"io/ioutil"
	"os/exec"
	"os/user"
)

var urlListener chan string

// TODOS
// cocoa alert instead of panic
// default config file copied when missing

func main() {
	config := loadConfig()

	urlListener = make(chan string)
	go C.RunApp()
	url := <-urlListener

	browser := config.GetBrowserForUrl(url)
	cmd := exec.Command("open", "-a", browser.Path, url)

	cmd.Run()
}

func loadConfig() Config {
	content, err := ioutil.ReadFile(homeDir() + "/.config/b-r-o-w-s-e/config.json")

	if err != nil {
		panic(err)
	}

	return ParseConfig(string(content))
}

func homeDir() string {
	currentUser, err := user.Current()
	if err != nil {
		panic(err)
	}

	return currentUser.HomeDir
}

//export HandleURL
func HandleURL(u *C.char) {
	urlListener <- C.GoString(u)
}
