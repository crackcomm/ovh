package ovh

import (
	"encoding/json"
	"net/http"

	"golang.org/x/net/context"
)

// User - OVH User info API Client.
type User struct {
	// Options - Client options.
	*Options
}

// Info - Gets information about authenticated consumer.
func (user *User) Info(ctx context.Context) (result *UserInfo, err error) {
	response, err := httpDo(ctx, user.Options, "GET", apiURL("/me"), nil)
	if err != nil {
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, unexpectedStatusError(response)
	}

	res := new(tempUserInfo)
	err = json.NewDecoder(response.Body).Decode(res)
	if err != nil {
		return
	}

	return res.toUserInfo(), nil
}
