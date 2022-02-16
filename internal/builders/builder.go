package builders

import (
	"github.com/peakier/gen-gql-model/internal/model"
)

type GQLBuilder struct {
	QuerySchema *model.QuerySchema
}

func NewGQLBuilder(querySchema *model.QuerySchema) *GQLBuilder {
	return &GQLBuilder{QuerySchema: querySchema}
}

func (g *GQLBuilder) Run(outputDir string, packageName string) error {

	enumTemps := []*EnumTemp{}
	inputTypeTemps := []*TypeTemp{}
	typeTemps := []*TypeTemp{}

	for _, gqlType := range g.QuerySchema.GraphQLSchema.Types {
		switch *gqlType.Kind {
		case inputObjectKind:
			gqlFields := []*model.Field{}
			for _, inputField := range gqlType.InputFields {
				gqlFields = append(gqlFields, inputField.ToField())
			}

			fields := getGqlField(gqlFields)
			inputTypeTemps = append(inputTypeTemps, &TypeTemp{
				TypeName: gqlType.Name,
				GoFields: fields,
			})
		case enumKind:
			enumTemps = append(enumTemps, &EnumTemp{
				EnumName:   gqlType.Name,
				EnumValues: gqlType.EnumValues,
			})

		case objectKind:
			fields := getGqlField(gqlType.Fields)
			typeTemps = append(typeTemps, &TypeTemp{
				TypeName: gqlType.Name,
				GoFields: fields,
			})
		}

	}

	err := genGqlEnum(enumTemps, outputDir, packageName)
	if err != nil {
		return err
	}

	err = genGqlType(inputTypeTemps, outputDir, packageName, "input_gen.go")
	if err != nil {
		return err
	}

	err = genGqlType(typeTemps, outputDir, packageName, "type_gen.go")
	if err != nil {
		return err
	}

	return nil
}
