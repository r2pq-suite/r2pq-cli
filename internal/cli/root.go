package cli

import (
    "fmt"
)

func Execute(args []string) int {
    if len(args) == 0 {
        fmt.Println("r2pq-cli: no command provided")
        return 1
    }

    switch args[0] {
    case "keygen":
        return runKeygen()
    default:
        fmt.Println("r2pq-cli: unknown command:", args[0])
        return 1
    }
}
