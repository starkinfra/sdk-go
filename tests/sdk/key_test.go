package sdk

import (
	"github.com/starkinfra/core-go/starkcore/key"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreatePrivateKey(t *testing.T) {

	privateKey, publicKey := key.Create("")
	assert.NotNil(t, privateKey)
	assert.NotNil(t, publicKey)
}

func TestPathPrivateKey(t *testing.T) {

	privateKey, publicKey := key.Create("sample")
	assert.NotNil(t, privateKey)
	assert.NotNil(t, publicKey)
}
