package auth_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mikeyscode/nexmo/auth"
)

func TestNew(t *testing.T) {
	auth := auth.New("foo", "bar")

	assert.Equal(t, "foo", auth.Key())
	assert.Equal(t, "bar", auth.Secret())
}
