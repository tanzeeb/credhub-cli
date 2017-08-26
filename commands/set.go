package commands

import (
	"encoding/json"
	"fmt"

	yaml "gopkg.in/yaml.v2"

	"net/http"

	"github.com/cloudfoundry-incubator/credhub-cli/client"
	"github.com/cloudfoundry-incubator/credhub-cli/config"
	"github.com/cloudfoundry-incubator/credhub-cli/credhub"
	"github.com/cloudfoundry-incubator/credhub-cli/credhub/credentials/values"

	"bufio"
	"os"
	"strings"

	"github.com/cloudfoundry-incubator/credhub-cli/errors"
	"github.com/cloudfoundry-incubator/credhub-cli/util"
)

type SetCommand struct {
	CredentialIdentifier string `short:"n" required:"yes" long:"name" description:"Name of the credential to set"`
	Type                 string `short:"t" long:"type" description:"Sets the credential type. Valid types include 'value', 'json', 'password', 'user', 'certificate', 'ssh' and 'rsa'. Type-specific values are set with the following flags (supported types prefixed)."`
	NoOverwrite          bool   `short:"O" long:"no-overwrite" description:"Credential is not modified if stored value already exists"`
	Value                string `short:"v" long:"value" description:"[Value, JSON] Sets the value for the credential"`
	CaName               string `short:"m" long:"ca-name" description:"[Certificate] Sets the root CA to a stored CA credential"`
	Root                 string `short:"r" long:"root" description:"[Certificate] Sets the root CA from file or value"`
	Certificate          string `short:"c" long:"certificate" description:"[Certificate] Sets the certificate from file or value"`
	Private              string `short:"p" long:"private" description:"[Certificate, SSH, RSA] Sets the private key from file or value"`
	Public               string `short:"u" long:"public" description:"[SSH, RSA] Sets the public key from file or value"`
	Username             string `short:"z" long:"username" description:"[User] Sets the username value of the credential"`
	Password             string `short:"w" long:"password" description:"[Password, User] Sets the password value of the credential"`
	OutputJson           bool   `          long:"output-json" description:"Return response in JSON format"`
}

func (cmd SetCommand) Execute([]string) error {
	cmd.Type = strings.ToLower(cmd.Type)

	if cmd.Type == "" {
		return errors.NewSetEmptyTypeError()
	}

	if cmd.Value == "" && (cmd.Type == "value" || cmd.Type == "json") {
		promptForInput("value: ", &cmd.Value)
	}

	if cmd.Password == "" && (cmd.Type == "password" || cmd.Type == "user") {
		promptForInput("password: ", &cmd.Password)
	}

	cfg := config.ReadConfig()
	ch, err := cfg.ApiClient()
	if err != nil {
		return err
	}
	//repository := repositories.NewCredentialRepository(client.NewHttpClient(cfg))

	//action := actions.NewAction(repository, &cfg)
	//request, err := MakeRequest(cmd, cfg)
	//if err != nil {
	//return err
	//}

	//credential, err := action.DoAction(request, cmd.CredentialIdentifier)
	credential, err := setCredential(ch, cmd)
	if err != nil {
		return err
	}

	prettyPrint(credential, cmd.OutputJson)

	return nil
}

