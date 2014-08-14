package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func main() {
    var (
        cmd *exec.Cmd
        out bytes.Buffer
    )
    //cmd = exec.Command("telnet", "www.google.com", "80")
    //cmd = exec.Command("ls")
    cmd = exec.Command("nmap -A scanme.nmap.org")
    //cmd.Env = os.Environ()
    cmd.Stdout = &out
    if err := cmd.Run(); err != nil {
        fmt.Println(err)  //exit status -1
        return
    }
    fmt.Println(out.String())
}
