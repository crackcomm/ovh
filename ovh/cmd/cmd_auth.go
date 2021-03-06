package cmd

import (
	"log"

	"golang.org/x/net/context"

	"github.com/codegangsta/cli"
	"github.com/crackcomm/ovh"
)

var cmdAuth = cli.Command{
	Name:  "auth",
	Usage: "requests authentication",
	Action: func(c *cli.Context) (err error) {
		resp, err := client(c).Authenticate(context.Background())
		if err != nil {
			return
		}

		log.Printf("Validate token: %s", resp.ValidationURL)
		log.Printf("export OVH_CONSUMER_KEY=%q", resp.ConsumerKey)
		return
	},
}

func client(c *cli.Context) *ovh.Client {
	return ovh.New(&ovh.Options{
		AppKey:      c.GlobalString("app-key"),
		AppSecret:   c.GlobalString("app-secret"),
		ConsumerKey: c.GlobalString("consumer-key"),
	})
}
