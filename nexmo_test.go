package nexmo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetupWithKeyAndSecret(t *testing.T) {
	n := Nexmo{}
	n.Setup(Auth{Key: "foo", Secret: "bar"})

	assert.Equal(t, "foo", n.Auth.Key)
	assert.Equal(t, "bar", n.Auth.Secret)
}

func TestSetupWithKeyAndSignature(t *testing.T) {
	n := Nexmo{}
	n.Setup(Auth{Key: "foo", Signature: "bar"})

	assert.Equal(t, "foo", n.Auth.Key)
	assert.Equal(t, "bar", n.Auth.Signature)
}
