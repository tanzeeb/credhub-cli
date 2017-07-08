package models

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Credential struct {
	Name  string      `yaml:"name"`
	Type  string      `yaml:"type"`
	Value interface{} `yaml:"value"`
}

type CredentialBulkImport struct {
	Credentials []Credential `yaml:"credentials"`
}

func (credentialBulkImport *CredentialBulkImport) ReadFile(filepath string) error {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return err
	}
	return credentialBulkImport.ReadBytes(data)
}

func (credentialBulkImport *CredentialBulkImport) ReadBytes(data []byte) error {
	err := yaml.Unmarshal(data, credentialBulkImport)

	// Having trouble because we're trying to convert in place and
	// "credential" here is a copy, not a reference
	for i, credential := range credentialBulkImport.Credentials {
		switch valueAsMap := credential.Value.(type) {
		case map[interface{}]interface{}:
			credentialBulkImport.Credentials[i].Value = convertToMapStringInterface(valueAsMap)
		default:
		}
	}

	return err
}

func convertToMapStringInterface(valueAsMap map[interface{}]interface{}) map[string]interface{} {
	mapStringInterface := make(map[string]interface{})
	for key, value := range valueAsMap {
		var desiredMapValue interface{}
		switch nestedValue := value.(type) {
		case string:
			desiredMapValue = nestedValue
		case map[interface{}]interface{}:
			desiredMapValue = convertToMapStringInterface(nestedValue)
		case []interface{}:
			desiredMapValue = convertInterfaceArrayValues(nestedValue)
		}

		mapStringInterface[key.(string)] = desiredMapValue
	}
	return mapStringInterface
}

func convertInterfaceArrayValues(array []interface{}) []interface{} {
	for i, value := range array {
		switch typedValue := value.(type) {
		case string:
			array[i] = typedValue
		case map[interface{}]interface{}:
			array[i] = convertToMapStringInterface(typedValue)
			// need case for array
		}
	}
	return array
}
