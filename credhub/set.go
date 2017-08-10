package credhub

import (
	"net/http"
	"encoding/json"
	"io/ioutil"

	"github.com/cloudfoundry-incubator/credhub-cli/credhub/credentials"
	"github.com/cloudfoundry-incubator/credhub-cli/credhub/credentials/values"
)

// Sets a Value credential with a user-provided value.
func (ch *CredHub) SetValue(name string, value values.Value, overwrite bool) (credentials.Value, error) {
	panic("Not implemented")
}

// Sets a JSON credential with a user-provided value.
func (ch *CredHub) SetJSON(name string, value values.JSON, overwrite bool) (credentials.JSON, error) {
	panic("Not implemented")
}

// Sets a Password credential with a user-provided value.
func (ch *CredHub) SetPassword(name string, value values.Password, overwrite bool) (credentials.Password, error) {
	panic("Not implemented")
}

// Sets a User credential with a user-provided value.
func (ch *CredHub) SetUser(name string, value values.User, overwrite bool) (credentials.User, error) {
	panic("Not implemented")
}

// Sets a Certificate credential with a user-provided value.
func (ch *CredHub) SetCertificate(name string, value values.Certificate, overwrite bool) (credentials.Certificate, error) {
	var cred credentials.Certificate

	requestBody := map[string]interface{}{}
	requestBody["name"] = name
	requestBody["type"] = "certificate"
	requestBody["value"] = value
	resp, err := ch.Request(http.MethodPut, "/api/v1/data", requestBody)

	if err != nil {
		return cred, err
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &cred)

	if err != nil {
		return cred, err
	}

	return cred, nil
}

// Sets an RSA credential with a user-provided value.
func (ch *CredHub) SetRSA(name string, value values.RSA, overwrite bool) (credentials.RSA, error) {
	panic("Not implemented")
}

// Sets an SSH credential with a user-provided value.
func (ch *CredHub) SetSSH(name string, value values.SSH, overwrite bool) (credentials.SSH, error) {
	panic("Not implemented")
}
