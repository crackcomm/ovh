package ovh

// Options - OVH API Client Options.
type Options struct {
	AppKey      string
	AppSecret   string
	ConsumerKey string
}

// Client - OVH API Client.
type Client struct {
	// User - OVH user info API client.
	*User
	// Domains - OVH domains API client.
	*Domains
	// NameServers - OVH name servers API client.
	*NameServers
	// Options - OVH API client options.
	*Options
}

// New - Creates a new OVH client.
func New(opts *Options) *Client {
	return &Client{
		User:        &User{Options: opts},
		Domains:     &Domains{Options: opts},
		NameServers: &NameServers{Options: opts},
		Options:     opts,
	}
}
