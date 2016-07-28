package cmd

import (
	"errors"
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
			Name:  "name-server-type",
			Usage: "changes domain name server type (external or hosted)",
		},
		cli.StringFlag{
			Name:  "transfer-lock-status",
			Usage: "changes domain transfer lock status (locked or unlocked)",
		},
	},
	Action: func(c *cli.Context) (err error) {
		if len(c.Args()) == 0 && !c.Bool("all") {
			return errors.New("You have to use --all or provide list of domains in arguments.")
		} else if len(c.Args()) > 0 && c.Bool("all") {
			return errors.New("Cannot use --all with list of domains")
		}

		client := client(c)
		ctx := context.Background()
		patch := &ovh.DomainPatch{
			NameServerType:     c.String("name-server-type"),
			TransferLockStatus: c.String("transfer-lock-status"),
		}

		if v := patch.NameServerType; v != "" {
			log.Printf("Changing name server type to %q", v)
		}
		if v := patch.TransferLockStatus; v != "" {
			log.Printf("Changing transfer lock status to %q", v)
		}

		for _, domain := range domainsFromCtx(c) {
			log.Printf("Updating %s", domain)
			err = client.Domains.Patch(ctx, domain, patch)
			if err != nil {
				return
			}
		}

		log.Println("OK")
		return
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
