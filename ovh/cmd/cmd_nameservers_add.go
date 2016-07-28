package cmd

import (
	"errors"
	"log"
	"strings"

	"github.com/codegangsta/cli"
	"golang.org/x/net/context"
)

var cmdNSAdd = cli.Command{
	Name:      "add",
	Usage:     "adds domain name servers",
	ArgsUsage: "<domain> <nameserver> [<nameserver> ...]",
	Action: func(c *cli.Context) (err error) {
		if len(c.Args()) < 2 {
			return errors.New("Usage: ovh ns delete <domain> <nameserver> [<nameserver> ...]")
		}

		domain := c.Args().First()
		ns := []string(c.Args())[1:]
		log.Printf("Adding nameservers to %q: %s", domain, strings.Join(ns, ", "))

		err = client(c).NameServers.Insert(context.Background(), domain, ns...)
		if err != nil {
			return
		}
		log.Println("OK")
		return
	},
}
