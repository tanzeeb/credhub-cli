// CredHub credential types
package credentials

import "github.com/cloudfoundry-incubator/credhub-cli/credhub/credentials/values"

// Base fields of a credential
type Base struct {
	Name             string `json:"name"`
	VersionCreatedAt string `json:"version_created_at" yaml:"version_created_at"`
}

type Metadata struct {
	Base `yaml:",inline"`
	Id   string `json:"id"`
	Type string `json:"type"`
}

// A generic credential
//
// Used when the Type of the credential is not known ahead of time.
//
// Value will be as unmarshalled by https://golang.org/pkg/encoding/json/#Unmarshal
type Credential struct {
	Metadata
	Value interface{} `json:"value"`
}

// A Value type credential
type Value struct {
	Metadata
	Value values.Value
}

// A JSON type credential
type JSON struct {
	Metadata
	Value values.JSON
}

// A Password type credential
type Password struct {
	Metadata
	Value values.Password
}

// A User type credential
type User struct {
	Metadata
	Value struct {
		values.User
		PasswordHash string
	}
}

// A Certificate type credential
type Certificate struct {
	Metadata `yaml:",inline"`
	Value    values.Certificate `json:"value"`
}

// An RSA type credential
type RSA struct {
	Metadata
	Value values.RSA
}

// An SSH type credential
type SSH struct {
	Metadata
	Value values.SSH
}