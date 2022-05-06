package v3_11

import (
	"encoding/json"
	"fmt"
	"log"
)

type DeploymentConfigs struct {
	Client     *Client
	Kind       string             `json:"kind"`
	ApiVersion string             `json:"apiVersion"`
	Metadata   Metadata           `json:"metadata"`
	Items      []DeploymentConfig `json:"items"`
}

func (p DeploymentConfigs) Output() {
	out, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Printf(string(out))
}
