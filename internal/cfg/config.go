package cfg

import (
	"github.com/go-yaml/yaml"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
)

type StoreType string

const (
	// StoreTypeMongoDB stores the document metadata and blob in MongoDB
	StoreTypeMongoDB StoreType = "mongodb"

	// StoreTypeMysql uses MySQL as the backing store and stores the documents as a blob
	StoreTypeMysql = "mysql"

	// StoreTypeAWS uses MySQL to store the metadata information and
	// AWS S3 to store the document content
	StoreTypeAWS = "aws"
)

type AppConfig struct {
	StoreType StoreType

	// StoreConfig can accept any structure. Parsing and validation
	// is delegated to the implementing engines
	//
	// Returns a map[interface{}]interface{}
	StoreConfig interface{}
}

func ParseConfigFromYamlFile(configFile string) AppConfig {
	fileContentAsBytes, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Fatal(err)
	}

	return parseConfigYaml(fileContentAsBytes)
}

func ParseConfigFromYamlString(yamlString string) AppConfig {
	return parseConfigYaml([]byte(yamlString))
}

func parseConfigYaml(yamlConfig []byte) AppConfig {
	appCfg := AppConfig{}

	err := yaml.Unmarshal(yamlConfig, &appCfg)
	if err != nil {
		log.Fatal(err)
	}

	return appCfg
}
