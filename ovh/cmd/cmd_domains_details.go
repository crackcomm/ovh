package cmd

import (
	"encoding/json"
	"fmt"

	"golang.org/x/net/context"

	"github.com/codegangsta/cli"
)

var cmdDomainsDetails = cli.Command{
	Name:      "details",
	Usage:     "prints domain details",
	ArgsUsage: "<domain>",
	Action: func(c *cli.Context) (err error) {
		if len(c.Args()) != 1 {
			cli.ShowSubcommandHelp(c)
			return
		}
		domain, err := client(c).Domains.Details(context.Background(), c.Args().First())
		if err != nil {
			return
		}

		body, err := json.MarshalIndent(domain, "", "  ")
		if err != nil {
			return
		}

		fmt.Printf("%s\n", body)
		return
	},
}
