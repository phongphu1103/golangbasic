package main

import (
	"fmt"
	"os"
	"runtime"
	"path/filepath"

	"github.com/joho/godotenv"
	// import local package
	"online-salon-cicd/utils"
)

/* global variable declaration */
var frontend_folder string
var theme_publish_path string
var publish_path string

func mergeToBranchPublish() {
	frontend_publish_path := filepath.Join(theme_publish_path, frontend_folder)
	fmt.Println(frontend_publish_path)

	fmt.Println("Checkout branch publish2")
	// run git checkout
	args := []string{"checkout", "publish2"}
	out, errout, abort := common.ExecuteCommand("git", args, publish_path)
	common.ConsoleLog(out, errout, abort)

	fmt.Println("Switch to branch master2")
	args = []string{"checkout", "master2"}
	out, errout, abort = common.ExecuteCommand("git", args, frontend_publish_path)
	common.ConsoleLog(out, errout, abort)

	// run git status
	args = []string{"status"}
	out, errout, abort = common.ExecuteCommand("git", args, frontend_publish_path)
	common.ConsoleLog(out, errout, abort)

	args = []string{"pull", "origin", "master2"}
	out, errout, abort = common.ExecuteCommand("git", args, frontend_publish_path)
	common.ConsoleLog(out, errout, abort)

	// if nothing to commit
	if(abort) {
		return
	}

	fmt.Println("Push submodule")
	// run git add .
	args = []string{"add", "."}
	out, errout, abort = common.ExecuteCommand("git", args, publish_path)
	common.ConsoleLog(out, errout, abort)

	// run git commit
	args = []string{"commit", "-m", "update home"}
	out, errout, abort = common.ExecuteCommand("git", args, publish_path)
	common.ConsoleLog(out, errout, abort)

	// run git push branch publish2
	args = []string{"push", "origin", "publish2"}
	out, errout, abort = common.ExecuteCommand("git", args, publish_path)
	common.ConsoleLog(out, errout, abort)
}

func mergeToBranchTest() {
	mergeToBranch("test")
}

func buildBranchTest() {
	buildBranch("test")
}

func mergeToBranchStaging() {
	branch_name := "staging/" + frontend_folder
	mergeToBranch(branch_name)
}

func buildBranchStaging() {
	buildBranch("staging")
}

func mergeToBranchRelease() {
	branch_name := "release/" + frontend_folder
	mergeToBranch(branch_name)
}

func mergeToBranch(branch_name string) {
	fmt.Printf("Checkout branch %s", branch_name)
	// run git checkout test
	args := []string{"checkout", branch_name}
	out, errout, abort := common.ExecuteCommand("git", args, publish_path)
	common.ConsoleLog(out, errout, abort)

	fmt.Printf("Pull branch publish2 to %s", branch_name)
	// run git pull
	args = []string{"pull", "origin", "publish2"}
	out, errout, abort = common.ExecuteCommand("git", args, publish_path)
	common.ConsoleLog(out, errout, abort)

	fmt.Printf("Push branch %s", branch_name)
	// run git push
	args = []string{"push", "origin", branch_name}
	out, errout, abort = common.ExecuteCommand("git", args, publish_path)
	common.ConsoleLog(out, errout, abort)
}

func buildBranch(branch_name string) {
	var convention string
	var tool string
	if runtime.GOOS == "windows" {
		convention = "cmd.exe"
		tool = "copy"
	} else {
		convention = "bash"
		tool = "cp"
	}

	env_file := fmt.Sprintf("env.backup.%s", branch_name)
	resource_path := os.Getenv("ENV_RESOURCE_PATH")
	source_path := filepath.Join(resource_path, frontend_folder, env_file)
	destination_path := filepath.Join(theme_publish_path, frontend_folder, ".env")

	args := []string{}
	args = append(args, "/c")
	args = append(args, tool)
	args = append(args, source_path)
	args = append(args, destination_path)

	out, errout, abort := common.ExecuteCommand(convention, args, publish_path)
	common.ConsoleLog(out, errout, abort)
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
	theme_publish_path = os.Getenv("THEME_PUBLISH_PATH")

	for {
		fmt.Println("Select step:")
		fmt.Println("1. Merge to branch publish2")
		fmt.Println("2. Merge to branch test")
		fmt.Println("3. Build branch test")
		fmt.Println("4. Merge to branch staging")
		fmt.Println("5. Build branch staging")
		fmt.Println("6. Merge to branch release")
		fmt.Println("0. Exit")
		fmt.Print("Your select: ")
		fmt.Scanf("%d\n", &step)

		// the break statement is provided automatically in Go
		switch step {
			case 0:
				os.Exit(1)
			case 1:
				mergeToBranchPublish()
			case 2:
				mergeToBranchTest()
			case 3:
				buildBranchTest()
			case 4:
				mergeToBranchStaging()
			case 5:
				buildBranchStaging()
			case 6:
				mergeToBranchRelease()
			default:
				fmt.Println("def")
		}
	}
}