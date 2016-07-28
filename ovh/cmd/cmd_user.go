package cmd

import (
	"encoding/json"
	"log"

	"golang.org/x/net/context"

	"github.com/codegangsta/cli"
)

var cmdUserInfo = cli.Command{
	Name:  "user",
	Usage: "shows user info",
	Action: func(c *cli.Context) (err error) {
		resp, err := client(c).User.Info(context.Background())
		if err != nil {
			return
		}

		body, err := json.MarshalIndent(resp, "", "  ")
		if err != nil {
			return
		}

		log.Printf("%s\n", body)
		return
	},
}
