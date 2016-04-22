package cmd

import (
	"encoding/json"
	"fmt"
	"log"

	"golang.org/x/net/context"

	"github.com/codegangsta/cli"
)

var cmdDomainsDetails = cli.Command{
	Name:      "details",
	Usage:     "prints domain details",
	ArgsUsage: "<domain>",
	Action: func(c *cli.Context) {
		if len(c.Args()) != 1 {
			cli.ShowSubcommandHelp(c)
			return
		}
		domain, err := client(c).Domains.Details(context.Background(), c.Args().First())
		if err != nil {
			log.Fatal(err)
		}

		body, err := json.MarshalIndent(domain, "", "  ")
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s\n", body)
	},
}
