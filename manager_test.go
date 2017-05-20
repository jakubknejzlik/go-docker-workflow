package main

import (
	"encoding/base64"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestManager(t *testing.T) {
	data, _ := ioutil.ReadFile("./test/jobs.json")

	man := NewManager(base64.StdEncoding.EncodeToString(data))

	assert.Equal(t, 1, len(man.Conf.Jobs), "should contain 1 job")
	assert.Equal(t, nil, man.RunJob("test"), "run job should not fail")
	assert.NotEqual(t, nil, man.RunJob("test2"), "run nonexisting job should fail")
}
