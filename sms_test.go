package nexmo_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mikeyscode/nexmo"
	gock "gopkg.in/h2non/gock.v1"
)

var response = nexmo.MessageDetail{
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

func TestSendSMSReturnsExpectedResponse(t *testing.T) {
	defer gock.Off()
	gock.New(nexmo.SendMessageEndpoint).Post("/").
		Reply(http.StatusOK).
		JSON(response)

	n := nexmo.Nexmo{}
	n.Setup(nexmo.Auth{Key: "foo", Secret: "bar"})

	resp, err := n.SendSMS("0", "0", nexmo.SMSOptions{})
	if err != nil {
		t.Errorf("response was not returned due to error: %v", err)
	}

	assert.Equal(t, &response, resp)
}
