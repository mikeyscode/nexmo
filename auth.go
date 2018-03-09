package nexmo

// Auth contains the necessary information to authenticate an API request to Nexmo
type Auth struct {
	Key       string `json:"api_key"`
	Secret    string `json:"api_secret,omitempty"`
	Signature string `json:"sig,omitempty"`
}
