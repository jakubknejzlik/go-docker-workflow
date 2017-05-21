package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/robfig/cron"
)

// Manager ...
type Manager struct {
	rootJob Job
}

// NewManager ...
func NewManager(config string) Manager {
	var rootJob Job
	decodedConfig, _ := base64.StdEncoding.DecodeString(config)
	if err := json.Unmarshal(decodedConfig, &rootJob); err != nil {
		fmt.Println(err)
		fmt.Println(config)
	}

	rootJob.IsRoot = true

	processJob(&rootJob)

	return Manager{rootJob}
}

// PullJobImage ...
func (m *Manager) PullJobImage(jobName string) error {
	job := m.FindJob(jobName)

	if job == nil {
		return fmt.Errorf("job %s not found", jobName)
	}

	return job.PullImage()
}

// FindJob ...
func (m *Manager) FindJob(name string) *Job {
	return m.rootJob.FindSubJob(strings.Split(name, "/"))
}

// RunJob ...
func (m *Manager) RunJob(jobName string) error {
	job := m.FindJob(jobName)

	if job == nil {
		return fmt.Errorf("job %s not found", jobName)
	}

	job.Run()
	return nil
}

// Start ...
func (m *Manager) Start() error {
	_, err := m.StartCrons()

	if err != nil {
		return err
	}

	fmt.Println("starting cronjobs")

	for _, job := range m.rootJob.Jobs {
		if job.Cron == "" {
			job.Run()
		}
	}

	for {
		time.Sleep(1 * time.Second)
	}
}

// Run ...
func (m *Manager) Run() error {
	fmt.Println("running jobs")

	m.rootJob.Run()

	return nil
}

// StartCrons ...
func (m *Manager) StartCrons() (*cron.Cron, error) {
	c := cron.New()

	for _, job := range m.rootJob.Jobs {
		if job.Cron != "" {
			fmt.Printf("Starting cron %s for job %s\n", job.Cron, job.GetFullname())
			c.AddJob(job.Cron, job)
		}
	}

	c.Start()

	return c, nil
}
