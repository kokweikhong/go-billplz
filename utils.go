package billplz

import (
	"encoding/base64"
	"strconv"
)

func basicAuth(apiKey string) string {
	return base64.StdEncoding.EncodeToString([]byte(apiKey + ":"))
}

// atoi is a helper to convert string to integer
func atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
