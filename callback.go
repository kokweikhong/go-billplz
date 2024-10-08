package billplz

import (
	"fmt"
	"net/http"
	"sort"
	"strings"
)

func (bp *Billplz) Callback(r *http.Request) error {
	// Check if the request method is POST
	if r.Method != http.MethodPost {
		return fmt.Errorf("Method Not Allowed")
	}

	fmt.Println("Received callback")

	// Parse the URL-encoded form data
	err := r.ParseForm()
	if err != nil {
		return err
	}

	fmt.Println("Form data:", r.Form)

	// Extract the x_signature
	xSignature := r.FormValue("x_signature")
	if xSignature == "" {
		return fmt.Errorf("X Signature is missing")
	}

	// Create a map for the form data without x_signature
	params := make(map[string]string)
	for key, values := range r.Form {
		if key != "x_signature" {
			// Use the first value from the slice
			params[key] = values[0]
		}
	}
	// Shared secret key (X Signature Key)
	secretKey := bp.config.xSignatureKey
	fmt.Println("secretKey: ", secretKey)

	// Generate the source string
	source := generateCallbackSource(params)
	fmt.Println("callback source: ", source)
	// Verify the X Signature
	if verifyXSignature(source, xSignature, secretKey) {
		fmt.Println("X Signature is valid")
	} else {
		fmt.Println("X Signature is invalid")
		return fmt.Errorf("Invalid signature")
	}
	return nil
}

func generateCallbackSource(query map[string]string) string {
	sources := make([]string, 0, len(query))
	for k, v := range query {
		sources = append(sources, fmt.Sprintf("%s%s", k, v))
	}
	sort.Strings(sources)
	source := strings.Join(sources, "|")

	return source
}
