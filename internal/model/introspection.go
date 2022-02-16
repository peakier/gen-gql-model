package model

type QuerySchema struct {
	GraphQLSchema struct {
		Types        []*Types     `json:"types"`
		QueryType    QueryType    `json:"queryType"`
		MutationType MutationType `json:"mutationType"`
	} `graphql:"__schema" json:"__schema"`
}

type QueryType struct {
	Name   string   `json:"name"`
	Fields []*Field `json:"field"`
}

type EnumValue struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type MutationType struct {
	Name   string   `json:"name"`
	Fields []*Field `json:"field"`
}

type Types struct {
	Fields      []*Field      `json:"fields"`
	InputFields []*InputField `json:"inputFields"`
	Kind        *string       `json:"kind"`
	Name        string        `json:"name"`
	EnumValues  []*EnumValue  `json:"enumValues"`
}

type InputField struct {
	Description string `json:"description"`
	Name        string `json:"name"`
	Type        *Type1 `json:"type"`
}

func (i *InputField) ToField() *Field {
	return &Field{
		Description: i.Description,
		Name:        i.Name,
		Type:        i.Type,
	}
}

type Field struct {
	Description string `json:"description"`
	Name        string `json:"name"`
	Args        []*Arg `json:"args"`
	Type        *Type1 `json:"type"`
}

type Arg struct {
	Name         string `json:"name"`
	DefaultValue string `json:"defaultValue"`
	Type         *Type1 `json:"type"`
}

type Type struct {
	Kind   *string `json:"kind"`
	Name   *string `json:"name"`
	OfType *Type   `json:"ofType"`
}

type Type1 struct {
	Kind   *string `json:"kind"`
	Name   *string `json:"name"`
	OfType *Type2  `json:"ofType"`
}

func (t *Type1) ToType() *Type {
	var ofType *Type
	if t.OfType != nil {
		ofType = t.OfType.toType()
	}
	return &Type{
		Kind:   t.Kind,
		Name:   t.Name,
		OfType: ofType,
	}
}

type Type2 struct {
	Kind   *string `json:"kind"`
	Name   *string `json:"name"`
	OfType *Type3  `json:"ofType"`
}

func (t *Type2) toType() *Type {
	var ofType *Type
	if t.OfType != nil {
		ofType = t.OfType.toType()
	}
	return &Type{
		Kind:   t.Kind,
		Name:   t.Name,
		OfType: ofType,
	}
}

type Type3 struct {
	Kind   *string `json:"kind"`
	Name   *string `json:"name"`
	OfType *Type4  `json:"ofType"`
}

func (t *Type3) toType() *Type {
	var ofType *Type
	if t.OfType != nil {
		ofType = t.OfType.toType()
	}
	return &Type{
		Kind:   t.Kind,
		Name:   t.Name,
		OfType: ofType,
	}
}

type Type4 struct {
	Kind *string `json:"kind"`
	Name *string `json:"name"`
}

func (t *Type4) toType() *Type {
	return &Type{
		Kind: t.Kind,
		Name: t.Name,
	}
}

// func (q *QuerySchema) ReStructType() {
// 	for _, field := range q.GraphQLSchema.QueryType.Fields {
// 		field.Type = field.TypeQuery.toType()
// 		for _, arg := range field.Args {
// 			arg.Type = arg.TypeQuery.toType()
// 		}
// 	}
// }

const QueryIntrospect = `
{
	__schema {
	 queryType {
	   name
	   fields {
		 name
		 type{
			 name
		 kind
		 ofType{
		   name
		   kind
		   ofType{
			 name
			 kind
				ofType{
			 name
			 kind            
				 ofType{
			 name
			 kind
		   }
		   }
		   }
		 }
		 }
		 args {
		   name
		   defaultValue
				type{
		   name
		   kind
		   ofType{
			 name
			 kind
			 ofType{
			   name
			   kind
			   ofType{
				 name
				 kind
				 ofType{
						 name
						 kind
					   }
			   }
			 }
		   }
		 }
		 }
	   }
	 }
	 mutationType {
	   name
	   fields {
		 name
				 type{
			 name
		 kind
		 ofType{
		   name
		   kind
		   ofType{
			 name
			 kind
		   }
			   ofType{
			 name
			 kind
							   ofType{
			 name
			 kind
		   }
		   }
		 }
		 }
		 args {
		   name
		   defaultValue
				type{
		   name
		   kind
		   ofType{
			 name
			 kind
			 ofType{
			   name
			   kind
			   ofType{
				 name
				 kind
			   }
			 }
		   }
		 }
		 }
	   }
	 }
	 types{
	   name
	   kind
	   enumValues{
		 name
		 description
	   }
	   ofType{
		 name
		 kind
		 ofType{
		   name
		   kind
		   ofType{
			 name
			 kind
		   }
		 }
	   }
	   fields{
		 name
		 description
		 type{
		   name
		   kind
		   ofType{
			 name
			 kind
			 ofType{
			   name
			   kind
			   ofType{
				 name
				 kind
			   }
			 }
		   }
		 }
	   }
	   inputFields{
		 name
		 description
		 type{
		   name
		   kind
		   ofType{
			 name
			 kind
			 ofType{
			   name
			   kind
			   ofType{
				 name
				 kind
			   }
			 }
		   }
		 }
	   }
	 }
   }
   
 }
 
`
