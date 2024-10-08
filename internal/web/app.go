package web

import (
	"net/http"
	"os"

	"github.com/kokweikhong/go-billplz"
)

type app struct {
	config   *config
	server   *http.ServeMux
	billplz  *billplz.Billplz
	template *WebTemplate
}

type config struct {
	billplz struct {
		apiKey        string
		apiURL        string
		collectionID  string
		callbackURL   string
		redirectURL   string
		xSignatureKey string
	}
}

func New() *app {
	a := new(app)
	a.server = http.NewServeMux()

	return a
}

func (a *app) Initialize() {
	a.InitConfig()
	a.InitBillPlz()
	a.InitTemplate()
	a.InitRoutes()
}

func (a *app) InitConfig() {
	cfg := new(config)
	cfg.billplz.apiKey = os.Getenv("BILLPLZ_API_KEY")
	cfg.billplz.apiURL = os.Getenv("BILLPLZ_API_URL")
	cfg.billplz.collectionID = os.Getenv("BILLPLZ_COLLECTION_ID")
	cfg.billplz.callbackURL = os.Getenv("BILLPLZ_CALLBACK_URL")
	cfg.billplz.redirectURL = os.Getenv("BILLPLZ_REDIRECT_URL")
	cfg.billplz.xSignatureKey = os.Getenv("BILLPLZ_X_SIGNATURE_KEY")

	a.config = cfg
}

func (a *app) InitBillPlz() {
	a.billplz = billplz.New(
		a.config.billplz.apiKey,
		a.config.billplz.apiURL,
		a.config.billplz.callbackURL,
		a.config.billplz.redirectURL,
		a.config.billplz.xSignatureKey,
	)
}

func (a *app) Start(addr string) error {
	return http.ListenAndServe(addr, a.server)
}
