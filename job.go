package main

import (
	"fmt"
	"os"
	"os/exec"
)

type Job struct {
	Name       string   `json:"name"`
	Image      string   `json:"image"`
	AlwaysPull bool     `json:"alwaysPull"`
	Cron       string   `json:"cron"`
	Steps      []string `json:"steps"`
	Env        []string `json:"environment"`
	Jobs       []*Job   `json:"jobs"`
	ParentJob  *Job
}

func processJob(j *Job) {
	for _, job := range j.Jobs {
		job.ParentJob = j
		processJob(job)
	}
}

func (j *Job) GetFullname() string {
	if j.ParentJob != nil {
		return fmt.Sprintf("%s/%s", j.ParentJob.GetFullname(), j.Name)
	}
	return j.Name
}

func (j *Job) GetFullEnv() []string {
	if j.ParentJob != nil {
		return append(j.ParentJob.GetFullEnv(), j.Env...)
	}
	return j.Env
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
	if j.AlwaysPull {
		j.PullImage()
	}

	fmt.Printf("Running job %s \n ======================================\n", j.GetFullname())
	fmt.Printf("Image:%s \nEnvs: %s \n ======================================\n", j.Image, j.Env)

	if j.Image != "" {
		args := []string{"run", "--rm"}
		// args = append(args, t.Arguments...)
		for _, env := range j.GetFullEnv() {
			args = append(args, "-e", env)
		}
		args = append(args, j.Image)
		cmd := exec.Command("docker", args...)

		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			fmt.Println(err)
		}
	}

	for _, job := range j.Jobs {
		job.Run()
	}
}
