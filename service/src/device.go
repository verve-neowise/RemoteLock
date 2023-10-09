package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func deviceModel() string {

	// for test
	if runtime.GOOS == "darwin" {
		return "test-model"
	}

	cmd := exec.Command("getprop", "ro.product.model")
	output, err := cmd.CombinedOutput()

	if err != nil {
		return ""
	}

	deviceModel := strings.TrimSpace(string(output))
	return deviceModel
}

func deviceId() string {

	if runtime.GOOS == "darwin" {
		return "test-id"
	}

	cmd := exec.Command("getprop", "ro.serialno")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return ""
	}
	deviceModel := strings.TrimSpace(string(output))
	return deviceModel
}

func verify() string {

	// for test
	if runtime.GOOS == "darwin" {
		return "Lock credential verified successfully"
	}

	cmd := exec.Command("locksettings", "verify", "--old", "1234")
	outPipe, _ := cmd.StdoutPipe()

	if err := cmd.Start(); err != nil {
		fmt.Println("Error starting command:", err)
		os.Exit(1)
	}

	reader := bufio.NewReader(outPipe)
	line, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Error reading result:", err)
	}

	if err := cmd.Wait(); err != nil {
		fmt.Println("Error waiting for command to complete:", err)
	}

	fmt.Printf("verify: %s\n", line)

	return line
}
