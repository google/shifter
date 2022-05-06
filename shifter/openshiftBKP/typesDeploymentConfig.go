package openshift

import (
	"encoding/json"
	"fmt"
	"log"
)

type DeploymentConfig struct {
	Kind       string   `json:"kind"`
	ApiVersion string   `json:"apiVersion"`
	Metadata   Metadata `json:"metadata"`
	Spec       Spec     `json:"spec"`
	Status     Status   `json:"status"`
}

func (p DeploymentConfig) Output() {
	out, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Printf(string(out))
}
