package ovh

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/net/context"
)

// AuthenticateRequest - Authentication request.
type AuthenticateRequest struct {
	ValidationURL string `json:"validation_url,omitempty"`
	ConsumerKey   string `json:"consumer_key,omitempty"`
}

// Authenticate - Requests OVH authentiction.
// It is ok to use one callback url or none.
func (client *Client) Authenticate(ctx context.Context, callbackURL ...string) (result *AuthenticateRequest, err error) {
	buffer := bytes.NewBufferString(fmt.Sprintf(`{
		"accessRules": [
			{"method": "GET", "path": "/*"},
			{"method": "PUT", "path": "/*"},
			{"method": "POST", "path": "/*"},
			{"method": "DELETE", "path": "/*"}
		],
		"redirection": "%s"
	}`, firstString(callbackURL)))
	response, err := httpDo(ctx, client.Options, "POST", apiURL("/auth/credential"), buffer)
	if err != nil {
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, unexpectedStatusError(response)
	}

	r := new(struct {
		ValidationURL string `json:"validationUrl"`
		ConsumerKey   string `json:"consumerKey"`
	})
	err = json.NewDecoder(response.Body).Decode(r)
	if err != nil {
		return
	}

	return &AuthenticateRequest{
		ValidationURL: r.ValidationURL,
		ConsumerKey:   r.ConsumerKey,
	}, nil
}

func firstString(s []string) (_ string) {
	if len(s) != 0 {
		return s[0]
	}
	return
}
