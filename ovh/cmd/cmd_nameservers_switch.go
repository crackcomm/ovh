package cmd

import (
	"errors"
	"log"
	"strings"

	"github.com/codegangsta/cli"
	"github.com/crackcomm/ovh"
	"golang.org/x/net/context"
)

var cmdNSSwitch = cli.Command{
	Name:  "switch",
	Usage: "switches domain name servers",
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "all",
			Usage: "applies settings to all domains on account",
		},
		cli.StringFlag{
			Name:  "domain",
			Usage: "comma separated list of domains to apply settings to",
		},
		cli.StringSliceFlag{
			Name:  "ns",
			Usage: "nameserver",
		},
	},
	Action: func(c *cli.Context) (err error) {
		if len(c.StringSlice("ns")) < 2 {
			return errors.New("At least two domain name servers are required.")
		}
		if !c.Bool("all") && c.String("domain") == "" {
			return errors.New("You have to use --all or --domain.")
		}

		ns := c.StringSlice("ns")
		for _, domain := range domainsFromCtx(c) {
			err = switchDomainNameservers(c, domain, ns)
			if err != nil {
				return
			}
		}

		log.Println("OK")
		return
	},
}

func switchDomainNameservers(c *cli.Context, domain string, ns []string) (err error) {
	log.Printf("Switching %q nameservers: %s", domain, strings.Join(ns, ", "))

	result, err := client(c).NameServers.List(context.Background(), domain)
	if err != nil {
		return
	}

	var (
		done   []string
		delete []*ovh.NameServer
	)
	for _, nameserver := range result {
		if nameserver.ToDelete {
			continue
		}
		if stringIn(nameserver.Host, ns) {
			done = append(done, nameserver.Host)
		} else {
			delete = append(delete, nameserver)
		}
	}

	var todo []string
	for _, nameserver := range ns {
		if stringIn(nameserver, done) {
			continue
		}
		log.Printf("Inserting new nameserver %q", nameserver)
		todo = append(todo, nameserver)
	}

	err = client(c).NameServers.Insert(context.Background(), domain, todo...)
	if err != nil {
		return
	}

	for _, nameserver := range delete {
		log.Printf("Deleting %q nameserver %q.", domain, nameserver.Host)
		err = client(c).NameServers.Delete(context.Background(), domain, nameserver.ID)
		if err != nil {
			return
		}
	}
	return
}
