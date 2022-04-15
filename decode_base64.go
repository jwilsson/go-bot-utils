package utils

import "encoding/base64"

func decodeBase64(data string) string {
	result, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return ""
	}

	return string(result)
}
