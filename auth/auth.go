package auth

// APIKey is a holder of key and secret details
type APIKey struct {
	key    string
	secret string
}

// New returns a new APIKey authentication struct
func New(key, secret string) APIKey {
	return APIKey{key, secret}
}

// Key returns an api key
func (a APIKey) Key() string {
	return a.key
}

// Secret returns a secret
func (a APIKey) Secret() string {
	return a.secret
}
