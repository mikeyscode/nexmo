package nexmo

import "fmt"

// Auth contains the necessary information to authenticate an API request to Nexmo
type Auth struct {
	Key       string `json:"api_key"`
	Secret    string `json:"api_secret,omitempty"`
	Signature string `json:"sig,omitempty"`
}

// buildQueryString will build a query string to be appended based on the configuration of the Auth struct
func (a *Auth) buildQueryString() (string, error) {
	if a.isValid() == false {
		return "", fmt.Errorf("Auth must specify at least one of [Secret] or [Signature]")
	}

	if a.Secret != "" {
		return fmt.Sprintf("?api_key=%s&api_secret=%s", a.Key, a.Secret), nil
	}

	return fmt.Sprintf("?api_key=%s&sig=%s", a.Key, a.Signature), nil
}

// isValid will check whether a given Auth struct contains at least one of the required properties: Secret or Signature
func (a *Auth) isValid() bool {
	return a.Secret == "" && a.Signature == ""
}
