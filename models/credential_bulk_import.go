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

	for i, credential := range credentialBulkImport.Credentials {
		switch valueAsMap := credential.Value.(type) {
		case map[interface{}]interface{}:
			credentialBulkImport.Credentials[i].Value = unpackMapIntoStringToInterface(valueAsMap)
		default:
		}
	}

	return err
}

func unpackValueIntoStringToInterface(value interface{}) interface{} {
	var unpackedValue interface{}
	switch typedValue := value.(type) {
	case map[interface{}]interface{}:
		unpackedValue = unpackMapIntoStringToInterface(typedValue)
	case []interface{}:
		unpackedValue = unpackArrayIntoStringToInterface(typedValue)
	default:
		unpackedValue = value
	}
	return unpackedValue
}

func unpackMapIntoStringToInterface(interfaceToInterfaceMap map[interface{}]interface{}) map[string]interface{} {
	stringToInterfaceMap := make(map[string]interface{})
	for key, value := range interfaceToInterfaceMap {
		stringToInterfaceMap[key.(string)] = unpackValueIntoStringToInterface(value)
	}
	return stringToInterfaceMap
}

func unpackArrayIntoStringToInterface(array []interface{}) []interface{} {
	for i, value := range array {
		array[i] = unpackValueIntoStringToInterface(value)
	}
	return array
}
