package gopoet

type StructSpec struct {
	// CodeBlock

	Name    string
	Comment string
	Fields  []IdentifierField
	Methods []MethodSpec
}

var _ TypeReference = (*StructSpec)(nil)

func NewStructSpec(name string) *StructSpec {
	return &StructSpec{
		Name: name,
	}
}

func (s *StructSpec) GetImports() []Import {
	return []Import{}
}

func (s *StructSpec) GetName() string {
	return s.Name
}

func (s *StructSpec) String() string {
	writer := NewCodeWriter()
	writer.WriteStatement(Statement{
		Format:      "type $L struct {",
		Arguments:   []interface{}{s.Name},
		AfterIndent: 1,
	})

	for _, field := range s.Fields {
		writer.WriteStatement(Statement{
			Format:    "$L $T",
			Arguments: []interface{}{field.Name, field.Type},
		})
	}

	writer.WriteStatement(Statement{
		Format:       "}",
		BeforeIndent: -1,
	})

	if len(s.Methods) != 0 {
		writer.WriteCode("\n")
	}

	for _, method := range s.Methods {
		writer.WriteCode(method.String() + "\n")
	}
	return writer.String()
}

func (s *StructSpec) Packages() []TypeReference {
	packages := []TypeReference{}

	for _, field := range s.Fields {
		packages = append(packages, field.Type)
	}

	return packages
}

func (s *StructSpec) StructComment(comment string) *StructSpec {
	s.Comment = comment
	return s
}

func (s *StructSpec) Field(name string, typeRef TypeReference) *StructSpec {
	s.Fields = append(s.Fields, IdentifierField{
		Identifier: Identifier{
			Type: typeRef,
			Name: name,
		},
	})
	return s
}

func (s *StructSpec) FieldWithTag(name string, typeRef TypeReference, tag string) *StructSpec {
	s.Fields = append(s.Fields, IdentifierField{
		Identifier: Identifier{
			Type: typeRef,
			Name: name,
		},
		Tag: tag,
	})
	return s
}

func (s *StructSpec) Method(name, receiverName string, isValueReceiver bool) *MethodSpec {
	return NewMethodSpec(name, receiverName, isValueReceiver, s)
}

func (s *StructSpec) MethodAndAttach(name, receiverName string, isValueReceiver bool) *MethodSpec {
	method := NewMethodSpec(name, receiverName, isValueReceiver, s)
	s.Methods = append(s.Methods, *method)
	return method
}
