// cmd/etherealnexus/main.go
package main

import (
"flag"
"log"
"os"

"etherealnexus/internal/etherealnexus"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := etherealnexus.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
