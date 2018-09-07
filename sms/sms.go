package sms

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	messageEndpoint = "https://rest.nexmo.com/sms/json"
)

var auth Authable

type Authable interface {
	Key() string
	Secret() string
}

type Options struct {
	Text                 string `json:"text,omitempty"`
	TTL                  int    `json:"ttl,omitempty"`
	StatusReportRequired bool   `json:"status-report-req,omitempty"`
	Callback             string `json:"callback,omitempty"`
	MessageClass         int    `json:"message-class,omitempty"`
}

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

type request struct {
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

func Auth(a Authable) {
	auth = a
}

func Send(to, from string, options Options) (interface{}, error) {
	requestBody := request{
		Key:                  auth.Key(),
		Secret:               auth.Secret(),
		To:                   to,
		From:                 from,
		Text:                 options.Text,
		TTL:                  options.TTL,
		StatusReportRequired: options.StatusReportRequired,
		Callback:             options.Callback,
		MessageClass:         options.MessageClass,
	}

	payload, err := json.Marshal(requestBody)
	if err != nil {
		return nil, fmt.Errorf("unable to encode payload as json: %v", err)
	}

	req, err := http.NewRequest(http.MethodPost, messageEndpoint, bytes.NewBuffer(payload))
	if err != nil {
		return nil, fmt.Errorf("unable to create post request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: time.Second * 10}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("unable to process post request: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("unable to read response body: %v", err)
	}

	var messageDetail = &MessageDetail{}
	if err = json.Unmarshal([]byte(body), messageDetail); err != nil {
		return nil, fmt.Errorf("unable to decode response body: %v", err)
	}

	return messageDetail, nil
}
