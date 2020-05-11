package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/davidMGS/codenames-pictures"
)

const DEFAULT_PORT = "9001"

func main() {
	if len(os.Args) > 2 {
		fmt.Fprintf(os.Stderr, "Too many arguments\n")
		os.Exit(1)
	}

	var port string
	if len(os.Args) == 2 {
		port = os.Args[1]
	} else {
		port = DEFAULT_PORT
	}

	rand.Seed(time.Now().UnixNano())

	server := &codenames.Server{
		Server: http.Server{
			Addr: ":" + port,
		},
        AssetsPath: "assets_codenames",
	}
	fmt.Printf("Starting server on port %s...\n", port)
	if err := server.Start(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
	}
}
