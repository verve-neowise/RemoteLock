package main

import (
	"fmt"
	"runtime"
)

func main() {

	fmt.Println(runtime.GOOS)

	model := deviceModel()
	id := deviceId()

	fmt.Printf("Android Device ID: %s\n", id)
	fmt.Printf("Android Device Model: %s\n", model)

	listenStatus()
}

func listenStatus() {
	fmt.Println("Listen to Status")
	listen("unknown", status)
}

func listenLock() {
	fmt.Println("Listen to Lock")
	listen("unlocked", lock)
}

func listenUnlock() {
	fmt.Println("Listen to Unlock")
	listen("locked", unlock)
}

func status(result string) bool {
	if result == "locked" {
		lock(result)
	} else if result == "unlocked" {
		unlock(result)
	}
	return true
}

func lock(result string) bool {

	if result == "locked" {
		
		cmd("cmd", "package", "set-home-activity", "\"ae.axcapital.lockapp/.MainActivity\"")
		cmd("am", "start", "-n", "\"ae.axcapital.lockapp/.MainActivity\"")

		fmt.Println("Locked")

		listenUnlock()
		return true
	}
	return false
}

func unlock(result string) bool {

	if result == "unlocked" {
		
		cmd("am", "force-stop", "ae.axcapital.lockapp")
		cmd("cmd", "package", "set-home-activity", "com.sec.android.app.launcher/com.sec.android.app.launcher.activities.LauncherActivity")

		listenLock()
		fmt.Println("UnLocked")
		return true
	}
	return false
}
