package main

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestManager(t *testing.T) {
	data, _ := ioutil.ReadFile("./test/jobs.yaml")

	man := NewManager(string(data))
	assert.Equal(t, nil, man.RunJob("test"), "run job should not fail")
	assert.NotEqual(t, nil, man.RunJob("test2"), "run nonexisting job should fail")
}
