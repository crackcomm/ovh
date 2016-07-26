package cmd

import (
	"encoding/json"
	"fmt"
	"log"

	"golang.org/x/net/context"

	"github.com/codegangsta/cli"
)

var cmdUserInfo = cli.Command{
	Name:  "user",
	Usage: "shows user info",
	Action: func(c *cli.Context) {
		resp, err := client(c).User.Info(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		body, err := json.MarshalIndent(resp, "", "  ")
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s\n", body)
	},
}
