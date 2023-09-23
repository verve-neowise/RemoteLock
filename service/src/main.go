package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func main() {

	fmt.Println(runtime.GOOS)

	model := deviceModel()
	id := deviceId()

	fmt.Printf("Android Device ID: %s\n", id)
	fmt.Printf("Android Device Model: %s\n", model)

	// already locked
	// Lock credential verified successfully

	// unlocked
	// Old password provided but user has no password

	status := verify()

	fmt.Printf("Device status: %s\n", status)

	if strings.HasPrefix(status, "Lock credential verified successfully") {
		listenUnlock(id, model)
	} else if strings.HasPrefix(status, "Old password provided but user has no password") {
		listenLock(id, model)
	}
}

func listenLock(id string, model string) {
	fmt.Println("Listen to Lock")
	listen(id, model, "unlocked", lock)
}

func listenUnlock(id string, model string) {
	fmt.Println("Listen to Unlock")
	listen(id, model, "locked", unlock)
}

func lock(id string, model string, result string) bool {

	if result == "locked" {
		cmd := exec.Command("locksettings", "set-pin", "1234")

		if err := cmd.Start(); err != nil {
			fmt.Println("Error starting set-pin:", err)
			os.Exit(1)
		}

		fmt.Println("Locked")

		listenUnlock(id, model)
		return true
	}
	return false
}

func unlock(id string, model string, result string) bool {

	if result == "unlocked" {
		cmd := exec.Command("locksettings", "clear", "--old", "1234")

		if err := cmd.Start(); err != nil {
			fmt.Println("Error starting set-pin:", err)
			os.Exit(1)
		}

		listenLock(id, model)
		fmt.Println("UnLocked")
		return true
	}
	return false
}
