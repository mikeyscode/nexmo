package nexmo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

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

// SendSMS ... @todo query parameters
func SendSMS(payload OutboundSMSPayload) (messageDetail *MessageDetail, err error) {
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("unable to encode payload as json: %v", err)
	}

	client := &http.Client{Timeout: time.Second * 15}
	request, err := http.NewRequest("POST", SendMessageEndpoint, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, fmt.Errorf("unable to create post request: %v", err)
	}
	request.Header.Set("Content-Type", "application/json")

	response, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("unable to process post request: %v", err)
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&messageDetail)
	if err != nil {
		return nil, fmt.Errorf("unable to decode response body: %v", err)
	}

	return messageDetail, nil
}

/**

Methods to Handle
- Send an SMS
- SendSMS()
- Path: https://rest.nexmo.com /sms/:format
- Method: POST
- Documentation: https://developer.nexmo.com/api/sms#send-an-sms
- Query Parameters
	- api_key (required)
	- api_secret (required if no sig)
	- sig (required if no api_secret)
- Request Body
	-
*/
