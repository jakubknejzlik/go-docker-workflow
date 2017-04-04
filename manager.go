package main

import (
	"fmt"
	"log"
	"time"

	"github.com/robfig/cron"

	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	Jobs []*Job `yaml:"jobs"`
}

type Manager struct {
	Conf Config
}

func NewManager(config string) Manager {
	var conf Config
	if err := yaml.Unmarshal([]byte(config), &conf); err != nil {
		log.Fatal(err)
	}

	for _, job := range conf.Jobs {
		processJob(job)
	}

	return Manager{conf}
}

func (m *Manager) PullJobImage(jobName string) error {
	job := m.GetJob(jobName)

	if job == nil {
		return fmt.Errorf("job %s not found", jobName)
	}

	return job.PullImage()
}

func (m *Manager) GetJob(name string) *Job {
	for _, job := range m.Conf.Jobs {
		if job.Name == name {
			return job
		}
	}
	return nil
}

func (m *Manager) RunJob(jobName string) error {
	job := m.GetJob(jobName)

	if job == nil {
		return fmt.Errorf("job %s not found", jobName)
	}

	job.Run()
	return nil
}

func (m *Manager) Start() error {
	_, err := m.StartCrons()

	if err != nil {
		return err
	}

	log.Println("starting manager")

	for _, job := range m.Conf.Jobs {
		if job.Cron == "" {
			job.Run()
		}
	}

	for {
		time.Sleep(1 * time.Second)
	}
}

func (m *Manager) StartCrons() (*cron.Cron, error) {
	c := cron.New()

	for _, job := range m.Conf.Jobs {
		if job.Cron != "" {
			c.AddJob(job.Cron, job)
		}
	}

	c.Start()

	return c, nil
}
