package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

func getSQLDumpParams(day string) []string {
	db_host := os.Getenv("DB_HOST")
	db_port := os.Getenv("DB_PORT")
	db_name := os.Getenv("DB_NAME")
	db_user := os.Getenv("DB_USER")
	db_pass := os.Getenv("DB_PASS")

	args := []string{}
	ignoreTables := [...]string{"sessions", "t_log_history"}
	args = append(args, "-u", db_user)
	args = append(args, "-p"+db_pass)
	args = append(args, "-h", db_host)
	args = append(args, "--port="+db_port)
	args = append(args, "--protocol=tcp")
	args = append(args, "--default-character-set=utf8")
	// Do not dump table contents
	args = append(args, "--no-data")
	// Do not write CREATE TABLE statements
	// args = append(args, "--no-create-info=TRUE")
	args = append(args, "--skip-triggers")
	// set this option when using mysql older version 5.7
	// args = append(args, "--column-statistics=0")
	args = append(args, "--databases", db_name)

	for _, table := range ignoreTables {
		args = append(args, "--ignore-table="+db_name+"."+table)
	}

	return args
}

func getForFilesParams(times int) []string {
	args := []string{}
	args = append(args, "/P")
	args = append(args, "/S")
	args = append(args, "/D")
	args = append(args, "-"+strconv.Itoa(times))

	return args
}

func main() {
	var formatDate string = "2006-01-02"
	// shorthand: formatDate := "2006-01-02"
	formatDateTime := "2006-01-02 15:04:05"
	formatDateName := "20060102"
	pathSeparator := string(os.PathSeparator)
	// get variables from .env
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}

	local, err := time.LoadLocation("Asia/Saigon")
	if err != nil {
		fmt.Println(err)
	}
	now := time.Now().In(local)
	beforeNow := now.AddDate(0, 0, -7)
	today := now.Format(formatDate)
	dayLastWeek := beforeNow.Format(formatDate)
	fmt.Println(now.Format(formatDate))
	fmt.Println(dayLastWeek)

	// create new folder
	storage_path := os.Getenv("STORAGE_PATH")
	storage_path += formatDateName
	if _, err := os.Stat(storage_path); os.IsNotExist(err) {
		if err := os.Mkdir(storage_path, 0777); err != nil {
			fmt.Println(err)
			return
		}
	}

	// move to mysql directory
	mysql_path := os.Getenv("MYSQL_PATH")
	if err := os.Chdir(mysql_path); err != nil {
		fmt.Println(err)
		return
	}
	// get current directory
	s, _ := os.Getwd()
	fmt.Println(s)
	// making display version
	cmd := exec.Command("mysqldump", "--version")
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return
	}
	fmt.Println(out.String())

	// making dump db
	args := getSQLDumpParams((today))
	fmt.Println(args)
	cmd = exec.Command("mysqldump", args...)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println(err)
	}

	if err := cmd.Start(); err != nil {
		fmt.Println(err)
	}
	// read
	bytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		fmt.Println(err)
	}
	// write
	storageFile := storage_path + pathSeparator + "Dump-" + formatDateName + ".sql"
	err = ioutil.WriteFile(storageFile, bytes, 0777)
	if err != nil {
		fmt.Println(err)
	}

	end := time.Now().In(local)
	fmt.Println("Completed at " + end.Format(formatDateTime))

}
