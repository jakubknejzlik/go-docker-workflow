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
				var man Manager

				config := c.Args().First()
				if config != "" {
					man = NewManagerFromBase64(config)
				} else {
					man = NewManagerFromYamlFile("./config.yml")
				}

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

				var man Manager

				config := c.Args().First()
				if config != "" {
					man = NewManagerFromBase64(config)
				} else {
					man = NewManagerFromYamlFile("./config.yml")
				}

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
