package cmd

import (
	"log"
	"strings"

	"github.com/codegangsta/cli"
	"github.com/crackcomm/ovh"
	"golang.org/x/net/context"
)

var cmdNSSwitch = cli.Command{
	Name:      "switch",
	Usage:     "switches domain name servers",
	ArgsUsage: "<nameserver> <nameserver> [<nameserver> ...]",
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
	Action: func(c *cli.Context) {
		if len(c.Args()) == 0 {
			log.Fatal("Usage: ovh ns switch <nameserver> <nameserver> [<nameserver> ...]")
		}
		if len(c.Args()) < 2 {
			log.Fatal("At least two domain name servers are required.")
		}
		if !c.Bool("all") && c.String("domain") == "" {
			log.Fatal("You have to use --all or --domain.")
		}

		ns := []string(c.Args())
		for _, domain := range domainsFromCtx(c) {
			switchDomainNameservers(c, domain, ns)
		}

		log.Println("Done")
	},
}

func switchDomainNameservers(c *cli.Context, domain string, ns []string) {
	log.Printf("Switching %q nameservers: %s", domain, strings.Join(ns, ", "))

	result, err := client(c).NameServers.List(context.Background(), domain)
	if err != nil {
		log.Fatal(err)
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
		log.Fatal(err)
	}

	for _, nameserver := range delete {
		log.Printf("Deleting %q nameserver %q.", domain, nameserver.Host)
		err := client(c).NameServers.Delete(context.Background(), domain, nameserver.ID)
		if err != nil {
			log.Fatal(err)
		}
	}
}
