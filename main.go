package main

import (
    "os"

    "github.com/r2pq-suite/r2pq-cli/internal/cli"
)

func main() {
    os.Exit(cli.Execute(os.Args[1:]))
}
