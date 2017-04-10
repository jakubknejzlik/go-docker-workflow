package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {

	app := cli.NewApp()
	app.Name = "go docker workflow"
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
			Name: "run",
			Action: func(c *cli.Context) error {
				man := NewManager(c.Args().First())
				man.Run()
				return nil
			},
		},
		{
			Name: "version",
			Action: func(c *cli.Context) error {
				fmt.Println("0.0.1")
				return nil
			},
		},
	}

	app.Run(os.Args)
}
