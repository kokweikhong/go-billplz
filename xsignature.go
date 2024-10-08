package billplz

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

// computeHMACSHA256 generates the HMAC-SHA256 signature
func computeHMACSHA256(data, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

// verifyXSignature verifies the X Signature for the redirect URL
func verifyXSignature(source string, xSignature string, secret string) bool {

	// Compute the HMAC-SHA256 signature
	computedXSignature := computeHMACSHA256(source, secret)

	// Compare the computed signature with the provided x_signature
	return hmac.Equal([]byte(computedXSignature), []byte(xSignature))
}
