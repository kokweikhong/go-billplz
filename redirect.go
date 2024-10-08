package billplz

import (
	"errors"
	"fmt"
	"net/http"
	"sort"
	"strings"
)

type billplzRedirectResponse struct {
	ID     string `json:"id"`
	Paid   bool   `json:"paid"`
	PaidAt string `json:"paid_at"`
}

func (bp *Billplz) Redirect(r *http.Request) (*billplzRedirectResponse, error) {
	// Parse the URL query parameters
	r.ParseForm()

	// Extract all relevant parameters except `x_signature`
	params := map[string]string{
		"billplz[id]":      r.FormValue("billplz[id]"),
		"billplz[paid]":    r.FormValue("billplz[paid]"),
		"billplz[paid_at]": r.FormValue("billplz[paid_at]"),
	}

	// Get x_signature from query parameters
	xSignature := r.FormValue("billplz[x_signature]")

	// Secret key (X Signature Key)
	secretKey := bp.config.xSignatureKey

	fmt.Println("secretKey: ", secretKey)

	source := generateRedirectSource(params)

	// Verify the X Signature
	if !verifyXSignature(source, xSignature, secretKey) {
		return nil, errors.New("Invalid signature")
	}

	res := &billplzRedirectResponse{
		ID:     params["billplz[id]"],
		Paid:   params["billplz[paid]"] == "true",
		PaidAt: params["billplz[paid_at]"],
	}

	return res, nil
}

func generateRedirectSource(query map[string]string) string {
	// Extract keys and sort them alphabetically (case-insensitive)
	keys := make([]string, 0, len(query))
	for k := range query {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// Build the concatenated source string
	var payloads []string
	for _, k := range keys {
		if k == "billplz[x_signature]" {
			continue
		}
		// change billplz[id]=abc -> billplzidabc
		key := strings.Replace(k, "[", "", -1)
		key = strings.Replace(key, "]", "", -1)
		payloads = append(payloads, key+query[k])

	}
	// sort the payloads and case-insensitive
	sort.Slice(payloads, func(i, j int) bool {
		return strings.ToLower(payloads[i]) < strings.ToLower(payloads[j])
	})

	sourceString := strings.Join(payloads, "|")

	fmt.Println("redirect source: ", sourceString)

	return sourceString
}
