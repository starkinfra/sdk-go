package utils

import (
	"github.com/starkinfra/core-go/starkcore/user/user"
	"github.com/starkinfra/core-go/starkcore/utils/parse"
	"github.com/starkinfra/sdk-go/starkinfra"
)

func ParseAndVerify(content string, signature string, key string, user user.User) string {
	if user == nil {
		return parse.ParseAndVerify(content, signature, starkinfra.SdkVersion, starkinfra.ApiVersion, starkinfra.Language, starkinfra.Timeout, starkinfra.Host, starkinfra.User, key).(string)
	}
	return parse.ParseAndVerify(content, signature, starkinfra.SdkVersion, starkinfra.ApiVersion, starkinfra.Language, starkinfra.Timeout, starkinfra.Host, user, key).(string)
}

func Verify(content string, signature string, user user.User) string {
	if user == nil {
		return parse.Verify(content, signature, starkinfra.SdkVersion, starkinfra.ApiVersion, starkinfra.Language, starkinfra.Timeout, starkinfra.Host, starkinfra.User).(string)
	}
	return parse.Verify(content, signature, starkinfra.SdkVersion, starkinfra.ApiVersion, starkinfra.Language, starkinfra.Timeout, starkinfra.Host, user).(string)
}
