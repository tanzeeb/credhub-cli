package commands

import (
	"fmt"

	"github.com/pivotal-cf/cm-cli/actions"
	"github.com/pivotal-cf/cm-cli/client"
	"github.com/pivotal-cf/cm-cli/config"
	"github.com/pivotal-cf/cm-cli/models"
	"github.com/pivotal-cf/cm-cli/repositories"
)

type GenerateCommand struct {
	SecretIdentifier string `short:"n" required:"yes" long:"name" description:"Selects the secret being set"`
	ContentType      string `short:"t" long:"type" description:"Sets the type of secret to store or generate. Default: 'value'"`
	Length           int    `short:"l" long:"length" description:"Sets length of generated value (Default: 20)"`
	ExcludeSpecial   bool   `long:"exclude-special" description:"Exclude special characters from generated value"`
	ExcludeNumber    bool   `long:"exclude-number" description:"Exclude number characters from generated value"`
	ExcludeUpper     bool   `long:"exclude-upper" description:"Exclude upper alpha characters from generated value"`
	ExcludeLower     bool   `long:"exclude-lower" description:"Exclude lower alpha characters from generated value"`
}

func (cmd GenerateCommand) Execute([]string) error {
	if cmd.ContentType == "" {
		cmd.ContentType = "value"
	}

	secretRepository := repositories.NewSecretRepository(client.NewHttpClient())

	parameters := models.SecretParameters{
		ExcludeSpecial: cmd.ExcludeSpecial,
		ExcludeNumber:  cmd.ExcludeNumber,
		ExcludeUpper:   cmd.ExcludeUpper,
		ExcludeLower:   cmd.ExcludeLower,
		Length:         cmd.Length,
	}

	config := config.ReadConfig()
	action := actions.NewAction(secretRepository, config)
	request := client.NewGenerateSecretRequest(config.ApiURL, cmd.SecretIdentifier, parameters, cmd.ContentType)
	secret, err := action.DoAction(request, cmd.SecretIdentifier)

	if err != nil {
		return err
	}

	fmt.Println(secret)

	return nil
}