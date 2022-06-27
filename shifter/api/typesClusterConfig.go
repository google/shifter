package api

type ClusterConfig struct {
	ConnectionName string `json:"connectionName"`
	BaseUrl        string `json:"baseUrl"`
	BearerToken    string `json:"bearerToken"`
	Username       string `json:"username"`
	Password       string `json:"password"`
}
