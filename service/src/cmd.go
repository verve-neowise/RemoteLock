package main

import (
	"fmt"
	"os/exec"
)

func cmd(name string, arg ...string) {
	cmdEcho := exec.Command(name, arg...)

	if err := cmdEcho.Start(); err != nil {
		fmt.Printf("Error run command script %s %s", name, err)
	}
}
