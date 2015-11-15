package cmd

import "github.com/codegangsta/cli"

var cmdNS = cli.Command{
	Name:  "ns",
	Usage: "domain name servers",
	Subcommands: []cli.Command{
		cmdNSAdd,
		cmdNSList,
		cmdNSDelete,
		cmdNSSwitch,
	},
}
