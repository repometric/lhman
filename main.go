package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/repometric/lhman/catalog"
	"github.com/repometric/lhman/install"
	"github.com/urfave/cli"
)

// AppVersion is the version of lhman
const AppVersion = "0.1.0"

func main() {
	app := cli.NewApp()

	app.Version = AppVersion
	app.Usage = "Linterhub Manager Core Component"
	app.Commands = []cli.Command{
		{
			Name:    "catalog",
			Aliases: []string{"c"},
			Usage:   "This strategy generates list of engines using filters or specific keys and propose recommendations.",
			Action: func(c *cli.Context) error {

				var (
					engine = c.StringSlice("engine")
					//project = c.String("project")
					res []byte
				)

				var stringInSlice = func(a string, list []string) bool {
					for _, b := range list {
						if b == a {
							return true
						}
					}
					return false
				}

				if len(engine) > 0 {
					engines := make([]catalog.Engine, 0)
					for _, v := range catalog.Get() {
						if stringInSlice(v.Meta.Name, engine) {
							engines = append(engines, v)
						}
					}
					res, _ = json.MarshalIndent(engines, "", "    ")
				} else {
					engines := make([]catalog.Meta, 0)
					for _, e := range catalog.Get() {
						engines = append(engines, e.Meta)
					}
					res, _ = json.MarshalIndent(engines, "", "    ")
				}

				fmt.Println(string(res))
				return nil
			},
			Flags: []cli.Flag{
				cli.StringSliceFlag{
					Name:  "engine,e",
					Usage: "Engine name to return metadata, arguments and dependencies",
				},
				cli.StringFlag{
					Name:  "project,p",
					Usage: "Project path to return list of recommended for installing engines",
				},
			},
		},
		{
			Name:    "install",
			Aliases: []string{"i"},
			Usage:   "The install strategy allows to install the engine and its dependencies.",
			Action: func(c *cli.Context) error {
				var context = install.Context{
					Version:     c.StringSlice("version"),
					Project:     c.String("project"),
					Environment: c.String("environment"),
				}

				for _, engineName := range c.StringSlice("engine") {
					for _, engine := range catalog.Get() {
						if engine.Meta.Name == engineName || engine.Meta.ID == engineName {
							context.Engine = append(context.Engine)
						}
					}
				}

				res, _ := json.MarshalIndent(install.Run(context), "", "    ")
				fmt.Println(string(res))

				return nil
			},
			Flags: []cli.Flag{
				cli.StringSliceFlag{
					Name:  "engine,e",
					Usage: "Engine name to install",
				},
				cli.StringSliceFlag{
					Name:  "version,v",
					Usage: "Engine version (latest version by default)",
				},
				cli.StringFlag{
					Name:  "project,p",
					Usage: "Project to associate with",
				},
				cli.StringFlag{
					Name:  "environment,env",
					Usage: "The way how to install engine. Allowed values: local, global, container. local is used by default.",
					Value: "local",
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
