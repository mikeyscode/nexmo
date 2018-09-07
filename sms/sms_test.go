package sms_test

import (
	"testing"

	"github.com/mikeyscode/nexmo/sms"
)

type MockAuth struct{}

func (MockAuth) Key() string    { return "foo" }
func (MockAuth) Secret() string { return "bar" }

func Test_Auth(t *testing.T) {
	mockAuth := MockAuth{}
	sms.Auth(mockAuth)
}
