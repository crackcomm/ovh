package cmd

import "github.com/codegangsta/cli"

var cmdDomains = cli.Command{
	Name:  "domains",
	Usage: "domains",
	Subcommands: []cli.Command{
		cmdDomainsDetails,
		cmdDomainsList,
		cmdDomainsSet,
	},
}
