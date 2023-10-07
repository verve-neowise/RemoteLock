package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

func getStatus() string {
	if Exists("/data/local/tmp/.locked") {
		return "locked"
	} else {
		return "unlocked"
	}
}

func setLocked() {
	err := ioutil.WriteFile("/data/local/tmp/.locked", []byte("locked"), 0644) //create a new file
	if err != nil {
		fmt.Println("Error create locked status", err)
		return
	}

	fmt.Println("Created locked status!")
}

func setUnLocked() {
	cmd := exec.Command("rm", "\"/data/local/tmp/.locked\"")
	if err := cmd.Start(); err != nil {
		fmt.Println("Error remove locked status", err)
		return
	}
	fmt.Println("Remove locked status!")
}

func Exists(name string) (bool) {
    _, err := os.Stat(name)
    if err == nil {
        return true
    }
    if errors.Is(err, os.ErrNotExist) {
        return false
    }
    return false
}