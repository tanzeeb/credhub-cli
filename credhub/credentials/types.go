// CredHub credential types
package credentials

import (
	"encoding/json"

	yaml "gopkg.in/yaml.v2"

	"github.com/cloudfoundry-incubator/credhub-cli/credhub/credentials/values"
)

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
	Metadata `yaml:",inline"`
	Value    values.Credential `json:"value"`
}

func (c Credential) MarshalYAML() (interface{}, error) {
	return mapSlice(c.Id, c.Name, c.Type, c.Value, c.VersionCreatedAt), nil
}

// A Value type credential
type Value struct {
	Metadata `yaml:",inline"`
	Value    values.Value `json:"value"`
}

func (v Value) MarshalYAML() (interface{}, error) {
	return mapSlice(v.Id, v.Name, v.Type, v.Value, v.VersionCreatedAt), nil
}

// A JSON type credential
//
// Value will need to be further unmarshalled by https://golang.org/pkg/encoding/json/#Unmarshal
type JSON struct {
	Metadata `yaml:",inline"`
	Value    json.RawMessage `json:"value"`
}

func (j JSON) MarshalYAML() (interface{}, error) {
	var value interface{}

	if err := json.Unmarshal(j.Value, &value); err != nil {
		return nil, err
	}

	return mapSlice(j.Id, j.Name, j.Type, value, j.VersionCreatedAt), nil
}

// A Password type credential
type Password struct {
	Metadata `yaml:",inline"`
	Value    values.Password `json:"value"`
}

func (p Password) MarshalYAML() (interface{}, error) {
	return mapSlice(p.Id, p.Name, p.Type, p.Value, p.VersionCreatedAt), nil
}

// A User type credential
type User struct {
	Metadata `yaml:",inline"`
	Value    struct {
		values.User  `yaml:",inline"`
		PasswordHash string `json:"password_hash" yaml:"password_hash"`
	} `json:"value"`
}

func (u User) MarshalYAML() (interface{}, error) {
	value := yaml.MapSlice{
		yaml.MapItem{Key: "password", Value: u.Value.Password},
		yaml.MapItem{Key: "password_hash", Value: u.Value.PasswordHash},
	}

	if u.Value.Username == "" {
		value = append(value, yaml.MapItem{Key: "username", Value: nil})
	} else {
		value = append(value, yaml.MapItem{Key: "username", Value: u.Value.Username})
	}

	return mapSlice(u.Id, u.Name, u.Type, value, u.VersionCreatedAt), nil
}

// A Certificate type credential
type Certificate struct {
	Metadata `yaml:",inline"`
	Value    values.Certificate `json:"value"`
}

func (c Certificate) MarshalYAML() (interface{}, error) {
	return mapSlice(c.Id, c.Name, c.Type, c.Value, c.VersionCreatedAt), nil
}

// An RSA type credential
type RSA struct {
	Metadata `yaml:",inline"`
	Value    values.RSA `json:"value"`
}

func (r RSA) MarshalYAML() (interface{}, error) {
	return mapSlice(r.Id, r.Name, r.Type, r.Value, r.VersionCreatedAt), nil
}

// An SSH type credential
type SSH struct {
	Metadata `yaml:",inline"`
	Value    struct {
		values.SSH           `yaml:",inline"`
		PublicKeyFingerprint string `json:"public_key_fingerprint,omitempty" yaml:"public_key_fingerprint,omitempty"`
	} `json:"value"`
}

func (s SSH) MarshalYAML() (interface{}, error) {
	return mapSlice(s.Id, s.Name, s.Type, s.Value, s.VersionCreatedAt), nil
}

func mapSlice(i string, n string, t string, v interface{}, c string) yaml.MapSlice {
	return yaml.MapSlice{
		yaml.MapItem{Key: "id", Value: i},
		yaml.MapItem{Key: "name", Value: n},
		yaml.MapItem{Key: "type", Value: t},
		yaml.MapItem{Key: "value", Value: v},
		yaml.MapItem{Key: "version_created_at", Value: c},
	}
}
