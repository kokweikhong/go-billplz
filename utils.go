package billplz

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"net/url"
	"sort"
	"strconv"
	"strings"
)

func basicAuth(apiKey string) string {
	return base64.StdEncoding.EncodeToString([]byte(apiKey + ":"))
}

func verifyXSignature(data url.Values, xSignature string, secret string) bool {
	// Sort the keys to ensure they are in the correct order
	keys := make([]string, 0, len(data))
	for k := range data {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// Build the query string from the sorted keys
	var payload strings.Builder
	for _, k := range keys {
		payload.WriteString(fmt.Sprintf("%s=%s&", k, data.Get(k)))
	}
	// Remove the trailing '&'
	payloadStr := strings.TrimSuffix(payload.String(), "&")

	// Generate HMAC-SHA256 signature
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(payloadStr))
	expectedSignature := hex.EncodeToString(h.Sum(nil))

	// Compare the generated signature with the received x_signature
	return hmac.Equal([]byte(expectedSignature), []byte(xSignature))
}

// atoi is a helper to convert string to integer
func atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
