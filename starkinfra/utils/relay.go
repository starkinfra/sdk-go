package utils

import (
	Errors "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	"github.com/starkinfra/core-go/starkcore/utils/rest"
	"github.com/starkinfra/sdk-go/starkinfra"
	"github.com/starkinfra/core-go/starkcore/utils/request"
)

func Page(resource map[string]string, params map[string]interface{}, user user.User) ([]byte, string, Errors.StarkErrors) {
	if user == nil {
		return rest.GetPage(starkinfra.SdkVersion, starkinfra.Host, starkinfra.ApiVersion, starkinfra.Language, starkinfra.Timeout, starkinfra.User, resource, params)
	}
	return rest.GetPage(starkinfra.SdkVersion, starkinfra.Host, starkinfra.ApiVersion, starkinfra.Language, starkinfra.Timeout, user, resource, params)
}

func Query(resource map[string]string, params map[string]interface{}, user user.User) chan map[string]interface{} {
	if user == nil {
		return rest.GetStream(starkinfra.SdkVersion, starkinfra.Host, starkinfra.ApiVersion, starkinfra.Language, starkinfra.Timeout, starkinfra.User, resource, params)
	}
	return rest.GetStream(starkinfra.SdkVersion, starkinfra.Host, starkinfra.ApiVersion, starkinfra.Language, starkinfra.Timeout, user, resource, params)
}

func Get(resource map[string]string, id string, query map[string]interface{}, user user.User) ([]byte, Errors.StarkErrors) {
	if user == nil {
		return rest.GetId(starkinfra.SdkVersion, starkinfra.Host, starkinfra.ApiVersion, starkinfra.Language, starkinfra.Timeout, starkinfra.User, resource, id, query)
	}
	return rest.GetId(starkinfra.SdkVersion, starkinfra.Host, starkinfra.ApiVersion, starkinfra.Language, starkinfra.Timeout, user, resource, id, query)
}

func GetContent(resource map[string]string, id string, params map[string]interface{}, user user.User, content string) ([]byte, Errors.StarkErrors) {
	if user == nil {
		return rest.GetContent(starkinfra.SdkVersion, starkinfra.Host, starkinfra.ApiVersion, starkinfra.Language, starkinfra.Timeout, starkinfra.User, resource, id, content, params)
	}
	return rest.GetContent(starkinfra.SdkVersion, starkinfra.Host, starkinfra.ApiVersion, starkinfra.Language, starkinfra.Timeout, user, resource, id, content, params)
}

func SubResource(resource map[string]string, id string, user user.User, subResource map[string]string) ([]byte, Errors.StarkErrors) {
	if user == nil {
		return rest.GetSubResource(starkinfra.SdkVersion, starkinfra.Host, starkinfra.ApiVersion, starkinfra.Language, starkinfra.Timeout, starkinfra.User, resource, id, subResource, nil)
	}
	return rest.GetSubResource(starkinfra.SdkVersion, starkinfra.Host, starkinfra.ApiVersion, starkinfra.Language, starkinfra.Timeout, user, resource, id, subResource, nil)
}

func Multi(resource map[string]string, entities interface{}, query map[string]interface{}, user user.User) ([]byte, Errors.StarkErrors) {
	if user == nil {
		return rest.PostMulti(starkinfra.SdkVersion, starkinfra.Host, starkinfra.ApiVersion, starkinfra.Language, starkinfra.Timeout, starkinfra.User, resource, entities, query)
	}
	return rest.PostMulti(starkinfra.SdkVersion, starkinfra.Host, starkinfra.ApiVersion, starkinfra.Language, starkinfra.Timeout, user, resource, entities, query)
}

func Single(resource map[string]string, entity interface{}, user user.User) ([]byte, Errors.StarkErrors) {
	if user == nil {
		return rest.PostSingle(starkinfra.SdkVersion, starkinfra.Host, starkinfra.ApiVersion, starkinfra.Language, starkinfra.Timeout, starkinfra.User, resource, entity, nil)
	}
	return rest.PostSingle(starkinfra.SdkVersion, starkinfra.Host, starkinfra.ApiVersion, starkinfra.Language, starkinfra.Timeout, user, resource, entity, nil)
}

