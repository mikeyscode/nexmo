package nexmo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Nexmo is the top level structure that interfaces with the API.
type Nexmo struct {
	Client        *http.Client
	Authorisation Auth
}

// Setup sets up the Client and Authorisation layer and returns a Nexmo instance
func Setup(client *http.Client, auth Auth) Nexmo {
	return Nexmo{client, auth}
}

// SendSMS will send a text message to a specified phone
func (n *Nexmo) SendSMS(to, from string, options SMSOptions) (messageDetail *MessageDetail, err error) {
	requestBody := OutboundSMSPayload{
		n.Authorisation.Key,
		n.Authorisation.Secret,
		n.Authorisation.Signature,
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

	if err != nil {
		return nil, fmt.Errorf("unable to create query string: %v", err)
	}

	req, err := http.NewRequest("POST", SendMessageEndpoint, bytes.NewBuffer(payload))
	if err != nil {
		return nil, fmt.Errorf("unable to create post request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

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
		return nil, fmt.Errorf("unable to decode response body: %v", err)
	}

	return messageDetail, nil
}
