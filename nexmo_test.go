package nexmo_test

import (
	"testing"

	"github.com/mikeyscode/nexmo"
	"github.com/stretchr/testify/assert"
)

func TestSetupWithKeyAndSecret(t *testing.T) {
	n := nexmo.Nexmo{}
	n.Setup(nexmo.Auth{Key: "foo", Secret: "bar"})

	assert.Equal(t, "foo", n.Auth.Key)
	assert.Equal(t, "bar", n.Auth.Secret)
}

func TestSetupWithKeyAndSignature(t *testing.T) {
	n := nexmo.Nexmo{}
	n.Setup(nexmo.Auth{Key: "foo", Signature: "bar"})

	assert.Equal(t, "foo", n.Auth.Key)
	assert.Equal(t, "bar", n.Auth.Signature)
}
