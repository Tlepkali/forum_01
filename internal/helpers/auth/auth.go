package auth

import (
	"encoding/json"
	"net/url"
)

func ExtractValueFromBody(body []byte, key string) string {
	var response map[string]interface{}
	err := json.Unmarshal(body, &response)
	if err != nil {
		return ""
	}

	value, ok := response[key].(string)
	if !ok {
		return ""
	}

	return value
}

func ExtractAccessTokenFromResponse(response string) (string, error) {
	params, err := url.ParseQuery(response)
	if err != nil {
		return "", err
	}

	accessToken := params.Get("access_token")
	return accessToken, nil
}
