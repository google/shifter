package openshift

import (
	"net/http"
	"net/url"
)

type AuthOptions struct {
	BearerToken string
}

type Client struct {
	httpClient *http.Client

	// Configurations
	BaseURL     *url.URL
	UserAgent   string
	AuthOptions *AuthOptions

	// Service Objects
	Apis *Apis
}

type Apis struct {
	client *Client
}

type Metadata struct {
	Name              string      `json:"name"`
	SelfLink          string      `json:"selfLink"`
	Uid               string      `json:"uid"`
	ResourceVersion   string      `json:"resourceVersion"`
	Generation        int         `json:"generation"`
	CreationTimestamp string      `json:"creationTimestamp"`
	Labels            Labels      `json:"labels"`
	Annotations       Annotations `json:"annotations"`
}

type Labels struct {
}

type Spec struct {
	Finalizers []string `json:"finalizers"`
}

type Status struct {
	Phase string `json:"phase"`
}

type Annotations struct {
	Annotation1 string `json:"openshift.io/description"`
	Annotation2 string `json:"openshift.io/display-name"`
	Annotation3 string `json:"openshift.io/requester"`
	Annotation4 string `json:"openshift.io/sa.scc.mcs"`
	Annotation5 string `json:"openshift.io/sa.scc.supplemental-groups"`
	Annotation6 string `json:"openshift.io/sa.scc.uid-range"`
}
