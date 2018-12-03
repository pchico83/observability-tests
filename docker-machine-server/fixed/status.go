package main

import (
	"os/exec"
)

func status() error {
	cmd := exec.Command("docker-machine", "version", "elastest")
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
