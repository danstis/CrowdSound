package main

import (
	"log"

	"github.com/danstis/CrowdSound/internal/version"
)

// Main entry point for the app.
func main() {
	log.Printf("Version %q", version.Version)
}
