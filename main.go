package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

const APP_VERSION = "0.0.1-alpha"

func main() {
	app := cli.NewApp()

	app.Version = APP_VERSION

	app.Commands = []cli.Command{
		{
			Name:    "catalog",
			Aliases: []string{"c"},
			Usage:   "This strategy generates list of engines using filters or specific keys and propose recommendations.",
			Action: func(c *cli.Context) error {
				fmt.Println("TODO")
				return nil
			},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "filter,f",
					Usage: "Filter results",
				},
				cli.StringFlag{
					Name:  "keys,k",
					Usage: "Show only specified keys",
				},
				cli.StringFlag{
					Name:  "engine,e",
					Usage: "Engine name to return metadata, arguments and dependencies",
				},
				cli.StringFlag{
					Name:  "project,p",
					Usage: "Project path to return list of recommended for installing engines",
				},
			},
			Subcommands: []cli.Command{
				{
					Name:  "bundle",
					Usage: "Generate a bundle with metadata, arguments and dependencies for all engines. ",
					Action: func(c *cli.Context) error {
						fmt.Println("TODO")
						return nil
					},
				},
			},
		},
		{
			Name:    "install",
			Aliases: []string{"i"},
			Usage:   "The install strategy allows to install the engine and its dependencies.",
			Action: func(c *cli.Context) error {
				fmt.Println("TODO")
				return nil
			},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "engine,e",
					Usage: "Engine name to install",
				},
				cli.StringFlag{
					Name:  "version,v",
					Usage: "Engine version (latest version by default)",
				},
				cli.StringFlag{
					Name:  "project,p",
					Usage: "Project to associate with",
				},
				cli.StringFlag{ // convert to enum
					Name:  "environment,env",
					Usage: "The way how to install engine. Allowed values: local, global, container. local is used by default.",
				},
			},
		},
		{
			Name:    "version",
			Aliases: []string{"v"},
			Usage:   "Returns current component version.",
			Action: func(c *cli.Context) error {
				fmt.Println(c.App.Version)
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
