package sdk

import (
	"fmt"
	"github.com/starkinfra/core-go/starkcore/key"
	"testing"
)

func TestCreatePrivateKey(t *testing.T) {

	privateKey, publicKey := key.Create("")
	fmt.Println("PRIVATE KEY", privateKey)
	fmt.Println("PUBLIC KEY", publicKey)
}

func TestPathPrivateKey(t *testing.T) {

	privateKey, publicKey := key.Create("sample")
	fmt.Println("PRIVATE KEY", privateKey)
	fmt.Println("PUBLIC KEY", publicKey)
}
