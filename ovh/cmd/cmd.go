// Package cmd implements OVH cli commands.
package cmd

import "github.com/codegangsta/cli"

// New - Returns OVH cli commands.
func New() []cli.Command {
	return []cli.Command{
		cmdAuth,
		cmdUserInfo,
		cmdDomains,
		cmdNS,
	}
}

func yesOrNo(b bool) string {
	if b {
		return "yes"
	}
	return "no"
}

func stringIn(s string, l []string) bool {
	for _, v := range l {
		if v == s {
			return true
		}
	}
	return false
}
