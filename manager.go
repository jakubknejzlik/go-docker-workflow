package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/ghodss/yaml"
	"github.com/robfig/cron"
)

// Manager ...
type Manager struct {
	rootJob Job
}

// NewManagerFromBase64 ...
func NewManagerFromBase64(config string) Manager {
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

// NewManagerFromYamlFile ...
func NewManagerFromYamlFile(file string) Manager {
	var rootJob Job
	data, _ := ioutil.ReadFile(file)

	if err := yaml.Unmarshal(data, &rootJob); err != nil {
		fmt.Println(err)
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

	return job.RunStrict()
}

// Start ...
func (m *Manager) Start() error {
	fmt.Println("starting cronjobs")

	_, err := m.StartCrons()
	if err != nil {
		return err
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

	timezone := os.Getenv("TIMEZONE")
	if timezone == "" {
		timezone = "UTC"
	}
	location, _ := time.LoadLocation(timezone)

	c := cron.NewWithLocation(location)

	if err := m.startCronsJob(&m.rootJob, c); err != nil {
		return c, err
	}

	c.Start()

	return c, nil
}

func (m *Manager) startCronsJob(j *Job, c *cron.Cron) error {
	if j.Cron != "" {
		fmt.Printf("Starting cron %s for job %s\n", j.Cron, j.GetFullname())
		c.AddJob(j.Cron, j)
	}
	for _, job := range j.Jobs {
		if err := m.startCronsJob(job, c); err != nil {
			return err
		}
	}
	return nil
}
