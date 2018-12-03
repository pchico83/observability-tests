package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func restart() error {
	cmd := exec.Command("docker-machine", "restart", "elastest")
	err := cmd.Run()
	if err != nil {
		return err
	}
	cmd = exec.Command("docker-machine", "regenerate-certs", "elastest")
	err = cmd.Run()
	if err != nil {
		return err
	}
	cmd = exec.Command("docker-machine", "ip", "elastest")
	var out bytes.Buffer
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		return err
	}
	fmt.Printf("elastest IP: %s\n", out.Bytes())
	return nil
}
