package main

import (
	"io/ioutil"
	"testing"
)

func TestManager(*testing.T) {
	data, _ := ioutil.ReadFile("./test/jobs.yaml")

	man := NewManager(string(data))
	man.RunJob("test")
	man.RunJob("test2")
}
