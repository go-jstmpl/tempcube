package main

import (
	"log"
	"os"

	_ "github.com/pkg/errors"
	"github.com/urfave/cli"
)

func CmdNew() *cli.App {
	app := cli.NewApp()

	app.Name = "tmpcube"
	app.Usage = "tmplate generater and testing framework"

	app.Commands = []cli.Command{
		{
			Name:   "init",
			Usage:  "create init project",
			Action: Init,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "project name",
					Usage: "Template project name you want to begin.",
				},
			},
		},
		{
			Name:   "build",
			Usage:  "building from template and JSON Schema file.",
			Action: Build,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "project",
					Usage: "Template project folder.",
				},
				cli.StringFlag{
					Name:  "schema, s",
					Usage: "Source JSON Schema file.",
				},
				cli.StringFlag{
					Name:  "destination, d",
					Usage: "Destionation file path want to build.",
				},
			},
		},
		{
			Name:   "test",
			Usage:  "testing validation and golang standard testing.",
			Action: Test,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "project",
					Usage: "Template project folder.",
				},
				cli.StringFlag{
					Name:  "schema, s",
					Usage: "Source JSON Schema file.",
				},
			},
		},
	}

	return app
}

func Init(c *cli.Context) error {
	return nil
}

func Build(c *cli.Context) error {
	return nil
}

func Test(c *cli.Context) error {
	return nil
}

func _cmd() int {
	cmd := CmdNew()
	err := cmd.Run(os.Args)
	if err != nil {
		log.Println(err)
		return 1
	}
	return 0
}

func main() {
	os.Exit(_cmd())
}
