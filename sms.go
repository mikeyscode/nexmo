package nexmo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// SendMessageEndpoint ...
const SendMessageEndpoint = "https://rest.nexmo.com/sms/json"

// SMSResponseInterface is the response interface for SMS
type SMSResponseInterface interface{}

// MessageDetail contains information about a sent message
type MessageDetail struct {
	To               string `json:"to"`
	MessageID        string `json:"message-id"`
	Status           string `json:"status"`
	RemainingBalance string `json:"remaining-balance"`
	MessagePrice     string `json:"message-price"`
	Network          string `json:"network"`
	MessageCount     string `json:"message-count"`
	Messages         []struct {
		Status    string `json:"status"`
		ErrorText string `json:"error-text"`
	} `json:"messages"`
}

// SMS contains information about the request and details of message information
type SMS struct {
	MessageCount int             `json:"message-count"`
	Messages     []MessageDetail `json:"messages"`
}

const (
	// Flash ...
	Flash int = iota
	// Standard ...
	Standard
	// SIMData ...;
	SIMData
	// Forward ...
	Forward
)

// OutboundSMSPayload ...
type OutboundSMSPayload struct {
	Key                  string `json:"api_key"`
	Secret               string `json:"api_secret,omitempty"`
	Signature            string `json:"sig,omitempty"`
	To                   string `json:"to"`
	From                 string `json:"from"`
	Text                 string `json:"text,omitempty"`
	TTL                  int    `json:"ttl,omitempty"`
	StatusReportRequired bool   `json:"status-report-req,omitempty"`
	Callback             string `json:"callback,omitempty"`
	MessageClass         int    `json:"message-class,omitempty"`
}

// SMSOptions ...
type SMSOptions struct {
	Text                 string `json:"text,omitempty"`
	TTL                  int    `json:"ttl,omitempty"`
	StatusReportRequired bool   `json:"status-report-req,omitempty"`
	Callback             string `json:"callback,omitempty"`
	MessageClass         int    `json:"message-class,omitempty"`
}

// SendSMS will send a text message to a specified phone
func (n *Nexmo) SendSMS(from, to string, options SMSOptions) (SMSResponseInterface, error) {
	var messageDetail = &MessageDetail{}

	requestBody := OutboundSMSPayload{
		n.Auth.Key,
		n.Auth.Secret,
		n.Auth.Signature,
		to,
		from,
		options.Text,
		options.TTL,
		options.StatusReportRequired,
		options.Callback,
		options.MessageClass,
	}

	payload, err := json.Marshal(requestBody)
	if err != nil {
		return nil, fmt.Errorf("unable to encode payload as json: %v", err)
	}

	req, err := http.NewRequest("POST", SendMessageEndpoint, bytes.NewBuffer(payload))
	if err != nil {
		return nil, fmt.Errorf("unable to create post request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: time.Second * 15}
	resp, err := client.Do(req)

	if err != nil {
		return nil, fmt.Errorf("unable to process post request: %v", err)
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("unable to read response body: %v", err)
	}

	if err = json.Unmarshal([]byte(respBody), messageDetail); err != nil {
		return nil, fmt.Errorf("unable to decode response body: %v", err)
	}

	return messageDetail, nil
}
