package gopoet

type StructSpec struct {
	// CodeBlock

	Name    string
	Comment string
	Fields  []IdentifierField
}

func NewStructSpec(name string) *StructSpec {
	return &StructSpec{
		Name: name,
	}
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

	return writer.String()
}

func (s *StructSpec) Packages() []ImportSpec {
	packages := []ImportSpec{}

	for _, field := range s.Fields {
		packages = append(packages, field.Type)
	}

	return packages
}

func (s *StructSpec) StructComment(comment string) *StructSpec {
	s.Comment = comment
	return s
}

func (s *StructSpec) Field(name string, spec ImportSpec) *StructSpec {
	s.Fields = append(s.Fields, IdentifierField{
		Identifier: Identifier{
			Type: spec,
			Name: name,
		},
	})
	return s
}

func (s *StructSpec) FieldWithTag(name string, spec ImportSpec, tag string) *StructSpec {
	s.Fields = append(s.Fields, IdentifierField{
		Identifier: Identifier{
			Type: spec,
			Name: name,
		},
		Tag: tag,
	})
	return s
}
