package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/joho/godotenv"
)

func executeCommand(tool string, args []string, path string) (string, string, bool) {
	cmd := exec.Command(tool, args...)
	cmd.Dir = path
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	var errout string
	var abort bool = false

	if err := cmd.Run(); err != nil {
		errout = fmt.Sprint(err) + ": " + stderr.String()
		abort = true
	}

	return stdout.String(), errout, abort
}

func consoleLog(out string, errout string, abort bool) {
	fmt.Println(out)
	fmt.Println(errout)
	if(abort) {
		os.Exit(1)
	}
}

func main() {
	// frontend_folder := os.Args[1]
	var frontend_folder string
	var value int

	fmt.Println("Select submodule:")
	fmt.Println("1. home")
	fmt.Println("2. admin")
	fmt.Print("Your select: ")
	fmt.Scanf("%d", &value)

	if value == 1 {
		frontend_folder = "home"
	} else if value == 2 {
		frontend_folder = "admin"
	} else {
		fmt.Println("Invalid value")
		os.Exit(1)
	}

	// get variables from .env
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}

	publish_path := os.Getenv("PUBLISH_PATH")
	theme_publish_path := os.Getenv("THEME_PUBLISH_PATH")

	frontend_publish_path := filepath.Join(theme_publish_path, frontend_folder)
	fmt.Println(frontend_publish_path)

	args := []string{"checkout", "master2"}
	out, errout, abort := executeCommand("git", args, frontend_publish_path)
	consoleLog(out, errout, abort)

	// run git status
	args = []string{"status"}
	out, errout, abort = executeCommand("git", args, frontend_publish_path)
	consoleLog(out, errout, abort)

	args = []string{"pull", "origin", "master2"}
	out, errout, abort = executeCommand("git", args, frontend_publish_path)
	consoleLog(out, errout, abort)

	// run git push branch publish2
	args = []string{"push", "origin", "publish2"}
	out, errout, abort = executeCommand("git", args, publish_path)
	consoleLog(out, errout, abort)

}