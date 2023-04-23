package main

import (
	"log"

	"github.com/jcyamacho/gh-stars/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatalf("Error executing CLI: %v", err)
	}
}
