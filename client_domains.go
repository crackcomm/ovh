package ovh

import (
	"bytes"
	"encoding/json"
	"net/http"

	"golang.org/x/net/context"
)

// Domains - OVH Domains API Client.
type Domains struct {
	opts *Options
}

// List - Lists domains.
func (domains *Domains) List(ctx context.Context) (result []string, err error) {
	response, err := httpDo(ctx, domains.opts, "GET", apiURL("/domain"), nil)
	if err != nil {
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, unexpectedStatusError(response)
	}

	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		return
	}

	return
}

// Details - Gets domain details by domain name.
func (domains *Domains) Details(ctx context.Context, domain string) (result *Domain, err error) {
	response, err := httpDo(ctx, domains.opts, "GET", apiURL("/domain/%s", domain), nil)
	if err != nil {
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, unexpectedStatusError(response)
	}

	r := new(tempDomain)
	err = json.NewDecoder(response.Body).Decode(&r)
	if err != nil {
		return
	}

	return r.getDomain(), nil
}

// Patch - Patches a domain details by domain name in domain.Name.
func (domains *Domains) Patch(ctx context.Context, domain string, p *DomainPatch) (err error) {
	buffer := new(bytes.Buffer)
	err = json.NewEncoder(buffer).Encode(&struct {
		NameServerType     string `json:"nameServerType,omitempty"`
		TransferLockStatus string `json:"transferLockStatus,omitempty"`
	}{
		NameServerType:     p.NameServerType,
		TransferLockStatus: p.TransferLockStatus,
	})
	if err != nil {
		return
	}
	response, err := httpDo(ctx, domains.opts, "PUT", apiURL("/domain/%s", domain), buffer)
	if err != nil {
		return
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return unexpectedStatusError(response)
	}
	return
}
