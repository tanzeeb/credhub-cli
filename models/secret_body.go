package models

import (
	"encoding/json"

	"github.com/pivotal-cf/credhub-cli/util"
)

type SecretBody struct {
	ContentType string            `json:"type" binding:"required"`
	Value       interface{}       `json:"value,omitempty"`
	Overwrite   bool              `json:"overwrite"`
	Parameters  *SecretParameters `json:"parameters,omitempty"`
	UpdatedAt   string            `json:"updated_at,omitempty"`
}

func (body SecretBody) Terminal() string {
	result := ""
	switch body.ContentType {
	case "value", "password":
		result = util.BuildLineOfFixedLength("Value:", body.Value.(string)) + "\n"
		break
	case "certificate":
		cert := Certificate{}
		json.Unmarshal(marshalBackIntoJson(body.Value.(map[string]interface{})), &cert)
		result = cert.Terminal()
		break
	case "ssh", "rsa":
		rsaSsh := RsaSsh{}
		json.Unmarshal(marshalBackIntoJson(body.Value.(map[string]interface{})), &rsaSsh)
		result = rsaSsh.Terminal()
	}

	return result
}

func marshalBackIntoJson(value map[string]interface{}) []byte {
	item, _ := json.Marshal(value)
	return item
}
