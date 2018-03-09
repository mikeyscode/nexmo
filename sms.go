package nexmo

// SendMessageEndpoint ...
const SendMessageEndpoint = "https://rest.nexmo.com /sms/json"

// MessageDetail contains information about a sent message
type MessageDetail struct {
	To               string `json:"to"`
	MessageID        string `json:"message-id"`
	Status           string `json:"status"`
	RemainingBalance string `json:"remaining-balance"`
	MessagePrice     string `json:"message-price"`
	Network          string `json:"network"`
}

// SMS contains information about the request and details of message information
type SMS struct {
	MessageCount int             `json:"message-count"`
	Messages     []MessageDetail `json:"messages"`
}

// DataCodingScheme is a SMS message class used by Nexmo
type DataCodingScheme int

const (
	// Flash ...
	Flash DataCodingScheme = iota
	// Standard ...
	Standard
	// SIMData ...;
	SIMData
	// Forward ...
	Forward
)

// OutboundSMSPayload ...
type OutboundSMSPayload struct {
	From                 string           `json:"from"`
	To                   string           `json:"to"`
	Text                 string           `json:"text,omitempty"`
	TTL                  int              `json:"ttl,omitempty"`
	StatusReportRequired bool             `json:"status-report-req,omitempty"`
	Callback             string           `json:"callback,omitempty"`
	MessageClass         DataCodingScheme `json:"message-class,omitempty"`
}

// SMSOptions ...
type SMSOptions struct {
	Text                 string           `json:"text,omitempty"`
	TTL                  int              `json:"ttl,omitempty"`
	StatusReportRequired bool             `json:"status-report-req,omitempty"`
	Callback             string           `json:"callback,omitempty"`
	MessageClass         DataCodingScheme `json:"message-class,omitempty"`
}
