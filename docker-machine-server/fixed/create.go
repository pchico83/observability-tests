package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func create() error {
	cmd := exec.Command("docker-machine", "create", "--driver", "aws")
	err := cmd.Run()
	if err != nil {
		return err
	}
	cmd = exec.Command("docker-machine", "ip")
	var out bytes.Buffer
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		return err
	}
	fmt.Printf("IP: %s\n", out.Bytes())
	return nil
}
