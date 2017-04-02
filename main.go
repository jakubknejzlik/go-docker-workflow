package main

import "os"

func main() {
	man := NewManager(os.Getenv("CONFIG"))

	man.Start()
}
