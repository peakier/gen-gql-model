.DEFAULT_GOAL := run

run: 
	go run github.com/peakier/gen-gql-model jsonmode --config-file _example/configs.yaml

