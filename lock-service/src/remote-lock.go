package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

func main() {

	model, err := getAndroidDeviceModel()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	id, err := getAndroidDeviceId()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Printf("Android Device ID: %s\n", id)
	fmt.Printf("Android Device Model: %s\n", model)


	// already locked
	// Lock credential verified successfully

	// unlocked 
	// Old password provided but user has no password

	status:= verify()

	fmt.Printf("Device status: %s\n", status)

	if strings.HasPrefix(status, "Lock credential verified successfully") {
		listenForUnLock(id, model)
	} else if strings.HasPrefix(status, "Old password provided but user has no password") {
		listenForLock(id, model)
	}
}

func listenForLock(id string, model string) {

	interval := 10 * time.Second
	ticker := time.NewTicker(interval)

	fmt.Println("listen lock")

	go func() {
		for {
			select {
			case <-ticker.C:
				result := fetchAPI(id, model)
				fmt.Printf("result: %s\n", result)

				if (result == "locked") {
					cmd := exec.Command("locksettings", "set-pin", "1234")

					if err := cmd.Start(); err != nil {
						fmt.Println("Error starting set-pin:", err)
						os.Exit(1)
					}
				
					fmt.Println("Locked")

					listenForUnLock(id, model)
					break
				}
			}
		}
	}()
	select {}
}

func listenForUnLock(id string, model string) {

	interval := 10 * time.Second
	ticker := time.NewTicker(interval)

	fmt.Println("listen unlock")

	go func() {
		for {
			select {
			case <-ticker.C:
				result := fetchAPI(id, model)
				fmt.Printf("result: %s\n", result)

				if (result == "unlocked") {
					cmd := exec.Command("locksettings", "clear", "--old", "1234")

					if err := cmd.Start(); err != nil {
						fmt.Println("Error starting set-pin:", err)
						os.Exit(1)
					}

					listenForLock(id, model)
					fmt.Println("UnLocked")
					break
				}
			}
		}
	}()
	select {}
}

func verify() (string) {
	cmd := exec.Command("locksettings", "verify", "--old", "1234")
	outPipe, _ := cmd.StdoutPipe()

	if err := cmd.Start(); err != nil {
		fmt.Println("Error starting command:", err)
		os.Exit(1)
	}

	reader := bufio.NewReader(outPipe)
	line, err := reader.ReadString('\n')

	if (err != nil) {
		fmt.Println("Error reading result:", err)
	}

	if err := cmd.Wait(); err != nil {
		fmt.Println("Error waiting for command to complete:", err)
	}

	fmt.Printf("verify: %s\n", line)

	return line
}

func fetchAPI(id string, model string) (string) {
		httpRequest := "GET /status?id=" + id + "&model=" + model + " HTTP/1.0\r\n\r\n"

		fmt.Println("Request: " + httpRequest)

		cmd := exec.Command("nc", "5.181.109.130", "3001")
	
		outPipe, _ := cmd.StdoutPipe()
		inPipe, _ := cmd.StdinPipe()

		if err := cmd.Start(); err != nil {
			fmt.Println("Error starting command:", err)
			os.Exit(1)
		}

		_, e := inPipe.Write([]byte(httpRequest))

		if (e != nil) {
			fmt.Println("Error starting command:", e)
		}

		inPipe.Close()

		var lastLine string

		go func(p io.ReadCloser) {
			reader := bufio.NewReader(outPipe)
			line, err := reader.ReadString('\n')
			for err == nil {
				line, err = reader.ReadString('\n')
			}
			lastLine = line
		}(outPipe)
		

		if err := cmd.Wait(); err != nil {
			fmt.Println("Error waiting for command to complete:", err)
		}

		fmt.Println("Success completed:", lastLine)

		return lastLine
}

func getAndroidDeviceModel() (string, error) {
	cmd := exec.Command("getprop", "ro.product.model")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	deviceModel := strings.TrimSpace(string(output))
	return deviceModel, nil
}

func getAndroidDeviceId() (string, error) {
	cmd := exec.Command("settings", "get", "secure", "android_id")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	deviceModel := strings.TrimSpace(string(output))
	return deviceModel, nil
}
