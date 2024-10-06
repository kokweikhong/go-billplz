package billplz

import (
	"log"
	"net/http"
)

func (bp *Billplz) HandleCallback(w http.ResponseWriter, r *http.Request) {
	// Parse the form data
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Extract the x_signature from the form
	xSignature := r.FormValue("x_signature")

	// Verify X-Signature to ensure the integrity of the request
	if !verifyXSignature(r.Form, xSignature, bp.config.xSignatureKey) {
		http.Error(w, "Invalid signature", http.StatusUnauthorized)
		return
	}

	// Read the form data into BillplzCallback structure
	callback := BillplzCallback{
		ID:                r.FormValue("id"),
		CollectionID:      r.FormValue("collection_id"),
		Paid:              r.FormValue("paid") == "true",
		State:             r.FormValue("state"),
		Amount:            atoi(r.FormValue("amount")),
		PaidAmount:        atoi(r.FormValue("paid_amount")),
		DueAt:             r.FormValue("due_at"),
		Email:             r.FormValue("email"),
		Mobile:            r.FormValue("mobile"),
		Name:              r.FormValue("name"),
		URL:               r.FormValue("url"),
		PaidAt:            r.FormValue("paid_at"),
		TransactionID:     r.FormValue("transaction_id"),
		TransactionStatus: r.FormValue("transaction_status"),
		XSignature:        xSignature,
	}

	// Process payment status
	if callback.Paid {
		log.Printf("Payment successful for bill ID: %s, Transaction ID: %s\n", callback.ID, callback.TransactionID)
		// Here you can update your database to mark the payment as successful
	} else {
		log.Printf("Payment failed or due for bill ID: %s\n", callback.ID)
		// Handle the case where payment failed
	}

	// Respond to Billplz server
	w.WriteHeader(http.StatusOK)
}