func Delete(resource map[string]string, id string, user user.User) ([]byte, Errors.StarkErrors) {
	if user == nil {
		return rest.DeleteId(starkinfra.SdkVersion, starkinfra.Host, starkinfra.ApiVersion, starkinfra.Language, starkinfra.Timeout, starkinfra.User, resource, id, nil)
	}
	return rest.DeleteId(starkinfra.SdkVersion, starkinfra.Host, starkinfra.ApiVersion, starkinfra.Language, starkinfra.Timeout, user, resource, id, nil)
}

func Patch(resource map[string]string, id string, payload map[string]interface{}, user user.User) ([]byte, Errors.StarkErrors) {
	if user == nil {
		return rest.PatchId(starkinfra.SdkVersion, starkinfra.Host, starkinfra.ApiVersion, starkinfra.Language, starkinfra.Timeout, starkinfra.User, resource, id, payload, nil)
	}
	return rest.PatchId(starkinfra.SdkVersion, starkinfra.Host, starkinfra.ApiVersion, starkinfra.Language, starkinfra.Timeout, user, resource, id, payload, nil)
}

func GetRaw(path string, query map[string]interface{}, user user.User, prefix string, throwError bool) (request.Response, Errors.StarkErrors) {
	if user == nil {
		return rest.GetRaw(starkinfra.SdkVersion, starkinfra.Host, starkinfra.ApiVersion, starkinfra.Language, starkinfra.Timeout, path, starkinfra.User, query, prefix, throwError)
	}
	return rest.GetRaw(starkinfra.SdkVersion, starkinfra.Host, starkinfra.ApiVersion, starkinfra.Language, starkinfra.Timeout, path, user, query, prefix, throwError)
}

func PostRaw(path string, body interface{}, user user.User, query map[string]interface{}, prefix string, throwError bool) (request.Response, Errors.StarkErrors) {
	if user == nil {
		return rest.PostRaw(starkinfra.SdkVersion, starkinfra.Host, starkinfra.ApiVersion, starkinfra.Language, starkinfra.Timeout, path, body, starkinfra.User, query, prefix, throwError)
	}
	return rest.PostRaw(starkinfra.SdkVersion, starkinfra.Host, starkinfra.ApiVersion, starkinfra.Language, starkinfra.Timeout, path, body, user, query, prefix, throwError)
}

func PatchRaw(path string, body interface{}, user user.User, query map[string]interface{}, prefix string, throwError bool) (request.Response, Errors.StarkErrors) {
	if user == nil {
		return rest.PatchRaw(starkinfra.SdkVersion, starkinfra.Host, starkinfra.ApiVersion, starkinfra.Language, starkinfra.Timeout, path, body, starkinfra.User, query, prefix, throwError)
	}
	return rest.PatchRaw(starkinfra.SdkVersion, starkinfra.Host, starkinfra.ApiVersion, starkinfra.Language, starkinfra.Timeout, path, body, user, query, prefix, throwError)
}

func PutRaw(path string, body interface{}, user user.User, query map[string]interface{}, prefix string, throwError bool) (request.Response, Errors.StarkErrors) {
	if user == nil {
		return rest.PutRaw(starkinfra.SdkVersion, starkinfra.Host, starkinfra.ApiVersion, starkinfra.Language, starkinfra.Timeout, path, body, starkinfra.User, query, prefix, throwError)
	}
	return rest.PutRaw(starkinfra.SdkVersion, starkinfra.Host, starkinfra.ApiVersion, starkinfra.Language, starkinfra.Timeout, path, body, user, query, prefix, throwError)
}

func DeleteRaw(path string, user user.User, prefix string, throwError bool) (request.Response, Errors.StarkErrors) {
	if user == nil {
		return rest.DeleteRaw(starkinfra.SdkVersion, starkinfra.Host, starkinfra.ApiVersion, starkinfra.Language, starkinfra.Timeout, path, starkinfra.User, prefix, throwError)
	}
	return rest.DeleteRaw(starkinfra.SdkVersion, starkinfra.Host, starkinfra.ApiVersion, starkinfra.Language, starkinfra.Timeout, path, user, prefix, throwError)
}
