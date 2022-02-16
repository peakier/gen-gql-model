package builders

import (
	"fmt"

	"github.com/iancoleman/strcase"
	"github.com/peakier/gen-gql-model/internal/model"
)

const nonNullKind = "NON_NULL"
const scalarKind = "SCALAR"
const enumKind = "ENUM"
const interfaceKind = "INTERFACE"
const objectKind = "OBJECT"
const inputObjectKind = "INPUT_OBJECT"
const listKind = "LIST"

type GqlTypeModel struct {
	Name  string
	Field []GoField
}

type GoField struct {
	Name            string
	GoName          string
	Description     string
	TypeName        string
	FieldTypeKind   FieldTypeKind
	FieldTypeIsNull bool
	ListType        *listType
}

func getGqlField(fields []*model.Field) []*GoField {

	goFields := []*GoField{}

	for _, field := range fields {
		var fieldListType *listType
		isNull := true
		typeName := ""
		fieldName := field.Name

		fieldType := field.Type.ToType()
		if fieldType.Kind != nil {
			if *fieldType.Kind == nonNullKind {
				fieldType = fieldType.OfType
				isNull = false
			}
		}

		if fieldType.Kind != nil {

			if *fieldType.Kind == listKind {
				fieldType = fieldType.OfType
				if isNull {
					n := nullList
					fieldListType = &n

				} else {
					n := nonNullList
					fieldListType = &n
					isNull = true
				}

			}
		}

		if fieldType.Kind != nil {
			if *fieldType.Kind == nonNullKind {
				fieldType = fieldType.OfType
				isNull = false
			}
		}

		fieldTypeKind := scalarFieldType
		if fieldType.Name == nil {
			fmt.Println("error : field type is empty")
			continue
		}

		typeName = *fieldType.Name

		switch *fieldType.Kind {
		case scalarKind:
			fieldTypeKind = scalarFieldType
			// convert scalar
			typeConverter := NewTypeConverter()
			typeName = typeConverter.GetGoTypeFromGqlType(typeName)

		case objectKind:
			fieldTypeKind = objectFieldType

		case inputObjectKind:
			fieldTypeKind = objectFieldType

		case interfaceKind:
			fieldTypeKind = interfaceFieldType

		case enumKind:
			fieldTypeKind = enumFieldType

		default:
			panic("")

		}

		goFields = append(goFields, &GoField{
			Name:            fieldName,
			GoName:          strcase.ToCamel(fieldName),
			Description:     field.Description,
			TypeName:        typeName,
			FieldTypeKind:   fieldTypeKind,
			FieldTypeIsNull: isNull,
			ListType:        fieldListType,
		})
	}

	return goFields
}

type FieldTypeKind int

const (
	objectFieldType FieldTypeKind = iota
	scalarFieldType
	interfaceFieldType
	enumFieldType
)

type listType int

const (
	nullList listType = iota
	nonNullList
)
