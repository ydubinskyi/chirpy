package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetAPIKey(headers http.Header) (string, error) {
	authHeaderValue := headers.Get("Authorization")

	if authHeaderValue == "" {
		return "", errors.New("Authorization header missing")
	}

	parts := strings.Fields(authHeaderValue)

	if len(parts) != 2 || parts[0] != "ApiKey" {
		return "", errors.New("Authorization header invalid")
	}

	tokenString := parts[1]

	return tokenString, nil
}
