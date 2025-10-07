package utils

import (
	"github.com/starkinfra/core-go/starkcore/user/user"
	Errors "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/utils/parse"
	"github.com/starkinfra/sdk-go/starkinfra"
)

func ParseAndVerify(content string, signature string, key string, user user.User) (string, Errors.StarkErrors) {
	if user == nil {
		response, err := parse.ParseAndVerify(content, signature, starkinfra.SdkVersion, starkinfra.ApiVersion, starkinfra.Language, starkinfra.Timeout, starkinfra.Host, starkinfra.User, key)
		if err.Errors != nil {
			return "", err
		}
		return response.(string), err
	}
	response, err := parse.ParseAndVerify(content, signature, starkinfra.SdkVersion, starkinfra.ApiVersion, starkinfra.Language, starkinfra.Timeout, starkinfra.Host, user, key)
	if err.Errors != nil {
		return "", err
	}
	return response.(string), err
}

func Verify(content string, signature string, user user.User) (string, Errors.StarkErrors) {
	if user == nil {
		response, err := parse.Verify(content, signature, starkinfra.SdkVersion, starkinfra.ApiVersion, starkinfra.Language, starkinfra.Timeout, starkinfra.Host, starkinfra.User)
		if err.Errors != nil {
			return "", err
		}
		return response.(string), err
	}
	response, err := parse.Verify(content, signature, starkinfra.SdkVersion, starkinfra.ApiVersion, starkinfra.Language, starkinfra.Timeout, starkinfra.Host, user)
	if err.Errors != nil {
		return "", err
	}
	return response.(string), err
}
