package main

import (
	"os"

	"github.com/urfave/cli"
)

func main() {

	app := cli.NewApp()
	app.Name = "go docker workflow"
	app.Version = "0.0.1"
	app.Usage = "..."

	app.Commands = []cli.Command{
		{
			Name: "start",
			Action: func(c *cli.Context) error {
				man := NewManager(c.Args().First())
				man.Start()
				return nil
			},
		},
		{
			Name:      "run",
			ArgsUsage: "[JOB_NAME]",
			Action: func(c *cli.Context) error {
				man := NewManager(c.Args().First())
				jobName := c.Args().Get(1)
				if jobName == "" {
					man.Run()
				} else {
					man.RunJob(jobName)
				}
				return nil
			},
		},
	}

	app.Run(os.Args)
}
