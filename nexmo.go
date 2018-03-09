package nexmo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Nexmo ...
type Nexmo struct {
	Client        *http.Client
	Authorisation Auth
}

// Setup ...
func Setup(client *http.Client, auth Auth) Nexmo {
	return Nexmo{client, auth}
}

// SendSMS ...
func (n *Nexmo) SendSMS(to, from string, options SMSOptions) (messageDetail *MessageDetail, err error) {
	requestBody := OutboundSMSPayload{
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

	resp, err := n.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("unable to process post request: %v", err)
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("unable to read response body: %v", err)
	}

	err = json.Unmarshal([]byte(respBody), &messageDetail)
	if err != nil {
		return nil, fmt.Errorf("unablet o decode response body: %v", err)
	}

	return messageDetail, nil
}
