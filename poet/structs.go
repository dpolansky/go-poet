package poet

type StructSpec struct {
	Name    string
	Comment string
	Fields  []IdentifierField
	Methods []*MethodSpec
}

var _ TypeReference = (*StructSpec)(nil)
var _ CodeBlock = (*StructSpec)(nil)

func NewStructSpec(name string) *StructSpec {
	return &StructSpec{
		Name: name,
	}
}

func (s *StructSpec) GetImports() []Import {
	imports := []Import{}

	for _, f := range s.Fields {
		imports = append(imports, f.Type.GetImports()...)
	}

	return imports
}

func (s *StructSpec) GetName() string {
	return s.Name
}

func (s *StructSpec) String() string {
	writer := newCodeWriter()

	if s.Comment != "" {
		writer.WriteCode("// " + s.Comment + "\n")
	}

	writer.WriteStatement(statement{
		Format:      "type $L struct {",
		Arguments:   []interface{}{s.Name},
		AfterIndent: 1,
	})

	for _, field := range s.Fields {
		var format string
		arguments := []interface{}{field.Name, field.Type}

		if field.Tag != "" {
			format = "$L $T `$L`"
			arguments = append(arguments, field.Tag)
		} else {
			format = "$L $T"
		}

		writer.WriteStatement(statement{
			Format:    format,
			Arguments: arguments,
		})
	}

	writer.WriteStatement(statement{
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

func (s *StructSpec) AttachFunction(receiverName string, funcSpec *FuncSpec) *StructSpec {
	method := &MethodSpec{
		FuncSpec:     *funcSpec,
		ReceiverName: receiverName,
		Receiver:     s.TypeReferenceAsPointer(),
	}
	s.Methods = append(s.Methods, method)
	return s
}

func (s *StructSpec) Method(name, receiverName string) *MethodSpec {
	return NewMethodSpec(name, receiverName, s)
}

func (s *StructSpec) MethodAndAttach(name, receiverName string) *MethodSpec {
	method := NewMethodSpec(name, receiverName, s)
	s.Methods = append(s.Methods, method)
	return method
}

type structSpecAsPointer struct {
	StructSpec
}

func (sP *structSpecAsPointer) GetName() string {
	return "*" + sP.Name
}

func (s *StructSpec) TypeReferenceAsPointer() TypeReference {
	return &structSpecAsPointer{*s}
}
