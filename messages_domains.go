package ovh

// Domain - Domain structure.
type Domain struct {
	Name string `json:"name,omitempty"`

	Offer              string `json:"offer,omitempty"`
	WhoisOwner         string `json:"whois_owner,omitempty"`
	LastUpdate         string `json:"last_update,omitempty"`
	NameServerType     string `json:"name_server_type,omitempty"`
	TransferLockStatus string `json:"transfer_lock_status,omitempty"`

	OwoSupported               bool `json:"owo_supported,omitempty"`
	DnssecSupported            bool `json:"dnssec_supported,omitempty"`
	GlueRecordIpv6Supported    bool `json:"glue_record_ipv6_supported,omitempty"`
	GlueRecordMultiIPSupported bool `json:"glue_record_multi_ip_supported,omitempty"`

	NameServers []*NameServer `json:"name_servers,omitempty"`
}

// DomainPatch - Domain patch request.
type DomainPatch struct {
	NameServerType     string `json:"name_server_type,omitempty"`
	TransferLockStatus string `json:"transfer_lock_status,omitempty"`
}

// NameServer - Name Server structure.
type NameServer struct {
	ID       string `json:"id,omitempty"`
	Host     string `json:"host,omitempty"`
	ToDelete bool   `json:"to_delete,omitempty"`
	IsUsed   bool   `json:"is_used,omitempty"`
}

type tempDomain struct {
	Offer              string `json:"offer,omitempty"`
	WhoisOwner         string `json:"whoisOwner,omitempty"`
	Domain             string `json:"domain,omitempty"`
	LastUpdate         string `json:"lastUpdate,omitempty"`
	NameServerType     string `json:"nameServerType,omitempty"`
	TransferLockStatus string `json:"transferLockStatus,omitempty"`

	OwoSupported               bool `json:"owoSupported,omitempty"`
	DnssecSupported            bool `json:"dnssecSupported,omitempty"`
	GlueRecordIpv6Supported    bool `json:"glueRecordIpv6Supported,omitempty"`
	GlueRecordMultiIPSupported bool `json:"glueRecordMultiIpSupported,omitempty"`
}

func makeTempDomain(domain *Domain) *tempDomain {
	return &tempDomain{
		Offer:                      domain.Offer,
		WhoisOwner:                 domain.WhoisOwner,
		Domain:                     domain.Name,
		LastUpdate:                 domain.LastUpdate,
		NameServerType:             domain.NameServerType,
		TransferLockStatus:         domain.TransferLockStatus,
		OwoSupported:               domain.OwoSupported,
		DnssecSupported:            domain.DnssecSupported,
		GlueRecordIpv6Supported:    domain.GlueRecordIpv6Supported,
		GlueRecordMultiIPSupported: domain.GlueRecordMultiIPSupported,
	}
}

func (r *tempDomain) getDomain() *Domain {
	return &Domain{
		Name:                       r.Domain,
		Offer:                      r.Offer,
		WhoisOwner:                 r.WhoisOwner,
		LastUpdate:                 r.LastUpdate,
		NameServerType:             r.NameServerType,
		TransferLockStatus:         r.TransferLockStatus,
		OwoSupported:               r.OwoSupported,
		DnssecSupported:            r.DnssecSupported,
		GlueRecordIpv6Supported:    r.GlueRecordIpv6Supported,
		GlueRecordMultiIPSupported: r.GlueRecordMultiIPSupported,
	}
}
