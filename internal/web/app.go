package web

import (
	"net/http"

	"github.com/kokweikhong/go-billplz"
)

type app struct {
	server  *http.ServeMux
	billplz *billplz.Billplz
}

func New() *app {
	a := new(app)
	a.server = http.NewServeMux()

	return a
}

func (a *app) InitRoutes() {
	a.server.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	a.server.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})
}

func (a *app) Start(addr string) error {
	return http.ListenAndServe(addr, a.server)
}
