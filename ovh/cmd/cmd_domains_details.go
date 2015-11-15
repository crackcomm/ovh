package cmd

import (
	"log"
	"os"

	"golang.org/x/net/context"

	"github.com/codegangsta/cli"
	"github.com/olekukonko/tablewriter"
)

var cmdDomainsDetails = cli.Command{
	Name:  "details",
	Usage: "print domain details",
	Action: func(c *cli.Context) {
		domains, err := client(c).Domains.List(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Domain"})
		for _, domain := range domains {
			table.Append([]string{domain})
		}
		table.Render()
	},
}
