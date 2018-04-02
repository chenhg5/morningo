package main

import (
  "bytes"
  "fmt"
  "log"
  "os/exec"
)

func execShell(s string) {
    cmd := exec.Command("/bin/bash", "-c", s)
    var out bytes.Buffer

    cmd.Stdout = &out
    err := cmd.Run()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("%s", out.String())
}

func main() {
    execShell("uname -a")
}