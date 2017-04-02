package main

import (
	"fmt"
	"log"

	"github.com/robfig/cron"

	yaml "gopkg.in/yaml.v2"
)

type Manager struct {
	Conf Config
}

func NewManager(config string) Manager {
	var conf Config
	if err := yaml.Unmarshal([]byte(config), &conf); err != nil {
		log.Fatal(err)
	}

	return Manager{conf}
}

func (m *Manager) PullJobImage(jobName string) error {
	job := m.Conf.Jobs[jobName]

	if job == nil {
		return fmt.Errorf("job %s not found", jobName)
	}

	return job.PullImage()
}

func (m *Manager) RunJob(jobName string) error {
	job := m.Conf.Jobs[jobName]

	if job == nil {
		return fmt.Errorf("job %s not found", jobName)
	}

	job.Run()
	return nil
}

func (m *Manager) StartCrons() (*cron.Cron, error) {
	c := cron.New()

	for _, job := range m.Conf.Jobs {
		if job.Cron != "" {
			c.AddJob(job.Cron, job)
		}
	}

	return c, nil
}
