package ovh

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/net/context"
)

// NameServers - OVH Domains Name Servers API Client.
type NameServers struct {
	opts *Options
}

// List - Gets domain nameservers info by domain name.
func (nameservers *NameServers) List(ctx context.Context, domain string) (result []*NameServer, err error) {
	response, err := httpDo(ctx, nameservers.opts, "GET", apiURL("/domain/%s/nameServer", domain), nil)
	if err != nil {
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, unexpectedStatusError(response)
	}

	var r []int64
	err = json.NewDecoder(response.Body).Decode(&r)
	if err != nil {
		return
	}

	for _, id := range r {
		ns, err := nameservers.Details(ctx, domain, fmt.Sprintf("%d", id))
		if err != nil {
			return nil, err
		}
		result = append(result, ns)
	}

	return
}

// Details - Gets domain nameserver info by domain name and name server ID.
func (nameservers *NameServers) Details(ctx context.Context, domain, id string) (result *NameServer, err error) {
	response, err := httpDo(ctx, nameservers.opts, "GET", apiURL("/domain/%s/nameServer/%s", domain, id), nil)
	if err != nil {
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, unexpectedStatusError(response)
	}

	r := new(struct {
		ID       int64  `json:"id,omitempty"`
		Host     string `json:"host,omitempty"`
		ToDelete bool   `json:"toDelete,omitempty"`
		IsUsed   bool   `json:"isUsed,omitempty"`
	})
	err = json.NewDecoder(response.Body).Decode(&r)
	if err != nil {
		return
	}

	return &NameServer{
		ID:       fmt.Sprintf("%d", r.ID),
		Host:     r.Host,
		ToDelete: r.ToDelete,
		IsUsed:   r.IsUsed,
	}, nil
}

// Delete - Deletes domain nameserver by domain name and name server ID.
func (nameservers *NameServers) Delete(ctx context.Context, domain, id string) (err error) {
	response, err := httpDo(ctx, nameservers.opts, "DELETE", apiURL("/domain/%s/nameServer/%s", domain, id), nil)
	if err != nil {
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return unexpectedStatusError(response)
	}

	return
}

// Insert - Inserts domain nameservers to domain.
func (nameservers *NameServers) Insert(ctx context.Context, domain string, nameserver ...string) (err error) {
	if len(nameserver) == 0 {
		return
	}
	buffer := new(bytes.Buffer)
	req := new(struct {
		NameServer []*struct {
			Host string `json:"host,omitempty"`
		} `json:"nameServer,omitempty"`
	})
	for _, ns := range nameserver {
		req.NameServer = append(req.NameServer, &struct {
			Host string `json:"host,omitempty"`
		}{
			Host: ns,
		})
	}
	err = json.NewEncoder(buffer).Encode(req)
	if err != nil {
		return
	}
	response, err := httpDo(ctx, nameservers.opts, "POST", apiURL("/domain/%s/nameServer", domain), buffer)
	if err != nil {
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return unexpectedStatusError(response)
	}

	return
}
