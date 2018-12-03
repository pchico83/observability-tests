package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

func create() error {
	awsRegion := os.Getenv("AWS_REGION")
	awsAccessKey := os.Getenv("AWS_ACCESS_KEY")
	awsSecretKey := os.Getenv("AWS_SECRET_KEY")
	cmd := exec.Command(
		"docker-machine",
		"create",
		"--driver",
		"amazonec2",
		"--amazonec2-region",
		awsRegion,
		"--amazonec2-access-key",
		awsAccessKey,
		"--amazonec2-secret-key",
		awsSecretKey,
		"elastest",
	)
	err := cmd.Run()
	if err != nil {
		fmt.Println("ERROR", err)
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
