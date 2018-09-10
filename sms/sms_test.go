package sms_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/h2non/gock"
	"github.com/mikeyscode/nexmo/sms"
)

const (
	messageEndpoint = "https://rest.nexmo.com/sms/json"
)

var mockResponse = sms.MessageDetail{
	To:               "0",
	MessageID:        "1",
	Status:           "200",
	RemainingBalance: "1.00",
	MessagePrice:     "0.01",
	Network:          "",
	MessageCount:     "1",
	Messages: []struct {
		Status    string `json:"status"`
		ErrorText string `json:"error-text"`
	}{
		{"200", ""},
	},
}

type mockAuth struct{}

func (mockAuth) Key() string    { return "foo" }
func (mockAuth) Secret() string { return "bar" }

func TestSend(t *testing.T) {
	defer gock.Off()
	gock.New(messageEndpoint).
		Post("/").
		Reply(http.StatusOK).
		JSON(mockResponse)

	sms.Auth(&mockAuth{})

	resp, err := sms.Send("07000000000", "Test", sms.Options{})
	if err != nil {
		t.Errorf("response was not returned due to error: %v", err)
	}

	assert.Equal(t, mockResponse, resp)
}
