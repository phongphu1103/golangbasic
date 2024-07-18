package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/joho/godotenv"
)

/* global variable declaration */
var frontend_folder string
var publish_path string

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
}

func mergeToBranchPublish() {
	theme_publish_path := os.Getenv("THEME_PUBLISH_PATH")

	frontend_publish_path := filepath.Join(theme_publish_path, frontend_folder)
	fmt.Println(frontend_publish_path)

	fmt.Println("Checkout branch publish2")
	// run git checkout
	args := []string{"checkout", "publish2"}
	out, errout, abort := executeCommand("git", args, publish_path)
	consoleLog(out, errout, abort)

	fmt.Println("Switch to branch master2")
	args = []string{"checkout", "master2"}
	out, errout, abort = executeCommand("git", args, frontend_publish_path)
	consoleLog(out, errout, abort)

	// run git status
	args = []string{"status"}
	out, errout, abort = executeCommand("git", args, frontend_publish_path)
	consoleLog(out, errout, abort)

	args = []string{"pull", "origin", "master2"}
	out, errout, abort = executeCommand("git", args, frontend_publish_path)
	consoleLog(out, errout, abort)

	// if nothing to commit
	if(abort) {
		return
	}

	fmt.Println("Push submodule")
	// run git add .
	args = []string{"add", "."}
	out, errout, abort = executeCommand("git", args, publish_path)
	consoleLog(out, errout, abort)

	// run git commit
	args = []string{"commit", "-m", "update home"}
	out, errout, abort = executeCommand("git", args, publish_path)
	consoleLog(out, errout, abort)

	// run git push branch publish2
	args = []string{"push", "origin", "publish2"}
	out, errout, abort = executeCommand("git", args, publish_path)
	consoleLog(out, errout, abort)
}

func mergeToBranchTest() {
	mergeToBranch("test")
}

func mergeToBranchStaging() {
	branch_name := "staging/" + frontend_folder
	mergeToBranch(branch_name)
}

func mergeToBranchRelease() {
	branch_name := "release/" + frontend_folder
	mergeToBranch(branch_name)
}

func mergeToBranch(branch_name string) {
	fmt.Printf("Checkout branch %s", branch_name)
	// run git checkout test
	args := []string{"checkout", branch_name}
	out, errout, abort := executeCommand("git", args, publish_path)
	consoleLog(out, errout, abort)

	fmt.Println("Pull branch publish2 to %s", branch_name)
	// run git pull
	args = []string{"pull", "origin", "publish2"}
	out, errout, abort = executeCommand("git", args, publish_path)
	consoleLog(out, errout, abort)
}

func main() {
	// frontend_folder := os.Args[1]
	var step int
	var submodule int
	// get variables from .env
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Select submodule:")
	fmt.Println("1. home")
	fmt.Println("2. admin")
	fmt.Print("Your select: ")
	fmt.Scanf("%d\n", &submodule)

	if submodule == 1 {
		frontend_folder = "home"
	} else if submodule == 2 {
		frontend_folder = "admin"
	} else {
		fmt.Println("Invalid value")
		os.Exit(1)
	}

	publish_path = os.Getenv("PUBLISH_PATH")

	for {
		fmt.Println("Select step:")
		fmt.Println("1. Merge to branch publish2")
		fmt.Println("2. Merge to branch test")
		fmt.Println("3. Merge to branch staging")
		fmt.Println("4. Merge to branch release")
		fmt.Println("0. Exit")
		fmt.Print("Your select: ")
		fmt.Scanf("%d\n", &step)

		switch step {
			case 0:
				os.Exit(1)
			case 1:
				mergeToBranchPublish()
			case 2:
				mergeToBranchTest()
			case 3:
				mergeToBranchStaging()
			case 4:
				mergeToBranchRelease()
			default:
				fmt.Println("def")
		}
	}
}