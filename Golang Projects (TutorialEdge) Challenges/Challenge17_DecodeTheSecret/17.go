package Challenge17_DecodeTheSecret

import (
	"encoding/base64"
)

func DecodeSecret(message string) string {
	data, _ := base64.StdEncoding.DecodeString(message)
	var secret []rune

	for _, c := range data {
		secret = append(secret, rune(c-1))
	}

	return string(secret)
}
