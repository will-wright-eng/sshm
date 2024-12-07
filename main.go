package main

import (
    "log"
    "os"

    "github.com/will-wright-eng/sshm/cmd"
)

func main() {
    if err := cmd.Execute(); err != nil {
        log.Printf("Error: %v\n", err)
        os.Exit(1)
    }
}
