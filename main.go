package main

import (
	"log"
	"os"
)

func main() {
	man := NewManager(os.Getenv("CONFIG"))

	log.Print(man)
}
