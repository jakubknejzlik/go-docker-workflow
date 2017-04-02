package main

import (
	"log"
	"os"
	"os/exec"
)

type Job struct {
	Name  string   `yaml:"name"`
	Image string   `yaml:"image"`
	Cron  string   `yaml:"cron"`
	Steps []string `yaml:"steps"`
	Env   []string `yaml:"environment"`
}

type Config struct {
	Jobs map[string]*Job `yaml:"jobs"`
}

func (j *Job) PullImage() error {

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

	log.Printf("Running job %s", j.Name)

	args := []string{"run", "--rm"}
	// args = append(args, t.Arguments...)
	for _, env := range j.Env {
		args = append(args, "-e", env)
	}
	args = append(args, j.Image)
	cmd := exec.Command("docker", args...)
	log.Print(cmd)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
