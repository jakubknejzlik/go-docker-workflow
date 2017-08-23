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
				if err := man.Start(); err != nil {
					return cli.NewExitError(err, 1)
				}
				return nil
			},
		},
		{
			Name: "run",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "job,j",
					Value: "",
				},
			},
			Action: func(c *cli.Context) error {
				man := NewManager(c.Args().First())
				jobName := c.String("job")
				if jobName == "" {
					if err := man.Run(); err != nil {
						return cli.NewExitError(err, 1)
					}
				} else {
					if err := man.RunJob(jobName); err != nil {
						return cli.NewExitError(err, 1)
					}
				}
				return nil
			},
		},
	}

	app.Run(os.Args)
}
