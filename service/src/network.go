package main

import (
	"bufio"
	"fmt"
	"os/exec"
	"time"
)

func listen(id string, model string, status string, callback func(string, string, string) bool) {

	interval := 10 * time.Second
	ticker := time.NewTicker(interval)

	go func() {
		for {
			select {
			case <-ticker.C:
				result := fetchAPI(id, model, status)
				fmt.Printf("result: %s\n", result)

				if callback(id, model, result) {
					break
				}
			}
		}
	}()
	select {}
}

func fetchAPI(id string, model string, status string) string {
	httpRequest := "GET /device/status?id=" + id + "&model=" + model + "&status=" + status + " HTTP/1.0"

	fmt.Println("Request: " + httpRequest)

	cmd := exec.Command("nc", "5.181.109.130", "3001")

	outPipe, _ := cmd.StdoutPipe()
	inPipe, _ := cmd.StdinPipe()

	if err := cmd.Start(); err != nil {
		fmt.Println("Error starting command:", err)
	}

	_, e := inPipe.Write([]byte(httpRequest + "\r\n\r\n"))

	if e != nil {
		fmt.Println("Error starting command:", e)
	}

	lastLine := ""

	scanner := bufio.NewScanner(outPipe)
	for scanner.Scan() {
		lastLine = scanner.Text()
	}

	inPipe.Close()

	if err := cmd.Wait(); err != nil {
		fmt.Println("Error waiting for command to complete:", err)
	}

	cmdEcho := exec.Command("sh", "process.sh", httpRequest + "\n")

	if err := cmdEcho.Start(); err != nil {
		fmt.Println("Error waiting for write to process.txt to complete:", err)
	}

	fmt.Println("Request completed:", lastLine)

	return lastLine
}
