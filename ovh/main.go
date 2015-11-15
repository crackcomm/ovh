package main

import (
	"log"
	"os"

	"github.com/codegangsta/cli"
	"github.com/crackcomm/ovh/ovh/cmd"
)

func init() {
	log.SetFlags(0)
}

func main() {
	app := cli.NewApp()
	app.Name = "ovh"
	app.HelpName = app.Name
	app.Usage = "OVH command line tool"
	app.Version = "1.0.0"
	app.Commands = cmd.New()
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "app-key",
			Usage:  "OVH API Application Key",
			EnvVar: "OVH_APP_KEY",
		},
		cli.StringFlag{
			Name:   "app-secret",
			Usage:  "OVH API Application Secret",
			EnvVar: "OVH_APP_SECRET",
		},
		cli.StringFlag{
			Name:   "consumer-key",
			Usage:  "OVH API Consumer Key",
			EnvVar: "OVH_CONSUMER_KEY",
		},
	}
	app.Run(os.Args)
}
