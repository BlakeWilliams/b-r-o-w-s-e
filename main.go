package main

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Cocoa
#include "b-r-o-w-s-e.h"
*/
import "C"

import (
	"io/ioutil"
	"os"
	"os/exec"
	"os/user"
	"time"
)

var urlListener chan string = make(chan string)

// TODOS
// default config file copied when missing

func main() {
	config := loadConfig()

	go func() {
		timeout := time.After(4 * time.Second)
		select {
		case url := <-urlListener:
			browser := config.GetBrowserForUrl(url)
			cmd := exec.Command("open", "-a", browser.Path, url)

			cmd.Run()
			os.Exit(0)
		case <-timeout:
			os.Exit(1)
		}
	}()

	C.RunApp()
}

func loadConfig() Config {
	content, err := ioutil.ReadFile(homeDir() + "/.config/b-r-o-w-s-e/config.json")

	if err != nil {
		ShowError(
			"Could not load config",
			"Try creating a config file in ~/.config/b-r-o-w-s-e/config.json",
		)

		panic(err)
	}

	return ParseConfig(string(content))
}

func homeDir() string {
	currentUser, err := user.Current()
	if err != nil {
		ShowError(
			"Could not load home directory",
			err.Error(),
		)
		panic(err)
	}

	return currentUser.HomeDir
}

func ShowError(title string, details string) {
	C.ShowAlert(
		C.CString(title),
		C.CString(details),
	)
}

//export HandleURL
func HandleURL(u *C.char) {
	urlListener <- C.GoString(u)
}
