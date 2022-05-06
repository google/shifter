package openshift

import (
	"encoding/json"
	"fmt"
	"log"
)

type Projects struct {
	Kind       string    `json:"kind"`
	ApiVersion string    `json:"apiVersion"`
	Metadata   Metadata  `json:"metadata"`
	Items      []Project `json:"items"`
}

func (p Projects) Output() {
	out, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Printf(string(out))
}
