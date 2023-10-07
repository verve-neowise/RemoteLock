package main

import (
	"fmt"
	"os/exec"
	"runtime"
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

	status := getStatus()

	fmt.Printf("Device status: %s\n", status)

	if status == "locked" {
		listenUnlock(id, model)
	} else if status == "unlocked" {
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
		
		cmdEcho := exec.Command("sh", "unlock_script.sh")

		if err := cmdEcho.Start(); err != nil {
			fmt.Println("Error waiting complete unlock script:", err)
		}

		setLocked()

		fmt.Println("Locked")

		listenUnlock(id, model)
		return true
	}
	return false
}

func unlock(id string, model string, result string) bool {

	if result == "unlocked" {
		
		cmdEcho := exec.Command("sh", "unlock_script.sh")

		if err := cmdEcho.Start(); err != nil {
			fmt.Println("Error waiting complete unlock script:", err)
		}

		setUnLocked()

		listenLock(id, model)
		fmt.Println("UnLocked")
		return true
	}
	return false
}
