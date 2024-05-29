package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

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
	// get variables from .env
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}

	home_dev_path := os.Getenv("HOME_DEV_PATH")
	fmt.Println("Check git branch status")

	// run git status
	args := []string{"status"}
	out, errout, abort := executeCommand("git", args, home_dev_path)
	consoleLog(out, errout, abort)

	var answer string
	fmt.Print("Run npm install (y/n): ")
	fmt.Scan(&answer)

	if runtime.GOOS == "windows" {
		answer = strings.TrimRight(answer, "\r\n")
	} else {
		answer = strings.TrimRight(answer, "\n")
	}

	if strings.ToLower(answer) == "y" || strings.ToLower(answer) == "yes" {
		args = []string{"i"}
		out, errout, abort = executeCommand("npm", args, home_dev_path)
		consoleLog(out, errout, abort)
	}

	fmt.Println("Update what will be committed")
	// run git add .
	args = []string{"add", "."}
	out, errout, abort = executeCommand("git", args, home_dev_path)
	consoleLog(out, errout, abort)

	// run git commit
	args = []string{"commit", "-m", "'update recruit'"}
	out, errout, abort = executeCommand("git", args, home_dev_path)
	consoleLog(out, errout, abort)

	// run git push
	args = []string{"push", "origin", "ft/new-recruit"}
	out, errout, abort = executeCommand("git", args, home_dev_path)
	consoleLog(out, errout, abort)

	fmt.Println("Updated branch ft/new-recruit")

	fmt.Println("Switch branch dev/main2")
	// run git checkout dev/main2
	args = []string{"checkout", "dev/main2"}
	out, errout, abort = executeCommand("git", args, home_dev_path)
	consoleLog(out, errout, abort)

	fmt.Println("Pull branch ft/new-recruit to dev/main2")
	// run git pull
	args = []string{"pull", "origin", "ft/new-recruit"}
	out, errout, abort = executeCommand("git", args, home_dev_path)
	consoleLog(out, errout, abort)

	fmt.Println("Push branch dev/main2")
	// run git push
	args = []string{"push", "origin", "dev/main2"}
	out, errout, abort = executeCommand("git", args, home_dev_path)
	consoleLog(out, errout, abort)

	fmt.Println("Switch branch master2")
	// run git checkout master2
	args = []string{"checkout", "master2"}
	out, errout, abort = executeCommand("git", args, home_dev_path)
	consoleLog(out, errout, abort)

	fmt.Println("Pull branch dev/main2 to master2")
	// run git pull
	args = []string{"pull", "origin", "dev/main2"}
	out, errout, abort = executeCommand("git", args, home_dev_path)
	consoleLog(out, errout, abort)

	fmt.Println("Push branch master2")

	args = []string{"push", "origin", "master2"}
	out, errout, abort = executeCommand("git", args, home_dev_path)
	consoleLog(out, errout, abort)
}