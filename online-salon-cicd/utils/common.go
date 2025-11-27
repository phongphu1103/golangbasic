package common

import (
    "bytes"
    "fmt"
    "os"
    "os/exec"
)

func ExecuteCommand(tool string, args []string, path string) (string, string, bool) {
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

func ConsoleLog(out string, errout string, abort bool) {
    fmt.Println(out)
    fmt.Println(errout)
    if (abort) {
        os.Exit(1)
    }
}