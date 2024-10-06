package billplz

type Billplz struct {
	config *config
}

type config struct {
	apiKey        string
	apiURL        string
	callbackURL   string
	redirectURL   string
	xSignatureKey string
}

func New(apiKey, apiURL, callbackURL, redirectURL, xSignatureKey string) *Billplz {
	bp := new(Billplz)
	config := new(config)
	config.apiKey = apiKey
	config.apiURL = apiURL
	config.callbackURL = callbackURL
	config.redirectURL = redirectURL
	config.xSignatureKey = xSignatureKey
	bp.config = config
	return bp
}
