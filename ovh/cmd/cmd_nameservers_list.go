package cmd

import (
	"errors"
	"os"

	"github.com/codegangsta/cli"
	"github.com/olekukonko/tablewriter"
	"golang.org/x/net/context"
)

var cmdNSList = cli.Command{
	Name:      "list",
	Usage:     "lists domain name servers",
	ArgsUsage: "[--all] [<domain> [<domain> ...]]",
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "all",
			Usage: "applies settings to all domains on account",
		},
		cli.StringFlag{
			Name:  "domain",
			Usage: "comma separated list of domains to apply settings to",
		},
	},
	Action: func(c *cli.Context) (err error) {
		if len(c.Args()) == 0 && !c.Bool("all") {
			return errors.New("You have to use --all or provide list of domains in arguments.")
		} else if len(c.Args()) > 0 && c.Bool("all") {
			return errors.New("Cannot use --all with list of domains")
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{
			"Domain",
			"ID",
			"Host",
			"Is Used",
			"To Delete",
		})

		for _, domain := range domainsFromCtx(c) {
			result, err := client(c).NameServers.List(context.Background(), domain)
			if err != nil {
				return err
			}
			for _, ns := range result {
				table.Append([]string{
					domain,
					ns.ID,
					ns.Host,
					yesOrNo(ns.IsUsed),
					yesOrNo(ns.ToDelete),
				})
			}
		}
		table.Render()
		return
	},
}
