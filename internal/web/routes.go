package web

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/kokweikhong/go-billplz"
)

func (a *app) InitRoutes() {
	a.server.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		a.RenderTemplate(w, "index", nil)
	})

	a.server.HandleFunc("/create-bill", a.handleCreateBill)

	a.server.HandleFunc("POST /bill/callback", func(w http.ResponseWriter, r *http.Request) {
		if err := a.billplz.Callback(r); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Write([]byte("Callback Success"))
	})

	a.server.HandleFunc("/payment-success", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Payment Success"))
	})

	a.server.HandleFunc("/bill/redirect", func(w http.ResponseWriter, r *http.Request) {
		data, err := a.billplz.Redirect(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		fmt.Println(data)

		http.Redirect(w, r, "/payment-success", http.StatusSeeOther)
	})

	a.server.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})
}

func (a *app) handleCreateBill(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	r.ParseForm()
	name := r.Form.Get("name")
	email := r.Form.Get("email")

	params := new(billplz.CreateBillParams)
	params.CollectionID = a.config.billplz.collectionID
	params.Email = email
	params.Name = name
	amount, _ := strconv.ParseFloat(r.Form.Get("amount"), 64)
	params.Amount = int(amount * 100)
	params.Mobile = r.Form.Get("mobile")
	params.CallbackURL = a.config.billplz.callbackURL
	params.Description = r.Form.Get("description")
	params.DueAt = "2024-10-10"
	params.RedirectURL = a.config.billplz.redirectURL

	fmt.Println(params)

	bill, err := a.billplz.CreateBill(params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, bill.URL, http.StatusSeeOther)
}
