package builders

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"

	"github.com/peakier/gen-gql-model/internal/util"
)

func genGqlType(typeTemps []*TypeTemp, outputDir string, packageName string, fileName string) error {
	err := os.MkdirAll(outputDir, os.ModePerm)
	if err != nil {
		return err
	}

	outputFile := filepath.Join(outputDir, fileName)

	resultStr, err := genTypeFileString(typeTemps, packageName)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(outputFile, []byte(util.FormatGoFile(resultStr)), 0755)
	if err != nil {
		return err
	}
	return nil
}

func genTypeFileString(typeTemps []*TypeTemp, packageName string) (string, error) {
	buf := []byte{}
	result := bytes.NewBuffer(buf)

	template := template.New("type-object")

	t, err := template.Parse(buildTypeTemplateStr(packageName))
	if err != nil {
		return "", err
	}

	err = t.Execute(result, typeTemps)
	if err != nil {
		return "", err
	}
	return result.String(), nil
}

type TypeTemp struct {
	TypeName string
	GoFields []*GoField
}

const typeTemplateStr = `
package %s
{{ range . }}
type {{ .TypeName }}  struct {
	{{ with .GoFields }}{{ range . }}{{ .GoName }} {{if .ListType}}[]{{end}}{{if .FieldTypeIsNull}}*{{end}}{{ .TypeName }} %sjson:"{{ .Name }}"%s
		{{ end }} {{ end }}
}
{{ end }} 
`

func buildTypeTemplateStr(packageName string) string {
	return fmt.Sprintf(typeTemplateStr, packageName, "`", "`")
}
