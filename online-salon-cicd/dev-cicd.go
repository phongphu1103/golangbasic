package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/joho/godotenv"
	// import local package
	"online-salon-cicd/utils"
)


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
	out, errout, abort := common.ExecuteCommand("git", args, home_dev_path)
	common.ConsoleLog(out, errout, abort)

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
		out, errout, abort = common.ExecuteCommand("npm", args, home_dev_path)
		common.ConsoleLog(out, errout, abort)
	}

	fmt.Println("Update what will be committed")
	// run git add .
	args = []string{"add", "."}
	out, errout, abort = common.ExecuteCommand("git", args, home_dev_path)
	common.ConsoleLog(out, errout, abort)

	// run git commit
	args = []string{"commit", "-m", "update recruit"}
	out, errout, abort = common.ExecuteCommand("git", args, home_dev_path)
	common.ConsoleLog(out, errout, abort)

	// run git push
	args = []string{"push", "origin", "ft/new-recruit"}
	out, errout, abort = common.ExecuteCommand("git", args, home_dev_path)
	common.ConsoleLog(out, errout, abort)

	fmt.Println("Updated branch ft/new-recruit")

	fmt.Println("Switch branch dev/main2")
	// run git checkout dev/main2
	args = []string{"checkout", "dev/main2"}
	out, errout, abort = common.ExecuteCommand("git", args, home_dev_path)
	common.ConsoleLog(out, errout, abort)

	fmt.Println("Pull branch ft/new-recruit to dev/main2")
	// run git pull
	args = []string{"pull", "origin", "ft/new-recruit"}
	out, errout, abort = common.ExecuteCommand("git", args, home_dev_path)
	common.ConsoleLog(out, errout, abort)

	fmt.Println("Push branch dev/main2")
	// run git push
	args = []string{"push", "origin", "dev/main2"}
	out, errout, abort = common.ExecuteCommand("git", args, home_dev_path)
	common.ConsoleLog(out, errout, abort)

	fmt.Println("Switch branch master2")
	// run git checkout master2
	args = []string{"checkout", "master2"}
	out, errout, abort = common.ExecuteCommand("git", args, home_dev_path)
	common.ConsoleLog(out, errout, abort)

	fmt.Println("Pull branch dev/main2 to master2")
	// run git pull
	args = []string{"pull", "origin", "dev/main2"}
	out, errout, abort = common.ExecuteCommand("git", args, home_dev_path)
	common.ConsoleLog(out, errout, abort)

	fmt.Println("Push branch master2")

	args = []string{"push", "origin", "master2"}
	out, errout, abort = common.ExecuteCommand("git", args, home_dev_path)
	common.ConsoleLog(out, errout, abort)
}