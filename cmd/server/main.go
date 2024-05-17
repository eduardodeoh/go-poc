package main

import (
	"fmt"
	"os"

	"github.com/eduardodeoh/go-proc/internal/core/config"
)

func main() {
	// Initialize Config
	appConfig, err := config.New()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to initialize App Config: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(appConfig)

}
