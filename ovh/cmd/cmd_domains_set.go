package cmd

import (
	"log"
	"strings"

	"github.com/codegangsta/cli"
	"github.com/crackcomm/ovh"
	"golang.org/x/net/context"
)

var cmdDomainsSet = cli.Command{
	Name:      "set",
	Usage:     "sets domain name server option",
	ArgsUsage: "[--all] [<domain> [<domain> ...]]",
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "all",
			Usage: "applies settings to all domains on account",
		},
		cli.StringFlag{
			Name:  "name_server_type",
			Usage: "changes domain name server type (external or hosted)",
		},
		cli.StringFlag{
			Name:  "transfer_lock_status",
			Usage: "changes domain transfer lock status (locked or unlocked)",
		},
	},
	Action: func(c *cli.Context) {
		if len(c.Args()) == 0 && !c.Bool("all") {
			log.Fatal("You have to use --all or provide list of domains in arguments.")
		} else if len(c.Args()) > 0 && c.Bool("all") {
			log.Fatal("Cannot use --all with list of domains")
		}

		client := client(c)
		ctx := context.Background()
		patch := &ovh.DomainPatch{
			NameServerType:     c.String("name_server_type"),
			TransferLockStatus: c.String("transfer_lock_status"),
		}

		if v := patch.NameServerType; v != "" {
			log.Printf("Changing name server type to %q", v)
		}
		if v := patch.TransferLockStatus; v != "" {
			log.Printf("Changing transfer lock status to %q", v)
		}

		for _, domain := range domainsFromCtx(c) {
			log.Printf("Updating %s", domain)
			err := client.Domains.Patch(ctx, domain, patch)
			if err != nil {
				log.Fatal(err)
			}
		}

		log.Println("Done")
	},
}

func domainsFromCtx(c *cli.Context) (domains []string) {
	if c.Bool("all") {
		domains, err := client(c).Domains.List(context.Background())
		if err != nil {
			log.Fatalf("Error retrieving list of domains: %v", err)
		}
		return domains
	}
	if l := c.String("domain"); l != "" {
		return strings.Split(l, ",")
	}
	return []string(c.Args())
}
