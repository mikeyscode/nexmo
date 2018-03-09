package nexmo

// APIInterface for API reference
type APIInterface interface {
	// Setup sets up the underlying library configuration
	Setup(Auth) error
	// SendSMS sends an SMS using the adaptors library layer
	SendSMS(string, string, SMSOptions) (SMSResponseInterface, error)
	// DispatchTextCall initiates a phone call and uses options in TextCallOptions to convert text to speech to speek to the recipient
	DispatchTextCall(from string, to string, optins TextCallOptions) (TextCallResponseInterface, error)
}

// Nexmo is the top level structure that interfaces with the API.
type Nexmo struct {
	Auth
}

// Setup sets up a nexmo instance for dispatching SMS
func (n *Nexmo) Setup(config Auth) error {
	n.Auth = config

	return nil
}
