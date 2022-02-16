package generator

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/hasura/go-graphql-client"
	"github.com/peakier/gen-gql-model/internal/builders"
	"github.com/peakier/gen-gql-model/internal/model"
	"github.com/peakier/gen-gql-model/internal/parsers"
	"gopkg.in/yaml.v2"
)

func Generate(cfgPath string, fromUrl bool) error {
	conf := getConf(cfgPath)
	var err error

	for _, graqhqlConfig := range conf.GraqhqlConfigs {
		var querySchema *model.QuerySchema
		if fromUrl {
			querySchema, err = readConfigFromGqlEndpoint(graqhqlConfig.Path)

		} else {
			querySchema, err = readConfigFromJsonFile(graqhqlConfig.Path)
		}
		gqlBuilder := builders.NewGQLBuilder(querySchema)
		if err != nil {
			return err
		}
		err = gqlBuilder.Run(graqhqlConfig.OutputDir, graqhqlConfig.PackageName)
		if err != nil {
			return err
		}
	}

	return err
}

func getConf(filePath string) parsers.Config {
	yamlFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Panicf("Unable to read config file: '%s'\n#%v ", filePath, err)
	}

	conf := parsers.Config{}
	err = yaml.Unmarshal(yamlFile, &conf)
	if err != nil {
		log.Panicf("Unmarshal: %v", err)
	}

	return conf
}

func readConfigFromJsonFile(path string) (*model.QuerySchema, error) {

	type _data struct {
		Q model.QuerySchema `json:"data"`
	}
	var querySchema _data

	jsonFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(byteValue, &querySchema)
	if err != nil {
		return nil, err
	}

	return &querySchema.Q, nil
}

func readConfigFromGqlEndpoint(endpoint string) (*model.QuerySchema, error) {
	querySchema := &model.QuerySchema{}
	var httpClient = http.Client{Timeout: 90 * time.Second}
	var gqlClient = graphql.NewClient(endpoint, &httpClient)
	err := gqlClient.Query(context.Background(), querySchema, nil)
	if err != nil {
		return nil, err
	}

	return querySchema, nil
}