func setCredential(ch *credhub.CredHub, cmd SetCommand) (interface{}, error) {
	switch cmd.Type {
	case "ssh":
		publicKey, err := util.ReadFileOrStringFromField(cmd.Public)
		if err != nil {
			return nil, err
		}

		privateKey, err := util.ReadFileOrStringFromField(cmd.Private)
		if err != nil {
			return nil, err
		}

		value := values.SSH{
			PublicKey:  publicKey,
			PrivateKey: privateKey,
		}

		return ch.SetSSH(cmd.CredentialIdentifier, value, !cmd.NoOverwrite)
	case "rsa":
		publicKey, err := util.ReadFileOrStringFromField(cmd.Public)
		if err != nil {
			return nil, err
		}

		privateKey, err := util.ReadFileOrStringFromField(cmd.Private)
		if err != nil {
			return nil, err
		}

		value := values.RSA{
			PublicKey:  publicKey,
			PrivateKey: privateKey,
		}

		return ch.SetRSA(cmd.CredentialIdentifier, value, !cmd.NoOverwrite)
	case "certificate":
		root, err := util.ReadFileOrStringFromField(cmd.Root)
		if err != nil {
			return nil, err
		}

		certificate, err := util.ReadFileOrStringFromField(cmd.Certificate)
		if err != nil {
			return nil, err
		}

		privateKey, err := util.ReadFileOrStringFromField(cmd.Private)
		if err != nil {
			return nil, err
		}

		value := values.Certificate{
			Ca:          root,
			CaName:      cmd.CaName,
			Certificate: certificate,
			PrivateKey:  privateKey,
		}

		return ch.SetCertificate(cmd.CredentialIdentifier, value, !cmd.NoOverwrite)
	case "user":
		value := values.User{
			Username: cmd.Username,
			Password: cmd.Password,
		}
		return ch.SetUser(cmd.CredentialIdentifier, value, !cmd.NoOverwrite)
	case "password":
		value := values.Password(cmd.Password)
		return ch.SetPassword(cmd.CredentialIdentifier, value, !cmd.NoOverwrite)
	case "json":
		value := json.RawMessage(cmd.Value)
		return ch.SetJSON(cmd.CredentialIdentifier, value, !cmd.NoOverwrite)
	default:
		value := values.Value(cmd.Value)
		return ch.SetValue(cmd.CredentialIdentifier, value, !cmd.NoOverwrite)
	}
}

func MakeRequest(cmd SetCommand, config config.Config) (*http.Request, error) {
	var request *http.Request
	if cmd.Type == "ssh" || cmd.Type == "rsa" {
		var err error

		publicKey, err := util.ReadFileOrStringFromField(cmd.Public)
		if err != nil {
			return nil, err
		}

		privateKey, err := util.ReadFileOrStringFromField(cmd.Private)
		if err != nil {
			return nil, err
		}

		request = client.NewSetRsaSshRequest(config, cmd.CredentialIdentifier, cmd.Type, publicKey, privateKey, !cmd.NoOverwrite)
	} else if cmd.Type == "certificate" {
		var err error

		root, err := util.ReadFileOrStringFromField(cmd.Root)
		if err != nil {
			return nil, err
		}

		certificate, err := util.ReadFileOrStringFromField(cmd.Certificate)
		if err != nil {
			return nil, err
		}

		privateKey, err := util.ReadFileOrStringFromField(cmd.Private)
		if err != nil {
			return nil, err
		}

		request = client.NewSetCertificateRequest(config, cmd.CredentialIdentifier, root, cmd.CaName, certificate, privateKey, !cmd.NoOverwrite)
	} else if cmd.Type == "user" {
		request = client.NewSetUserRequest(config, cmd.CredentialIdentifier, cmd.Username, cmd.Password, !cmd.NoOverwrite)
	} else if cmd.Type == "password" {
		request = client.NewSetCredentialRequest(config, cmd.Type, cmd.CredentialIdentifier, cmd.Password, !cmd.NoOverwrite)
	} else if cmd.Type == "json" {
		request = client.NewSetJsonCredentialRequest(config, cmd.Type, cmd.CredentialIdentifier, cmd.Value, !cmd.NoOverwrite)
	} else {
		request = client.NewSetCredentialRequest(config, cmd.Type, cmd.CredentialIdentifier, cmd.Value, !cmd.NoOverwrite)
	}

	return request, nil
}

func promptForInput(prompt string, value *string) {
	fmt.Printf(prompt)
	reader := bufio.NewReader(os.Stdin)
	val, _ := reader.ReadString('\n')
	*value = string(strings.TrimSpace(val))
}

func prettyPrint(v interface{}, asJson bool) {
	var b []byte

	if asJson {
		b, _ = json.MarshalIndent(v, "", "\t")
	} else {
		b, _ = yaml.Marshal(v)
	}

	fmt.Println(string(b))
}
