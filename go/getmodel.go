package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
	"net/http"
	"time"
	"io/ioutil"
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
}


func startFetching() {
	fetchAPI()

	interval := 30 * time.Second
	ticker := time.NewTicker(interval)

	go func() {
		for {
			select {
			case <-ticker.C:
				fetchAPI()
			}
		}
	}()
	select {}
}

func fetchAPI() {
	apiURL := "YourAPIURL"

	resp, err := http.Get(apiURL)
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}
		fmt.Printf("Data fetched successfully:\n%s\n", string(body))
	} else {
		fmt.Printf("Error: Received a non-success status code: %d\n", resp.StatusCode)
	}
}

func getAndroidDeviceModel() (string, error) {
	cmd := exec.Command( "getprop", "ro.product.model")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	deviceModel := strings.TrimSpace(string(output))
	return deviceModel, nil
}

func getAndroidDeviceId() (string, error) {
	cmd := exec.Command( "settings", "get", "secure", "android_id")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	deviceModel := strings.TrimSpace(string(output))
	return deviceModel, nil
}