package main

import (
	"os/exec"
)

func version() error {
	cmd := exec.Command("docker-machine", "version")
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
