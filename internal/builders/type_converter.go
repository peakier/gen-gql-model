package builders

import "fmt"

type TypeConverter struct {
	gotypeBuilderByGqlType map[string]*GotypeBuilder
}

func NewTypeConverter() *TypeConverter {
	gotypeBuilderByGqlType := map[string]*GotypeBuilder{
		"Boolean": {
			Name: "bool",
		},
		"String": {
			Name: "string",
		},
		"Int": {
			Name: "int",
		},
		"ID": {
			Name: "int",
		},
		"UUID": {
			Name: "string",
		},

		"Time": {
			Name: "time.Time",
		},
		"Decimal": {
			Name: "decimal.Decimal",
		},
		"Float": {
			Name: "float64",
		},
	}
	return &TypeConverter{
		gotypeBuilderByGqlType: gotypeBuilderByGqlType,
	}
}

func (t *TypeConverter) GetGoTypeFromGqlType(gqlType string) string {
	goType := "interface{}"

	gotypeBuilder := t.gotypeBuilderByGqlType[gqlType]

	if gotypeBuilder != nil {
		goType = gotypeBuilder.Name
	} else {
		warningStr := fmt.Sprintf("warning : can not find %s gqltype", gqlType)
		fmt.Println(warningStr)
	}
	return goType
}

type GotypeBuilder struct {
	Name   string
	Import string
}
