package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("curl", "-fsSL", "https://get.deta.dev/space-cli.sh")
	o, _ := cmd.Output()
	os.OpenFile("install.sh", os.O_CREATE|os.O_WRONLY, 0777)
	os.WriteFile("install.sh", o, 0777)
	cmd = exec.Command("sh", "install.sh")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
	os.Remove("install.sh")
	bin := os.Getenv("HOME") + "/.detaspace/bin/space"
	SPACE_ACCESS_TOKEN := os.Getenv("SPACE_ACCESS_TOKEN")
	SPACE_PROJECT_ID := os.Getenv("SPACE_PROJECT_ID")
	if SPACE_ACCESS_TOKEN == "" {
		fmt.Println("panic: set `SPACE_ACCESS_TOKEN` in environment")
		os.Exit(1)
	}
	if SPACE_PROJECT_ID == "" {
		fmt.Println("panic: set `SPACE_PROJECT_ID` in environment")
		os.Exit(1)
	}
	cmd = exec.Command(bin, "link", fmt.Sprintf("--id=%s", SPACE_PROJECT_ID))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
	cmd = exec.Command(bin, "push")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
