package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestManager(t *testing.T) {
	man := NewManagerFromYamlFile("./test/jobs.yml")
	assert.Equal(t, 1, len(man.rootJob.Jobs), "should contain 1 job")
	assert.NotNil(t, man.FindJob("test"), "should find root job")
	assert.NotNil(t, man.FindJob("test/subtest1"), "should find child job")
	assert.NotNil(t, man.FindJob("test/subtest1/subsubtest2"), "should find child job")
	assert.NotEqual(t, nil, man.RunJob("test2"), "run nonexisting job should fail")
	assert.Nil(t, man.RunJob("test"), "run job should not fail")
	assert.Nil(t, man.Run(), "run root job should not fail")
}
