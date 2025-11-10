package cli

import (
    "crypto/rand"
    "encoding/hex"
    "fmt"
)

func runKeygen() int {
    buf := make([]byte, 32)

    _, err := rand.Read(buf)
    if err != nil {
        fmt.Println("error generating key:", err)
        return 1
    }

    fmt.Println("Generated R2PQ key:", hex.EncodeToString(buf))
    return 0
}
