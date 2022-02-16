package builders

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"

	"github.com/peakier/gen-gql-model/internal/model"
	"github.com/peakier/gen-gql-model/internal/util"
)

func genGqlEnum(enumTemps []*EnumTemp, outputDir string, packageName string) error {
	err := os.MkdirAll(outputDir, os.ModePerm)
	if err != nil {
		return err
	}

	outputFile := filepath.Join(outputDir, "enum_gen.go")

	resultStr, err := genEnumFileString(enumTemps, packageName)
	if err != nil {
		return err
	}

	ioutil.WriteFile(outputFile, []byte(util.FormatGoFile(resultStr)), 0755)

	return nil
}

func genEnumFileString(enumTemps []*EnumTemp, packageName string) (string, error) {
	buf := []byte{}
	result := bytes.NewBuffer(buf)

	template := template.New("enum-object")

	t, err := template.Parse(fmt.Sprintf(enumTemplateStr, packageName))
	if err != nil {
		return "", err
	}

	err = t.Execute(result, enumTemps)
	if err != nil {
		return "", err
	}

	return result.String(), nil
}

type EnumTemp struct {
	EnumName   string
	EnumValues []*model.EnumValue
}

const enumTemplateStr = `
package %s

{{ range . }}
	{{$enumName := .EnumName}}
	type {{ $enumName }}  string
	const (
		{{ with .EnumValues }}{{ range . }}{{$enumName}}{{ .Name }} {{$enumName}} = "{{ .Name }}"
		{{ end }} {{ end }}
	)
{{ end }} 
`
