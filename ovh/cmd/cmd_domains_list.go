package cmd

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/net/context"

	"github.com/codegangsta/cli"
	"github.com/olekukonko/tablewriter"
)

var cmdDomainsList = cli.Command{
	Name:  "list",
	Usage: "lists domains",
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "list",
			Usage: "print list of domains instead of table",
		},
	},
	Action: func(c *cli.Context) {
		domains, err := client(c).Domains.List(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		if c.Bool("list") {
			for _, domain := range domains {
				fmt.Println(domain)
			}
			return
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Domain"})
		for _, domain := range domains {
			table.Append([]string{domain})
		}
		table.Render()
	},
}
