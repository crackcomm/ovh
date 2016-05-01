package ovh

// NameServer - Name Server structure.
type NameServer struct {
	ID       string `json:"id,omitempty"`
	Host     string `json:"host,omitempty"`
	ToDelete bool   `json:"to_delete,omitempty"`
	IsUsed   bool   `json:"is_used,omitempty"`
}
