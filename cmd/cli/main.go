package main

import (
	"context"
	"flag"
	"log"

	"go-hasher/internal"
)

var (
	path string
	algo string
)

func init() {
	flag.StringVar(&path, "p", "", "a specific file or directory")
	flag.StringVar(&algo, "a", "sha256", "algorithm")
}

func main() {
	flag.Parse()

	hasher, err := internal.NewAction(algo)
	if err != nil {
		log.Fatalf("can't init action: %s", err)
	}

	manager, err := internal.NewManager(context.Background(), path, hasher)
	if err != nil {
		log.Fatalf("can't init action: %s", err)
	}

	manager.Loop()
}
