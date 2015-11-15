# Golang OVH API client

[![GoDoc](https://godoc.org/github.com/crackcomm/ovh?status.svg)](https://godoc.org/github.com/crackcomm/ovh)

Golang API Client for OVH APi.

## Command Line Tool

```sh
$ go install github.com/crackcomm/ovh/ovh
$ ovh
NAME:
   ovh - OVH command line tool

USAGE:
   ovh [global options] command [command options] [arguments...]

VERSION:
   1.0.0

COMMANDS:
   auth		requests authentication
   domains	domains
   ns		domain name servers
   help, h	Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --app-key 		OVH API Application Key [$OVH_APP_KEY]
   --app-secret 	OVH API Application Secret [$OVH_APP_SECRET]
   --consumer-key 	OVH API Consumer Key [$OVH_CONSUMER_KEY]
   --help, -h		show help
   --version, -v	print the version
```

## License

Apache 2.0 License.
