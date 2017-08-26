package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"

	"github.com/cloudfoundry-incubator/credhub-cli/credhub"
	"github.com/cloudfoundry-incubator/credhub-cli/credhub/auth"
	"github.com/cloudfoundry-incubator/credhub-cli/util"
)

const AuthClient = "credhub_cli"
const AuthPassword = ""

type Config struct {
	ApiURL             string
	AuthURL            string
	AccessToken        string
	RefreshToken       string
	InsecureSkipVerify bool
	CaCerts            []string
}

// FIXME Tests!
func (c Config) ApiClient() (*credhub.CredHub, error) {
	var options []credhub.Option

	accessToken := c.AccessToken
	if accessToken == "revoked" {
		accessToken = ""
	}

	refreshToken := c.RefreshToken
	if refreshToken == "revoked" {
		refreshToken = ""
	}

	clientId := AuthClient
	if val := os.Getenv("CREDHUB_CLIENT"); val != "" {
		clientId = val
	}

	clientSecret := AuthPassword
	if val := os.Getenv("CREDHUB_SECRET"); val != "" {
		clientSecret = val
	}

	if c.InsecureSkipVerify {
		options = append(options, credhub.SkipTLSValidation())
	}
	options = append(options, credhub.AuthURL(c.AuthURL))
	options = append(options, credhub.CaCerts(c.CaCerts...))
	options = append(options, credhub.Auth(auth.Uaa(clientId, clientSecret, "", "", accessToken, refreshToken)))

	return credhub.New(c.ApiURL, options...)
}

func ConfigDir() string {
	return path.Join(userHomeDir(), ".credhub")
}

func ConfigPath() string {
	return path.Join(ConfigDir(), "config.json")
}

func ReadConfig() Config {
	c := Config{}

	data, err := ioutil.ReadFile(ConfigPath())
	if err != nil {
		return c
	}

	json.Unmarshal(data, &c)

	return c
}

func WriteConfig(c Config) error {
	err := makeDirectory()
	if err != nil {
		return err
	}

	data, err := json.Marshal(c)
	if err != nil {
		return err
	}

	configPath := ConfigPath()
	return ioutil.WriteFile(configPath, data, 0600)
}

func RemoveConfig() error {
	return os.Remove(ConfigPath())
}

func (cfg *Config) UpdateTrustedCAs(caCerts []string) error {
	certs := []string{}

	for _, cert := range caCerts {
		certContents, err := util.ReadFileOrStringFromField(cert)
		if err != nil {
			return err
		}
		certs = append(certs, certContents)
	}

	cfg.CaCerts = certs

	return nil
}
