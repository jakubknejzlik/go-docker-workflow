package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

type Job struct {
	Name      string          `yaml:"name"`
	Image     string          `yaml:"image"`
	Cron      string          `yaml:"cron"`
	Steps     []string        `yaml:"steps"`
	Env       []string        `yaml:"environment"`
	Jobs      map[string]*Job `yaml:"jobs"`
	ParentJob *Job
}

func processJob(j *Job, name string) {
	j.Name = name
	for name, job := range j.Jobs {
		job.ParentJob = j
		job.Env = append(job.Env, j.Env...)
		processJob(job, name)
	}
}

func (j *Job) GetFullname() string {
	if j.ParentJob != nil {
		return fmt.Sprintf("%s/%s", j.ParentJob.GetFullname(), j.Name)
	}
	return j.Name
}

func (j *Job) PullImage() error {

	if j.Image == "" {
		return nil
	}

	args := []string{"pull", j.Image}
	// args = append(args, t.Arguments...)
	cmd := exec.Command("docker", args...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func (j *Job) Run() {
	j.PullImage()

	log.Printf("Running job %s", j.GetFullname())

	if j.Image != "" {
		args := []string{"run", "--rm"}
		// args = append(args, t.Arguments...)
		for _, env := range j.Env {
			args = append(args, "-e", env)
		}
		args = append(args, j.Image)
		cmd := exec.Command("docker", args...)

		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}
	}

	for _, job := range j.Jobs {
		job.Run()
	}
}
