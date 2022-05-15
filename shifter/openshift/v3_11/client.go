package v3_11

import (
	"crypto/tls"
	"net/http"
)

func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	return ClientInit(httpClient)
}

func ClientInit(httpClient *http.Client) *Client {

	//Handle x509: certificate signed by unknown authority Error for http instead of https
	httpClient.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	c := &Client{httpClient: httpClient}
	c.Apis = &Apis{
		Project:           &Project{Client: c},
		Projects:          &Projects{Client: c},
		DeploymentConfig:  &DeploymentConfig{Client: c},
		DeploymentConfigs: &DeploymentConfigs{Client: c},
	}
	// Add Apis Service Object
	//c.Apis = &Apis{client: c}
	return c
}
