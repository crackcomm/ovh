package ovh

// UserInfo - User info structure.
type UserInfo struct {
	Name      string `json:"name,omitempty"`
	Email     string `json:"email,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
}

type tempUserInfo struct {
	Name      string `json:"nichandle"`
	Email     string `json:"email"`
	FirstName string `json:"firstname"`
	LastName  string `json:"name"`
}

func (profile *tempUserInfo) toUserInfo() *UserInfo {
	return &UserInfo{
		Name:      profile.Name,
		Email:     profile.Email,
		FirstName: profile.FirstName,
		LastName:  profile.LastName,
	}
}
