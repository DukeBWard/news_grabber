package auth

import (
	"errors"
	"net/http"
	"strings"
)

// Gets the api key from the headers of http request
// if not, returns error
// ex. Authorization: APIKey {aljuk;sdhqi28e123123812512}
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no auth found")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("bad auth header")
	}

	if vals[0] != "ApiKey" {
		return "", errors.New("bad auth header")
	}

	return vals[1], nil
}
