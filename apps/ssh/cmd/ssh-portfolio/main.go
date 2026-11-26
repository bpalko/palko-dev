package main

import (
	"os"

	"github.com/charmbracelet/log"

	"github.com/bpalko/palko-dev/ssh/internal/server"
)

func main() {
	if err := server.Run(server.FromEnv()); err != nil {
		log.Error("ssh portfolio exited", "error", err)
		os.Exit(1)
	}
}
