package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/repometric/lhman/catalog"
	"github.com/repometric/lhman/install"
	"github.com/urfave/cli"
)

const appVersion = "0.1.1"
const bundleURL = "https://repometric.github.io/linterhub/engine/bundle.json"

func updateBundle() {
	os.Mkdir("tmp", 0777)

	client := &http.Client{}
	request, _ := http.NewRequest(http.MethodHead, bundleURL, nil)
	response, _ := client.Do(request)
	lastModified := response.Header.Get("Last-Modified")

	dat, _ := ioutil.ReadFile("tmp/last.dat")
	lastModifiedLocal := string(dat)

	if lastModified != lastModifiedLocal {
		response, _ := http.Get("https://repometric.github.io/linterhub/engine/bundle.json")
		body, _ := ioutil.ReadAll(response.Body)
		ioutil.WriteFile("tmp/last.dat", []byte(lastModified), 0644)
		ioutil.WriteFile("tmp/bundle.json", body, 0644)
	}
}

func main() {
	updateBundle()
	app := cli.NewApp()

	app.Version = appVersion
	app.Usage = "Linterhub Manager Core Component"
	app.Commands = []cli.Command{
		{
			Name:    "catalog",
			Aliases: []string{"c"},
			Usage:   "This strategy generates list of engines using filters or specific keys and propose recommendations.",
			Action: func(c *cli.Context) error {
				if c.IsSet("version") {
					dat, _ := ioutil.ReadFile("tmp/last.dat")
					lastModifiedLocal := string(dat)
					fmt.Println("Linterhub bundle: " + lastModifiedLocal)
				} else {
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
				}
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
				cli.BoolFlag{
					Name:  "version,v",
					Usage: "Show bundle version",
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
					Folder:      c.String("folder"),
					Environment: c.String("environment"),
				}

				enginesArg := c.StringSlice("engine")

				if len(enginesArg) == 0 {
					cli.ShowCommandHelp(c, "install")
				} else {
					for _, engineName := range enginesArg {
						for _, engine := range catalog.Get() {
							if engine.Meta.Name == engineName || engine.Meta.ID == engineName {
								context.Engine = append(context.Engine, engine)
							}
						}
					}
					res, _ := json.MarshalIndent(install.Run(context), "", "    ")
					fmt.Println(string(res))
				}

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
					Name:  "folder,f",
					Usage: "Folder path for local installation",
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
