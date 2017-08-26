// CredHub credential value types
package values

type Value string

type Credential interface{}

type JSON interface{}

type Password string

type User struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password"`
}

type Certificate struct {
	Ca          string `json:"ca,omitempty" yaml:",omitempty"`
	CaName      string `json:"ca_name,omitempty" yaml:"ca_name,omitempty"`
	Certificate string `json:"certificate"`
	PrivateKey  string `json:"private_key" yaml:"private_key"`
}

type RSA struct {
	PrivateKey string `json:"private_key" yaml:"private_key"`
	PublicKey  string `json:"public_key" yaml:"public_key"`
}

type SSH struct {
	PrivateKey string `json:"private_key" yaml:"private_key"`
	PublicKey  string `json:"public_key" yaml:"public_key"`
}
