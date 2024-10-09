package billplz

type CreateBillParams struct {
	CollectionID string `json:"collection_id"`
	Email        string `json:"email"`
	Mobile       string `json:"mobile"`
	Name         string `json:"name"`
	Amount       int    `json:"amount"`
	CallbackURL  string `json:"callback_url"`
	Description  string `json:"description"`
	DueAt        string `json:"due_at"`
	RedirectURL  string `json:"redirect_url"`
}

type CreateBillResponse struct {
	ID           string `json:"id"`
	CollectionID string `json:"collection_id"`
	Paid         bool   `json:"paid"`
	State        string `json:"state"`
	Amount       int    `json:"amount"`
	PaidAmount   int    `json:"paid_amount"`
	DueAt        string `json:"due_at"`
	Email        string `json:"email"`
	Mobile       string `json:"mobile"`
	Name         string `json:"name"`
	URL          string `json:"url"`
	RedirectURL  string `json:"redirect_url"`
	CallbackURL  string `json:"callback_url"`
	Description  string `json:"description"`
}

type BillplzCallback struct {
	ID                string `json:"id"`
	CollectionID      string `json:"collection_id"`
	Paid              bool   `json:"paid"`
	State             string `json:"state"`
	Amount            int    `json:"amount"`
	PaidAmount        int    `json:"paid_amount"`
	DueAt             string `json:"due_at"`
	Email             string `json:"email"`
	Mobile            string `json:"mobile"`
	Name              string `json:"name"`
	URL               string `json:"url"`
	PaidAt            string `json:"paid_at"`
	TransactionID     string `json:"transaction_id"`
	TransactionStatus string `json:"transaction_status"`
	XSignature        string `json:"x_signature"`
}

type GetBillResponse struct {
	ID    string `json:"id"`
	Paid  bool   `json:"paid"`
	State string `json:"state"`
}
