package ovh

// Options - OVH API Client Options.
type Options struct {
	AppKey      string
	AppSecret   string
	ConsumerKey string
}

// Client - OVH API Client.
type Client struct {
	*Domains
	*NameServers
	opts *Options
}

// New - Creates a new OVH client.
func New(opts *Options) *Client {
	return &Client{
		Domains:     &Domains{opts: opts},
		NameServers: &NameServers{opts: opts},
		opts:        opts,
	}
}
