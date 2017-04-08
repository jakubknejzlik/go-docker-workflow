package main

import "os"

func main() {
	config := os.Getenv("CONFIG")
	if len(os.Args) > 1 {
		config = os.Args[1]
	}

	man := NewManager(config)

	man.Start()
}
