package v3_11

import (
	"net/http"
	"net/url"
)

type Client struct {
	httpClient *http.Client

	// Configurations
	BaseURL     *url.URL
	UserAgent   string
	AuthOptions *AuthOptions

	// APIs
	Apis *Apis
}
