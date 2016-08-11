package poet

// StructSpec represents a struct
type StructSpec struct {
	Name    string
	Comment string
	Fields  []IdentifierField
	Methods []*MethodSpec
}

var _ TypeReference = (*StructSpec)(nil)
var _ CodeBlock = (*StructSpec)(nil)

// NewStructSpec creates a new struct with the given type name
func NewStructSpec(name string) *StructSpec {
	return &StructSpec{
		Name: name,
	}
}

// GetImports returns a slice of imports needed by this struct
func (s *StructSpec) GetImports() []Import {
	imports := []Import{}

	for _, f := range s.Fields {
		imports = append(imports, f.Type.GetImports()...)
	}

	return imports
}

// GetName returns the name of this struct's type
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

// StructComment adds a comment to this struct.
func (s *StructSpec) StructComment(comment string) *StructSpec {
	s.Comment = comment
	return s
}

// Field adds a field to this struct.
func (s *StructSpec) Field(name string, typeRef TypeReference) *StructSpec {
	s.Fields = append(s.Fields, IdentifierField{
		Identifier: Identifier{
			Type: typeRef,
			Name: name,
		},
	})
	return s
}

// FieldWithTag adds a field to this struct with a tag on the field.
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

// MethodFromFunction creates a method from a FuncSpec and adds this struct as the receiver.
func (s *StructSpec) MethodFromFunction(receiverName string, receiverIsPtr bool, funcSpec *FuncSpec) *MethodSpec {
	return &MethodSpec{
		FuncSpec:     *funcSpec,
		ReceiverName: receiverName,
		Receiver:     s.getTypeReference(receiverIsPtr),
	}
}

// Method creates a new method spec with this struct as the receiver.
func (s *StructSpec) Method(name, receiverName string, receiverIsPtr bool) *MethodSpec {
	return NewMethodSpec(name, receiverName, s.getTypeReference(receiverIsPtr))
}

// AttachMethod attaches a MethodSpec to this struct, such that a call to String() on this struct
// will output attached methods next to this struct. This is useful for having a method placed
// next to the struct it belongs to.
func (s *StructSpec) AttachMethod(m *MethodSpec) *StructSpec {
	s.Methods = append(s.Methods, m)
	return s
}

func (s *StructSpec) getTypeReference(isPtr bool) TypeReference {
	if isPtr {
		return s.typeReferenceAsPointer()
	}
	return s
}

type structSpecAsPointer struct {
	StructSpec
}

func (sP *structSpecAsPointer) GetName() string {
	return "*" + sP.Name
}

func (s *StructSpec) typeReferenceAsPointer() TypeReference {
	return &structSpecAsPointer{*s}
}
