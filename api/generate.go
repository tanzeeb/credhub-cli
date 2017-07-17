package api

import (
	"github.com/cloudfoundry-incubator/credhub-cli/actions"
	"github.com/cloudfoundry-incubator/credhub-cli/client"
	"github.com/cloudfoundry-incubator/credhub-cli/config"
	"github.com/cloudfoundry-incubator/credhub-cli/models"
	"github.com/cloudfoundry-incubator/credhub-cli/repositories"
)

func Generate(
	credentialIdentifier string,
	credentialType string,
	noOverwrite bool,
	outputJson bool,
	username string,
	length int,
	includeSpecial bool,
	excludeNumber bool,
	excludeUpper bool,
	excludeLower bool,
	sshComment string,
	keyLength int,
	duration int,
	commonName string,
	organization string,
	organizationUnit string,
	locality string,
	state string,
	country string,
	alternativeName []string,
	keyUsage []string,
	extendedKeyUsage []string,
	ca string,
	isCA bool,
	selfSign bool,
) (models.CredentialResponse, error) {

	cfg := config.ReadConfig()
	repository := repositories.NewCredentialRepository(client.NewHttpClient(cfg))

	parameters := models.GenerationParameters{
		IncludeSpecial:   includeSpecial,
		ExcludeNumber:    excludeNumber,
		ExcludeUpper:     excludeUpper,
		ExcludeLower:     excludeLower,
		Length:           length,
		CommonName:       commonName,
		Organization:     organization,
		OrganizationUnit: organizationUnit,
		Locality:         locality,
		State:            state,
		Country:          country,
		AlternativeName:  alternativeName,
		ExtendedKeyUsage: extendedKeyUsage,
		KeyUsage:         keyUsage,
		KeyLength:        keyLength,
		Duration:         duration,
		Ca:               ca,
		SelfSign:         selfSign,
		IsCA:             isCA,
		SshComment:       sshComment,
	}

	var value *models.ProvidedValue
	if len(username) > 0 {
		value = &models.ProvidedValue{
			Username: username,
		}
	}

	action := actions.NewAction(repository, &cfg)
	request := client.NewGenerateCredentialRequest(cfg, credentialIdentifier, parameters, value, credentialType, !noOverwrite)
	credential, err := action.DoAction(request, credentialIdentifier)

	return credential.(models.CredentialResponse), err
}