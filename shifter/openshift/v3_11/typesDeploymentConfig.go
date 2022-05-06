package v3_11

import (
	"encoding/json"
	"fmt"
	"log"

	osTpyes "github.com/openshift/api/apps/v1"
)

type DeploymentConfig struct {
	Client     *Client
	Kind       string                         `json:"kind"`
	ApiVersion string                         `json:"apiVersion"`
	Metadata   Metadata                       `json:"metadata"`
	Spec       osTpyes.DeploymentConfigSpec   `json:"spec"`
	Status     osTpyes.DeploymentConfigStatus `json:"status"`
}

func (p DeploymentConfig) Output() {
	out, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Printf(string(out))
}
