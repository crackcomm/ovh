package cmd

import (
	"log"
	"strings"

	"github.com/codegangsta/cli"
	"golang.org/x/net/context"
)

var cmdNSDelete = cli.Command{
	Name:      "delete",
	Usage:     "deletes domain name server by id",
	ArgsUsage: "<domain> <nameserver> [<nameserver> ...]",
	Action: func(c *cli.Context) {
		if len(c.Args()) < 2 {
			log.Fatal("Usage: ovh ns delete <domain> <nameserver> [<nameserver> ...]")
		}

		domain := c.Args().First()
		ns := []string(c.Args())[1:]
		log.Printf("Deleting nameservers to %q: %s", domain, strings.Join(ns, ", "))

		for _, id := range ns {
			err := client(c).NameServers.Delete(context.Background(), domain, id)
			if err != nil {
				log.Fatal(err)
			}
			log.Println("Done")
		}
	},
}
